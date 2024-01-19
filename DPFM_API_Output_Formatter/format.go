package dpfm_api_output_formatter

import (
	"data-platform-api-new-contract-reads-rmq-kube/DPFM_API_Caller/requests"
	api_input_reader "data-platform-api-new-contract-reads-rmq-kube/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Header{}

		err := rows.Scan(
			&pm.Contract,
			&pm.ContractDate,
			&pm.ContractType,
			&pm.ContractStatus,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipBillingID,
			&pm.SupplyChainRelationshipPaymentID,
			&pm.Buyer,
			&pm.Seller,
			&pm.BillToParty,
			&pm.BillFromParty,
			&pm.BillToCountry,
			&pm.BillFromCountry,
			&pm.Payer,
			&pm.Payee,
			&pm.ContractValidityStartDate,
			&pm.ContractValidityEndDate,
			&pm.InvoicePeriodStartDate,
			&pm.InvoicePeriodEndDate,
			&pm.TotalNetAmount,
			&pm.TotalTaxAmount,
			&pm.TotalGrossAmount,
			&pm.TransactionCurrency,
			&pm.PricingDate,
			&pm.PriceDetnExchangeRate,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.AccountAssignmentGroup,
			&pm.AccountingExchangeRate,
			&pm.InvoiceDocumentDate,
			&pm.IsExportImport,
			&pm.HeaderText,
			&pm.HeaderBlockStatus,
			&pm.ExternalReferenceDocument,
			&pm.CertificateAuthorityChain,
			&pm.UsageControlChain,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &header, err
		}

		data := pm
		header = append(header, Header{
			Contract:                         data.Contract,
			ContractDate:                     data.ContractDate,
			ContractType:                     data.ContractType,
			ContractStatus:                   data.ContractStatus,
			SupplyChainRelationshipID:        data.SupplyChainRelationshipID,
			SupplyChainRelationshipBillingID: data.SupplyChainRelationshipBillingID,
			SupplyChainRelationshipPaymentID: data.SupplyChainRelationshipPaymentID,
			Buyer:                            data.Buyer,
			Seller:                           data.Seller,
			BillToParty:                      data.BillToParty,
			BillFromParty:                    data.BillFromParty,
			BillToCountry:                    data.BillToCountry,
			BillFromCountry:                  data.BillFromCountry,
			Payer:                            data.Payer,
			Payee:                            data.Payee,
			ContractValidityStartDate:        data.ContractValidityStartDate,
			ContractValidityEndDate:          data.ContractValidityEndDate,
			InvoicePeriodStartDate:           data.InvoicePeriodStartDate,
			InvoicePeriodEndDate:             data.InvoicePeriodEndDate,
			TotalNetAmount:                   data.TotalNetAmount,
			TotalTaxAmount:                   data.TotalTaxAmount,
			TotalGrossAmount:                 data.TotalGrossAmount,
			TransactionCurrency:              data.TransactionCurrency,
			PricingDate:                      data.PricingDate,
			PriceDetnExchangeRate:            data.PriceDetnExchangeRate,
			PaymentTerms:                     data.PaymentTerms,
			PaymentMethod:                    data.PaymentMethod,
			AccountAssignmentGroup:           data.AccountAssignmentGroup,
			AccountingExchangeRate:           data.AccountingExchangeRate,
			InvoiceDocumentDate:              data.InvoiceDocumentDate,
			IsExportImport:                   data.IsExportImport,
			HeaderText:                       data.HeaderText,
			HeaderBlockStatus:                data.HeaderBlockStatus,
			ExternalReferenceDocument:		  data.ExternalReferenceDocument,
			CertificateAuthorityChain:		  data.CertificateAuthorityChain,
			UsageControlChain:		  		  data.UsageControlChain,
			CreationDate:                     data.CreationDate,
			CreationTime:                     data.CreationTime,
			LastChangeDate:                   data.LastChangeDate,
			LastChangeTime:                   data.LastChangeTime,
			IsCancelled:                      data.IsCancelled,
			IsMarkedForDeletion:              data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}

func ConvertToPartner(rows *sql.Rows) (*[]Partner, error) {
	defer rows.Close()
	partner := make([]Partner, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Partner{}

		err := rows.Scan(
			&pm.Contract,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.Organization,
			&pm.Country,
			&pm.Language,
			&pm.Currency,
			&pm.ExternalDocumentID,
			&pm.AddressID,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &partner, err
		}

		data := pm
		partner = append(partner, Partner{
			Contract:                data.Contract,
			PartnerFunction:         data.PartnerFunction,
			BusinessPartner:         data.BusinessPartner,
			BusinessPartnerFullName: data.BusinessPartnerFullName,
			BusinessPartnerName:     data.BusinessPartnerName,
			Organization:            data.Organization,
			Country:                 data.Country,
			Language:                data.Language,
			Currency:                data.Currency,
			ExternalDocumentID:      data.ExternalDocumentID,
			AddressID:               data.AddressID,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &partner, nil
	}

	return &partner, nil
}

func ConvertToAddress(rows *sql.Rows) (*[]Address, error) {
	defer rows.Close()
	address := make([]Address, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Address{}

		err := rows.Scan(
			&pm.Contract,
			&pm.AddressID,
			&pm.PostalCode,
			&pm.LocalRegion,
			&pm.Country,
			&pm.District,
			&pm.StreetName,
			&pm.CityName,
			&pm.Building,
			&pm.Floor,
			&pm.Room,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &address, err
		}

		data := pm
		address = append(address, Address{
			Contract:    data.Contract,
			AddressID:   data.AddressID,
			PostalCode:  data.PostalCode,
			LocalRegion: data.LocalRegion,
			Country:     data.Country,
			District:    data.District,
			StreetName:  data.StreetName,
			CityName:    data.CityName,
			Building:    data.Building,
			Floor:       data.Floor,
			Room:        data.Room,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &address, nil
	}

	return &address, nil
}

func ConvertToItem(rows *sql.Rows) (*[]Item, error) {
	defer rows.Close()
	item := make([]Item, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Item{}

		err := rows.Scan(
			&pm.Contract,
			&pm.ContractItem,
			&pm.ContractItemCategory,
			&pm.ContractStatus,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipProductionPlantID,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.DeliverToPlant,
			&pm.DeliverFromPlant,
			&pm.ContractItemText,
			&pm.ContractItemTextByBuyer,
			&pm.ContractItemTextBySeller,
			&pm.Product,
			&pm.SizeOrDimensionText,
			&pm.ProductStandardID,
			&pm.ProductGroup,
			&pm.ProductSpecification,
			&pm.MarkingOfMaterial,
			&pm.BaseUnit,
			&pm.DeliveryUnit,
			&pm.ProductionVersion,
			&pm.ProductionVersionItem,
			&pm.BillOfMaterial,
			&pm.BillOfMaterialItem,
			&pm.PricingDate,
			&pm.PriceDetnExchangeRate,
			&pm.ServicesRenderingDate,
			&pm.ContractQuantityInBaseUnit,
			&pm.ContractQuantityInDeliveryUnit,
			&pm.QuantityPerPackage,
			&pm.InternalCapacityQuantity,
			&pm.InternalCapacityQuantityUnit,
			&pm.NetAmount,
			&pm.TaxAmount,
			&pm.GrossAmount,
			&pm.InvoiceDocumentDate,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.InspectionPlantBusinessPartner,
			&pm.InspectionPlant,
			&pm.InspectionPlan,
			&pm.InspectionLot,
			&pm.TransactionTaxClassification,
			&pm.ProductTaxClassificationBillToCountry,
			&pm.ProductTaxClassificationBillFromCountry,
			&pm.DefinedTaxClassification,
			&pm.AccountAssignmentGroup,
			&pm.ProductAccountAssignmentGroup,
			&pm.PaymentTerms,
			&pm.DueCalculationBaseDate,
			&pm.PaymentDueDate,
			&pm.NetPaymentDays,
			&pm.PaymentMethod,
			&pm.Project,
			&pm.WBSElement,
			&pm.AccountingExchangeRate,
			&pm.TaxCode,
			&pm.TaxRate,
			&pm.CountryOfOrigin,
			&pm.CountryOfOriginLanguage,
			&pm.Equipment,
			&pm.ItemBlockStatus,
			&pm.ExternalReferenceDocument,
			&pm.ExternalReferenceDocumentItem,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &item, nil
		}

		data := pm
		item = append(item, Item{
			Contract:                                      data.Contract,
			ContractItem:                                  data.ContractItem,
			ContractItemCategory:                          data.ContractItemCategory,
			ContractStatus:                      		   data.ContractStatus,
			SupplyChainRelationshipID:                     data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:             data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:        data.SupplyChainRelationshipDeliveryPlantID,
			SupplyChainRelationshipProductionPlantID:      data.SupplyChainRelationshipProductionPlantID,
			Buyer:                            			   data.Buyer,
			Seller:                           			   data.Seller,
			DeliverToParty:                                data.DeliverToParty,
			DeliverFromParty:                              data.DeliverFromParty,
			DeliverToPlant:                                data.DeliverToPlant,
			DeliverFromPlant:                              data.DeliverFromPlant,
			ContractItemText:                              data.ContractItemText,
			ContractItemTextByBuyer:                       data.ContractItemTextByBuyer,
			ContractItemTextBySeller:                      data.ContractItemTextBySeller,
			Product:                                       data.Product,
			SizeOrDimensionText:                           data.SizeOrDimensionText,
			ProductStandardID:                             data.ProductStandardID,
			ProductGroup:                                  data.ProductGroup,
			ProductSpecification:                          data.ProductSpecification,
			MarkingOfMaterial:                             data.MarkingOfMaterial,
			BaseUnit:                                      data.BaseUnit,
			DeliveryUnit:                                  data.DeliveryUnit,
			ProductionVersion:                             data.ProductionVersion,
			ProductionVersionItem:                         data.ProductionVersionItem,
			BillOfMaterial:                                data.BillOfMaterial,
			BillOfMaterialItem:                            data.BillOfMaterialItem,
			PricingDate:                                   data.PricingDate,
			PriceDetnExchangeRate:                         data.PriceDetnExchangeRate,
			ServicesRenderingDate:                         data.ServicesRenderingDate,
			ContractQuantityInBaseUnit:                    data.ContractQuantityInBaseUnit,
			ContractQuantityInDeliveryUnit:                data.ContractQuantityInDeliveryUnit,
			QuantityPerPackage:                            data.QuantityPerPackage,
			InternalCapacityQuantity:                      data.InternalCapacityQuantity,
			InternalCapacityQuantityUnit:                  data.InternalCapacityQuantityUnit,
			NetAmount:                                     data.NetAmount,
			TaxAmount:                                     data.TaxAmount,
			GrossAmount:                                   data.GrossAmount,
			InvoiceDocumentDate:                           data.InvoiceDocumentDate,
			ProductionPlantBusinessPartner:                data.ProductionPlantBusinessPartner,
			ProductionPlant:                               data.ProductionPlant,
			InspectionPlantBusinessPartner:                data.InspectionPlantBusinessPartner,
			InspectionPlant:                               data.InspectionPlant,
			InspectionPlan:                                data.InspectionPlan,
			InspectionLot:                           	   data.InspectionLot,
			TransactionTaxClassification:                  data.TransactionTaxClassification,
			ProductTaxClassificationBillToCountry:         data.ProductTaxClassificationBillToCountry,
			ProductTaxClassificationBillFromCountry:       data.ProductTaxClassificationBillFromCountry,
			DefinedTaxClassification:                      data.DefinedTaxClassification,
			AccountAssignmentGroup:                        data.AccountAssignmentGroup,
			ProductAccountAssignmentGroup:                 data.ProductAccountAssignmentGroup,
			PaymentTerms:                                  data.PaymentTerms,
			DueCalculationBaseDate:                        data.DueCalculationBaseDate,
			PaymentDueDate:                                data.PaymentDueDate,
			NetPaymentDays:                                data.NetPaymentDays,
			PaymentMethod:                                 data.PaymentMethod,
			Project:                                       data.Project,
			WBSElement:                                    data.WBSElement,
			AccountingExchangeRate:                        data.AccountingExchangeRate,
			TaxCode:                                       data.TaxCode,
			TaxRate:                                       data.TaxRate,
			CountryOfOrigin:                               data.CountryOfOrigin,
			CountryOfOriginLanguage:                       data.CountryOfOriginLanguage,
			Equipment:                                     data.Equipment,
			ItemBlockStatus:                               data.ItemBlockStatus,
			ExternalReferenceDocument:		  			   data.ExternalReferenceDocument,
			ExternalReferenceDocumentItem:	  			   data.ExternalReferenceDocumentItem,
			CreationDate:                                  data.CreationDate,
			CreationTime:                                  data.CreationTime,
			LastChangeDate:                                data.LastChangeDate,
			LastChangeTime:                                data.LastChangeTime,
			IsCancelled:                                   data.IsCancelled,
			IsMarkedForDeletion:                           data.IsMarkedForDeletion,
		})
	}

	return &item, nil
}

func ConvertToItemPricingElement(rows *sql.Rows) (*[]ItemPricingElement, error) {
	defer rows.Close()
	itemPricingElement := make([]ItemPricingElement, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ItemPricingElement{}

		err := rows.Scan(
			&pm.Contract,
			&pm.ContractItem,
			&pm.PricingProcedureCounter,
			&pm.SupplyChainRelationshipID,
			&pm.Buyer,
			&pm.Seller,
			&pm.ConditionRecord,
			&pm.ConditionSequentialNumber,
			&pm.ConditionType,
			&pm.PricingDate,
			&pm.ConditionRateValue,
			&pm.ConditionRateValueUnit,
			&pm.ConditionScaleQuantity,
			&pm.ConditionCurrency,
			&pm.ConditionQuantity,
			&pm.TaxCode,
			&pm.ConditionAmount,
			&pm.TransactionCurrency,
			&pm.ConditionIsManuallyChanged,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &itemPricingElement, err
		}

		data := pm
		itemPricingElement = append(itemPricingElement, ItemPricingElement{
			Contract:                   data.Contract,
			ContractItem:               data.ContractItem,
			PricingProcedureCounter:    data.PricingProcedureCounter,
			SupplyChainRelationshipID:  data.SupplyChainRelationshipID,
			Buyer:                      data.Buyer,
			Seller:                     data.Seller,
			ConditionRecord:            data.ConditionRecord,
			ConditionSequentialNumber:  data.ConditionSequentialNumber,
			ConditionType:              data.ConditionType,
			PricingDate:                data.PricingDate,
			ConditionRateValue:         data.ConditionRateValue,
			ConditionRateValueUnit:     data.ConditionRateValueUnit,
			ConditionScaleQuantity:     data.ConditionScaleQuantity,
			ConditionCurrency:          data.ConditionCurrency,
			ConditionQuantity:          data.ConditionQuantity,
			TaxCode:                    data.TaxCode,
			ConditionAmount:            data.ConditionAmount,
			TransactionCurrency:        data.TransactionCurrency,
			ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
			CreationDate:               data.CreationDate,
			CreationTime:               data.CreationTime,
			LastChangeDate:             data.LastChangeDate,
			LastChangeTime:             data.LastChangeTime,
			IsCancelled:                data.IsCancelled,
			IsMarkedForDeletion:        data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &itemPricingElement, nil
	}

	return &itemPricingElement, nil
}
