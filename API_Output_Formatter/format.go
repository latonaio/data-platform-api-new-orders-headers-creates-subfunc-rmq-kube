package dpfm_api_output_formatter

import (
	api_input_reader "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToHeader(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*Header, error) {
	var err error

	header := &Header{}
	inputHeader := sdc.Header

	// 入力ファイル
	header, err = jsonTypeConversion(header, inputHeader)
	if err != nil {
		return nil, xerrors.Errorf("request create error: %w", err)
	}

	// 1-0
	header, err = jsonTypeConversion(header, psdc.SupplyChainRelationshipGeneral[0])
	if err != nil {
		return nil, xerrors.Errorf("request create error: %w", err)
	}

	// 1-3
	header, err = jsonTypeConversion(header, psdc.SupplyChainRelationshipTransaction[0])
	if err != nil {
		return nil, xerrors.Errorf("request create error: %w", err)
	}

	// 1-4
	header, err = jsonTypeConversion(header, psdc.SupplyChainRelationshipBillingRelation[0])
	if err != nil {
		return nil, xerrors.Errorf("request create error: %w", err)
	}

	// 1-5
	header, err = jsonTypeConversion(header, psdc.SupplyChainRelationshipPaymentRelation[0])
	if err != nil {
		return nil, xerrors.Errorf("request create error: %w", err)
	}

	header.CreationDate = psdc.CreationDateHeader.CreationDate
	header.LastChangeDate = psdc.LastChangeDateHeader.LastChangeDate
	header.TotalNetAmount = psdc.TotalNetAmount.TotalNetAmount
	header.TotalTaxAmount = psdc.TotalTaxAmount.TotalTaxAmount
	header.TotalGrossAmount = psdc.TotalGrossAmount.TotalGrossAmount
	header.HeaderDeliveryStatus = "NP"
	header.HeaderBillingStatus = "NP"
	header.PricingDate = psdc.PricingDate.PricingDate
	header.PriceDetnExchangeRate = psdc.PriceDetnExchangeRate.PriceDetnExchangeRate
	header.HeaderCompleteDeliveryIsDefined = getBoolPtr(false)
	header.InvoiceDocumentDate = psdc.InvoiceDocumentDate.InvoiceDocumentDate
	header.HeaderBlockStatus = getBoolPtr(false)
	header.HeaderBillingBlockStatus = getBoolPtr(false)
	header.HeaderDeliveryBlockStatus = getBoolPtr(false)
	header.HeaderDeliveryBlockStatus = getBoolPtr(false)
	header.HeaderDeliveryBlockStatus = getBoolPtr(false)
	header.IsCancelled = getBoolPtr(false)
	header.IsMarkedForDeletion = getBoolPtr(false)

	return header, nil
}

func getBoolPtr(b bool) *bool {
	return &b
}

func jsonTypeConversion[T any](dist T, data interface{}) (T, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
