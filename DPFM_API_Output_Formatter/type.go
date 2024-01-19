package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	Header      		*[]Header				`json:"Header"`
	Item        		*[]Item					`json:"Item"`
	ItemPricingElement	*[]ItemPricingElement	`json:"ItemPricingElement"`
	Address				*[]Address				`json:"Address"`
	Partner				*[]Partner				`json:"Partner"`
}

type Header struct {
	Contract                         int      `json:"Contract"`
	ContractDate                     string   `json:"ContractDate"`
	ContractType                     string   `json:"ContractType"`
	ContractStatus                   string   `json:"ContractStatus"`
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
	ContractValidityStartDate        *string  `json:"ContractValidityStartDate"`
	ContractValidityEndDate          *string  `json:"ContractValidityEndDate"`
	InvoicePeriodStartDate           *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate             *string  `json:"InvoicePeriodEndDate"`
	TotalNetAmount                   *float32 `json:"TotalNetAmount"`
	TotalTaxAmount                   *float32 `json:"TotalTaxAmount"`
	TotalGrossAmount                 *float32 `json:"TotalGrossAmount"`
	TransactionCurrency              *string  `json:"TransactionCurrency"`
	PricingDate                      *string  `json:"PricingDate"`
	PriceDetnExchangeRate            *float32 `json:"PriceDetnExchangeRate"`
	PaymentTerms                     *string  `json:"PaymentTerms"`
	PaymentMethod                    *string  `json:"PaymentMethod"`
	ReferenceDocument                *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem            *int     `json:"ReferenceDocumentItem"`
	AccountAssignmentGroup           *string  `json:"AccountAssignmentGroup"`
	AccountingExchangeRate           *float32 `json:"AccountingExchangeRate"`
	InvoiceDocumentDate              *string  `json:"InvoiceDocumentDate"`
	IsExportImport                   *bool    `json:"IsExportImport"`
	HeaderText                       *string  `json:"HeaderText"`
	HeaderBlockStatus                *bool    `json:"HeaderBlockStatus"`
	ExternalReferenceDocument        *string  `json:"ExternalReferenceDocument"`
	CertificateAuthorityChain        *string  `json:"CertificateAuthorityChain"`
	UsageControlChain        		 *string  `json:"UsageControlChain"`
	CreationDate                     string   `json:"CreationDate"`
	CreationTime                     string   `json:"CreationTime"`
	LastChangeDate                   string   `json:"LastChangeDate"`
	LastChangeTime                   string   `json:"LastChangeTime"`
	IsCancelled                      *bool    `json:"IsCancelled"`
	IsMarkedForDeletion              *bool    `json:"IsMarkedForDeletion"`
}

