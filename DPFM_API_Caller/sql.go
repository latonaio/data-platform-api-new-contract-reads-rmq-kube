package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-new-contract-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-new-contract-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *[]dpfm_api_output_formatter.Header
	var item *[]dpfm_api_output_formatter.Item
	var itemPricingElement *[]dpfm_api_output_formatter.ItemPricingElement
	var address *[]dpfm_api_output_formatter.Address
	var partner *[]dpfm_api_output_formatter.Partner

	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				header = c.Header(mtx, input, output, errs, log)
			}()
			//		case "Headers":
			//			func() {
			//				header = c.Headers(mtx, input, output, errs, log)
			//			}()
		case "HeadersByBuyer":
			func() {
				header = c.HeadersByBuyer(mtx, input, output, errs, log)
			}()
		case "HeadersBySeller":
			func() {
				header = c.HeadersBySeller(mtx, input, output, errs, log)
			}()
		case "Item":
			func() {
				item = c.Item(mtx, input, output, errs, log)
			}()
		case "Items":
			func() {
				item = c.Items(mtx, input, output, errs, log)
			}()
		case "ItemPricingElement":
			func() {
				itemPricingElement = c.ItemPricingElement(mtx, input, output, errs, log)
			}()
		case "ItemPricingElements":
			func() {
				itemPricingElement = c.ItemPricingElements(mtx, input, output, errs, log)
			}()
		case "Address":
			func() {
				address = c.Address(mtx, input, output, errs, log)
			}()
		case "Partner":
			func() {
				partner = c.Partner(mtx, input, output, errs, log)
			}()

		default:
		}
		if len(*errs) != 0 {
			break
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:             header,
		Item:               item,
		ItemPricingElement: itemPricingElement,
		Address:            address,
		Partner:            partner,
	}

	return data
}

