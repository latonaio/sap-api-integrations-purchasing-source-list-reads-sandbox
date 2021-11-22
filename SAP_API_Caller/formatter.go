package sap_api_caller

type PurchasingSourceListReads struct {
	ConnectionKey        string `json:"connection_key"`
	Result               bool   `json:"result"`
	RedisKey             string `json:"redis_key"`
	Filepath             string `json:"filepath"`
	APISchema            string `json:"api_schema"`
	Material             string `json:"material_code"`
	Plant                string `json:"plant"`
	SourceListRecord     int    `json:"source_list_record"`
	Deleted              string `json:"deleted"`
}

PurchasingSourceList struct {
	Material                   string   `json:"Material"`
	Plant                      string   `json:"Plant"`
	SourceListRecord           int      `json:"SourceListRecord"`
	ValidityStartDate          string   `json:"ValidityStartDate"`
	ValidityEndDate            string   `json:"ValidityEndDate"`
	Supplier                   string   `json:"Supplier"`
	PurchasingOrganization     string   `json:"PurchasingOrganization"`
	SupplyingPlant             string   `json:"SupplyingPlant"`
	OrderQuantityUnit          string   `json:"OrderQuantityUnit"`
	PurchaseOutlineAgreement   string   `json:"PurchaseOutlineAgreement"`
	SupplierIsFixed            string   `json:"SupplierIsFixed"`
	SourceOfSupplyIsBlocked    string   `json:"SourceOfSupplyIsBlocked"`
	MRPSourcingControl         string   `json:"MRPSourcingControl"`
	LastChangeDateTime         string   `json:"LastChangeDateTime"`
	IssgPlantIsFixed           string   `json:"IssgPlantIsFixed"`
	PurOutlineAgreementIsFixed string   `json:"PurOutlineAgreementIsFixed"`
	SourceOfSupplyIsFixed      string   `json:"SourceOfSupplyIsFixed"`
}