package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-purchasing-source-list-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetPurchasingSourceList(material, plant, sourceListRecord, supplier, supplyingPlant string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "List":
			func() {
				c.List(material, plant, sourceListRecord)
				wg.Done()
			}()
		case "Supplier":
			func() {
				c.Supplier(material, plant, sourceListRecord, supplier)
				wg.Done()
			}()
		case "SupplyingPlant":
			func() {
				c.SupplyingPlant(material, plant, sourceListRecord, supplyingPlant)
				wg.Done()
			}()

		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) List(material, plant, sourceListRecord string) {
	data, err := c.callPurchasingSourceListSrvAPIRequirementList("A_PurchasingSource", material, plant, sourceListRecord)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPurchasingSourceListSrvAPIRequirementList(api, material, plant, sourceListRecord string) (*sap_api_output_formatter.List, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASING_SOURCE_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithList(req, material, plant, sourceListRecord)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToList(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Supplier(material, plant, sourceListRecord, supplier string) {
	data, err := c.callPurchasingSourceListSrvAPIRequirementSupplier("A_PurchasingSource", material, plant, sourceListRecord, supplier)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPurchasingSourceListSrvAPIRequirementSupplier(api, material, plant, sourceListRecord, supplier string) (*sap_api_output_formatter.List, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASING_SOURCE_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithSupplier(req, material, plant, sourceListRecord, supplier)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToList(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) SupplyingPlant(material, plant, sourceListRecord, supplyingPlant string) {
	data, err := c.callPurchasingSourceListSrvAPIRequirementSupplyingPlant("A_PurchasingSource", material, plant, sourceListRecord, supplyingPlant)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPurchasingSourceListSrvAPIRequirementSupplyingPlant(api, material, plant, sourceListRecord, supplyingPlant string) (*sap_api_output_formatter.List, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASING_SOURCE_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithSupplyingPlant(req, material, plant, sourceListRecord, supplyingPlant)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToList(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithList(req *http.Request, material, plant, sourceListRecord string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Material eq '%s' and Plant eq '%s' and SourceListRecord eq '%s'", material, plant, sourceListRecord))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithSupplier(req *http.Request, material, plant, sourceListRecord, supplier string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Material eq '%s' and Plant eq '%s' and SourceListRecord eq '%s' and Supplier eq '%s'", material, plant, sourceListRecord, supplier))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithSupplyingPlant(req *http.Request, material, plant, sourceListRecord, supplyingPlant string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Material eq '%s' and Plant eq '%s' and SourceListRecord eq '%s' and SupplyingPlant eq '%s'", material, plant, sourceListRecord, supplyingPlant))
	req.URL.RawQuery = params.Encode()
}
