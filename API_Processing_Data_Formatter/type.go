package api_processing_data_formatter

type SDC struct {
	MetaData                                     *MetaData                                       `json:"MetaData"`
	OrderRegistrationType                        *OrderRegistrationType                          `json:"OrderRegistrationType"`
	OrderReferenceDocumentType                   *OrderReferenceDocumentType                     `json:"OrderReferenceDocumentType"`
	SupplyChainRelationshipGeneral               []*SupplyChainRelationshipGeneral               `json:"SupplyChainRelationshipGeneral"`
	SupplyChainRelationshipDeliveryRelation      []*SupplyChainRelationshipDeliveryRelation      `json:"SupplyChainRelationshipDeliveryRelation"`
	SupplyChainRelationshipDeliveryPlantRelation []*SupplyChainRelationshipDeliveryPlantRelation `json:"SupplyChainRelationshipDeliveryPlantRelation"`
	SupplyChainRelationshipTransaction           []*SupplyChainRelationshipTransaction           `json:"SupplyChainRelationshipTransaction"`
	SupplyChainRelationshipBillingRelation       []*SupplyChainRelationshipBillingRelation       `json:"SupplyChainRelationshipBillingRelation"`
	SupplyChainRelationshipPaymentRelation       []*SupplyChainRelationshipPaymentRelation       `json:"SupplyChainRelationshipPaymentRelation"`
	PaymentTerms                                 []*PaymentTerms                                 `json:"PaymentTerms"`
	InvoiceDocumentDate                          *InvoiceDocumentDate                            `json:"InvoiceDocumentDate"`
	PaymentDueDate                               []*PaymentDueDate                               `json:"PaymentDueDate"`
	NetPaymentDays                               []*NetPaymentDays                               `json:"NetPaymentDays"`
	HeaderDocReferenceStatus                     *HeaderDocReferenceStatus                       `json:"HeaderDocReferenceStatus"`
	PricingDate                                  *PricingDate                                    `json:"PricingDate"`
	PriceDetnExchangeRate                        *PriceDetnExchangeRate                          `json:"PriceDetnExchangeRate"`
	AccountingExchangeRate                       *AccountingExchangeRate                         `json:"AccountingExchangeRate"`
	TotalNetAmount                               *TotalNetAmount                                 `json:"TotalNetAmount"`
	TotalTaxAmount                               *TotalTaxAmount                                 `json:"TotalTaxAmount"`
	TotalGrossAmount                             *TotalGrossAmount                               `json:"TotalGrossAmount"`
	CreationDateHeader                           *CreationDateHeader                             `json:"CreationDateHeader"`
	LastChangeDateHeader                         *LastChangeDateHeader                           `json:"LastChangeDateHeader"`
	ProductTaxClassificationBillToCountry        []*ProductTaxClassificationBillToCountry        `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry      []*ProductTaxClassificationBillFromCountry      `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                     []*DefinedTaxClassification                     `json:"DefinedTaxClassification"`
	TaxCode                                      []*TaxCode                                      `json:"TaxCode"`
	TaxRate                                      []*TaxRate                                      `json:"TaxRate"`
	NetAmount                                    []*NetAmount                                    `json:"NetAmount"`
	TaxAmount                                    []*TaxAmount                                    `json:"TaxAmount"`
	GrossAmount                                  []*GrossAmount                                  `json:"GrossAmount"`
	PriceMaster                                  []*PriceMaster                                  `json:"PriceMaster"`
	ConditionAmount                              []*ConditionAmount                              `json:"ConditionAmount"`
}

type MetaData struct {
	BusinessPartnerID *int   `json:"business_partner"`
	ServiceLabel      string `json:"service_label"`
}

type OrderRegistrationType struct {
	ReferenceDocument     *int   `json:"ReferenceDocument"`
	ReferenceDocumentItem *int   `json:"ReferenceDocumentItem"`
	RegistrationType      string `json:"RegistrationType"`
}

type OrderReferenceDocumentTypeQueryGets struct {
	ServiceLabel             string `json:"ServiceLabel"`
	FieldNameWithNumberRange string `json:"FieldNameWithNumberRange"`
	NumberRangeFrom          *int   `json:"NumberRangeFrom"`
	NumberRangeTo            *int   `json:"NumberRangeTo"`
}

type OrderReferenceDocumentType struct {
	ServiceLabel string `json:"ServiceLabel"`
}

// Header
type SupplyChainRelationshipGeneral struct {
	SupplyChainRelationshipID int `json:"SupplyChainRelationshipID"`
	Buyer                     int `json:"Buyer"`
	Seller                    int `json:"Seller"`
}

type SupplyChainRelationshipDeliveryRelationKey struct {
	SupplyChainRelationshipID []int `json:"SupplyChainRelationshipID"`
	Buyer                     []int `json:"Buyer"`
	Seller                    []int `json:"Seller"`
	DeliverToParty            []int `json:"DeliverToParty"`
	DeliverFromParty          []int `json:"DeliverFromParty"`
}

