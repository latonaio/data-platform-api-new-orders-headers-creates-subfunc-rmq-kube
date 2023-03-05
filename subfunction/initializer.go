package subfunction

import (
	"context"
	api_input_reader "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Output_Formatter"
	api_processing_data_formatter "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"sync"

	database "github.com/latonaio/golang-mysql-network-connector"
	"golang.org/x/xerrors"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type SubFunction struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewSubFunction(ctx context.Context, db *database.Mysql, l *logger.Logger) *SubFunction {
	return &SubFunction{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (f *SubFunction) MetaData(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.MetaData {
	metaData := psdc.ConvertToMetaData(sdc)

	return metaData
}

func (f *SubFunction) OrderRegistrationType(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.OrderRegistrationType {
	registrationType := "Direct Registration"
	if sdc.InputParameters.ReferenceDocument != nil && sdc.InputParameters.ReferenceDocumentItem != nil {
		registrationType = "Reference Registration"
	}

	data := psdc.ConvertToOrderRegistrationType(sdc, registrationType)

	return data
}

func (f *SubFunction) OrderReferenceDocumentType(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.OrderReferenceDocumentType, error) {
	rows, err := f.db.Query(
		`SELECT ServiceLabel, FieldNameWithNumberRange, NumberRangeFrom, NumberRangeTo
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_number_range_number_range_data`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dataQueryGets, err := psdc.ConvertToOrderReferenceDocumentTypeQueryGets(rows)
	if err != nil {
		return nil, err
	}

	data := &api_processing_data_formatter.OrderReferenceDocumentType{}

	for i := 0; i < len(dataQueryGets); i++ {
		if sdc.InputParameters.ReferenceDocument != nil && dataQueryGets[i].NumberRangeFrom != nil && dataQueryGets[i].NumberRangeTo != nil {
			if *sdc.InputParameters.ReferenceDocument >= *dataQueryGets[i].NumberRangeFrom && *sdc.InputParameters.ReferenceDocument <= *dataQueryGets[i].NumberRangeTo {
				data = psdc.ConvertToOrderReferenceDocumentType(dataQueryGets[i])
				break
			}
		}
	}

	return data, err
}

func (f *SubFunction) CreateSdc(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	var err error

	psdc.MetaData = f.MetaData(sdc, psdc)

	// 0-1. ReferenceDocumentおよびReferenceDocumentItemの値によるオーダー登録種別の判定
	psdc.OrderRegistrationType = f.OrderRegistrationType(sdc, psdc)

	if psdc.OrderRegistrationType.RegistrationType == "Direct Registration" {
		err = f.DirectRegistration(sdc, psdc, osdc)
	} else if psdc.OrderRegistrationType.RegistrationType == "Reference Registration" {
		err = f.ReferenceRegistration(sdc, psdc, osdc)
	} else {
		err = xerrors.Errorf("オーダー登録種別の判定でエラーが発生しました。")
	}

	if err != nil {
		return err
	}

	// 99-1-1. CreationDate(Header)
	psdc.CreationDateHeader = f.CreationDateHeader(sdc, psdc)

	// 99-2-1. LastChangeDate(Header)
	psdc.LastChangeDateHeader = f.LastChangeDateHeader(sdc, psdc)

	f.l.Info(psdc)
	err = f.SetValue(sdc, psdc, osdc)
	if err != nil {
		return err
	}

	return nil
}

func (f *SubFunction) DirectRegistration(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	var err error
	var e error

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-0. サプライチェーンリレーションシップマスタでの取引妥当性確認(一般データ)
		psdc.SupplyChainRelationshipGeneral, e = f.SupplyChainRelationshipGeneral(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-1. サプライチェーンリレーションシップマスタでの取引妥当性確認(入出荷関係データ)  //1-0
		psdc.SupplyChainRelationshipDeliveryRelation, e = f.SupplyChainRelationshipDeliveryRelation(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-2. サプライチェーンリレーションシップマスタ入出荷プラント関係データの取得  //1-1
		psdc.SupplyChainRelationshipDeliveryPlantRelation, e = f.SupplyChainRelationshipDeliveryPlantRelation(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-3. サプライチェーンリレーションシップマスタ取引データの取得  //1-0
		psdc.SupplyChainRelationshipTransaction, e = f.SupplyChainRelationshipTransaction(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-4. サプライチェーンリレーションシップマスタ請求関係データの取得  //1-0
		psdc.SupplyChainRelationshipBillingRelation, e = f.SupplyChainRelationshipBillingRelation(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-5. サプライチェーンリレーションシップマスタ支払関係データの取得  //1-4
		psdc.SupplyChainRelationshipPaymentRelation, e = f.SupplyChainRelationshipPaymentRelation(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-7. InvoiceDocumentDate  //1-3
		psdc.InvoiceDocumentDate, e = f.InvoiceDocumentDate(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-8. PaymentDueDate  //1-7
		psdc.PaymentDueDate, e = f.PaymentDueDate(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-9. NetPaymentDays  //1-7,1-8
		psdc.NetPaymentDays, e = f.NetPaymentDays(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-12. PricingDate
		psdc.PricingDate = f.PricingDate(sdc, psdc)

		// 2-1. ProductTaxClassificationBillToCountry  //1-4
		psdc.ProductTaxClassificationBillToCountry, e = f.ProductTaxClassificationBillToCountry(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 2-2. ProductTaxClassificationBillFromCountry  //1-4
		psdc.ProductTaxClassificationBillFromCountry, e = f.ProductTaxClassificationBillFromCountry(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 2-3. DefinedTaxClassification  //2-1,2-2
		psdc.DefinedTaxClassification, e = f.DefinedTaxClassification(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 2-20. TaxCode  //1-4,2-3
		psdc.TaxCode, e = f.TaxCode(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		//2-21. TaxRateの計算  //2-20
		psdc.TaxRate, e = f.TaxRate(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 8-1. 価格マスタデータの取得(入力ファイルの[ConditionAmount]がnullである場合)  //1-0,1-12
		psdc.PriceMaster, e = f.PriceMaster(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 8-2. 価格の計算(入力ファイルの[ConditionAmount]がnullである場合)  //8-1
		psdc.ConditionAmount, e = f.ConditionAmount(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 9-1. NetAmount  //8-2
		psdc.NetAmount = f.NetAmount(sdc, psdc)

		// 11-1. TotalNetAmount
		psdc.TotalNetAmount, e = f.TotalNetAmount(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 9-2. TaxAmount  //2-20,2-21,9-1
		psdc.TaxAmount, e = f.TaxAmount(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 11-2. TotalTaxAmount
		psdc.TotalTaxAmount, e = f.TotalTaxAmount(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 9-3. GrossAmount  // 9-1,9-2
		psdc.GrossAmount, e = f.GrossAmount(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 11-3. TotalGrossAmount  // 9-3
		psdc.TotalGrossAmount, e = f.TotalGrossAmount(sdc, psdc)
		if e != nil {
			err = e
			return
		}

	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-13. PriceDetnExchangeRate
		psdc.PriceDetnExchangeRate = f.PriceDetnExchangeRate(sdc, psdc)

		// 1-14. AccountingExchangeRate
		psdc.AccountingExchangeRate = f.AccountingExchangeRate(sdc, psdc)
	}(&wg)

	wg.Wait()

	return err
}

func (f *SubFunction) ReferenceRegistration(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	var err error

	// 0-2. ReferenceDocumentの値によるオーダー参照先伝票種別の判定
	psdc.OrderReferenceDocumentType, err = f.OrderReferenceDocumentType(sdc, psdc)
	if err != nil {
		return err
	}

	// // 1-11. HeaderDocReferenceStatus
	// psdc.HeaderDocReferenceStatus = f.HeaderDocReferenceStatus(sdc, psdc)

	return err
}
