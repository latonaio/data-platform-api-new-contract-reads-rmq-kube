package requests

type Item struct {
	Contract                                 int		`json:"Contract"`
	ContractItem                             int		`json:"ContractItem"`
	ContractItemCategory                     string		`json:"ContractItemCategory"`
	ContractStatus                           string		`json:"ContractStatus"`
	SupplyChainRelationshipID                int		`json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID        *int		`json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID   *int		`json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipProductionPlantID *int		`json:"SupplyChainRelationshipProductionPlantID"`
	Buyer                                    int		`json:"Buyer"`
	Seller                                   int		`json:"Seller"`
	DeliverToParty                           *int		`json:"DeliverToParty"`
	DeliverFromParty                         *int		`json:"DeliverFromParty"`
	DeliverToPlant                           *string	`json:"DeliverToPlant"`
	DeliverFromPlant                         *string	`json:"DeliverFromPlant"`
	ContractItemText                         string 	`json:"ContractItemText"`
	Product                                  string 	`json:"Product"`
	SizeOrDimensionText                      *string	`json:"SizeOrDimensionText"`
	ProductStandardID                        *string 	`json:"ProductStandardID"`
	ProductGroup                             *string  	`json:"ProductGroup"`
	ProductSpecification                     *string    `json:"ProductSpecification"`
	MarkingOfMaterial                        *string    `json:"MarkingOfMaterial"`
	BaseUnit                                 string		`json:"BaseUnit"`
	DeliveryUnit                             *string	`json:"DeliveryUnit"`
	ProductionVersion				 		 *int       `json:"ProductionVersion"`
	ProductionVersionItem			 		 *int       `json:"ProductionVersionItem"`
	BillOfMaterial                           *int		`json:"BillOfMaterial"`
	BillOfMaterialItem                       *int		`json:"BillOfMaterialItem"`
	PricingDate                              *string	`json:"PricingDate"`
	PriceDetnExchangeRate                    *float32	`json:"PriceDetnExchangeRate"`
	ServicesRenderingDate                    *string	`json:"ServicesRenderingDate"`
	OrderQuantityInBaseUnit                  *float32	`json:"OrderQuantityInBaseUnit"`
	OrderQuantityInDeliveryUnit              *float32	`json:"OrderQuantityInDeliveryUnit"`
	QuantityPerPackage                       *float32	`json:"QuantityPerPackage"`
	InternalCapacityQuantity                 *float32	`json:"InternalCapacityQuantity"`
	InternalCapacityQuantityUnit             *string	`json:"InternalCapacityQuantityUnit"`
	NetAmount                                float32	`json:"NetAmount"`
	TaxAmount                                float32	`json:"TaxAmount"`
	GrossAmount                              float32	`json:"GrossAmount"`
	InvoiceDocumentDate                      *string	`json:"InvoiceDocumentDate"`
	ProductionPlantBusinessPartner           *int		`json:"ProductionPlantBusinessPartner"`
	ProductionPlant                          *string	`json:"ProductionPlant"`
	InspectionPlantBusinessPartner           *int     	`json:"InspectionPlantBusinessPartner"`
	InspectionPlant                          *string  	`json:"InspectionPlant"`
	InspectionPlan                           *int     	`json:"InspectionPlan"`
	InspectionLot                            *int     	`json:"InspectionLot"`
	TransactionTaxClassification             string		`json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry    string		`json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry  string		`json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                 string		`json:"DefinedTaxClassification"`
	AccountAssignmentGroup                   string		`json:"AccountAssignmentGroup"`
	ProductAccountAssignmentGroup            string		`json:"ProductAccountAssignmentGroup"`
	PaymentTerms                             string		`json:"PaymentTerms"`
	DueCalculationBaseDate                   *string	`json:"DueCalculationBaseDate"`
	PaymentDueDate                           *string	`json:"PaymentDueDate"`
	NetPaymentDays                           *int		`json:"NetPaymentDays"`
	PaymentMethod                            string		`json:"PaymentMethod"`
	Project                                  *int		`json:"Project"`
	WBSElement                               *int		`json:"WBSElement"`
	AccountingExchangeRate                   *float32	`json:"AccountingExchangeRate"`
	TaxCode                                  *string	`json:"TaxCode"`
	TaxRate                                  *float32	`json:"TaxRate"`
	CountryOfOrigin                          *string	`json:"CountryOfOrigin"`
	CountryOfOriginLanguage                  *string	`json:"CountryOfOriginLanguage"`
	Equipment                                *int		`json:"Equipment"`
	ItemBlockStatus                          *bool		`json:"ItemBlockStatus"`
	ExternalReferenceDocument                *string	`json:"ExternalReferenceDocument"`
	ExternalReferenceDocumentItem            *string	`json:"ExternalReferenceDocumentItem"`
	CreationDate                             string		`json:"CreationDate"`
	CreationTime                             string		`json:"CreationTime"`
	LastChangeDate                           string		`json:"LastChangeDate"`
	LastChangeTime                           string		`json:"LastChangeTime"`
	IsCancelled                              *bool		`json:"IsCancelled"`
	IsMarkedForDeletion                      *bool		`json:"IsMarkedForDeletion"`
}
