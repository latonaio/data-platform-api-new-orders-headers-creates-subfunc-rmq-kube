package api_processing_data_formatter

import (
	api_input_reader "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	"data-platform-api-orders-headers-creates-subfunc-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"

	"golang.org/x/xerrors"
)

// Initializer
func (psdc *SDC) ConvertToMetaData(sdc *api_input_reader.SDC) *MetaData {
	pm := &requests.MetaData{
		BusinessPartnerID: sdc.BusinessPartnerID,
		ServiceLabel:      sdc.ServiceLabel,
	}

	data := pm
	res := MetaData{
		BusinessPartnerID: data.BusinessPartnerID,
		ServiceLabel:      data.ServiceLabel,
	}

	return &res
}

func (psdc *SDC) ConvertToOrderRegistrationType(sdc *api_input_reader.SDC, registrationType string) *OrderRegistrationType {
	pm := &requests.OrderRegistrationType{
		ReferenceDocument:     sdc.InputParameters.ReferenceDocument,
		ReferenceDocumentItem: sdc.InputParameters.ReferenceDocumentItem,
	}

	pm.RegistrationType = registrationType

	data := pm
	res := OrderRegistrationType{
		ReferenceDocument:     data.ReferenceDocument,
		ReferenceDocumentItem: data.ReferenceDocumentItem,
		RegistrationType:      data.RegistrationType,
	}

	return &res
}

