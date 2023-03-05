package subfunction

import (
	api_input_reader "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"math"
	"strconv"
	"strings"

	"golang.org/x/xerrors"
)

func (f *SubFunction) PriceMaster(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.PriceMaster, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToPriceMasterKey()

	for _, v := range sdc.Header.Item {
		if v.ItemPricingElement[0].ConditionAmount == nil {
			dataKey.Product = append(dataKey.Product, v.Product)
		}
	}

	if dataKey.Product == nil {
		return nil, nil
	}

	dataKey.SupplyChainRelationshipID = psdc.SupplyChainRelationshipGeneral[0].SupplyChainRelationshipID
	dataKey.Buyer = psdc.SupplyChainRelationshipGeneral[0].Buyer
	dataKey.Seller = psdc.SupplyChainRelationshipGeneral[0].Seller
	dataKey.ConditionValidityEndDate = psdc.PricingDate.PricingDate
	dataKey.ConditionValidityStartDate = psdc.PricingDate.PricingDate

	if len(dataKey.Product) == 0 {
		return nil, xerrors.Errorf("入力ファイルから取得した'Product'がありません。")
	}
	repeat := strings.Repeat("?,", len(dataKey.Product)-1) + "?"
	for _, v := range dataKey.Product {
		args = append(args, v)
	}

	args = append(args, dataKey.SupplyChainRelationshipID, dataKey.Buyer, dataKey.Seller, dataKey.ConditionValidityEndDate, dataKey.ConditionValidityStartDate)
	rows, err := f.db.Query(
		`SELECT SupplyChainRelationshipID, Buyer, Seller, ConditionRecord, ConditionSequentialNumber,
		ConditionValidityStartDate, ConditionValidityEndDate, Product, ConditionType, ConditionRateValue
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_price_master_price_master_data
		WHERE Product IN ( `+repeat+` )
		AND (SupplyChainRelationshipID, Buyer, Seller) = (?, ?, ?)
		AND ConditionValidityEndDate >= ?
		AND ConditionValidityStartDate <= ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToPriceMaster(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) ConditionAmount(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.ConditionAmount, error) {
	data := make([]*api_processing_data_formatter.ConditionAmount, 0)

	priceMaster := psdc.PriceMaster
	priceMasterMap := make(map[string]*api_processing_data_formatter.PriceMaster, len(priceMaster))
	for _, v := range priceMaster {
		priceMasterMap[v.Product] = v
	}

	for _, v := range sdc.Header.Item {
		orderItem := v.OrderItem
		product := v.Product
		conditionQuantity := v.OrderQuantityInBaseUnit
		if v.ItemPricingElement[0].ConditionAmount == nil && v.Product != nil {
			conditionRateValue := priceMasterMap[*v.Product].ConditionRateValue
			conditionAmount, err := calculateConditionAmount(conditionQuantity, conditionRateValue)
			if err != nil {
				return nil, err
			}

			if product == nil {
				return nil, xerrors.Errorf("入力ファイルから取得した'Product'がありません。")
			}
			datum := psdc.ConvertToConditionAmount(orderItem, *product, conditionQuantity, conditionAmount)
			data = append(data, datum)
		} else if v.ItemPricingElement[0].ConditionAmount != nil && v.Product != nil {
			conditionAmount := v.ItemPricingElement[0].ConditionAmount

			if product == nil {
				return nil, xerrors.Errorf("入力ファイルから取得した'Product'がありません。")
			}
			datum := psdc.ConvertToConditionAmount(orderItem, *product, conditionQuantity, conditionAmount)
			data = append(data, datum)
		}
	}

	return data, nil
}

func calculateConditionAmount(conditionQuantity, conditionRateValue *float32) (*float32, error) {
	if conditionQuantity == nil || conditionRateValue == nil {
		return nil, xerrors.Errorf("ConditionRateValueまたはConditionQuantityがnullです。")
	}

	digit := float32DecimalDigit(*conditionRateValue)
	mul := *conditionRateValue * *conditionQuantity

	s := math.Round(float64(mul)*math.Pow10(digit)) / math.Pow10(digit)
	res := parseFloat32Ptr(float32(s))

	return res, nil
}

func float32DecimalDigit(f float32) int {
	s := strconv.FormatFloat(float64(f), 'f', -1, 32)

	i := strings.Index(s, ".")
	if i == -1 {
		return 0
	}

	return len(s[i+1:])
}

func parseFloat32Ptr(f float32) *float32 {
	return &f
}
