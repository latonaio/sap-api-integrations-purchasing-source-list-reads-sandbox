package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-purchasing-source-list-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
)

func ConvertToPurchasingSourceList(raw []byte, l *logger.Logger) *PurchasingSourceList {
	pm := &responses.PurchasingSourceList{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		l.Error(err)
		return nil
	}
	if len(pm.D.Results) == 0 {
		l.Error("Result data is not exist.")
		return nil
	}
	if len(pm.D.Results) > 1 {
		l.Error("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &PurchasingSourceList{
		Material                   data.Material,
		Plant                      data.Plant,
		SourceListRecord           data.SourceListRecord,
		ValidityStartDate          data.ValidityStartDate,
		ValidityEndDate            data.ValidityEndDate,
		Supplier                   data.Supplier,
		PurchasingOrganization     data.PurchasingOrganization,
		SupplyingPlant             data.SupplyingPlant,
		OrderQuantityUnit          data.OrderQuantityUnit,
		PurchaseOutlineAgreement   data.PurchaseOutlineAgreement,
		SupplierIsFixed            data.SupplierIsFixed,
		SourceOfSupplyIsBlocked    data.SourceOfSupplyIsBlocked,
		MRPSourcingControl         data.MRPSourcingControl,
		LastChangeDateTime         data.LastChangeDateTime,
		IssgPlantIsFixed           data.IssgPlantIsFixed,
		PurOutlineAgreementIsFixed data.PurOutlineAgreementIsFixed,
		SourceOfSupplyIsFixed      data.SourceOfSupplyIsFixed,
	}
}
