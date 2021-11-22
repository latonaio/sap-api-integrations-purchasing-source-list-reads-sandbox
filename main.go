package main

import (
	sap_api_caller "sap-api-integrations-purchasing-source-list-reads/SAP_API_Caller"
	"sap-api-integrations-purchasing-source-list-reads/file_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := file_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs//SDC_Purchasing_Source_List_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	caller.AsyncGetPurchasingSourceList(
		inoutSDC.Material.Plant.VaridityEndDate,
		inoutSDC.Material.Plant.Supplier.VaridityEndDate,
		inoutSDC.Material.Plant.SupplyingPlant.VaridityEndDate,
	)
}