type SupplyChainRelationshipDeliveryRelation struct {
	SupplyChainRelationshipID         int `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID int `json:"SupplyChainRelationshipDeliveryID"`
	Buyer                             int `json:"Buyer"`
	Seller                            int `json:"Seller"`
	DeliverToParty                    int `json:"DeliverToParty"`
	DeliverFromParty                  int `json:"DeliverFromParty"`
}

type SupplyChainRelationshipDeliveryPlantRelationKey struct {
	SupplyChainRelationshipID         []int `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID []int `json:"SupplyChainRelationshipDeliveryID"`
	Buyer                             []int `json:"Buyer"`
	Seller                            []int `json:"Seller"`
	DeliverToParty                    []int `json:"DeliverToParty"`
	DeliverFromParty                  []int `json:"DeliverFromParty"`
	DefaultRelation                   bool  `json:"DefaultRelation"`
}

type SupplyChainRelationshipDeliveryPlantRelation struct {
	SupplyChainRelationshipID              int    `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int    `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int    `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int    `json:"Buyer"`
	Seller                                 int    `json:"Seller"`
	DeliverToParty                         int    `json:"DeliverToParty"`
	DeliverFromParty                       int    `json:"DeliverFromParty"`
	DeliverToPlant                         string `json:"DeliverToPlant"`
	DeliverFromPlant                       string `json:"DeliverFromPlant"`
	DefaultRelation                        *bool  `json:"DefaultRelation"`
}

type SupplyChainRelationshipTransaction struct {
	SupplyChainRelationshipID int     `json:"SupplyChainRelationshipID"`
	Buyer                     int     `json:"Buyer"`
	Seller                    int     `json:"Seller"`
	TransactionCurrency       *string `json:"TransactionCurrency"`
	Incoterms                 *string `json:"Incoterms"`
	PaymentTerms              *string `json:"PaymentTerms"`
	PaymentMethod             *string `json:"PaymentMethod"`
	AccountAssignmentGroup    *string `json:"AccountAssignmentGroup"`
}

type SupplyChainRelationshipBillingRelationKey struct {
	SupplyChainRelationshipID []int `json:"SupplyChainRelationshipID"`
	Buyer                     []int `json:"Buyer"`
	Seller                    []int `json:"Seller"`
	DefaultRelation           bool  `json:"DefaultRelation"`
}

type SupplyChainRelationshipBillingRelation struct {
	SupplyChainRelationshipID        int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID int     `json:"SupplyChainRelationshipBillingID"`
	Buyer                            int     `json:"Buyer"`
	Seller                           int     `json:"Seller"`
	BillToParty                      int     `json:"BillToParty"`
	BillFromParty                    int     `json:"BillFromParty"`
	DefaultRelation                  *bool   `json:"DefaultRelation"`
	BillToCountry                    string  `json:"BillToCountry"`
	BillFromCountry                  string  `json:"BillFromCountry"`
	IsExportImport                   *bool   `json:"IsExportImport"`
	TransactionTaxCategory           *string `json:"TransactionTaxCategory"`
	TransactionTaxClassification     *string `json:"TransactionTaxClassification"`
}

type SupplyChainRelationshipPaymentRelationKey struct {
	SupplyChainRelationshipID []int `json:"SupplyChainRelationshipID"`
	Buyer                     []int `json:"Buyer"`
	Seller                    []int `json:"Seller"`
	BillToParty               []int `json:"BillToParty"`
	BillFromParty             []int `json:"BillFromParty"`
	DefaultRelation           bool  `json:"DefaultRelation"`
}

type SupplyChainRelationshipPaymentRelation struct {
	SupplyChainRelationshipID        int   `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID int   `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID int   `json:"SupplyChainRelationshipPaymentID"`
	Buyer                            int   `json:"Buyer"`
	Seller                           int   `json:"Seller"`
	BillToParty                      int   `json:"BillToParty"`
	BillFromParty                    int   `json:"BillFromParty"`
	Payer                            int   `json:"Payer"`
	Payee                            int   `json:"Payee"`
	DefaultRelation                  *bool `json:"DefaultRelation"`
}

type PaymentTerms struct {
	PaymentTerms                string `json:"PaymentTerms"`
	BaseDate                    int    `json:"BaseDate"`
	BaseDateCalcAddMonth        *int   `json:"BaseDateCalcAddMonth"`
	BaseDateCalcFixedDate       *int   `json:"BaseDateCalcFixedDate"`
	PaymentDueDateCalcAddMonth  *int   `json:"PaymentDueDateCalcAddMonth"`
	PaymentDueDateCalcFixedDate *int   `json:"PaymentDueDateCalcFixedDate"`
}

type InvoiceDocumentDate struct {
	RequestedDeliveryDate string `json:"RequestedDeliveryDate"`
	InvoiceDocumentDate   string `json:"InvoiceDocumentDate"`
}

type PaymentDueDate struct {
	InvoiceDocumentDate string `json:"InvoiceDocumentDate"`
	PaymentDueDate      string `json:"PaymentDueDate"`
}

type NetPaymentDays struct {
	InvoiceDocumentDate string `json:"InvoiceDocumentDate"`
	PaymentDueDate      string `json:"PaymentDueDate"`
	NetPaymentDays      *int   `json:"NetPaymentDays"`
}

type HeaderDocReferenceStatus struct {
	HeaderDocReferenceStatus string `json:"HeaderDocReferenceStatus"`
}

type PricingDate struct {
	PricingDate string `json:"PricingDate"`
}

type PriceDetnExchangeRate struct {
	PriceDetnExchangeRate *float32 `json:"PriceDetnExchangeRate"`
}

type AccountingExchangeRate struct {
	AccountingExchangeRate *float32 `json:"AccountingExchangeRate"`
}

type TotalNetAmount struct {
	TotalNetAmount float32 `json:"TotalNetAmount"`
}

type TotalTaxAmount struct {
	TotalTaxAmount float32 `json:"TotalTaxAmount"`
}

type TotalGrossAmount struct {
	TotalGrossAmount float32 `json:"TotalGrossAmount"`
}

type CreationDateHeader struct {
	CreationDate string `json:"CreationDate"`
}

type LastChangeDateHeader struct {
	LastChangeDate string `json:"LastChangeDate"`
}

// Item
type ProductTaxClassificationKey struct {
	Product            []*string `json:"Product"`
	Country            string    `json:"Country"`
	ProductTaxCategory string    `json:"ProductTaxCategory"`
}

type ProductTaxClassificationBillToCountry struct {
	Product                   string  `json:"Product"`
	Country                   string  `json:"Country"`
	ProductTaxCategory        string  `json:"ProductTaxCategory"`
	ProductTaxClassifiication *string `json:"ProductTaxClassification"`
}

type ProductTaxClassificationBillFromCountry struct {
	Product                   string  `json:"Product"`
	Country                   string  `json:"Country"`
	ProductTaxCategory        string  `json:"ProductTaxCategory"`
	ProductTaxClassifiication *string `json:"ProductTaxClassification"`
}

type DefinedTaxClassification struct {
	Product                                 string  `json:"Product"`
	TransactionTaxClassification            *string `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry   *string `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry *string `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                string  `json:"DefinedTaxClassification"`
}