func (c *DPFMAPICaller) Header(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := fmt.Sprintf("WHERE header.Contract = %d ", input.Header.Contract)

	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND header.IsCancelled = %v ", where, *input.Header.IsCancelled)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND header.IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_contract_header_data AS header
		` + where + ` ORDER BY header.IsMarkedForDeletion ASC, header.IsCancelled ASC, header.Contract DESC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersByBuyer(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := "WHERE 1 = 1"
	if input.Header.Buyer != nil {
		where = fmt.Sprintf("%s\nAND Buyer = %v", where, *input.Header.Buyer)
	}
	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND IsCancelled = %t", where, *input.Header.IsCancelled)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %t", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_contract_header_data
		` + where + `;`,
	)

	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersBySeller(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := "WHERE 1 = 1"
	if input.Header.Seller != nil {
		where = fmt.Sprintf("%s\nAND Seller = %v", where, *input.Header.Seller)
	}
	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND IsCancelled = %t", where, *input.Header.IsCancelled)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %t", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_contract_header_data
		` + where + `;`,
	)

	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Item(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	var args []interface{}
	where := fmt.Sprintf("WHERE Contract = %d ", input.Header.Contract)

	itemIDs := ""
	for _, v := range input.Header.Item {
		itemIDs = fmt.Sprintf("%s, %d", itemIDs, v.ContractItem)
	}

	if len(itemIDs) != 0 {
		where = fmt.Sprintf("%s\nAND ContractItem IN ( %s ) ", where, itemIDs[1:])
	}
	rows, err := c.db.Query(
		`SELECT 
		Contract, ContractItem, ContractItemCategory, SupplyChainRelationshipID, SupplyChainRelationshipDeliveryID,
		SupplyChainRelationshipDeliveryPlantID, SupplyChainRelationshipProductionPlantID, 
		ContractItemText, ContractItemTextByBuyer, ContractItemTextBySeller, Product, ProductStandardID, ProductGroup, BaseUnit, PricingDate, 
		PriceDetnExchangeRate, DeliverToParty, DeliverFromParty, CreationDate, CreationTime, 
		LastChangeDate, LastChangeTime, DeliverToPlant, DeliverFromPlant, DeliveryUnit, 
		ServicesRenderingDate, 
		ContractQuantityInBaseUnit, ContractQuantityInDeliveryUnit, QuantityPerPackage, 
		InternalCapacityQuantity, InternalCapacityQuantityUnit, NetAmount, TaxAmount, GrossAmount, InvoiceDocumentDate,
		ProductionPlantBusinessPartner, ProductionPlant,
		InspectionPlan, InspectionPlant, InspectionOrder, 
		TransactionTaxClassification, ProductTaxClassificationBillToCountry, ProductTaxClassificationBillFromCountry, 
		DefinedTaxClassification, AccountAssignmentGroup, ProductAccountAssignmentGroup, PaymentTerms, DueCalculationBaseDate,
		PaymentDueDate, NetPaymentDays, PaymentMethod, Project, AccountingExchangeRate,
		TaxCode, TaxRate, 
		CountryOfOrigin, CountryOfOriginLanguage, ItemBlockStatus, 
		IsCancelled, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_contract_item_data
		`+where+` ORDER BY IsMarkedForDeletion ASC, IsCancelled ASC, Contract DESC, ContractItem ASC ;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItem(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
func (c *DPFMAPICaller) Items(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	item := &dpfm_api_input_reader.Item{}
	if len(input.Header.Item) > 0 {
		item = &input.Header.Item[0]
	}
	where := "WHERE 1 = 1"

	if item != nil {
		if item.IsCancelled != nil {
			where = fmt.Sprintf("%s\nAND item.IsCancelled = %v", where, *item.IsCancelled)
		}
		if item.IsMarkedForDeletion != nil {
			where = fmt.Sprintf("%s\nAND item.IsMarkedForDeletion = %v", where, *item.IsMarkedForDeletion)
		}
	}

	rows, err := c.db.Query(
		`SELECT 
			*
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_contract_item_data as item
		` + where + ` ORDER BY item.IsMarkedForDeletion ASC, item.IsCancelled ASC, item.Contract DESC, item.ContractItem ASC ;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItem(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemPricingElement(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemPricingElement {
	var args []interface{}
	contract := input.Header.Contract
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		args = append(args, contract, v.ContractItem)
		cnt++
	}
	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_contract_item_pricing_element_data
		WHERE (Contract, ContractItem) IN ( `+repeat+` )
		ORDER BY Contract DESC, ContractItem DESC, PricingProcedureCounter DESC ;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemPricingElement(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemPricingElements(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemPricingElement {
	where := fmt.Sprintf("WHERE Contract = %d", input.Header.Contract)
	if input.Header.Buyer != nil && *input.Header.Buyer != 0 {
		where = fmt.Sprintf("%s\nAND Buyer = %d", where, *input.Header.Buyer)
	}
	if input.Header.Seller != nil && *input.Header.Seller != 0 {
		where = fmt.Sprintf("%s\nAND Seller = %d", where, *input.Header.Seller)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_contarct_item_pricing_element_data
		` + where + ` ORDER BY Contract DESC, ContractItem DESC, PricingProcedureCounter DESC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemPricingElement(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Address(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Address {
	var args []interface{}
	contract := input.Header.Contract
	address := input.Header.Address

	cnt := 0
	for _, v := range address {
		args = append(args, contract, v.AddressID)
		cnt++
	}
	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_contract_address_data
		WHERE (Contract, AddressID) IN ( `+repeat+` ) 
		ORDER BY Contract DESC, AddressID DESC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToAddress(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Partner(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Partner {
	var args []interface{}
	contract := input.Header.Contract
	partner := input.Header.Partner

	cnt := 0
	for _, v := range partner {
		args = append(args, contract, v.PartnerFunction, v.BusinessPartner)
		cnt++
	}
	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_contract_partner_data
		WHERE (Contract, PartnerFunction, BusinessPartner) IN ( `+repeat+` ) 
		ORDER BY Contract DESC, BusinessPartner DESC, AddressID DESC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToPartner(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
