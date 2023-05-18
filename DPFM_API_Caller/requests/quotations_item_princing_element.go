package requests

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
