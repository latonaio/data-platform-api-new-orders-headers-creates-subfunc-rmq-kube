package subfunction

import (
	api_input_reader "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"math"
	"strings"

	"golang.org/x/xerrors"
)

func (f *SubFunction) ProductTaxClassificationBillToCountry(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.ProductTaxClassificationBillToCountry, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToProductTaxClassificationKey()

	for _, v := range sdc.Header.Item {
		dataKey.Product = append(dataKey.Product, v.Product)
	}

	dataKey.Country = psdc.SupplyChainRelationshipBillingRelation[0].BillToCountry

	if len(dataKey.Product) == 0 {
		return nil, xerrors.Errorf("入力ファイルから取得した'Product'がありません。")
	}
	repeat := strings.Repeat("?,", len(dataKey.Product)-1) + "?"
	for _, v := range dataKey.Product {
		args = append(args, v)
	}

	args = append(args, dataKey.Country, dataKey.ProductTaxCategory)
	rows, err := f.db.Query(
		`SELECT Product, Country, ProductTaxCategory, ProductTaxClassification
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_tax_data
		WHERE Product IN ( `+repeat+` )
		AND (Country, ProductTaxCategory) = (?, ?);`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToProductTaxClassificationBillToCountry(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) ProductTaxClassificationBillFromCountry(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.ProductTaxClassificationBillFromCountry, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToProductTaxClassificationKey()

	for _, v := range sdc.Header.Item {
		dataKey.Product = append(dataKey.Product, v.Product)
	}

	dataKey.Country = psdc.SupplyChainRelationshipBillingRelation[0].BillFromCountry

	if len(dataKey.Product) == 0 {
		return nil, xerrors.Errorf("入力ファイルから取得した'Product'がありません。")
	}
	repeat := strings.Repeat("?,", len(dataKey.Product)-1) + "?"
	for _, v := range dataKey.Product {
		args = append(args, v)
	}

	args = append(args, dataKey.Country, dataKey.ProductTaxCategory)
	rows, err := f.db.Query(
		`SELECT Product, Country, ProductTaxCategory, ProductTaxClassification
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_tax_data
		WHERE Product IN ( `+repeat+` )
		AND (Country, ProductTaxCategory) = (?, ?);`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToProductTaxClassificationBillFromCountry(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) DefinedTaxClassification(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.DefinedTaxClassification, error) {
	data := make([]*api_processing_data_formatter.DefinedTaxClassification, 0)
	var err error

	transactionTaxClassification := psdc.SupplyChainRelationshipBillingRelation[0].TransactionTaxClassification

	productTaxClassificationBillFromCountry := psdc.ProductTaxClassificationBillFromCountry

	productTaxClassificationBillFromCountryMap := make(map[string]*api_processing_data_formatter.ProductTaxClassificationBillFromCountry, len(productTaxClassificationBillFromCountry))
	for _, v := range productTaxClassificationBillFromCountry {
		productTaxClassificationBillFromCountryMap[v.Product] = v
	}

	for _, v := range psdc.ProductTaxClassificationBillToCountry {
		var definedTaxClassification string

		product := v.Product
		productTaxClassificationBillToCountry := v.ProductTaxClassifiication
		productTaxClassificationBillFromCountry := productTaxClassificationBillFromCountryMap[v.Product].ProductTaxClassifiication

		if transactionTaxClassification == nil || productTaxClassificationBillToCountry == nil || productTaxClassificationBillFromCountry == nil {
			return nil, xerrors.Errorf("TransactionTaxClassificationまたはProductTaxClassificationBillToCountryまたはProductTaxClassificationBillFromCountryがnullです。")
		}
		if *transactionTaxClassification == "1" && *productTaxClassificationBillToCountry == "1" && *productTaxClassificationBillFromCountry == "1" {
			definedTaxClassification = "1"
		} else {
			definedTaxClassification = "0"
		}

		datum := psdc.ConvertToDefinedTaxClassification(product, transactionTaxClassification, productTaxClassificationBillToCountry, productTaxClassificationBillFromCountry, definedTaxClassification)
		data = append(data, datum)
	}

	return data, err
}

func (f *SubFunction) TaxCode(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.TaxCode, error) {
	data := make([]*api_processing_data_formatter.TaxCode, 0)
	var err error

	isExportImport := psdc.SupplyChainRelationshipBillingRelation[0].IsExportImport

	for _, v := range psdc.DefinedTaxClassification {
		taxCode := new(string)

		product := v.Product
		definedTaxClassification := v.DefinedTaxClassification

		if isExportImport == nil {
			return nil, xerrors.Errorf("IsExportImportがnullです。")
		}
		if definedTaxClassification == "1" && !*isExportImport {
			taxCode = getStringPtr("1")
		} else if definedTaxClassification == "0" && !*isExportImport {
			taxCode = getStringPtr("8")
		} else if definedTaxClassification == "0" && *isExportImport {
			taxCode = getStringPtr("9")
		} else {
			taxCode = getStringPtr("0")
		}

		datum := psdc.ConvertToTaxCode(product, definedTaxClassification, isExportImport, taxCode)
		data = append(data, datum)
	}

	return data, err

}

func (f *SubFunction) TaxRate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.TaxRate, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToTaxRateKey()

	for _, v := range psdc.TaxCode {
		dataKey.TaxCode = append(dataKey.TaxCode, v.TaxCode)
	}

	dataKey.ValidityEndDate = getSystemDate()
	dataKey.ValidityStartDate = getSystemDate()

	if len(dataKey.TaxCode) == 0 {
		return nil, xerrors.Errorf("2-20-1でセットした'TaxCode'がありません。")
	}
	repeat := strings.Repeat("?,", len(dataKey.TaxCode)-1) + "?"
	args = append(args, dataKey.Country)
	for _, v := range dataKey.TaxCode {
		args = append(args, v)
	}

	args = append(args, dataKey.ValidityEndDate, dataKey.ValidityStartDate)
	rows, err := f.db.Query(
		`SELECT Country, TaxCode, ValidityEndDate, ValidityStartDate, TaxRate
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_tax_code_tax_rate_data
		WHERE Country = ?
		AND TaxCode IN ( `+repeat+` )
		AND ValidityEndDate >= ?
		AND ValidityStartDate <= ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToTaxRate(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) NetAmount(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.NetAmount {
	data := psdc.ConvertToNetAmount(psdc.ConditionAmount)

	return data
}

func (f *SubFunction) TaxAmount(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.TaxAmount, error) {
	data := make([]*api_processing_data_formatter.TaxAmount, 0)

	item := sdc.Header.Item
	itemMap := make(map[int]api_input_reader.Item, len(item))
	for _, v := range item {
		itemMap[v.OrderItem] = v
	}

	taxRate := psdc.TaxRate
	taxRateMap := make(map[string]*api_processing_data_formatter.TaxRate, len(taxRate))
	for _, v := range taxRate {
		taxRateMap[v.TaxCode] = v
	}

	netAmount := psdc.NetAmount
	netAmountMap := make(map[string]*api_processing_data_formatter.NetAmount, len(netAmount))
	for _, v := range netAmount {
		netAmountMap[v.Product] = v
	}

	taxCode := psdc.TaxCode
	taxCodeMap := make(map[string]*api_processing_data_formatter.TaxCode, len(taxCode))
	for _, v := range taxCode {
		taxCodeMap[v.Product] = v
	}
	for _, v := range psdc.NetAmount {
		taxAmount := new(float32)
		product := v.Product

		if taxCodeMap[product].TaxCode == nil {
			continue
		}
		taxCode := *taxCodeMap[product].TaxCode
		if taxCode == "1" {
			taxAmount, _ = calculateTaxAmount(taxRateMap[taxCode].TaxRate, netAmountMap[v.Product].NetAmount)
		} else {
			taxAmount = parseFloat32Ptr(0)
		}

		if itemMap[v.OrderItem].TaxAmount == nil {
			datum := psdc.ConvertToTaxAmount(v.OrderItem, v.Product, taxCode, taxRateMap[taxCode].TaxRate, netAmountMap[v.Product].NetAmount, taxAmount)
			data = append(data, datum)
		} else {
			datum := psdc.ConvertToTaxAmount(v.OrderItem, v.Product, taxCode, taxRateMap[taxCode].TaxRate, netAmountMap[v.Product].NetAmount, itemMap[v.OrderItem].TaxAmount)
			data = append(data, datum)
			if differenceIsOver(*taxAmount, *itemMap[v.OrderItem].TaxAmount, 2) {
				return nil, xerrors.Errorf("TaxAmountについて入力ファイルの値と計算結果の差の絶対値が2以上の明細が一つ以上存在します。")
			}
		}
	}

	return data, nil
}

func (f *SubFunction) GrossAmount(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.GrossAmount, error) {
	data := make([]*api_processing_data_formatter.GrossAmount, 0)

	item := sdc.Header.Item
	itemMap := make(map[int]api_input_reader.Item, len(item))
	for _, v := range item {
		itemMap[v.OrderItem] = v
	}

	for _, v := range psdc.TaxAmount {
		if v.NetAmount == nil || v.TaxAmount == nil {
			return nil, xerrors.Errorf("NetAmountまたはTaxAmountがnullです。")
		}
		grossAmount := parseFloat32Ptr(*v.NetAmount + *v.TaxAmount)

		if itemMap[v.OrderItem].GrossAmount == nil {
			datum := psdc.ConvertToGrossAmount(v.OrderItem, v.Product, v.NetAmount, v.TaxAmount, grossAmount)
			data = append(data, datum)
		} else {
			datum := psdc.ConvertToGrossAmount(v.OrderItem, v.Product, v.NetAmount, v.TaxAmount, itemMap[v.OrderItem].GrossAmount)
			data = append(data, datum)
			if differenceIsOver(*grossAmount, *itemMap[v.OrderItem].GrossAmount, 2) {
				return nil, xerrors.Errorf("GrossAmountについて入力ファイルの値と計算結果の差の絶対値が2以上の明細が一つ以上存在します。")
			}
		}
	}

	return data, nil
}

func calculateTaxAmount(taxRate *float32, netAmount *float32) (*float32, error) {
	if taxRate == nil || netAmount == nil {
		return nil, xerrors.Errorf("TaxRateまたはNetAmountがnullです。")
	}

	digit := float32DecimalDigit(*netAmount)
	mul := *netAmount * *taxRate / 100

	s := math.Round(float64(mul)*math.Pow10(digit)) / math.Pow10(digit)
	res := parseFloat32Ptr(float32(s))

	return res, nil
}

func differenceIsOver(inputValue, calculatedValue float32, baseValue int) bool {
	return math.Abs(float64(inputValue-calculatedValue)) >= float64(baseValue)
}

func getStringPtr(s string) *string {
	return &s
}