type TaxCode struct {
	Product                  string  `json:"Product"`
	DefinedTaxClassification string  `json:"DefinedTaxClassification"`
	IsExportImport           *bool   `json:"IsExportImport"`
	TaxCode                  *string `json:"TaxCode"`
}

type TaxRateKey struct {
	Country           string    `json:"Country"`
	TaxCode           []*string `json:"TaxCode"`
	ValidityEndDate   string    `json:"ValidityEndDate"`
	ValidityStartDate string    `json:"ValidityStartDate"`
}

type TaxRate struct {
	Country           string   `json:"Country"`
	TaxCode           string   `json:"TaxCode"`
	ValidityEndDate   string   `json:"ValidityEndDate"`
	ValidityStartDate string   `json:"ValidityStartDate"`
	TaxRate           *float32 `json:"TaxRate"`
}

type NetAmount struct {
	OrderItem int      `json:"OrderItem"`
	Product   string   `json:"Product"`
	NetAmount *float32 `json:"NetAmount"`
}

type TaxAmount struct {
	OrderItem int      `json:"OrderItem"`
	Product   string   `json:"Product"`
	TaxCode   string   `json:"TaxCode"`
	TaxRate   *float32 `json:"TaxRate"`
	NetAmount *float32 `json:"NetAmount"`
	TaxAmount *float32 `json:"TaxAmount"`
}

type GrossAmount struct {
	OrderItem   int      `json:"OrderItem"`
	Product     string   `json:"Product"`
	NetAmount   *float32 `json:"NetAmount"`
	TaxAmount   *float32 `json:"TaxAmount"`
	GrossAmount *float32 `json:"GrossAmount"`
}

// Item Pricing Element
type PriceMasterKey struct {
	Product                    []*string `json:"Product"`
	SupplyChainRelationshipID  int       `json:"SupplyChainRelationshipID"`
	Buyer                      int       `json:"Buyer"`
	Seller                     int       `json:"Seller"`
	ConditionValidityEndDate   string    `json:"ConditionValidityEndDate"`
	ConditionValidityStartDate string    `json:"ConditionValidityStartDate"`
}

type PriceMaster struct {
	SupplyChainRelationshipID  int      `json:"SupplyChainRelationshipID"`
	Buyer                      int      `json:"Buyer"`
	Seller                     int      `json:"Seller"`
	ConditionRecord            int      `json:"ConditionRecord"`
	ConditionSequentialNumber  int      `json:"ConditionSequentialNumber"`
	ConditionValidityStartDate string   `json:"ConditionValidityStartDate"`
	ConditionValidityEndDate   string   `json:"ConditionValidityEndDate"`
	Product                    string   `json:"Product"`
	ConditionType              string   `json:"ConditionType"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
}

type ConditionAmount struct {
	OrderItem                  int      `json:"OrderItem"`
	Product                    string   `json:"Product"`
	ConditionQuantity          *float32 `json:"ConditionQuantity"`
	ConditionAmount            *float32 `json:"ConditionAmount"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged"`
}
