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
	QuotationsHeader                             *QuotationsHeader                               `json:"QuotationsHeader"`
	QuotationsItem                               []*QuotationsItem                               `json:"QuotationsItem"`
	QuotationsItemPricingElement                 []*QuotationsItemPricingElement                 `json:"QuotationsItemPricingElement"`
	QuotationsPartner                            []*QuotationsPartner                            `json:"QuotationsPartner"`
	QuotationsAddress                            []*QuotationsAddress                            `json:"QuotationsAddress"`
	OrdersHeader                                 []*OrdersHeader                                 `json:"OrdersHeader"`
	OrdersItem                                   []*OrdersItem                                   `json:"OrdersItem"`
	OrdersItemPricingElement                     []*OrdersItemPricingElement                     `json:"OrdersItemPricingElement"`
	OrdersItemScheduleLine                       []*OrdersItemScheduleLine                       `json:"OrdersItemScheduleLine"`
	OrdersPartner                                []*OrdersPartner                                `json:"OrdersPartner"`
	OrdersAddress                                []*OrdersAddress                                `json:"OrdersAddress"`
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

type OrdersHeader struct {
	OrderID                          int      `json:"OrderID"`
	OrderDate                        string   `json:"OrderDate"`
	OrderType                        string   `json:"OrderType"`
	SupplyChainRelationshipID        int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID *int     `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID *int     `json:"SupplyChainRelationshipPaymentID"`
	Buyer                            int      `json:"Buyer"`
	Seller                           int      `json:"Seller"`
	BillToParty                      *int     `json:"BillToParty"`
	BillFromParty                    *int     `json:"BillFromParty"`
	BillToCountry                    *string  `json:"BillToCountry"`
	BillFromCountry                  *string  `json:"BillFromCountry"`
	Payer                            *int     `json:"Payer"`
	Payee                            *int     `json:"Payee"`
	CreationDate                     string   `json:"CreationDate"`
	LastChangeDate                   string   `json:"LastChangeDate"`
	ContractType                     *string  `json:"ContractType"`
	OrderValidityStartDate           *string  `json:"OrderValidityStartDate"`
	OrderValidityEndDate             *string  `json:"OrderValidityEndDate"`
	InvoicePeriodStartDate           *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate             *string  `json:"InvoicePeriodEndDate"`
	TotalNetAmount                   float32  `json:"TotalNetAmount"`
	TotalTaxAmount                   float32  `json:"TotalTaxAmount"`
	TotalGrossAmount                 float32  `json:"TotalGrossAmount"`
	HeaderDeliveryStatus             string   `json:"HeaderDeliveryStatus"`
	HeaderBillingStatus              string   `json:"HeaderBillingStatus"`
	HeaderDocReferenceStatus         string   `json:"HeaderDocReferenceStatus"`
	TransactionCurrency              string   `json:"TransactionCurrency"`
	PricingDate                      string   `json:"PricingDate"`
	PriceDetnExchangeRate            *bool    `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate            string   `json:"RequestedDeliveryDate"`
	HeaderCompleteDeliveryIsDefined  *bool    `json:"HeaderCompleteDeliveryIsDefined"`
	Incoterms                        *string  `json:"Incoterms"`
	PaymentTerms                     string   `json:"PaymentTerms"`
	PaymentMethod                    string   `json:"PaymentMethod"`
	ReferenceDocument                *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem            *int     `json:"ReferenceDocumentItem"`
	AccountAssignmentGroup           string   `json:"AccountAssignmentGroup"`
	AccountingExchangeRate           *float32 `json:"AccountingExchangeRate"`
	InvoiceDocumentDate              string   `json:"InvoiceDocumentDate"`
	IsExportImport                   *bool    `json:"IsExportImport"`
	HeaderText                       *string  `json:"HeaderText"`
	HeaderBlockStatus                *bool    `json:"HeaderBlockStatus"`
	HeaderDeliveryBlockStatus        *bool    `json:"HeaderDeliveryBlockStatus"`
	HeaderBillingBlockStatus         *bool    `json:"HeaderBillingBlockStatus"`
	IsCancelled                      *bool    `json:"IsCancelled"`
	IsMarkedForDeletion              *bool    `json:"IsMarkedForDeletion"`
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

type QuotationsItem struct {
	Quotation                               int     `json:"Quotation"`
	QuotationItem                           int     `json:"QuotationItem"`
	QuotationItemCategory                   string  `json:"QuotationItemCategory"`
	SupplyChainRelationshipID               int     `json:"SupplyChainRelationshipID"`
	QuotationItemText                       string  `json:"QuotationItemText"`
	QuotationItemTextByBuyer                string  `json:"QuotationItemTextByBuyer"`
	QuotationItemTextBySeller               string  `json:"QuotationItemTextBySeller"`
	Product                                 string  `json:"Product"`
	ProductStandardID                       string  `json:"ProductStandardID"`
	ProductGroup                            string  `json:"ProductGroup"`
	BaseUnit                                string  `json:"BaseUnit"`
	PricingDate                             string  `json:"PricingDate"`
	PriceDetnExchangeRate                   float32 `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate                   string  `json:"RequestedDeliveryDate"`
	CreationDate                            string  `json:"CreationDate"`
	LastChangeDate                          string  `json:"LastChangeDate"`
	DeliveryUnit                            string  `json:"DeliveryUnit"`
	ServicesRenderingDate                   string  `json:"ServicesRenderingDate"`
	QuotationQuantityInBaseUnit             float32 `json:"QuotationQuantityInBaseUnit"`
	QuotationQuantityInDeliveryUnit         float32 `json:"QuotationQuantityInDeliveryUnit"`
	ItemWeightUnit                          string  `json:"ItemWeightUnit"`
	ProductGrossWeight                      float32 `json:"ProductGrossWeight"`
	ItemGrossWeight                         float32 `json:"ItemGrossWeight"`
	ProductNetWeight                        float32 `json:"ProductNetWeight"`
	ItemNetWeight                           float32 `json:"ItemNetWeight"`
	InternalCapacityQuantity                float32 `json:"InternalCapacityQuantity"`
	InternalCapacityQuantityUnit            string  `json:"InternalCapacityQuantityUnit"`
	NetAmount                               float32 `json:"NetAmount"`
	TaxAmount                               float32 `json:"TaxAmount"`
	GrossAmount                             float32 `json:"GrossAmount"`
	Incoterms                               string  `json:"Incoterms"`
	TransactionTaxClassification            string  `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry   string  `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry string  `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                string  `json:"DefinedTaxClassification"`
	AccountAssignmentGroup                  string  `json:"AccountAssignmentGroup"`
	ProductAccountAssignmentGroup           string  `json:"ProductAccountAssignmentGroup"`
	PaymentTerms                            string  `json:"PaymentTerms"`
	PaymentMethod                           string  `json:"PaymentMethod"`
	Project                                 string  `json:"Project"`
	AccountingExchangeRate                  float32 `json:"AccountingExchangeRate"`
	ReferenceDocument                       int     `json:"ReferenceDocument"`
	ReferenceDocumentItem                   int     `json:"ReferenceDocumentItem"`
	TaxCode                                 string  `json:"TaxCode"`
	TaxRate                                 float32 `json:"TaxRate"`
	CountryOfOrigin                         string  `json:"CountryOfOrigin"`
	CountryOfOriginLanguage                 string  `json:"CountryOfOriginLanguage"`
}
type OrdersItem struct {
	OrderID                                       int      `json:"OrderID"`
	OrderItem                                     int      `json:"OrderItem"`
	OrderItemCategory                             string   `json:"OrderItemCategory"`
	SupplyChainRelationshipID                     int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID             *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID        *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID       *int     `json:"SupplyChainRelationshipStockConfPlantID"`
	SupplyChainRelationshipProductionPlantID      *int     `json:"SupplyChainRelationshipProductionPlantID"`
	OrderItemText                                 string   `json:"OrderItemText"`
	OrderItemTextByBuyer                          string   `json:"OrderItemTextByBuyer"`
	OrderItemTextBySeller                         string   `json:"OrderItemTextBySeller"`
	Product                                       string   `json:"Product"`
	ProductStandardID                             string   `json:"ProductStandardID"`
	ProductGroup                                  *string  `json:"ProductGroup"`
	BaseUnit                                      string   `json:"BaseUnit"`
	PricingDate                                   float32  `json:"PricingDate"`
	PriceDetnExchangeRate                         *string  `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate                         int      `json:"RequestedDeliveryDate"`
	DeliverToParty                                *int     `json:"DeliverToParty"`
	DeliverFromParty                              *int     `json:"DeliverFromParty"`
	CreationDate                                  string   `json:"CreationDate"`
	LastChangeDate                                string   `json:"LastChangeDate"`
	DeliverToPlant                                *string  `json:"DeliverToPlant"`
	DeliverToPlantTimeZone                        *string  `json:"DeliverToPlantTimeZone"`
	DeliverToPlantStorageLocation                 *string  `json:"DeliverToPlantStorageLocation"`
	ProductIsBatchManagedInDeliverToPlant         *bool    `json:"ProductIsBatchManagedInDeliverToPlant"`
	BatchMgmtPolicyInDeliverToPlant               *string  `json:"BatchMgmtPolicyInDeliverToPlant"`
	DeliverToPlantBatch                           *string  `json:"DeliverToPlantBatch"`
	DeliverToPlantBatchValidityStartDate          *string  `json:"DeliverToPlantBatchValidityStartDate"`
	DeliverToPlantBatchValidityStartTime          *string  `json:"DeliverToPlantBatchValidityStartTime"`
	DeliverToPlantBatchValidityEndDate            *string  `json:"DeliverToPlantBatchValidityEndDate"`
	DeliverToPlantBatchValidityEndTime            *string  `json:"DeliverToPlantBatchValidityEndTime"`
	DeliverFromPlant                              *string  `json:"DeliverFromPlant"`
	DeliverFromPlantTimeZone                      *string  `json:"DeliverFromPlantTimeZone"`
	DeliverFromPlantStorageLocation               *string  `json:"DeliverFromPlantStorageLocation"`
	ProductIsBatchManagedInDeliverFromPlant       *bool    `json:"ProductIsBatchManagedInDeliverFromPlant"`
	BatchMgmtPolicyInDeliverFromPlant             *string  `json:"BatchMgmtPolicyInDeliverFromPlant"`
	DeliverFromPlantBatch                         *string  `json:"DeliverFromPlantBatch"`
	DeliverFromPlantBatchValidityStartDate        *string  `json:"DeliverFromPlantBatchValidityStartDate"`
	DeliverFromPlantBatchValidityStartTime        *string  `json:"DeliverFromPlantBatchValidityStartTime"`
	DeliverFromPlantBatchValidityEndDate          *string  `json:"DeliverFromPlantBatchValidityEndDate"`
	DeliverFromPlantBatchValidityEndTime          *string  `json:"DeliverFromPlantBatchValidityEndTime"`
	DeliveryUnit                                  string   `json:"DeliveryUnit"`
	StockConfirmationBusinessPartner              *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                        *string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantTimeZone                *string  `json:"StockConfirmationPlantTimeZone"`
	ProductIsBatchManagedInStockConfirmationPlant *bool    `json:"ProductIsBatchManagedInStockConfirmationPlant"`
	BatchMgmtPolicyInStockConfirmationPlant       *string  `json:"BatchMgmtPolicyInStockConfirmationPlant"`
	StockConfirmationPlantBatch                   *string  `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate  *string  `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityStartTime  *string  `json:"StockConfirmationPlantBatchValidityStartTime"`
	StockConfirmationPlantBatchValidityEndDate    *string  `json:"StockConfirmationPlantBatchValidityEndDate"`
	StockConfirmationPlantBatchValidityEndTime    *string  `json:"StockConfirmationPlantBatchValidityEndTime"`
	ServicesRenderingDate                         *string  `json:"ServicesRenderingDate"`
	OrderQuantityInBaseUnit                       float32  `json:"OrderQuantityInBaseUnit"`
	OrderQuantityInDeliveryUnit                   float32  `json:"OrderQuantityInDeliveryUnit"`
	StockConfirmationPolicy                       *string  `json:"StockConfirmationPolicy"`
	StockConfirmationStatus                       *string  `json:"StockConfirmationStatus"`
	ConfirmedOrderQuantityInBaseUnit              *float32 `json:"ConfirmedOrderQuantityInBaseUnit"`
	ItemWeightUnit                                *string  `json:"ItemWeightUnit"`
	ProductGrossWeight                            *float32 `json:"ProductGrossWeight"`
	ItemGrossWeight                               *float32 `json:"ItemGrossWeight"`
	ProductNetWeight                              *float32 `json:"ProductNetWeight"`
	ItemNetWeight                                 *float32 `json:"ItemNetWeight"`
	InternalCapacityQuantity                      *float32 `json:"InternalCapacityQuantity"`
	InternalCapacityQuantityUnit                  *string  `json:"InternalCapacityQuantityUnit"`
	NetAmount                                     *float32 `json:"NetAmount"`
	TaxAmount                                     *float32 `json:"TaxAmount"`
	GrossAmount                                   *float32 `json:"GrossAmount"`
	InvoiceDocumentDate                           *string  `json:"InvoiceDocumentDate"`
	ProductionPlantBusinessPartner                *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                               *string  `json:"ProductionPlant"`
	ProductionPlantTimeZone                       *string  `json:"ProductionPlantTimeZone"`
	ProductionPlantStorageLocation                *string  `json:"ProductionPlantStorageLocation"`
	ProductIsBatchManagedInProductionPlant        *bool    `json:"ProductIsBatchManagedInProductionPlant"`
	BatchMgmtPolicyInProductionPlant              *string  `json:"BatchMgmtPolicyInProductionPlant"`
	ProductionPlantBatch                          *string  `json:"ProductionPlantBatch"`
	ProductionPlantBatchValidityStartDate         *string  `json:"ProductionPlantBatchValidityStartDate"`
	ProductionPlantBatchValidityStartTime         *string  `json:"ProductionPlantBatchValidityStartTime"`
	ProductionPlantBatchValidityEndDate           *string  `json:"ProductionPlantBatchValidityEndDate"`
	ProductionPlantBatchValidityEndTime           *string  `json:"ProductionPlantBatchValidityEndTime"`
	Incoterms                                     *string  `json:"Incoterms"`
	TransactionTaxClassification                  string   `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry         string   `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry       string   `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                      string   `json:"DefinedTaxClassification"`
	AccountAssignmentGroup                        string   `json:"AccountAssignmentGroup"`
	ProductAccountAssignmentGroup                 string   `json:"ProductAccountAssignmentGroup"`
	PaymentTerms                                  string   `json:"PaymentTerms"`
	DueCalculationBaseDate                        *string  `json:"DueCalculationBaseDate"`
	PaymentDueDate                                *string  `json:"PaymentDueDate"`
	NetPaymentDays                                *int     `json:"NetPaymentDays"`
	PaymentMethod                                 string   `json:"PaymentMethod"`
	Project                                       *string  `json:"Project"`
	AccountingExchangeRate                        *float32 `json:"AccountingExchangeRate"`
	ReferenceDocument                             *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem                         *int     `json:"ReferenceDocumentItem"`
	ItemCompleteDeliveryIsDefined                 *bool    `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryStatus                            *string  `json:"ItemDeliveryStatus"`
	IssuingStatus                                 *string  `json:"IssuingStatus"`
	ReceivingStatus                               *string  `json:"ReceivingStatus"`
	ItemBillingStatus                             *string  `json:"ItemBillingStatus"`
	TaxCode                                       *string  `json:"TaxCode"`
	TaxRate                                       *float32 `json:"TaxRate"`
	CountryOfOrigin                               *string  `json:"CountryOfOrigin"`
	CountryOfOriginLanguage                       *string  `json:"CountryOfOriginLanguage"`
	ItemBlockStatus                               *bool    `json:"ItemBlockStatus"`
	ItemDeliveryBlockStatus                       *bool    `json:"ItemDeliveryBlockStatus"`
	ItemBillingBlockStatus                        *bool    `json:"ItemBillingBlockStatus"`
	IsCancelled                                   *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                           *bool    `json:"IsMarkedForDeletion"`
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

type QuotationsItemPricingElement struct {
	Quotation                  int     `json:"Quotation"`
	QuotationItem              int     `json:"QuotationItem"`
	SupplyChainRelationshipID  int     `json:"SupplyChainRelationshipID"`
	Buyer                      int     `json:"Buyer"`
	Seller                     int     `json:"Seller"`
	PricingProcedureCounter    int     `json:"PricingProcedureCounter"`
	ConditionRecord            int     `json:"ConditionRecord"`
	ConditionSequentialNumber  int     `json:"ConditionSequentialNumber"`
	ConditionType              string  `json:"ConditionType"`
	PricingDate                string  `json:"PricingDate"`
	ConditionRateValue         float32 `json:"ConditionRateValue"`
	ConditionCurrency          string  `json:"ConditionCurrency"`
	ConditionQuantity          float32 `json:"ConditionQuantity"`
	ConditionQuantityUnit      string  `json:"ConditionQuantityUnit"`
	TaxCode                    string  `json:"TaxCode"`
	ConditionAmount            float32 `json:"ConditionAmount"`
	TransactionCurrency        string  `json:"TransactionCurrency"`
	ConditionIsManuallyChanged bool    `json:"ConditionIsManuallyChanged"`
}

type OrdersItemPricingElement struct {
	OrderID                    int      `json:"OrderID"`
	OrderItem                  int      `json:"OrderItem"`
	SupplyChainRelationshipID  int      `json:"SupplyChainRelationshipID"`
	Buyer                      int      `json:"Buyer"`
	Seller                     int      `json:"Seller"`
	PricingProcedureCounter    int      `json:"PricingProcedureCounter"`
	ConditionRecord            *int     `json:"ConditionRecord"`
	ConditionSequentialNumber  *int     `json:"ConditionSequentialNumber"`
	ConditionType              *string  `json:"ConditionType"`
	PricingDate                *string  `json:"PricingDate"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	ConditionCurrency          *string  `json:"ConditionCurrency"`
	ConditionQuantity          *float32 `json:"ConditionQuantity"`
	ConditionQuantityUnit      *string  `json:"ConditionQuantityUnit"`
	TaxCode                    *string  `json:"TaxCode"`
	ConditionAmount            *float32 `json:"ConditionAmount"`
	TransactionCurrency        *string  `json:"TransactionCurrency"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged"`
}
type OrdersItemScheduleLine struct {
	OrderID                                         int      `json:"OrderID"`
	OrderItem                                       int      `json:"OrderItem"`
	ScheduleLine                                    int      `json:"ScheduleLine"`
	SupplyChainRelationshipID                       int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipStockConfPlantID         int      `json:"SupplyChainRelationshipStockConfPlantID"`
	Product                                         string   `json:"Product"`
	StockConfirmationBussinessPartner               int      `json:"StockConfirmationBussinessPartner"`
	StockConfirmationPlant                          string   `json:"StockConfirmationPlant"`
	StockConfirmationPlantTimeZone                  *string  `json:"StockConfirmationPlantTimeZone"`
	StockConfirmationPlantBatch                     *string  `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate    *string  `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityEndDate      *string  `json:"StockConfirmationPlantBatchValidityEndDate"`
	RequestedDeliveryDate                           string   `json:"RequestedDeliveryDate"`
	ConfirmedDeliveryDate                           string   `json:"ConfirmedDeliveryDate"`
	OrderQuantityInBaseUnit                         float32  `json:"OrderQuantityInBaseUnit"`
	ConfirmedOrderQuantityByPDTAvailCheck           float32  `json:"ConfirmedOrderQuantityByPDTAvailCheck"`
	ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit float32  `json:"ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit"`
	DeliveredQuantityInBaseUnit                     *float32 `json:"DeliveredQuantityInBaseUnit"`
	UndeliveredQuantityInBaseUnit                   *float32 `json:"UndeliveredQuantityInBaseUnit"`
	OpenConfirmedQuantityInBaseUnit                 *float32 `json:"OpenConfirmedQuantityInBaseUnit"`
	StockIsFullyConfirmed                           *bool    `json:"StockIsFullyConfirmed"`
	PlusMinusFlag                                   string   `json:"PlusMinusFlag"`
	ItemScheduleLineDeliveryBlockStatus             *bool    `json:"ItemScheduleLineDeliveryBlockStatus"`
	IsCancelled                                     *bool    `json:"IsCancelled"`
	IsDeleted                                       *bool    `json:"IsDeleted"`
}

// Partner
type QuotationsPartner struct {
	Quotation               int    `json:"Quotation"`
	PartnerFunction         string `json:"PartnerFunction"`
	BusinessPartner         int    `json:"BusinessPartner"`
	BusinessPartnerFullName string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     string `json:"BusinessPartnerName"`
	Organization            string `json:"Organization"`
	Country                 string `json:"Country"`
	Language                string `json:"Language"`
	Currency                string `json:"Currency"`
	ExternalDocumentID      string `json:"ExternalDocumentID"`
	AddressID               int    `json:"AddressID"`
}

type OrdersPartner struct {
	OrderID                 int     `json:"OrderID"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
}

// Address
type QuotationsAddress struct {
	Quotation   int    `json:"Quotation"`
	AddressID   int    `json:"AddressID"`
	PostalCode  string `json:"PostalCode"`
	LocalRegion string `json:"LocalRegion"`
	Country     string `json:"Country"`
	District    string `json:"District"`
	StreetName  string `json:"StreetName"`
	CityName    string `json:"CityName"`
	Building    string `json:"Building"`
	Floor       int    `json:"Floor"`
	Room        int    `json:"Room"`
}

type OrdersAddress struct {
	OrderID     int     `json:"OrderID"`
	AddressID   int     `json:"AddressID"`
	PostalCode  *string `json:"PostalCode"`
	LocalRegion *string `json:"LocalRegion"`
	Country     *string `json:"Country"`
	District    *string `json:"District"`
	StreetName  *string `json:"StreetName"`
	CityName    *string `json:"CityName"`
	Building    *string `json:"Building"`
	Floor       *int    `json:"Floor"`
	Room        *int    `json:"Room"`
}
