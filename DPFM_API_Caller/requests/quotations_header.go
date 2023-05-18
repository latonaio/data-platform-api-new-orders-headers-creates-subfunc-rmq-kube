package requests

type QuotationsHeader struct {
	Quotation                        int     `json:"Quotation"`
	QuotationType                    string  `json:"QuotationType"`
	QuotationStatus                  string  `json:"QuotationStatus"`
	SupplyChainRelationshipID        int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID int     `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID int     `json:"SupplyChainRelationshipPaymentID"`
	Buyer                            int     `json:"Buyer"`
	Seller                           int     `json:"Seller"`
	BillToParty                      int     `json:"BillToParty"`
	BillFromParty                    int     `json:"BillFromParty"`
	BillToCountry                    string  `json:"BillToCountry"`
	BillFromCountry                  string  `json:"BillFromCountry"`
	Payer                            int     `json:"Payer"`
	Payee                            int     `json:"Payee"`
	CreationDate                     string  `json:"CreationDate"`
	LastChangeDate                   string  `json:"LastChangeDate"`
	ContractType                     string  `json:"ContractType"`
	BindingPeriodValidityStartDate   string  `json:"BindingPeriodValidityStartDate"`
	BindingPeriodValidityEndDate     string  `json:"BindingPeriodValidityEndDate"`
	OrderVaridityStartDate           string  `json:"OrderVaridityStartDate"`
	OrderValidityEndDate             string  `json:"OrderValidityEndDate"`
	InvoicePeriodStartDate           string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate             string  `json:"InvoicePeriodEndDate"`
	TotalNetAmount                   float32 `json:"TotalNetAmount"`
	TotalTaxAmount                   float32 `json:"TotalTaxAmount"`
	TotalGrossAmount                 float32 `json:"TotalGrossAmount"`
	TransactionCurrency              string  `json:"TransactionCurrency"`
	PricingDate                      string  `json:"PricingDate"`
	PriceDetnExchangeRate            float32 `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate            string  `json:"RequestedDeliveryDate"`
	OrderProbabilityInPercent        float32 `json:"OrderProbabilityInPercent"`
	ExpectedOrderNetAmount           float32 `json:"ExpectedOrderNetAmount"`
	Incoterms                        string  `json:"Incoterms"`
	PaymentTerms                     string  `json:"PaymentTerms"`
	PaymentMethod                    string  `json:"PaymentMethod"`
	ReferenceDocument                int     `json:"ReferenceDocument"`
	ReferenceDocumentItem            int     `json:"ReferenceDocumentItem"`
	AccountAssignmentGroup           string  `json:"AccountAssignmentGroup"`
	AccountingExchangeRate           float32 `json:"AccountingExchangeRate"`
	InvoiceDocumentDate              string  `json:"InvoiceDocumentDate"`
	IsExportImport                   bool    `json:"IsExportImport"`
	HeaderText                       string  `json:"HeaderText"`
}