func (psdc *SDC) ConvertToOrderReferenceDocumentTypeQueryGets(rows *sql.Rows) ([]*OrderReferenceDocumentTypeQueryGets, error) {
	defer rows.Close()
	res := make([]*OrderReferenceDocumentTypeQueryGets, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.OrderReferenceDocumentTypeQueryGets{}

		err := rows.Scan(
			&pm.ServiceLabel,
			&pm.FieldNameWithNumberRange,
			&pm.NumberRangeFrom,
			&pm.NumberRangeTo)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &OrderReferenceDocumentTypeQueryGets{
			ServiceLabel:             data.ServiceLabel,
			FieldNameWithNumberRange: data.FieldNameWithNumberRange,
			NumberRangeFrom:          data.NumberRangeFrom,
			NumberRangeTo:            data.NumberRangeTo,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_number_range_number_range_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToOrderReferenceDocumentType(orderReferenceDocumentTypeQueryGets *OrderReferenceDocumentTypeQueryGets) *OrderReferenceDocumentType {
	pm := &requests.OrderReferenceDocumentType{}

	pm.ServiceLabel = orderReferenceDocumentTypeQueryGets.ServiceLabel

	data := pm
	res := OrderReferenceDocumentType{
		ServiceLabel: data.ServiceLabel,
	}

	return &res
}

// Header
func (psdc *SDC) ConvertToSupplyChainRelationshipGeneral(rows *sql.Rows) ([]*SupplyChainRelationshipGeneral, error) {
	defer rows.Close()
	res := make([]*SupplyChainRelationshipGeneral, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.SupplyChainRelationshipGeneral{}

		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.Buyer,
			&pm.Seller,
		)
		if err != nil {
			return nil, err
		}
		data := pm
		res = append(res, &SupplyChainRelationshipGeneral{
			SupplyChainRelationshipID: data.SupplyChainRelationshipID,
			Buyer:                     data.Buyer,
			Seller:                    data.Seller,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_supply_chain_relationship_general_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToSupplyChainRelationshipDeliveryRelationKey() *SupplyChainRelationshipDeliveryRelationKey {
	pm := &requests.SupplyChainRelationshipDeliveryRelationKey{}

	data := pm
	res := SupplyChainRelationshipDeliveryRelationKey{
		SupplyChainRelationshipID: data.SupplyChainRelationshipID,
		Buyer:                     data.Buyer,
		Seller:                    data.Seller,
		DeliverToParty:            data.DeliverToParty,
		DeliverFromParty:          data.DeliverFromParty,
	}

	return &res
}

func (psdc *SDC) ConvertToSupplyChainRelationshipDeliveryRelation(rows *sql.Rows) ([]*SupplyChainRelationshipDeliveryRelation, error) {
	defer rows.Close()
	res := make([]*SupplyChainRelationshipDeliveryRelation, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.SupplyChainRelationshipDeliveryRelation{}

		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &SupplyChainRelationshipDeliveryRelation{
			SupplyChainRelationshipID:         data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID: data.SupplyChainRelationshipDeliveryID,
			Buyer:                             data.Buyer,
			Seller:                            data.Seller,
			DeliverToParty:                    data.DeliverToParty,
			DeliverFromParty:                  data.DeliverFromParty,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_supply_chain_relationship_delivery_relation_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToSupplyChainRelationshipDeliveryPlantRelationKey() *SupplyChainRelationshipDeliveryPlantRelationKey {
	pm := &requests.SupplyChainRelationshipDeliveryPlantRelationKey{
		DefaultRelation: true,
	}

	data := pm
	res := SupplyChainRelationshipDeliveryPlantRelationKey{
		SupplyChainRelationshipID:         data.SupplyChainRelationshipID,
		SupplyChainRelationshipDeliveryID: data.SupplyChainRelationshipDeliveryID,
		Buyer:                             data.Buyer,
		Seller:                            data.Seller,
		DeliverToParty:                    data.DeliverToParty,
		DeliverFromParty:                  data.DeliverFromParty,
		DefaultRelation:                   data.DefaultRelation,
	}

	return &res
}

func (psdc *SDC) ConvertToSupplyChainRelationshipDeliveryPlantRelation(rows *sql.Rows) ([]*SupplyChainRelationshipDeliveryPlantRelation, error) {
	defer rows.Close()
	res := make([]*SupplyChainRelationshipDeliveryPlantRelation, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.SupplyChainRelationshipDeliveryPlantRelation{}

		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.DeliverToPlant,
			&pm.DeliverFromPlant,
			&pm.DefaultRelation,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &SupplyChainRelationshipDeliveryPlantRelation{
			SupplyChainRelationshipID:              data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:      data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID: data.SupplyChainRelationshipDeliveryPlantID,
			Buyer:                                  data.Buyer,
			Seller:                                 data.Seller,
			DeliverToParty:                         data.DeliverToParty,
			DeliverFromParty:                       data.DeliverFromParty,
			DeliverToPlant:                         data.DeliverToPlant,
			DeliverFromPlant:                       data.DeliverFromPlant,
			DefaultRelation:                        data.DefaultRelation,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_supply_chain_relationship_delivery_plant_relation_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToSupplyChainRelationshipTransaction(rows *sql.Rows) ([]*SupplyChainRelationshipTransaction, error) {
	defer rows.Close()
	res := make([]*SupplyChainRelationshipTransaction, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.SupplyChainRelationshipTransaction{}

		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.Buyer,
			&pm.Seller,
			&pm.TransactionCurrency,
			&pm.Incoterms,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.AccountAssignmentGroup,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &SupplyChainRelationshipTransaction{
			SupplyChainRelationshipID: data.SupplyChainRelationshipID,
			Buyer:                     data.Buyer,
			Seller:                    data.Seller,
			TransactionCurrency:       data.TransactionCurrency,
			Incoterms:                 data.Incoterms,
			PaymentTerms:              data.PaymentTerms,
			PaymentMethod:             data.PaymentMethod,
			AccountAssignmentGroup:    data.AccountAssignmentGroup,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_supply_chain_relationship_transaction_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToSupplyChainRelationshipBillingRelationKey() *SupplyChainRelationshipBillingRelationKey {
	pm := &requests.SupplyChainRelationshipBillingRelationKey{
		DefaultRelation: true,
	}

	data := pm
	res := SupplyChainRelationshipBillingRelationKey{
		SupplyChainRelationshipID: data.SupplyChainRelationshipID,
		Buyer:                     data.Buyer,
		Seller:                    data.Seller,
		DefaultRelation:           data.DefaultRelation,
	}

	return &res
}

func (psdc *SDC) ConvertToSupplyChainRelationshipBillingRelation(rows *sql.Rows) ([]*SupplyChainRelationshipBillingRelation, error) {
	defer rows.Close()
	res := make([]*SupplyChainRelationshipBillingRelation, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.SupplyChainRelationshipBillingRelation{}

		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipBillingID,
			&pm.Buyer,
			&pm.Seller,
			&pm.BillToParty,
			&pm.BillFromParty,
			&pm.DefaultRelation,
			&pm.BillToCountry,
			&pm.BillFromCountry,
			&pm.IsExportImport,
			&pm.TransactionTaxCategory,
			&pm.TransactionTaxClassification,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &SupplyChainRelationshipBillingRelation{
			SupplyChainRelationshipID:        data.SupplyChainRelationshipID,
			SupplyChainRelationshipBillingID: data.SupplyChainRelationshipBillingID,
			Buyer:                            data.Buyer,
			Seller:                           data.Seller,
			BillToParty:                      data.BillToParty,
			BillFromParty:                    data.BillFromParty,
			DefaultRelation:                  data.DefaultRelation,
			BillToCountry:                    data.BillToCountry,
			BillFromCountry:                  data.BillFromCountry,
			IsExportImport:                   data.IsExportImport,
			TransactionTaxCategory:           data.TransactionTaxCategory,
			TransactionTaxClassification:     data.TransactionTaxClassification,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_supply_chain_relationship_delivery_plant_relation_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToSupplyChainRelationshipPaymentRelationKey() *SupplyChainRelationshipPaymentRelationKey {
	pm := &requests.SupplyChainRelationshipPaymentRelationKey{
		DefaultRelation: true,
	}

	data := pm
	res := SupplyChainRelationshipPaymentRelationKey{
		SupplyChainRelationshipID: data.SupplyChainRelationshipID,
		Buyer:                     data.Buyer,
		Seller:                    data.Seller,
		BillToParty:               data.BillToParty,
		BillFromParty:             data.BillFromParty,
		DefaultRelation:           data.DefaultRelation,
	}

	return &res
}

func (psdc *SDC) ConvertToSupplyChainRelationshipPaymentRelation(rows *sql.Rows) ([]*SupplyChainRelationshipPaymentRelation, error) {
	defer rows.Close()
	res := make([]*SupplyChainRelationshipPaymentRelation, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.SupplyChainRelationshipPaymentRelation{}

		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipBillingID,
			&pm.SupplyChainRelationshipPaymentID,
			&pm.Buyer,
			&pm.Seller,
			&pm.BillToParty,
			&pm.BillFromParty,
			&pm.Payer,
			&pm.Payee,
			&pm.DefaultRelation,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &SupplyChainRelationshipPaymentRelation{
			SupplyChainRelationshipID:        data.SupplyChainRelationshipID,
			SupplyChainRelationshipBillingID: data.SupplyChainRelationshipBillingID,
			SupplyChainRelationshipPaymentID: data.SupplyChainRelationshipPaymentID,
			Buyer:                            data.Buyer,
			Seller:                           data.Seller,
			BillToParty:                      data.BillToParty,
			BillFromParty:                    data.BillFromParty,
			Payer:                            data.Payer,
			Payee:                            data.Payee,
			DefaultRelation:                  data.DefaultRelation,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_supply_chain_relationship_payment_relation_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToPaymentTerms(rows *sql.Rows) ([]*PaymentTerms, error) {
	defer rows.Close()
	res := make([]*PaymentTerms, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.PaymentTerms{}

		err := rows.Scan(
			&pm.PaymentTerms,
			&pm.BaseDate,
			&pm.BaseDateCalcAddMonth,
			&pm.BaseDateCalcFixedDate,
			&pm.PaymentDueDateCalcAddMonth,
			&pm.PaymentDueDateCalcFixedDate,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &PaymentTerms{
			PaymentTerms:                data.PaymentTerms,
			BaseDate:                    data.BaseDate,
			BaseDateCalcAddMonth:        data.BaseDateCalcAddMonth,
			BaseDateCalcFixedDate:       data.BaseDateCalcFixedDate,
			PaymentDueDateCalcAddMonth:  data.PaymentDueDateCalcAddMonth,
			PaymentDueDateCalcFixedDate: data.PaymentDueDateCalcFixedDate,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_payment_terms_payment_terms_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToInvoiceDocumentDate(sdc *api_input_reader.SDC) (*InvoiceDocumentDate, error) {
	pm := &requests.InvoiceDocumentDate{}

	if sdc.Header.InvoiceDocumentDate == nil {
		return nil, xerrors.Errorf("InvoiceDocumentDateがnullです。")
	}
	pm.InvoiceDocumentDate = *sdc.Header.InvoiceDocumentDate
	data := pm

	res := InvoiceDocumentDate{
		RequestedDeliveryDate: data.RequestedDeliveryDate,
		InvoiceDocumentDate:   data.InvoiceDocumentDate,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToRequestedDeliveryDate(sdc *api_input_reader.SDC) (*InvoiceDocumentDate, error) {
	if sdc.Header.RequestedDeliveryDate == nil {
		return nil, xerrors.Errorf("RequestedDeliveryDateがnullです。")
	}

	pm := &requests.InvoiceDocumentDate{
		RequestedDeliveryDate: *sdc.Header.RequestedDeliveryDate,
	}

	data := pm
	res := InvoiceDocumentDate{
		RequestedDeliveryDate: data.RequestedDeliveryDate,
		InvoiceDocumentDate:   data.InvoiceDocumentDate,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToCaluculateInvoiceDocumentDate(sdc *api_input_reader.SDC, invoiceDocumentDate string) (*InvoiceDocumentDate, error) {
	if sdc.Header.RequestedDeliveryDate == nil {
		return nil, xerrors.Errorf("RequestedDeliveryDateがnullです。")
	}

	pm := &requests.InvoiceDocumentDate{
		RequestedDeliveryDate: *sdc.Header.RequestedDeliveryDate,
	}

	pm.InvoiceDocumentDate = invoiceDocumentDate

	data := pm
	res := InvoiceDocumentDate{
		RequestedDeliveryDate: data.RequestedDeliveryDate,
		InvoiceDocumentDate:   data.InvoiceDocumentDate,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToPaymentDueDate(paymentDueDate string) *PaymentDueDate {
	pm := &requests.PaymentDueDate{}

	pm.InvoiceDocumentDate = psdc.InvoiceDocumentDate.InvoiceDocumentDate
	pm.PaymentDueDate = paymentDueDate

	data := pm
	res := PaymentDueDate{
		InvoiceDocumentDate: data.InvoiceDocumentDate,
		PaymentDueDate:      data.PaymentDueDate,
	}

	return &res
}

func (psdc *SDC) ConvertToNetPaymentDays(paymentDueDate string, netPaymentDays int) *NetPaymentDays {
	pm := &requests.NetPaymentDays{}

	pm.InvoiceDocumentDate = psdc.InvoiceDocumentDate.InvoiceDocumentDate
	pm.PaymentDueDate = paymentDueDate
	pm.NetPaymentDays = &netPaymentDays

	data := pm
	res := NetPaymentDays{
		InvoiceDocumentDate: data.InvoiceDocumentDate,
		PaymentDueDate:      data.PaymentDueDate,
		NetPaymentDays:      data.NetPaymentDays,
	}

	return &res
}

func (psdc *SDC) ConvertToHeaderDocReferenceStatus(headerDocReferenceStatus string) *HeaderDocReferenceStatus {
	pm := &requests.HeaderDocReferenceStatus{}

	pm.HeaderDocReferenceStatus = headerDocReferenceStatus

	data := pm
	res := HeaderDocReferenceStatus{
		HeaderDocReferenceStatus: data.HeaderDocReferenceStatus,
	}

	return &res
}

func (psdc *SDC) ConvertToQuotationsHeader(rows *sql.Rows) (*QuotationsHeader, error) {
	defer rows.Close()
	pm := &requests.QuotationsHeader{}

	i := 0
	for rows.Next() {
		i++
		err := rows.Scan(
			&pm.Quotation,
			&pm.QuotationType,
			&pm.QuotationStatus,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipBillingID,
			&pm.SupplyChainRelationshipPaymentID,
			&pm.Buyer,
			&pm.Seller,
			&pm.BillToParty,
			&pm.BillFromParty,
			&pm.BillToCountry,
			&pm.BillFromCountry,
			&pm.Payer,
			&pm.Payee,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.ContractType,
			&pm.BindingPeriodValidityStartDate,
			&pm.BindingPeriodValidityEndDate,
			&pm.OrderVaridityStartDate,
			&pm.OrderValidityEndDate,
			&pm.InvoicePeriodStartDate,
			&pm.InvoicePeriodEndDate,
			&pm.TotalNetAmount,
			&pm.TotalTaxAmount,
			&pm.TotalGrossAmount,
			&pm.TransactionCurrency,
			&pm.PricingDate,
			&pm.PriceDetnExchangeRate,
			&pm.RequestedDeliveryDate,
			&pm.OrderProbabilityInPercent,
			&pm.ExpectedOrderNetAmount,
			&pm.Incoterms,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.ReferenceDocument,
			&pm.ReferenceDocumentItem,
			&pm.AccountAssignmentGroup,
			&pm.AccountingExchangeRate,
			&pm.InvoiceDocumentDate,
			&pm.IsExportImport,
			&pm.HeaderText)
		if err != nil {
			return nil, err
		}
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_quotations_header_data'テーブルに対象のレコードが存在しません。")
	}

	data := pm
	res := QuotationsHeader{
		Quotation:                        data.Quotation,
		QuotationType:                    data.QuotationType,
		QuotationStatus:                  data.QuotationStatus,
		SupplyChainRelationshipID:        data.SupplyChainRelationshipID,
		SupplyChainRelationshipBillingID: data.SupplyChainRelationshipBillingID,
		SupplyChainRelationshipPaymentID: data.SupplyChainRelationshipPaymentID,
		Buyer:                            data.Buyer,
		Seller:                           data.Seller,
		BillToParty:                      data.BillToParty,
		BillFromParty:                    data.BillFromParty,
		BillToCountry:                    data.BillToCountry,
		BillFromCountry:                  data.BillFromCountry,
		Payer:                            data.Payer,
		Payee:                            data.Payee,
		CreationDate:                     data.CreationDate,
		LastChangeDate:                   data.LastChangeDate,
		ContractType:                     data.ContractType,
		BindingPeriodValidityStartDate:   data.BindingPeriodValidityStartDate,
		BindingPeriodValidityEndDate:     data.BindingPeriodValidityEndDate,
		OrderVaridityStartDate:           data.OrderVaridityStartDate,
		OrderValidityEndDate:             data.OrderValidityEndDate,
		InvoicePeriodStartDate:           data.InvoicePeriodStartDate,
		InvoicePeriodEndDate:             data.InvoicePeriodEndDate,
		TotalNetAmount:                   data.TotalNetAmount,
		TotalTaxAmount:                   data.TotalTaxAmount,
		TotalGrossAmount:                 data.TotalGrossAmount,
		TransactionCurrency:              data.TransactionCurrency,
		PricingDate:                      data.PricingDate,
		PriceDetnExchangeRate:            data.PriceDetnExchangeRate,
		RequestedDeliveryDate:            data.RequestedDeliveryDate,
		OrderProbabilityInPercent:        data.OrderProbabilityInPercent,
		ExpectedOrderNetAmount:           data.ExpectedOrderNetAmount,
		Incoterms:                        data.Incoterms,
		PaymentTerms:                     data.PaymentTerms,
		PaymentMethod:                    data.PaymentMethod,
		ReferenceDocument:                data.ReferenceDocument,
		ReferenceDocumentItem:            data.ReferenceDocumentItem,
		AccountAssignmentGroup:           data.AccountAssignmentGroup,
		AccountingExchangeRate:           data.AccountingExchangeRate,
		InvoiceDocumentDate:              data.InvoiceDocumentDate,
		IsExportImport:                   data.IsExportImport,
		HeaderText:                       data.HeaderText,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToOrdersHeader(rows *sql.Rows) ([]*OrdersHeader, error) {
	defer rows.Close()
	res := make([]*OrdersHeader, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.OrdersHeader{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderDate,
			&pm.OrderType,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipBillingID,
			&pm.SupplyChainRelationshipPaymentID,
			&pm.Buyer,
			&pm.Seller,
			&pm.BillToParty,
			&pm.BillFromParty,
			&pm.BillToCountry,
			&pm.BillFromCountry,
			&pm.Payer,
			&pm.Payee,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.ContractType,
			&pm.OrderValidityStartDate,
			&pm.OrderValidityEndDate,
			&pm.InvoicePeriodStartDate,
			&pm.InvoicePeriodEndDate,
			&pm.TotalNetAmount,
			&pm.TotalTaxAmount,
			&pm.TotalGrossAmount,
			&pm.HeaderDeliveryStatus,
			&pm.HeaderBillingStatus,
			&pm.HeaderDocReferenceStatus,
			&pm.TransactionCurrency,
			&pm.PricingDate,
			&pm.PriceDetnExchangeRate,
			&pm.RequestedDeliveryDate,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.Incoterms,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.ReferenceDocument,
			&pm.ReferenceDocumentItem,
			&pm.AccountAssignmentGroup,
			&pm.AccountingExchangeRate,
			&pm.InvoiceDocumentDate,
			&pm.IsExportImport,
			&pm.HeaderText,
			&pm.HeaderBlockStatus,
			&pm.HeaderDeliveryBlockStatus,
			&pm.HeaderBillingBlockStatus,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &OrdersHeader{
			OrderID:                          data.OrderID,
			OrderDate:                        data.OrderDate,
			OrderType:                        data.OrderType,
			SupplyChainRelationshipID:        data.SupplyChainRelationshipID,
			SupplyChainRelationshipBillingID: data.SupplyChainRelationshipBillingID,
			SupplyChainRelationshipPaymentID: data.SupplyChainRelationshipPaymentID,
			Buyer:                            data.Buyer,
			Seller:                           data.Seller,
			BillToParty:                      data.BillToParty,
			BillFromParty:                    data.BillFromParty,
			BillToCountry:                    data.BillToCountry,
			BillFromCountry:                  data.BillFromCountry,
			Payer:                            data.Payer,
			Payee:                            data.Payee,
			CreationDate:                     data.CreationDate,
			LastChangeDate:                   data.LastChangeDate,
			ContractType:                     data.ContractType,
			OrderValidityStartDate:           data.OrderValidityStartDate,
			OrderValidityEndDate:             data.OrderValidityEndDate,
			InvoicePeriodStartDate:           data.InvoicePeriodStartDate,
			InvoicePeriodEndDate:             data.InvoicePeriodEndDate,
			TotalNetAmount:                   data.TotalNetAmount,
			TotalTaxAmount:                   data.TotalTaxAmount,
			TotalGrossAmount:                 data.TotalGrossAmount,
			HeaderDeliveryStatus:             data.HeaderDeliveryStatus,
			HeaderBillingStatus:              data.HeaderBillingStatus,
			HeaderDocReferenceStatus:         data.HeaderDocReferenceStatus,
			TransactionCurrency:              data.TransactionCurrency,
			PricingDate:                      data.PricingDate,
			PriceDetnExchangeRate:            data.PriceDetnExchangeRate,
			RequestedDeliveryDate:            data.RequestedDeliveryDate,
			HeaderCompleteDeliveryIsDefined:  data.HeaderCompleteDeliveryIsDefined,
			Incoterms:                        data.Incoterms,
			PaymentTerms:                     data.PaymentTerms,
			PaymentMethod:                    data.PaymentMethod,
			ReferenceDocument:                data.ReferenceDocument,
			ReferenceDocumentItem:            data.ReferenceDocumentItem,
			AccountAssignmentGroup:           data.AccountAssignmentGroup,
			AccountingExchangeRate:           data.AccountingExchangeRate,
			InvoiceDocumentDate:              data.InvoiceDocumentDate,
			IsExportImport:                   data.IsExportImport,
			HeaderText:                       data.HeaderText,
			HeaderBlockStatus:                data.HeaderBlockStatus,
			HeaderDeliveryBlockStatus:        data.HeaderDeliveryBlockStatus,
			HeaderBillingBlockStatus:         data.HeaderBillingBlockStatus,
			IsCancelled:                      data.IsCancelled,
			IsMarkedForDeletion:              data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_orders_header_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToQuotationsPartner(rows *sql.Rows) ([]*QuotationsPartner, error) {
	defer rows.Close()
	res := make([]*QuotationsPartner, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.QuotationsPartner{}

		err := rows.Scan(
			&pm.Quotation,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.Organization,
			&pm.Country,
			&pm.Language,
			&pm.Currency,
			&pm.ExternalDocumentID,
			&pm.AddressID,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &QuotationsPartner{
			Quotation:               data.Quotation,
			PartnerFunction:         data.PartnerFunction,
			BusinessPartner:         data.BusinessPartner,
			BusinessPartnerFullName: data.BusinessPartnerFullName,
			BusinessPartnerName:     data.BusinessPartnerName,
			Organization:            data.Organization,
			Country:                 data.Country,
			Language:                data.Language,
			Currency:                data.Currency,
			ExternalDocumentID:      data.ExternalDocumentID,
			AddressID:               data.AddressID,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_quotations_partner_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToOrdersPartner(rows *sql.Rows) ([]*OrdersPartner, error) {
	defer rows.Close()
	res := make([]*OrdersPartner, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.OrdersPartner{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.Organization,
			&pm.Country,
			&pm.Language,
			&pm.Currency,
			&pm.ExternalDocumentID,
			&pm.AddressID,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &OrdersPartner{
			OrderID:                 data.OrderID,
			PartnerFunction:         data.PartnerFunction,
			BusinessPartner:         data.BusinessPartner,
			BusinessPartnerFullName: data.BusinessPartnerFullName,
			BusinessPartnerName:     data.BusinessPartnerName,
			Organization:            data.Organization,
			Country:                 data.Country,
			Language:                data.Language,
			Currency:                data.Currency,
			ExternalDocumentID:      data.ExternalDocumentID,
			AddressID:               data.AddressID,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_orders_partner_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToQuotationsAddress(rows *sql.Rows) ([]*QuotationsAddress, error) {
	defer rows.Close()
	res := make([]*QuotationsAddress, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.QuotationsAddress{}

		err := rows.Scan(
			&pm.Quotation,
			&pm.AddressID,
			&pm.PostalCode,
			&pm.LocalRegion,
			&pm.Country,
			&pm.District,
			&pm.StreetName,
			&pm.CityName,
			&pm.Building,
			&pm.Floor,
			&pm.Room,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &QuotationsAddress{
			Quotation:   data.Quotation,
			AddressID:   data.AddressID,
			PostalCode:  data.PostalCode,
			LocalRegion: data.LocalRegion,
			Country:     data.Country,
			District:    data.District,
			StreetName:  data.StreetName,
			CityName:    data.CityName,
			Building:    data.Building,
			Floor:       data.Floor,
			Room:        data.Room,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_quotations_address_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToOrdersAddress(rows *sql.Rows) ([]*OrdersAddress, error) {
	defer rows.Close()
	res := make([]*OrdersAddress, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.OrdersAddress{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.AddressID,
			&pm.PostalCode,
			&pm.LocalRegion,
			&pm.Country,
			&pm.District,
			&pm.StreetName,
			&pm.CityName,
			&pm.Building,
			&pm.Floor,
			&pm.Room,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &OrdersAddress{
			OrderID:     data.OrderID,
			AddressID:   data.AddressID,
			PostalCode:  data.PostalCode,
			LocalRegion: data.LocalRegion,
			Country:     data.Country,
			District:    data.District,
			StreetName:  data.StreetName,
			CityName:    data.CityName,
			Building:    data.Building,
			Floor:       data.Floor,
			Room:        data.Room,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_orders_address_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToPricingDate(inputPricingDate string) *PricingDate {
	pm := &requests.PricingDate{}

	pm.PricingDate = inputPricingDate

	data := pm
	res := PricingDate{
		PricingDate: data.PricingDate,
	}

	return &res
}

func (psdc *SDC) ConvertToPriceDetnExchangeRate(sdc *api_input_reader.SDC) *PriceDetnExchangeRate {
	pm := &requests.PriceDetnExchangeRate{
		PriceDetnExchangeRate: sdc.Header.PriceDetnExchangeRate,
	}

	data := pm
	res := PriceDetnExchangeRate{
		PriceDetnExchangeRate: data.PriceDetnExchangeRate,
	}

	return &res
}

func (psdc *SDC) ConvertToAccountingExchangeRate(sdc *api_input_reader.SDC) *AccountingExchangeRate {
	pm := &requests.AccountingExchangeRate{
		AccountingExchangeRate: sdc.Header.AccountingExchangeRate,
	}

	data := pm
	res := AccountingExchangeRate{
		AccountingExchangeRate: data.AccountingExchangeRate,
	}

	return &res
}

func (psdc *SDC) ConvertToTotalNetAmount(totalNetAmount float32) *TotalNetAmount {
	pm := &requests.TotalNetAmount{}

	pm.TotalNetAmount = totalNetAmount

	data := pm
	res := TotalNetAmount{
		TotalNetAmount: data.TotalNetAmount,
	}

	return &res
}

func (psdc *SDC) ConvertToTotalTaxAmount(totalTaxAmount float32) *TotalTaxAmount {
	pm := &requests.TotalTaxAmount{}

	pm.TotalTaxAmount = totalTaxAmount

	data := pm
	res := TotalTaxAmount{
		TotalTaxAmount: data.TotalTaxAmount,
	}

	return &res
}

func (psdc *SDC) ConvertToTotalGrossAmount(totalGrossAmount float32) *TotalGrossAmount {
	pm := &requests.TotalGrossAmount{}

	pm.TotalGrossAmount = totalGrossAmount

	data := pm
	res := TotalGrossAmount{
		TotalGrossAmount: data.TotalGrossAmount,
	}

	return &res
}

func (psdc *SDC) ConvertToCreationDateHeader(systemDate string) *CreationDateHeader {
	pm := &requests.CreationDateHeader{}

	pm.CreationDate = systemDate

	data := pm
	res := CreationDateHeader{
		CreationDate: data.CreationDate,
	}

	return &res
}

func (psdc *SDC) ConvertToLastChangeDateHeader(systemDate string) *LastChangeDateHeader {
	pm := &requests.LastChangeDateHeader{}

	pm.LastChangeDate = systemDate

	data := pm
	res := LastChangeDateHeader{
		LastChangeDate: data.LastChangeDate,
	}

	return &res
}

// Item
func (psdc *SDC) ConvertToProductTaxClassificationKey() *ProductTaxClassificationKey {
	pm := &requests.ProductTaxClassificationKey{
		ProductTaxCategory: "MWST",
	}

	data := pm
	res := ProductTaxClassificationKey{
		Product:            data.Product,
		Country:            data.Country,
		ProductTaxCategory: data.ProductTaxCategory,
	}

	return &res
}

func (psdc *SDC) ConvertToProductTaxClassificationBillToCountry(rows *sql.Rows) ([]*ProductTaxClassificationBillToCountry, error) {
	defer rows.Close()
	res := make([]*ProductTaxClassificationBillToCountry, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ProductTaxClassificationBillToCountry{}

		err := rows.Scan(
			&pm.Product,
			&pm.Country,
			&pm.ProductTaxCategory,
			&pm.ProductTaxClassifiication,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &ProductTaxClassificationBillToCountry{
			Product:                   data.Product,
			Country:                   data.Country,
			ProductTaxCategory:        data.ProductTaxCategory,
			ProductTaxClassifiication: data.ProductTaxClassifiication,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_product_master_tax_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToProductTaxClassificationBillFromCountry(rows *sql.Rows) ([]*ProductTaxClassificationBillFromCountry, error) {
	defer rows.Close()
	res := make([]*ProductTaxClassificationBillFromCountry, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ProductTaxClassificationBillFromCountry{}

		err := rows.Scan(
			&pm.Product,
			&pm.Country,
			&pm.ProductTaxCategory,
			&pm.ProductTaxClassifiication,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &ProductTaxClassificationBillFromCountry{
			Product:                   data.Product,
			Country:                   data.Country,
			ProductTaxCategory:        data.ProductTaxCategory,
			ProductTaxClassifiication: data.ProductTaxClassifiication,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_product_master_tax_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToDefinedTaxClassification(product string, transactionTaxClassification, productTaxClassificationBillToCountry, productTaxClassificationBillFromCountry *string, definedTaxClassification string) *DefinedTaxClassification {
	pm := &requests.DefinedTaxClassification{}

	pm.Product = product
	pm.TransactionTaxClassification = transactionTaxClassification
	pm.ProductTaxClassificationBillToCountry = productTaxClassificationBillToCountry
	pm.ProductTaxClassificationBillFromCountry = productTaxClassificationBillFromCountry
	pm.DefinedTaxClassification = definedTaxClassification

	data := pm
	res := DefinedTaxClassification{
		Product:                                 data.Product,
		TransactionTaxClassification:            data.TransactionTaxClassification,
		ProductTaxClassificationBillToCountry:   data.ProductTaxClassificationBillToCountry,
		ProductTaxClassificationBillFromCountry: data.ProductTaxClassificationBillFromCountry,
		DefinedTaxClassification:                data.DefinedTaxClassification,
	}

	return &res
}

func (psdc *SDC) ConvertToTaxCode(product, definedTaxClassification string, isExportImport *bool, taxCode *string) *TaxCode {
	pm := &requests.TaxCode{}

	pm.Product = product
	pm.DefinedTaxClassification = definedTaxClassification
	pm.IsExportImport = isExportImport
	pm.TaxCode = taxCode

	data := pm
	res := TaxCode{
		Product:                  data.Product,
		DefinedTaxClassification: data.DefinedTaxClassification,
		IsExportImport:           data.IsExportImport,
		TaxCode:                  data.TaxCode,
	}

	return &res
}

func (psdc *SDC) ConvertToTaxRateKey() *TaxRateKey {
	pm := &requests.TaxRateKey{
		Country: "JP",
	}

	data := pm
	res := TaxRateKey{
		Country:           data.Country,
		TaxCode:           data.TaxCode,
		ValidityEndDate:   data.ValidityEndDate,
		ValidityStartDate: data.ValidityStartDate,
	}

	return &res
}

func (psdc *SDC) ConvertToTaxRate(rows *sql.Rows) ([]*TaxRate, error) {
	defer rows.Close()
	res := make([]*TaxRate, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.TaxRate{}

		err := rows.Scan(
			&pm.Country,
			&pm.TaxCode,
			&pm.ValidityEndDate,
			&pm.ValidityStartDate,
			&pm.TaxRate,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &TaxRate{
			Country:           data.Country,
			TaxCode:           data.TaxCode,
			ValidityEndDate:   data.ValidityEndDate,
			ValidityStartDate: data.ValidityStartDate,
			TaxRate:           data.TaxRate,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_tax_code_tax_rate_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToNetAmount(conditionAmount []*ConditionAmount) []*NetAmount {
	res := make([]*NetAmount, 0)

	for _, v := range conditionAmount {
		pm := &requests.NetAmount{}

		pm.OrderItem = v.OrderItem
		pm.Product = v.Product
		pm.NetAmount = v.ConditionAmount

		data := pm
		res = append(res, &NetAmount{
			OrderItem: data.OrderItem,
			Product:   data.Product,
			NetAmount: data.NetAmount,
		})
	}

	return res
}

func (psdc *SDC) ConvertToTaxAmount(orderItem int, product string, taxCode string, taxRate, netAmount, taxAmount *float32) *TaxAmount {
	pm := &requests.TaxAmount{}

	pm.OrderItem = orderItem
	pm.Product = product
	pm.TaxCode = taxCode
	pm.TaxRate = taxRate
	pm.NetAmount = netAmount
	pm.TaxAmount = taxAmount

	data := pm
	res := TaxAmount{
		OrderItem: data.OrderItem,
		Product:   data.Product,
		TaxCode:   data.TaxCode,
		TaxRate:   data.TaxRate,
		NetAmount: data.NetAmount,
		TaxAmount: data.TaxAmount,
	}

	return &res
}

func (psdc *SDC) ConvertToGrossAmount(orderItem int, product string, netAmount, taxAmount, grossAmount *float32) *GrossAmount {
	pm := &requests.GrossAmount{}

	pm.OrderItem = orderItem
	pm.Product = product
	pm.NetAmount = netAmount
	pm.TaxAmount = taxAmount
	pm.GrossAmount = grossAmount

	data := pm
	res := GrossAmount{
		OrderItem:   data.OrderItem,
		Product:     data.Product,
		NetAmount:   data.NetAmount,
		TaxAmount:   data.TaxAmount,
		GrossAmount: data.GrossAmount,
	}

	return &res
}

func (psdc *SDC) ConvertToQuotationsItem(rows *sql.Rows) ([]*QuotationsItem, error) {
	defer rows.Close()
	res := make([]*QuotationsItem, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.QuotationsItem{}

		err := rows.Scan(
			&pm.Quotation,
			&pm.QuotationItem,
			&pm.QuotationItemCategory,
			&pm.SupplyChainRelationshipID,
			&pm.QuotationItemText,
			&pm.QuotationItemTextByBuyer,
			&pm.QuotationItemTextBySeller,
			&pm.Product,
			&pm.ProductStandardID,
			&pm.ProductGroup,
			&pm.BaseUnit,
			&pm.PricingDate,
			&pm.PriceDetnExchangeRate,
			&pm.RequestedDeliveryDate,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.DeliveryUnit,
			&pm.ServicesRenderingDate,
			&pm.QuotationQuantityInBaseUnit,
			&pm.QuotationQuantityInDeliveryUnit,
			&pm.ItemWeightUnit,
			&pm.ProductGrossWeight,
			&pm.ItemGrossWeight,
			&pm.ProductNetWeight,
			&pm.ItemNetWeight,
			&pm.InternalCapacityQuantity,
			&pm.InternalCapacityQuantityUnit,
			&pm.NetAmount,
			&pm.TaxAmount,
			&pm.GrossAmount,
			&pm.Incoterms,
			&pm.TransactionTaxClassification,
			&pm.ProductTaxClassificationBillToCountry,
			&pm.ProductTaxClassificationBillFromCountry,
			&pm.DefinedTaxClassification,
			&pm.AccountAssignmentGroup,
			&pm.ProductAccountAssignmentGroup,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.Project,
			&pm.AccountingExchangeRate,
			&pm.ReferenceDocument,
			&pm.ReferenceDocumentItem,
			&pm.TaxCode,
			&pm.TaxRate,
			&pm.CountryOfOrigin,
			&pm.CountryOfOriginLanguage,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &QuotationsItem{
			Quotation:                               data.Quotation,
			QuotationItem:                           data.QuotationItem,
			QuotationItemCategory:                   data.QuotationItemCategory,
			SupplyChainRelationshipID:               data.SupplyChainRelationshipID,
			QuotationItemText:                       data.QuotationItemText,
			QuotationItemTextByBuyer:                data.QuotationItemTextByBuyer,
			QuotationItemTextBySeller:               data.QuotationItemTextBySeller,
			Product:                                 data.Product,
			ProductStandardID:                       data.ProductStandardID,
			ProductGroup:                            data.ProductGroup,
			BaseUnit:                                data.BaseUnit,
			PricingDate:                             data.PricingDate,
			PriceDetnExchangeRate:                   data.PriceDetnExchangeRate,
			RequestedDeliveryDate:                   data.RequestedDeliveryDate,
			CreationDate:                            data.CreationDate,
			LastChangeDate:                          data.LastChangeDate,
			DeliveryUnit:                            data.DeliveryUnit,
			ServicesRenderingDate:                   data.ServicesRenderingDate,
			QuotationQuantityInBaseUnit:             data.QuotationQuantityInBaseUnit,
			QuotationQuantityInDeliveryUnit:         data.QuotationQuantityInDeliveryUnit,
			ItemWeightUnit:                          data.ItemWeightUnit,
			ProductGrossWeight:                      data.ProductGrossWeight,
			ItemGrossWeight:                         data.ItemGrossWeight,
			ProductNetWeight:                        data.ProductNetWeight,
			ItemNetWeight:                           data.ItemNetWeight,
			InternalCapacityQuantity:                data.InternalCapacityQuantity,
			InternalCapacityQuantityUnit:            data.InternalCapacityQuantityUnit,
			NetAmount:                               data.NetAmount,
			TaxAmount:                               data.TaxAmount,
			GrossAmount:                             data.GrossAmount,
			Incoterms:                               data.Incoterms,
			TransactionTaxClassification:            data.TransactionTaxClassification,
			ProductTaxClassificationBillToCountry:   data.ProductTaxClassificationBillToCountry,
			ProductTaxClassificationBillFromCountry: data.ProductTaxClassificationBillFromCountry,
			DefinedTaxClassification:                data.DefinedTaxClassification,
			AccountAssignmentGroup:                  data.AccountAssignmentGroup,
			ProductAccountAssignmentGroup:           data.ProductAccountAssignmentGroup,
			PaymentTerms:                            data.PaymentTerms,
			PaymentMethod:                           data.PaymentMethod,
			Project:                                 data.Project,
			AccountingExchangeRate:                  data.AccountingExchangeRate,
			ReferenceDocument:                       data.ReferenceDocument,
			ReferenceDocumentItem:                   data.ReferenceDocumentItem,
			TaxCode:                                 data.TaxCode,
			TaxRate:                                 data.TaxRate,
			CountryOfOrigin:                         data.CountryOfOrigin,
			CountryOfOriginLanguage:                 data.CountryOfOriginLanguage,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_quotations_item_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToOrdersItem(rows *sql.Rows) ([]*OrdersItem, error) {
	defer rows.Close()
	res := make([]*OrdersItem, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.OrdersItem{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.OrderItemCategory,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipStockConfPlantID,
			&pm.SupplyChainRelationshipProductionPlantID,
			&pm.OrderItemText,
			&pm.OrderItemTextByBuyer,
			&pm.OrderItemTextBySeller,
			&pm.Product,
			&pm.ProductStandardID,
			&pm.ProductGroup,
			&pm.BaseUnit,
			&pm.PricingDate,
			&pm.PriceDetnExchangeRate,
			&pm.RequestedDeliveryDate,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.DeliverToPlant,
			&pm.DeliverToPlantTimeZone,
			&pm.DeliverToPlantStorageLocation,
			&pm.ProductIsBatchManagedInDeliverToPlant,
			&pm.BatchMgmtPolicyInDeliverToPlant,
			&pm.DeliverToPlantBatch,
			&pm.DeliverToPlantBatchValidityStartDate,
			&pm.DeliverToPlantBatchValidityStartTime,
			&pm.DeliverToPlantBatchValidityEndDate,
			&pm.DeliverToPlantBatchValidityEndTime,
			&pm.DeliverFromPlant,
			&pm.DeliverFromPlantTimeZone,
			&pm.DeliverFromPlantStorageLocation,
			&pm.ProductIsBatchManagedInDeliverFromPlant,
			&pm.BatchMgmtPolicyInDeliverFromPlant,
			&pm.DeliverFromPlantBatch,
			&pm.DeliverFromPlantBatchValidityStartDate,
			&pm.DeliverFromPlantBatchValidityStartTime,
			&pm.DeliverFromPlantBatchValidityEndDate,
			&pm.DeliverFromPlantBatchValidityEndTime,
			&pm.DeliveryUnit,
			&pm.StockConfirmationBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.StockConfirmationPlantTimeZone,
			&pm.ProductIsBatchManagedInStockConfirmationPlant,
			&pm.BatchMgmtPolicyInStockConfirmationPlant,
			&pm.StockConfirmationPlantBatch,
			&pm.StockConfirmationPlantBatchValidityStartDate,
			&pm.StockConfirmationPlantBatchValidityStartTime,
			&pm.StockConfirmationPlantBatchValidityEndDate,
			&pm.StockConfirmationPlantBatchValidityEndTime,
			&pm.ServicesRenderingDate,
			&pm.OrderQuantityInBaseUnit,
			&pm.OrderQuantityInDeliveryUnit,
			&pm.StockConfirmationPolicy,
			&pm.StockConfirmationStatus,
			&pm.ConfirmedOrderQuantityInBaseUnit,
			&pm.ItemWeightUnit,
			&pm.ProductGrossWeight,
			&pm.ItemGrossWeight,
			&pm.ProductNetWeight,
			&pm.ItemNetWeight,
			&pm.InternalCapacityQuantity,
			&pm.InternalCapacityQuantityUnit,
			&pm.NetAmount,
			&pm.TaxAmount,
			&pm.GrossAmount,
			&pm.InvoiceDocumentDate,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ProductionPlantTimeZone,
			&pm.ProductionPlantStorageLocation,
			&pm.ProductIsBatchManagedInProductionPlant,
			&pm.BatchMgmtPolicyInProductionPlant,
			&pm.ProductionPlantBatch,
			&pm.ProductionPlantBatchValidityStartDate,
			&pm.ProductionPlantBatchValidityStartTime,
			&pm.ProductionPlantBatchValidityEndDate,
			&pm.ProductionPlantBatchValidityEndTime,
			&pm.Incoterms,
			&pm.TransactionTaxClassification,
			&pm.ProductTaxClassificationBillToCountry,
			&pm.ProductTaxClassificationBillFromCountry,
			&pm.DefinedTaxClassification,
			&pm.AccountAssignmentGroup,
			&pm.ProductAccountAssignmentGroup,
			&pm.PaymentTerms,
			&pm.DueCalculationBaseDate,
			&pm.PaymentDueDate,
			&pm.NetPaymentDays,
			&pm.PaymentMethod,
			&pm.Project,
			&pm.AccountingExchangeRate,
			&pm.ReferenceDocument,
			&pm.ReferenceDocumentItem,
			&pm.ItemCompleteDeliveryIsDefined,
			&pm.ItemDeliveryStatus,
			&pm.IssuingStatus,
			&pm.ReceivingStatus,
			&pm.ItemBillingStatus,
			&pm.TaxCode,
			&pm.TaxRate,
			&pm.CountryOfOrigin,
			&pm.CountryOfOriginLanguage,
			&pm.ItemBlockStatus,
			&pm.ItemDeliveryBlockStatus,
			&pm.ItemBillingBlockStatus,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &OrdersItem{
			OrderID:                                       data.OrderID,
			OrderItem:                                     data.OrderItem,
			OrderItemCategory:                             data.OrderItemCategory,
			SupplyChainRelationshipID:                     data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:             data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:        data.SupplyChainRelationshipDeliveryPlantID,
			SupplyChainRelationshipStockConfPlantID:       data.SupplyChainRelationshipStockConfPlantID,
			SupplyChainRelationshipProductionPlantID:      data.SupplyChainRelationshipProductionPlantID,
			OrderItemText:                                 data.OrderItemText,
			OrderItemTextByBuyer:                          data.OrderItemTextByBuyer,
			OrderItemTextBySeller:                         data.OrderItemTextBySeller,
			Product:                                       data.Product,
			ProductStandardID:                             data.ProductStandardID,
			ProductGroup:                                  data.ProductGroup,
			BaseUnit:                                      data.BaseUnit,
			PricingDate:                                   data.PricingDate,
			PriceDetnExchangeRate:                         data.PriceDetnExchangeRate,
			RequestedDeliveryDate:                         data.RequestedDeliveryDate,
			DeliverToParty:                                data.DeliverToParty,
			DeliverFromParty:                              data.DeliverFromParty,
			CreationDate:                                  data.CreationDate,
			LastChangeDate:                                data.LastChangeDate,
			DeliverToPlant:                                data.DeliverToPlant,
			DeliverToPlantTimeZone:                        data.DeliverToPlantTimeZone,
			DeliverToPlantStorageLocation:                 data.DeliverToPlantStorageLocation,
			ProductIsBatchManagedInDeliverToPlant:         data.ProductIsBatchManagedInDeliverToPlant,
			BatchMgmtPolicyInDeliverToPlant:               data.BatchMgmtPolicyInDeliverToPlant,
			DeliverToPlantBatch:                           data.DeliverToPlantBatch,
			DeliverToPlantBatchValidityStartDate:          data.DeliverToPlantBatchValidityStartDate,
			DeliverToPlantBatchValidityStartTime:          data.DeliverToPlantBatchValidityStartTime,
			DeliverToPlantBatchValidityEndDate:            data.DeliverToPlantBatchValidityEndDate,
			DeliverToPlantBatchValidityEndTime:            data.DeliverToPlantBatchValidityEndTime,
			DeliverFromPlant:                              data.DeliverFromPlant,
			DeliverFromPlantTimeZone:                      data.DeliverFromPlantTimeZone,
			DeliverFromPlantStorageLocation:               data.DeliverFromPlantStorageLocation,
			ProductIsBatchManagedInDeliverFromPlant:       data.ProductIsBatchManagedInDeliverFromPlant,
			BatchMgmtPolicyInDeliverFromPlant:             data.BatchMgmtPolicyInDeliverFromPlant,
			DeliverFromPlantBatch:                         data.DeliverFromPlantBatch,
			DeliverFromPlantBatchValidityStartDate:        data.DeliverFromPlantBatchValidityStartDate,
			DeliverFromPlantBatchValidityStartTime:        data.DeliverFromPlantBatchValidityStartTime,
			DeliverFromPlantBatchValidityEndDate:          data.DeliverFromPlantBatchValidityEndDate,
			DeliverFromPlantBatchValidityEndTime:          data.DeliverFromPlantBatchValidityEndTime,
			DeliveryUnit:                                  data.DeliveryUnit,
			StockConfirmationBusinessPartner:              data.StockConfirmationBusinessPartner,
			StockConfirmationPlant:                        data.StockConfirmationPlant,
			StockConfirmationPlantTimeZone:                data.StockConfirmationPlantTimeZone,
			ProductIsBatchManagedInStockConfirmationPlant: data.ProductIsBatchManagedInStockConfirmationPlant,
			BatchMgmtPolicyInStockConfirmationPlant:       data.BatchMgmtPolicyInStockConfirmationPlant,
			StockConfirmationPlantBatch:                   data.StockConfirmationPlantBatch,
			StockConfirmationPlantBatchValidityStartDate:  data.StockConfirmationPlantBatchValidityStartDate,
			StockConfirmationPlantBatchValidityStartTime:  data.StockConfirmationPlantBatchValidityStartTime,
			StockConfirmationPlantBatchValidityEndDate:    data.StockConfirmationPlantBatchValidityEndDate,
			StockConfirmationPlantBatchValidityEndTime:    data.StockConfirmationPlantBatchValidityEndTime,
			ServicesRenderingDate:                         data.ServicesRenderingDate,
			OrderQuantityInBaseUnit:                       data.OrderQuantityInBaseUnit,
			OrderQuantityInDeliveryUnit:                   data.OrderQuantityInDeliveryUnit,
			StockConfirmationPolicy:                       data.StockConfirmationPolicy,
			StockConfirmationStatus:                       data.StockConfirmationStatus,
			ConfirmedOrderQuantityInBaseUnit:              data.ConfirmedOrderQuantityInBaseUnit,
			ItemWeightUnit:                                data.ItemWeightUnit,
			ProductGrossWeight:                            data.ProductGrossWeight,
			ItemGrossWeight:                               data.ItemGrossWeight,
			ProductNetWeight:                              data.ProductNetWeight,
			ItemNetWeight:                                 data.ItemNetWeight,
			InternalCapacityQuantity:                      data.InternalCapacityQuantity,
			InternalCapacityQuantityUnit:                  data.InternalCapacityQuantityUnit,
			NetAmount:                                     data.NetAmount,
			TaxAmount:                                     data.TaxAmount,
			GrossAmount:                                   data.GrossAmount,
			InvoiceDocumentDate:                           data.InvoiceDocumentDate,
			ProductionPlantBusinessPartner:                data.ProductionPlantBusinessPartner,
			ProductionPlant:                               data.ProductionPlant,
			ProductionPlantTimeZone:                       data.ProductionPlantTimeZone,
			ProductionPlantStorageLocation:                data.ProductionPlantStorageLocation,
			ProductIsBatchManagedInProductionPlant:        data.ProductIsBatchManagedInProductionPlant,
			BatchMgmtPolicyInProductionPlant:              data.BatchMgmtPolicyInProductionPlant,
			ProductionPlantBatch:                          data.ProductionPlantBatch,
			ProductionPlantBatchValidityStartDate:         data.ProductionPlantBatchValidityStartDate,
			ProductionPlantBatchValidityStartTime:         data.ProductionPlantBatchValidityStartTime,
			ProductionPlantBatchValidityEndDate:           data.ProductionPlantBatchValidityEndDate,
			ProductionPlantBatchValidityEndTime:           data.ProductionPlantBatchValidityEndTime,
			Incoterms:                                     data.Incoterms,
			TransactionTaxClassification:                  data.TransactionTaxClassification,
			ProductTaxClassificationBillToCountry:         data.ProductTaxClassificationBillToCountry,
			ProductTaxClassificationBillFromCountry:       data.ProductTaxClassificationBillFromCountry,
			DefinedTaxClassification:                      data.DefinedTaxClassification,
			AccountAssignmentGroup:                        data.AccountAssignmentGroup,
			ProductAccountAssignmentGroup:                 data.ProductAccountAssignmentGroup,
			PaymentTerms:                                  data.PaymentTerms,
			DueCalculationBaseDate:                        data.DueCalculationBaseDate,
			PaymentDueDate:                                data.PaymentDueDate,
			NetPaymentDays:                                data.NetPaymentDays,
			PaymentMethod:                                 data.PaymentMethod,
			Project:                                       data.Project,
			AccountingExchangeRate:                        data.AccountingExchangeRate,
			ReferenceDocument:                             data.ReferenceDocument,
			ReferenceDocumentItem:                         data.ReferenceDocumentItem,
			ItemCompleteDeliveryIsDefined:                 data.ItemCompleteDeliveryIsDefined,
			ItemDeliveryStatus:                            data.ItemDeliveryStatus,
			IssuingStatus:                                 data.IssuingStatus,
			ReceivingStatus:                               data.ReceivingStatus,
			ItemBillingStatus:                             data.ItemBillingStatus,
			TaxCode:                                       data.TaxCode,
			TaxRate:                                       data.TaxRate,
			CountryOfOrigin:                               data.CountryOfOrigin,
			CountryOfOriginLanguage:                       data.CountryOfOriginLanguage,
			ItemBlockStatus:                               data.ItemBlockStatus,
			ItemDeliveryBlockStatus:                       data.ItemDeliveryBlockStatus,
			ItemBillingBlockStatus:                        data.ItemBillingBlockStatus,
			IsCancelled:                                   data.IsCancelled,
			IsMarkedForDeletion:                           data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_orders_item_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

// Item Pricing Element
func (psdc *SDC) ConvertToPriceMasterKey() *PriceMasterKey {
	pm := &requests.PriceMasterKey{}

	data := pm
	res := PriceMasterKey{
		Product:                    data.Product,
		SupplyChainRelationshipID:  data.SupplyChainRelationshipID,
		Buyer:                      data.Buyer,
		Seller:                     data.Seller,
		ConditionValidityEndDate:   data.ConditionValidityEndDate,
		ConditionValidityStartDate: data.ConditionValidityStartDate,
	}

	return &res
}

func (psdc *SDC) ConvertToPriceMaster(rows *sql.Rows) ([]*PriceMaster, error) {
	defer rows.Close()
	res := make([]*PriceMaster, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.PriceMaster{}

		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.Buyer,
			&pm.Seller,
			&pm.ConditionRecord,
			&pm.ConditionSequentialNumber,
			&pm.ConditionValidityStartDate,
			&pm.ConditionValidityEndDate,
			&pm.Product,
			&pm.ConditionType,
			&pm.ConditionRateValue,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &PriceMaster{
			SupplyChainRelationshipID:  data.SupplyChainRelationshipID,
			Buyer:                      data.Buyer,
			Seller:                     data.Seller,
			ConditionRecord:            data.ConditionRecord,
			ConditionSequentialNumber:  data.ConditionSequentialNumber,
			ConditionValidityStartDate: data.ConditionValidityStartDate,
			ConditionValidityEndDate:   data.ConditionValidityEndDate,
			Product:                    data.Product,
			ConditionType:              data.ConditionType,
			ConditionRateValue:         data.ConditionRateValue,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_price_master_price_master_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToConditionAmount(orderItem int, product string, conditionQuantity *float32, conditionAmount *float32) *ConditionAmount {
	pm := &requests.ConditionAmount{
		ConditionIsManuallyChanged: getBoolPtr(false),
	}

	pm.OrderItem = orderItem
	pm.Product = product
	pm.ConditionQuantity = conditionQuantity
	pm.ConditionAmount = conditionAmount

	data := pm
	res := ConditionAmount{
		OrderItem:                  data.OrderItem,
		Product:                    data.Product,
		ConditionQuantity:          data.ConditionQuantity,
		ConditionAmount:            data.ConditionAmount,
		ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
	}

	return &res
}

func (psdc *SDC) ConvertToQuotationsItemPricingElement(rows *sql.Rows) ([]*QuotationsItemPricingElement, error) {
	defer rows.Close()
	res := make([]*QuotationsItemPricingElement, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.QuotationsItemPricingElement{}

		err := rows.Scan(
			&pm.Quotation,
			&pm.QuotationItem,
			&pm.SupplyChainRelationshipID,
			&pm.Buyer,
			&pm.Seller,
			&pm.PricingProcedureCounter,
			&pm.ConditionRecord,
			&pm.ConditionSequentialNumber,
			&pm.ConditionType,
			&pm.PricingDate,
			&pm.ConditionRateValue,
			&pm.ConditionCurrency,
			&pm.ConditionQuantity,
			&pm.ConditionQuantityUnit,
			&pm.TaxCode,
			&pm.ConditionAmount,
			&pm.TransactionCurrency,
			&pm.ConditionIsManuallyChanged,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &QuotationsItemPricingElement{
			Quotation:                  data.Quotation,
			QuotationItem:              data.QuotationItem,
			SupplyChainRelationshipID:  data.SupplyChainRelationshipID,
			Buyer:                      data.Buyer,
			Seller:                     data.Seller,
			PricingProcedureCounter:    data.PricingProcedureCounter,
			ConditionRecord:            data.ConditionRecord,
			ConditionSequentialNumber:  data.ConditionSequentialNumber,
			ConditionType:              data.ConditionType,
			PricingDate:                data.PricingDate,
			ConditionRateValue:         data.ConditionRateValue,
			ConditionCurrency:          data.ConditionCurrency,
			ConditionQuantity:          data.ConditionQuantity,
			ConditionQuantityUnit:      data.ConditionQuantityUnit,
			TaxCode:                    data.TaxCode,
			ConditionAmount:            data.ConditionAmount,
			TransactionCurrency:        data.TransactionCurrency,
			ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_quotations_item_pricing_element'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToOrdersItemPricingElement(rows *sql.Rows) ([]*OrdersItemPricingElement, error) {
	defer rows.Close()
	res := make([]*OrdersItemPricingElement, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.OrdersItemPricingElement{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.SupplyChainRelationshipID,
			&pm.Buyer,
			&pm.Seller,
			&pm.PricingProcedureCounter,
			&pm.ConditionRecord,
			&pm.ConditionSequentialNumber,
			&pm.ConditionType,
			&pm.PricingDate,
			&pm.ConditionRateValue,
			&pm.ConditionCurrency,
			&pm.ConditionQuantity,
			&pm.ConditionQuantityUnit,
			&pm.TaxCode,
			&pm.ConditionAmount,
			&pm.TransactionCurrency,
			&pm.ConditionIsManuallyChanged,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &OrdersItemPricingElement{
			OrderID:                    data.OrderID,
			OrderItem:                  data.OrderItem,
			SupplyChainRelationshipID:  data.SupplyChainRelationshipID,
			Buyer:                      data.Buyer,
			Seller:                     data.Seller,
			PricingProcedureCounter:    data.PricingProcedureCounter,
			ConditionRecord:            data.ConditionRecord,
			ConditionSequentialNumber:  data.ConditionSequentialNumber,
			ConditionType:              data.ConditionType,
			PricingDate:                data.PricingDate,
			ConditionRateValue:         data.ConditionRateValue,
			ConditionCurrency:          data.ConditionCurrency,
			ConditionQuantity:          data.ConditionQuantity,
			ConditionQuantityUnit:      data.ConditionQuantityUnit,
			TaxCode:                    data.TaxCode,
			ConditionAmount:            data.ConditionAmount,
			TransactionCurrency:        data.TransactionCurrency,
			ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_orders_item_pricing_element'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToOrdersItemScheduleLine(rows *sql.Rows) ([]*OrdersItemScheduleLine, error) {
	defer rows.Close()
	res := make([]*OrdersItemScheduleLine, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.OrdersItemScheduleLine{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.ScheduleLine,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipStockConfPlantID,
			&pm.Product,
			&pm.StockConfirmationBussinessPartner,
			&pm.StockConfirmationPlant,
			&pm.StockConfirmationPlantTimeZone,
			&pm.StockConfirmationPlantBatch,
			&pm.StockConfirmationPlantBatchValidityStartDate,
			&pm.StockConfirmationPlantBatchValidityEndDate,
			&pm.RequestedDeliveryDate,
			&pm.ConfirmedDeliveryDate,
			&pm.OrderQuantityInBaseUnit,
			&pm.ConfirmedOrderQuantityByPDTAvailCheck,
			&pm.ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit,
			&pm.DeliveredQuantityInBaseUnit,
			&pm.UndeliveredQuantityInBaseUnit,
			&pm.OpenConfirmedQuantityInBaseUnit,
			&pm.StockIsFullyConfirmed,
			&pm.PlusMinusFlag,
			&pm.ItemScheduleLineDeliveryBlockStatus,
			&pm.IsCancelled,
			&pm.IsDeleted,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &OrdersItemScheduleLine{
			OrderID:                                         data.OrderID,
			OrderItem:                                       data.OrderItem,
			ScheduleLine:                                    data.ScheduleLine,
			SupplyChainRelationshipID:                       data.SupplyChainRelationshipID,
			SupplyChainRelationshipStockConfPlantID:         data.SupplyChainRelationshipStockConfPlantID,
			Product:                                         data.Product,
			StockConfirmationBussinessPartner:               data.StockConfirmationBussinessPartner,
			StockConfirmationPlant:                          data.StockConfirmationPlant,
			StockConfirmationPlantTimeZone:                  data.StockConfirmationPlantTimeZone,
			StockConfirmationPlantBatch:                     data.StockConfirmationPlantBatch,
			StockConfirmationPlantBatchValidityStartDate:    data.StockConfirmationPlantBatchValidityStartDate,
			StockConfirmationPlantBatchValidityEndDate:      data.StockConfirmationPlantBatchValidityEndDate,
			RequestedDeliveryDate:                           data.RequestedDeliveryDate,
			ConfirmedDeliveryDate:                           data.ConfirmedDeliveryDate,
			OrderQuantityInBaseUnit:                         data.OrderQuantityInBaseUnit,
			ConfirmedOrderQuantityByPDTAvailCheck:           data.ConfirmedOrderQuantityByPDTAvailCheck,
			ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit: data.ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit,
			DeliveredQuantityInBaseUnit:                     data.DeliveredQuantityInBaseUnit,
			UndeliveredQuantityInBaseUnit:                   data.UndeliveredQuantityInBaseUnit,
			OpenConfirmedQuantityInBaseUnit:                 data.OpenConfirmedQuantityInBaseUnit,
			StockIsFullyConfirmed:                           data.StockIsFullyConfirmed,
			PlusMinusFlag:                                   data.PlusMinusFlag,
			ItemScheduleLineDeliveryBlockStatus:             data.ItemScheduleLineDeliveryBlockStatus,
			IsCancelled:                                     data.IsCancelled,
			IsDeleted:                                       data.IsDeleted,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_orders_item_schedule_line_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func getBoolPtr(b bool) *bool {
	return &b
}
