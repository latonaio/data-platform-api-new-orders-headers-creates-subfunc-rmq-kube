package subfunction

import (
	api_input_reader "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"sort"
	"strings"
	"time"

	"golang.org/x/xerrors"
)

func (f *SubFunction) SupplyChainRelationshipGeneral(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.SupplyChainRelationshipGeneral, error) {
	rows, err := f.db.Query(
		`SELECT SupplyChainRelationshipID, Buyer, Seller
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_general_data
		WHERE (Buyer, Seller) = (?, ?);`, sdc.Header.Buyer, sdc.Header.Seller,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToSupplyChainRelationshipGeneral(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) SupplyChainRelationshipDeliveryRelation(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.SupplyChainRelationshipDeliveryRelation, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToSupplyChainRelationshipDeliveryRelationKey()

	for _, v := range psdc.SupplyChainRelationshipGeneral {
		dataKey.SupplyChainRelationshipID = append(dataKey.SupplyChainRelationshipID, v.SupplyChainRelationshipID)
		dataKey.Buyer = append(dataKey.Buyer, v.Buyer)
		dataKey.Seller = append(dataKey.Seller, v.Seller)
	}

	for _, v := range sdc.Header.Item {
		if v.DeliverToParty == nil {
			return nil, xerrors.Errorf("入力ファイルの'DeliverToParty'がnullです。")
		} else if v.DeliverFromParty == nil {
			return nil, xerrors.Errorf("入力ファイルの'DeliverToParty'がnullです。")
		}
		dataKey.DeliverToParty = append(dataKey.DeliverToParty, *v.DeliverToParty)
		dataKey.DeliverFromParty = append(dataKey.DeliverFromParty, *v.DeliverFromParty)
	}

	if len(dataKey.SupplyChainRelationshipID) == 0 {
		return nil, xerrors.Errorf("入力ファイルから取得した'SupplyChainRelationshipID'がありません。")
	}
	repeat1 := strings.Repeat("(?,?,?),", len(dataKey.SupplyChainRelationshipID)-1) + "(?,?,?)"
	for i := range dataKey.SupplyChainRelationshipID {
		args = append(args, dataKey.SupplyChainRelationshipID[i], dataKey.Buyer[i], dataKey.Seller[i])
	}

	if len(dataKey.DeliverToParty) == 0 {
		return nil, xerrors.Errorf("入力ファイルの'DeliverToParty'がありません。")
	}
	repeat2 := strings.Repeat("(?,?),", len(dataKey.DeliverToParty)-1) + "(?,?)"
	for i := range dataKey.DeliverToParty {
		args = append(args, dataKey.DeliverToParty[i], dataKey.DeliverFromParty[i])
	}

	rows, err := f.db.Query(
		`SELECT SupplyChainRelationshipID, SupplyChainRelationshipDeliveryID, Buyer, Seller, DeliverToParty, DeliverFromParty
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_delivery_relation_data
		WHERE (SupplyChainRelationshipID, Buyer, Seller) IN ( `+repeat1+` )
		AND (DeliverToParty, DeliverFromParty) IN ( `+repeat2+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToSupplyChainRelationshipDeliveryRelation(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) SupplyChainRelationshipDeliveryPlantRelation(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.SupplyChainRelationshipDeliveryPlantRelation, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToSupplyChainRelationshipDeliveryPlantRelationKey()

	for _, v := range psdc.SupplyChainRelationshipDeliveryRelation {
		dataKey.SupplyChainRelationshipID = append(dataKey.SupplyChainRelationshipID, v.SupplyChainRelationshipID)
		dataKey.SupplyChainRelationshipDeliveryID = append(dataKey.SupplyChainRelationshipDeliveryID, v.SupplyChainRelationshipDeliveryID)
		dataKey.Buyer = append(dataKey.Buyer, v.Buyer)
		dataKey.Seller = append(dataKey.Seller, v.Seller)
		dataKey.DeliverToParty = append(dataKey.DeliverToParty, v.DeliverToParty)
		dataKey.DeliverFromParty = append(dataKey.DeliverFromParty, v.DeliverFromParty)
	}

	if len(dataKey.DeliverToParty) == 0 {
		return nil, xerrors.Errorf("psdc.SupplyChainRelationshipDeliveryRelationの'DeliverToParty'がありません。")
	}
	repeat := strings.Repeat("(?,?,?,?,?,?),", len(dataKey.DeliverToParty)-1) + "(?,?,?,?,?,?)"
	for i := range dataKey.SupplyChainRelationshipID {
		args = append(
			args,
			dataKey.SupplyChainRelationshipID[i],
			dataKey.SupplyChainRelationshipDeliveryID[i],
			dataKey.Buyer[i],
			dataKey.Seller[i],
			dataKey.DeliverToParty[i],
			dataKey.DeliverFromParty[i],
		)
	}

	args = append(args, dataKey.DefaultRelation)

	rows, err := f.db.Query(
		`SELECT SupplyChainRelationshipID, SupplyChainRelationshipDeliveryID, SupplyChainRelationshipDeliveryPlantID,
		Buyer, Seller, DeliverToParty, DeliverFromParty, DeliverToPlant, DeliverFromPlant, DefaultRelation
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_delivery_plant_rel_data
		WHERE (SupplyChainRelationshipID, SupplyChainRelationshipDeliveryID, Buyer, Seller, DeliverToParty, DeliverFromParty) IN ( `+repeat+` )
		AND DefaultRelation = ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToSupplyChainRelationshipDeliveryPlantRelation(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) SupplyChainRelationshipTransaction(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.SupplyChainRelationshipTransaction, error) {
	args := make([]interface{}, 0)

	supplyChainRelationshipGeneral := psdc.SupplyChainRelationshipGeneral

	repeat := strings.Repeat("(?,?,?),", len(supplyChainRelationshipGeneral)-1) + "(?,?,?)"
	for _, v := range supplyChainRelationshipGeneral {
		args = append(args, v.SupplyChainRelationshipID, v.Buyer, v.Seller)
	}

	rows, err := f.db.Query(
		`SELECT SupplyChainRelationshipID, Buyer, Seller, TransactionCurrency, Incoterms, PaymentTerms, PaymentMethod, AccountAssignmentGroup
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_transaction_data
		WHERE (SupplyChainRelationshipID, Buyer, Seller) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToSupplyChainRelationshipTransaction(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) SupplyChainRelationshipBillingRelation(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.SupplyChainRelationshipBillingRelation, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToSupplyChainRelationshipBillingRelationKey()

	for _, v := range psdc.SupplyChainRelationshipGeneral {
		dataKey.SupplyChainRelationshipID = append(dataKey.SupplyChainRelationshipID, v.SupplyChainRelationshipID)
		dataKey.Buyer = append(dataKey.Buyer, v.Buyer)
		dataKey.Seller = append(dataKey.Seller, v.Seller)
	}

	if len(dataKey.SupplyChainRelationshipID) == 0 {
		return nil, xerrors.Errorf("psdc.SupplyChainRelationshipGeneralの'SupplyChainRelationshipID'がありません。")
	}
	repeat := strings.Repeat("(?,?,?),", len(dataKey.SupplyChainRelationshipID)-1) + "(?,?,?)"
	for i := range dataKey.SupplyChainRelationshipID {
		args = append(
			args,
			dataKey.SupplyChainRelationshipID[i],
			dataKey.Buyer[i],
			dataKey.Seller[i],
		)
	}

	args = append(args, dataKey.DefaultRelation)

	rows, err := f.db.Query(
		`SELECT SupplyChainRelationshipID, SupplyChainRelationshipBillingID, Buyer, Seller, BillToParty, BillFromParty, DefaultRelation, BillToCountry, BillFromCountry, IsExportImport, TransactionTaxCategory, TransactionTaxClassification
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_billing_relation_data
		WHERE (SupplyChainRelationshipID, Buyer, Seller) IN ( `+repeat+` )
		AND DefaultRelation = ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToSupplyChainRelationshipBillingRelation(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) SupplyChainRelationshipPaymentRelation(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.SupplyChainRelationshipPaymentRelation, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToSupplyChainRelationshipPaymentRelationKey()

	for _, v := range psdc.SupplyChainRelationshipBillingRelation {
		dataKey.SupplyChainRelationshipID = append(dataKey.SupplyChainRelationshipID, v.SupplyChainRelationshipID)
		dataKey.Buyer = append(dataKey.Buyer, v.Buyer)
		dataKey.Seller = append(dataKey.Seller, v.Seller)
		dataKey.BillToParty = append(dataKey.BillToParty, v.BillToParty)
		dataKey.BillFromParty = append(dataKey.BillFromParty, v.BillFromParty)
	}

	if len(dataKey.SupplyChainRelationshipID) == 0 {
		return nil, xerrors.Errorf("psdc.SupplyChainRelationshipBillingRelation'SupplyChainRelationshipID'がありません。")
	}
	repeat := strings.Repeat("(?,?,?,?,?),", len(dataKey.SupplyChainRelationshipID)-1) + "(?,?,?,?,?)"
	for i := range dataKey.SupplyChainRelationshipID {
		args = append(
			args,
			dataKey.SupplyChainRelationshipID[i],
			dataKey.Buyer[i],
			dataKey.Seller[i],
			dataKey.BillToParty[i],
			dataKey.BillFromParty[i],
		)
	}

	args = append(args, dataKey.DefaultRelation)

	rows, err := f.db.Query(
		`SELECT SupplyChainRelationshipID, SupplyChainRelationshipBillingID, SupplyChainRelationshipPaymentID, Buyer, Seller, BillToParty, BillFromParty, Payer, Payee, DefaultRelation
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_payment_relation_data
		WHERE (SupplyChainRelationshipID, Buyer, Seller, BillToParty, BillFromParty) IN ( `+repeat+` )
		AND DefaultRelation = ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToSupplyChainRelationshipPaymentRelation(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) InvoiceDocumentDate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.InvoiceDocumentDate, error) {
	rows, err := f.db.Query(
		`SELECT PaymentTerms, BaseDate, BaseDateCalcAddMonth, BaseDateCalcFixedDate, PaymentDueDateCalcAddMonth, PaymentDueDateCalcFixedDate
			FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_payment_terms_payment_terms_data
			WHERE PaymentTerms = ?;`, psdc.SupplyChainRelationshipTransaction[0].PaymentTerms,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	psdc.PaymentTerms, err = psdc.ConvertToPaymentTerms(rows)
	if err != nil {
		return nil, err
	}

	if sdc.Header.InvoiceDocumentDate != nil {
		if *sdc.Header.InvoiceDocumentDate != "" {
			data, err := psdc.ConvertToInvoiceDocumentDate(sdc)
			return data, err
		}
	}

	requestedDeliveryDate, err := psdc.ConvertToRequestedDeliveryDate(sdc)
	if err != nil {
		return nil, err
	}

	invoiceDocumentDate, err := calculateInvoiceDocumentDate(psdc, requestedDeliveryDate.RequestedDeliveryDate, psdc.PaymentTerms)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToCaluculateInvoiceDocumentDate(sdc, invoiceDocumentDate)

	return data, err
}

func calculateInvoiceDocumentDate(
	psdc *api_processing_data_formatter.SDC,
	requestedDeliveryDate string,
	paymentTerms []*api_processing_data_formatter.PaymentTerms,
) (string, error) {
	format := "2006-01-02"
	t, err := time.Parse(format, requestedDeliveryDate)
	if err != nil {
		return "", err
	}

	sort.Slice(paymentTerms, func(i, j int) bool {
		return paymentTerms[i].BaseDate < paymentTerms[j].BaseDate
	})

	day := t.Day()
	for i, v := range paymentTerms {
		if day <= v.BaseDate {
			if v.BaseDateCalcAddMonth == nil {
				return "", xerrors.Errorf("PaymentTermsの'BaseDateCalcAddMonth'がnullです。")
			}
			t = time.Date(t.Year(), t.Month()+time.Month(*v.BaseDateCalcAddMonth)+1, 0, 0, 0, 0, 0, time.UTC)
			if v.BaseDateCalcFixedDate == nil {
				return "", xerrors.Errorf("PaymentTermsの'BaseDateCalcFixedDate'がnullです。")
			}
			if *v.BaseDateCalcFixedDate == 31 {
				t = time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.UTC)
			} else {
				t = time.Date(t.Year(), t.Month(), *v.BaseDateCalcFixedDate, 0, 0, 0, 0, time.UTC)
			}
			break
		}
		if len(paymentTerms) == 0 {
			return "", xerrors.Errorf("'paymentTerms'がありません。")
		}
		if i == len(paymentTerms)-1 {
			return "", xerrors.Errorf("'data_platform_payment_terms_payment_terms_data'テーブルが不適切です。")
		}
	}

	res := t.Format(format)

	return res, nil
}

func (f *SubFunction) PaymentDueDate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.PaymentDueDate, error) {
	data := make([]*api_processing_data_formatter.PaymentDueDate, 0)

	calculatePaymentDueDate, err := calculatePaymentDueDate(psdc.InvoiceDocumentDate.InvoiceDocumentDate, psdc.PaymentTerms)
	if err != nil {
		return nil, err
	}

	for _, item := range sdc.Header.Item {
		paymentDueDate := calculatePaymentDueDate

		if item.PaymentDueDate != nil {
			if *item.PaymentDueDate != "" {
				paymentDueDate = *item.PaymentDueDate
			}
		}

		datum := psdc.ConvertToPaymentDueDate(paymentDueDate)
		data = append(data, datum)
	}

	return data, err
}

func calculatePaymentDueDate(
	invoiceDocumentDate string,
	paymentTerms []*api_processing_data_formatter.PaymentTerms,
) (string, error) {
	format := "2006-01-02"
	t, err := time.Parse(format, invoiceDocumentDate)
	if err != nil {
		return "", err
	}

	sort.Slice(paymentTerms, func(i, j int) bool {
		return paymentTerms[i].BaseDate < paymentTerms[j].BaseDate
	})

	day := t.Day()
	if day == time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day() {
		day = 31
	}
	for i, v := range paymentTerms {
		if v.BaseDateCalcFixedDate == nil {
			return "", xerrors.Errorf("PaymentTermsの'BaseDateCalcFixedDate'がnullです。")
		}
		if day <= *v.BaseDateCalcFixedDate {
			if v.PaymentDueDateCalcAddMonth == nil {
				return "", xerrors.Errorf("PaymentTermsの'PaymentDueDateCalcAddMonth'がnullです。")
			}
			t = time.Date(t.Year(), t.Month()+time.Month(*v.PaymentDueDateCalcAddMonth)+1, 0, 0, 0, 0, 0, time.UTC)
			if v.PaymentDueDateCalcFixedDate == nil {
				return "", xerrors.Errorf("PaymentTermsの'PaymentDueDateCalcFixedDate'がnullです。")
			}
			if *v.PaymentDueDateCalcFixedDate == 31 {
				t = time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.UTC)
			} else {
				t = time.Date(t.Year(), t.Month(), *v.PaymentDueDateCalcFixedDate, 0, 0, 0, 0, time.UTC)
			}
			break
		}
		if len(paymentTerms) == 0 {
			return "", xerrors.Errorf("'paymentTerms'がありません。")
		}
		if i == len(paymentTerms)-1 {
			t = time.Date(t.Year(), t.Month()+time.Month(*v.PaymentDueDateCalcAddMonth)+2, 0, 0, 0, 0, 0, time.UTC)
			if v.PaymentDueDateCalcFixedDate == nil {
				return "", xerrors.Errorf("PaymentTermsの'PaymentDueDateCalcFixedDate'がnullです。")
			}
			if *v.PaymentDueDateCalcFixedDate == 31 {
				t = time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.UTC)
			} else {
				t = time.Date(t.Year(), t.Month(), *v.PaymentDueDateCalcFixedDate, 0, 0, 0, 0, time.UTC)
			}
		}
	}

	res := t.Format(format)

	return res, nil
}

func (f *SubFunction) NetPaymentDays(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.NetPaymentDays, error) {
	var err error
	data := make([]*api_processing_data_formatter.NetPaymentDays, 0)

	for i, item := range sdc.Header.Item {
		paymentDueDate := (psdc.PaymentDueDate)[i].PaymentDueDate

		calculateNetPaymentDays, err := calculateNetPaymentDays(psdc.InvoiceDocumentDate.InvoiceDocumentDate, paymentDueDate)
		if err != nil {
			return nil, err
		}

		netPaymentDays := calculateNetPaymentDays

		if item.NetPaymentDays != nil {
			netPaymentDays = *item.NetPaymentDays
		}

		datum := psdc.ConvertToNetPaymentDays(paymentDueDate, netPaymentDays)
		data = append(data, datum)
	}

	return data, err
}

func calculateNetPaymentDays(
	invoiceDocumentDate string,
	paymentDueDate string,
) (int, error) {
	format := "2006-01-02"
	tb, err := time.Parse(format, invoiceDocumentDate)
	if err != nil {
		return 0, err
	}

	tp, err := time.Parse(format, paymentDueDate)
	if err != nil {
		return 0, err
	}

	res := int(tp.Sub(tb).Hours() / 24)

	return res, nil
}

func (f *SubFunction) HeaderDocReferenceStatus(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.HeaderDocReferenceStatus {
	var headerDocReferenceStatus string
	serviceLabel := psdc.OrderReferenceDocumentType.ServiceLabel

	if serviceLabel == "QUOTATIONS" {
		headerDocReferenceStatus = "QT"
	} else if serviceLabel == "INQUIRIES" {
		headerDocReferenceStatus = "IN"
	} else if serviceLabel == "PURCHASE_REQUISITION" {
		headerDocReferenceStatus = "PR"
	}

	data := psdc.ConvertToHeaderDocReferenceStatus(headerDocReferenceStatus)

	return data
}

func (f *SubFunction) QuotationsHeader(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.QuotationsHeader, error) {
	rows, err := f.db.Query(
		`SELECT Quotation, QuotationType, QuotationStatus, SupplyChainRelationshipID, SupplyChainRelationshipBillingID, SupplyChainRelationshipPaymentID,
		Buyer, Seller, BillToParty, BillFromParty, BillToCountry, BillFromCountry, Payer, Payee, CreationDate, LastChangeDate, ContractType,
		BindingPeriodValidityStartDate, BindingPeriodValidityEndDate, OrderVaridityStartDate, OrderValidityEndDate, InvoicePeriodStartDate,
		InvoicePeriodEndDate, TotalNetAmount, TotalTaxAmount, TotalGrossAmount, TransactionCurrency, PricingDate, PriceDetnExchangeRate,
		RequestedDeliveryDate, OrderProbabilityInPercent, ExpectedOrderNetAmount, Incoterms, PaymentTerms, PaymentMethod, ReferenceDocument,
		ReferenceDocumentItem, AccountAssignmentGroup, AccountingExchangeRate, InvoiceDocumentDate, IsExportImport, HeaderText
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_quotations_header_data
		WHERE ReferenceDocument =  ?;`, sdc.InputParameters.ReferenceDocument,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToQuotationsHeader(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrdersHeader(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrdersHeader, error) {
	rows, err := f.db.Query(
		`SELECT OrderID, OrderDate, OrderType, SupplyChainRelationshipID, SupplyChainRelationshipBillingID, SupplyChainRelationshipPaymentID,
		Buyer, Seller, BillToParty, BillFromParty, BillToCountry, BillFromCountry, Payer, Payee, CreationDate, LastChangeDate, ContractType,
		OrderValidityStartDate, OrderValidityEndDate, InvoicePeriodStartDate, InvoicePeriodEndDate, TotalNetAmount, TotalTaxAmount, TotalGrossAmount,
		HeaderDeliveryStatus, HeaderBillingStatus, HeaderDocReferenceStatus, TransactionCurrency, PricingDate, PriceDetnExchangeRate, 
		RequestedDeliveryDate, HeaderCompleteDeliveryIsDefined, Incoterms, PaymentTerms, PaymentMethod, ReferenceDocument, ReferenceDocumentItem,
		AccountAssignmentGroup, AccountingExchangeRate, InvoiceDocumentDate, IsExportImport, HeaderText, HeaderBlockStatus, HeaderDeliveryBlockStatus,
		HeaderBillingBlockStatus, IsCancelled, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data
		WHERE (ReferenceDocument, ReferenceDocumentItem) =  (?, ?);`, sdc.InputParameters.ReferenceDocument, sdc.InputParameters.ReferenceDocumentItem,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToOrdersHeader(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) QuotationsPartner(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.QuotationsPartner, error) {
	rows, err := f.db.Query(
		`SELECT Quotation, PartnerFunction, BusinessPartner, BusinessPartnerFullName, BusinessPartnerName, Organization,
		Country, Language, Currency, ExternalDocumentID, AddressID
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_quotations_partner_data
		WHERE ReferenceDocument =  ?;`, sdc.InputParameters.ReferenceDocument,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToQuotationsPartner(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrdersPartner(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrdersPartner, error) {
	rows, err := f.db.Query(
		`SELECT OrderID, PartnerFunction, BusinessPartner, BusinessPartnerFullName, BusinessPartnerName, Organization, Country,
		Language, Currency, ExternalDocumentID, AddressID
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_partner_data
		WHERE ReferenceDocument =  ?;`, sdc.InputParameters.ReferenceDocument,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToOrdersPartner(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) QuotationsAddress(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.QuotationsAddress, error) {
	rows, err := f.db.Query(
		`SELECT Quotation, AddressID, PostalCode, LocalRegion, Country, District, StreetName, CityName, Building, Floor, Room
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_quotations_address_data
		WHERE ReferenceDocument =  ?;`, sdc.InputParameters.ReferenceDocument,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToQuotationsAddress(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrdersAddress(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrdersAddress, error) {
	rows, err := f.db.Query(
		`SELECT OrderID, AddressID, PostalCode, LocalRegion, Country, District, StreetName, CityName, Building, Floor, Room
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_address_data
		WHERE ReferenceDocument =  ?;`, sdc.InputParameters.ReferenceDocument,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToOrdersAddress(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) PricingDate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.PricingDate {
	pricingDate := getSystemDate()

	if sdc.Header.PricingDate != nil {
		if *sdc.Header.PricingDate != "" {
			pricingDate = *sdc.Header.PricingDate
		}
	}

	data := psdc.ConvertToPricingDate(pricingDate)

	return data
}

func (f *SubFunction) PriceDetnExchangeRate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.PriceDetnExchangeRate {
	data := new(api_processing_data_formatter.PriceDetnExchangeRate)

	if sdc.Header.PriceDetnExchangeRate != nil {
		data = psdc.ConvertToPriceDetnExchangeRate(sdc)
	}

	return data
}

func (f *SubFunction) AccountingExchangeRate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.AccountingExchangeRate {
	data := new(api_processing_data_formatter.AccountingExchangeRate)

	if sdc.Header.AccountingExchangeRate != nil {
		data = psdc.ConvertToAccountingExchangeRate(sdc)
	}

	return data
}

func (f *SubFunction) TotalNetAmount(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.TotalNetAmount, error) {
	totalNetAmount := float32(0)

	netAmount := psdc.NetAmount

	for _, v := range netAmount {
		if v.NetAmount != nil {
			totalNetAmount += *v.NetAmount
		}
	}

	if sdc.Header.TotalNetAmount != nil {
		if *sdc.Header.TotalNetAmount != totalNetAmount {
			return nil, xerrors.Errorf("入力ファイルのTotalNetAmountと計算結果が一致しません。")
		}
	}

	data := psdc.ConvertToTotalNetAmount(totalNetAmount)

	return data, nil
}

func (f *SubFunction) TotalTaxAmount(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.TotalTaxAmount, error) {
	totalTaxAmount := float32(0)

	taxAmount := psdc.TaxAmount

	for _, v := range taxAmount {
		if v.TaxAmount != nil {
			totalTaxAmount += *v.TaxAmount
		}
	}

	if sdc.Header.TotalTaxAmount != nil {
		if *sdc.Header.TotalTaxAmount != totalTaxAmount {
			return nil, xerrors.Errorf("入力ファイルのTotalTaxAmountと計算結果が一致しません。")
		}
	}

	data := psdc.ConvertToTotalTaxAmount(totalTaxAmount)

	return data, nil
}

func (f *SubFunction) TotalGrossAmount(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.TotalGrossAmount, error) {
	totalGrossAmount := float32(0)

	grossAmount := psdc.GrossAmount

	for _, v := range grossAmount {
		if v.GrossAmount != nil {
			totalGrossAmount += *v.GrossAmount
		}
	}

	if sdc.Header.TotalGrossAmount != nil {
		if *sdc.Header.TotalGrossAmount != totalGrossAmount {
			return nil, xerrors.Errorf("入力ファイルのTotalGrossAmountと計算結果が一致しません。")
		}
	}

	data := psdc.ConvertToTotalGrossAmount(totalGrossAmount)

	return data, nil
}

func (f *SubFunction) CreationDateHeader(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.CreationDateHeader {
	data := psdc.ConvertToCreationDateHeader(getSystemDate())

	return data
}

func (f *SubFunction) LastChangeDateHeader(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.LastChangeDateHeader {
	data := psdc.ConvertToLastChangeDateHeader(getSystemDate())

	return data
}

func getSystemDate() string {
	day := time.Now()
	return day.Format("2006-01-02")
}