type Item struct {
	Contract                                 int      `json:"Contract"`
	ContractItem                             int      `json:"ContractItem"`
	ContractItemCategory                     string   `json:"ContractItemCategory"`
	ContractStatus                           string   `json:"ContractStatus"`
	SupplyChainRelationshipID                int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID        *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID   *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipProductionPlantID *int     `json:"SupplyChainRelationshipProductionPlantID"`
	Buyer                                    int      `json:"Buyer"`
	Seller                                   int      `json:"Seller"`
	DeliverToParty                           *int     `json:"DeliverToParty"`
	DeliverFromParty                         *int     `json:"DeliverFromParty"`
	DeliverToPlant                           *string  `json:"DeliverToPlant"`
	DeliverFromPlant                         *string  `json:"DeliverFromPlant"`
	ContractItemText                         string   `json:"ContractItemText"`
	Product                                  string   `json:"Product"`
	SizeOrDimensionText                      *string  `json:"SizeOrDimensionText"`
	ProductStandardID                        *string  `json:"ProductStandardID"`
	ProductGroup                             *string  `json:"ProductGroup"`
	ProductSpecification                     *string  `json:"ProductSpecification"`
	MarkingOfMaterial                        *string  `json:"MarkingOfMaterial"`
	BaseUnit                                 string   `json:"BaseUnit"`
	DeliveryUnit                             *string  `json:"DeliveryUnit"`
	ProductionVersion				 		 *int     `json:"ProductionVersion"`
	ProductionVersionItem			 		 *int     `json:"ProductionVersionItem"`
	BillOfMaterial                           *int     `json:"BillOfMaterial"`
	BillOfMaterialItem                       *int     `json:"BillOfMaterialItem"`
	PricingDate                              *string  `json:"PricingDate"`
	PriceDetnExchangeRate                    *float32 `json:"PriceDetnExchangeRate"`
	ServicesRenderingDate                    *string  `json:"ServicesRenderingDate"`
	OrderQuantityInBaseUnit                  *float32 `json:"OrderQuantityInBaseUnit"`
	OrderQuantityInDeliveryUnit              *float32 `json:"OrderQuantityInDeliveryUnit"`
	QuantityPerPackage                       *float32 `json:"QuantityPerPackage"`
	InternalCapacityQuantity                 *float32 `json:"InternalCapacityQuantity"`
	InternalCapacityQuantityUnit             *string  `json:"InternalCapacityQuantityUnit"`
	NetAmount                                float32  `json:"NetAmount"`
	TaxAmount                                float32  `json:"TaxAmount"`
	GrossAmount                              float32  `json:"GrossAmount"`
	InvoiceDocumentDate                      *string  `json:"InvoiceDocumentDate"`
	ProductionPlantBusinessPartner           *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                          *string  `json:"ProductionPlant"`
	InspectionPlantBusinessPartner           *int     `json:"InspectionPlantBusinessPartner"`
	InspectionPlant                          *string  `json:"InspectionPlant"`
	InspectionPlan                           *int     `json:"InspectionPlan"`
	InspectionLot                            *int     `json:"InspectionLot"`
	TransactionTaxClassification             string   `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry    string   `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry  string   `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                 string   `json:"DefinedTaxClassification"`
	AccountAssignmentGroup                   string   `json:"AccountAssignmentGroup"`
	ProductAccountAssignmentGroup            string   `json:"ProductAccountAssignmentGroup"`
	PaymentTerms                             string   `json:"PaymentTerms"`
	DueCalculationBaseDate                   *string  `json:"DueCalculationBaseDate"`
	PaymentDueDate                           *string  `json:"PaymentDueDate"`
	NetPaymentDays                           *int     `json:"NetPaymentDays"`
	PaymentMethod                            string   `json:"PaymentMethod"`
	Project                                  *int     `json:"Project"`
	WBSElement                               *int     `json:"WBSElement"`
	AccountingExchangeRate                   *float32 `json:"AccountingExchangeRate"`
	TaxCode                                  *string  `json:"TaxCode"`
	TaxRate                                  *float32 `json:"TaxRate"`
	CountryOfOrigin                          *string  `json:"CountryOfOrigin"`
	CountryOfOriginLanguage                  *string  `json:"CountryOfOriginLanguage"`
	Equipment                                *int     `json:"Equipment"`
	ItemBlockStatus                          *bool    `json:"ItemBlockStatus"`
	ExternalReferenceDocument                *string  `json:"ExternalReferenceDocument"`
	ExternalReferenceDocumentItem            *string  `json:"ExternalReferenceDocumentItem"`
	CreationDate                             string   `json:"CreationDate"`
	CreationTime                             string   `json:"CreationTime"`
	LastChangeDate                           string   `json:"LastChangeDate"`
	LastChangeTime                           string   `json:"LastChangeTime"`
	IsCancelled                              *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                      *bool    `json:"IsMarkedForDeletion"`
}

type ItemPricingElement struct {
	Contract                   int     `json:"Contract"`
	ContractItem               int     `json:"ContractItem"`
	PricingProcedureCounter    int     `json:"PricingProcedureCounter"`
	SupplyChainRelationshipID  int     `json:"SupplyChainRelationshipID"`
	Buyer                      int     `json:"Buyer"`
	Seller                     int     `json:"Seller"`
	ConditionRecord            int     `json:"ConditionRecord"`
	ConditionSequentialNumber  int     `json:"ConditionSequentialNumber"`
	ConditionType              string  `json:"ConditionType"`
	PricingDate                string  `json:"PricingDate"`
	ConditionRateValue         float32 `json:"ConditionRateValue"`
	ConditionRateValueUnit     int     `json:"ConditionRateValueUnit"`
	ConditionScaleQuantity     int     `json:"ConditionScaleQuantity"`
	ConditionCurrency          string  `json:"ConditionCurrency"`
	ConditionQuantity          float32 `json:"ConditionQuantity"`
	TaxCode                    *string `json:"TaxCode"`
	ConditionAmount            float32 `json:"ConditionAmount"`
	TransactionCurrency        string  `json:"TransactionCurrency"`
	ConditionIsManuallyChanged bool    `json:"ConditionIsManuallyChanged"`
	CreationDate               string  `json:"CreationDate"`
	CreationTime               string  `json:"CreationTime"`
	LastChangeDate             string  `json:"LastChangeDate"`
	LastChangeTime             string  `json:"LastChangeTime"`
	IsCancelled                *bool   `json:"IsCancelled"`
	IsMarkedForDeletion        *bool   `json:"IsMarkedForDeletion"`
}

type Address struct {
	Contract    int     `json:"Contract"`
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

type Partner struct {
	Contract                int     `json:"Contract"`
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
	EmailAddress            *string `json:"EmailAddress"`
}
