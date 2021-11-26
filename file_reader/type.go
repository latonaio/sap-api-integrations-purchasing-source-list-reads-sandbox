package file_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo      string      `json:"document_no"`
		DeliverTo       string      `json:"deliver_to"`
		Quantity        float64     `json:"quantity"`
		PickedQuantity  float64     `json:"picked_quantity"`
		Price           float64     `json:"price"`
	    Batch           string      `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string      `json:"document_no"`
		Status               string      `json:"status"`
		DeliverTo            string      `json:"deliver_to"`
		Quantity             float64     `json:"quantity"`
		CompletedQuantity    float64     `json:"completed_quantity"`
	    PlannedStartDate     string      `json:"planned_start_date"`
	    PlannedValidatedDate string      `json:"planned_validated_date"`
	    ActualStartDate      string      `json:"actual_start_date"`
	    ActualValidatedDate  string      `json:"actual_validated_date"`
	    Batch                string      `json:"batch"`
		Work              struct {
			WorkNo                   string      `json:"work_no"`
			Quantity                 float64     `json:"quantity"`
			CompletedQuantity        float64     `json:"completed_quantity"`
			ErroredQuantity          float64     `json:"errored_quantity"`
			Component                string      `json:"component"`
			PlannedComponentQuantity float64     `json:"planned_component_quantity"`
			PlannedStartDate         string      `json:"planned_start_date"`
			PlannedStartTime         string      `json:"planned_start_time"`
			PlannedValidatedDate     string      `json:"planned_validated_date"`
			PlannedValidatedTime     string      `json:"planned_validated_time"`
			ActualStartDate          string      `json:"actual_start_date"`
			ActualStartTime          string      `json:"actual_start_time"`
			ActualValidatedDate      string      `json:"actual_validated_date"`
			ActualValidatedTime      string      `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema     string `json:"api_schema"`
	MaterialCode  string `json:"material_code"`
	Plant         string `json:"plant/supplier"`
	Stock         float64 `json:"stock"`
	DocumentType  string `json:"document_type"`
	DocumentNo    string `json:"document_no"`
	PlannedDate   string `json:"planned_date"`
	ValidatedDate string `json:"validated_date"`
	Deleted       string `json:"deleted"`
}

type SDC struct {
	ConnectionKey        string `json:"connection_key"`
	Result               bool   `json:"result"`
	RedisKey             string `json:"redis_key"`
	Filepath             string `json:"filepath"`
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
	} `json:"PurchasingSourceList"`
	APISchema        string `json:"api_schema"`
	MaterialCode     string `json:"material_code"`
	Plant            string `json:"plant"`
	SourceListRecord int    `json:"source_list_record"`
	Deleted          string `json:"deleted"`
}
