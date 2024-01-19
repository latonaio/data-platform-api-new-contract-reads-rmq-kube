package requests

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
