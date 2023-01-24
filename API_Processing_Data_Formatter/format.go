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
		ReferenceDocument:     sdc.OrdersInputParameters.ReferenceDocument,
		ReferenceDocumentItem: sdc.OrdersInputParameters.ReferenceDocumentItem,
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

func (psdc *SDC) ConvertToCalculateOrderIDKey() *CalculateOrderIDKey {
	pm := &requests.CalculateOrderIDKey{
		FieldNameWithNumberRange: "OrderID",
	}

	data := pm
	res := CalculateOrderIDKey{
		ServiceLabel:             data.ServiceLabel,
		FieldNameWithNumberRange: data.FieldNameWithNumberRange,
	}

	return &res
}

func (psdc *SDC) ConvertToCalculateOrderIDQueryGets(rows *sql.Rows) (*CalculateOrderIDQueryGets, error) {
	defer rows.Close()
	pm := &requests.CalculateOrderIDQueryGets{}

	i := 0
	for rows.Next() {
		i++
		err := rows.Scan(
			&pm.ServiceLabel,
			&pm.FieldNameWithNumberRange,
			&pm.LatestNumber,
		)
		if err != nil {
			return nil, err
		}
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_number_range_latest_number_data'テーブルに対象のレコードが存在しません。")
	}

	data := pm
	res := CalculateOrderIDQueryGets{
		ServiceLabel:             data.ServiceLabel,
		FieldNameWithNumberRange: data.FieldNameWithNumberRange,
		LatestNumber:             data.LatestNumber,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToCalculateOrderID(orderIDLatestNumber *int, orderID int) *CalculateOrderID {
	pm := &requests.CalculateOrderID{}

	pm.OrderIDLatestNumber = orderIDLatestNumber
	pm.OrderID = orderID

	data := pm
	res := CalculateOrderID{
		OrderIDLatestNumber: data.OrderIDLatestNumber,
		OrderID:             data.OrderID,
	}

	return &res
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

func (psdc *SDC) ConvertToInvoiceDocumentDate(sdc *api_input_reader.SDC) *InvoiceDocumentDate {
	pm := &requests.InvoiceDocumentDate{}

	pm.InvoiceDocumentDate = *sdc.Header.InvoiceDocumentDate
	data := pm

	res := InvoiceDocumentDate{
		RequestedDeliveryDate: data.RequestedDeliveryDate,
		InvoiceDocumentDate:   data.InvoiceDocumentDate,
	}

	return &res
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

func (psdc *SDC) ConvertToCaluculateInvoiceDocumentDate(sdc *api_input_reader.SDC, invoiceDocumentDate string) *InvoiceDocumentDate {
	pm := &requests.InvoiceDocumentDate{
		RequestedDeliveryDate: *sdc.Header.RequestedDeliveryDate,
	}

	pm.InvoiceDocumentDate = invoiceDocumentDate

	data := pm
	res := InvoiceDocumentDate{
		RequestedDeliveryDate: data.RequestedDeliveryDate,
		InvoiceDocumentDate:   data.InvoiceDocumentDate,
	}

	return &res
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

		pm.Product = v.Product
		pm.NetAmount = v.ConditionAmount

		data := pm
		res = append(res, &NetAmount{
			Product:   data.Product,
			NetAmount: data.NetAmount,
		})
	}

	return res
}

func (psdc *SDC) ConvertToTaxAmount(product string, taxCode *string, taxRate, netAmount, taxAmount *float32) *TaxAmount {
	pm := &requests.TaxAmount{}

	pm.Product = product
	pm.TaxCode = taxCode
	pm.TaxRate = taxRate
	pm.NetAmount = netAmount
	pm.TaxAmount = taxAmount

	data := pm
	res := TaxAmount{
		Product:   data.Product,
		TaxCode:   data.TaxCode,
		TaxRate:   data.TaxRate,
		NetAmount: data.NetAmount,
		TaxAmount: data.TaxAmount,
	}

	return &res
}

func (psdc *SDC) ConvertToGrossAmount(product string, netAmount, taxAmount, grossAmount *float32) *GrossAmount {
	pm := &requests.GrossAmount{}

	pm.Product = product
	pm.NetAmount = netAmount
	pm.TaxAmount = taxAmount
	pm.GrossAmount = grossAmount

	data := pm
	res := GrossAmount{
		Product:     data.Product,
		NetAmount:   data.NetAmount,
		TaxAmount:   data.TaxAmount,
		GrossAmount: data.GrossAmount,
	}

	return &res
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

func (psdc *SDC) ConvertToConditionAmount(product string, conditionQuantity *float32, conditionAmount *float32) *ConditionAmount {
	pm := &requests.ConditionAmount{
		ConditionIsManuallyChanged: getBoolPtr(false),
	}

	pm.Product = product
	pm.ConditionQuantity = conditionQuantity
	pm.ConditionAmount = conditionAmount

	data := pm
	res := ConditionAmount{
		Product:                    data.Product,
		ConditionQuantity:          data.ConditionQuantity,
		ConditionAmount:            data.ConditionAmount,
		ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
	}

	return &res
}

func getBoolPtr(b bool) *bool {
	return &b
}
