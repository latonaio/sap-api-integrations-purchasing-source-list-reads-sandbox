package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
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
		
func (c *SAPAPICaller) AsyncGetPurchasingSourceList(Material, Plant, Supplier, SupplyingPlant, VaridityEndDate string) {
	wg := &sync.WaitGroup{}

	wg.Add(3)
	go func() {
		c.PurchasingSourceList(Material, Plant, VaridityEndDate)
		wg.Done()
	}()
	go func() {
		c.PurchasingSourceListSupplier(Material, Plant, Supplier, VaridityEndDate)
		wg.Done()
	}()
	go func() {
		c.PurchasingSourceListSupplyingPlant(Material, Plant, SupplyingPlant, VaridityEndDate)
		wg.Done()
	}()
	wg.Wait()
}

func (c *SAPAPICaller) PurchasingSourceList(Material, Plant, VaridityEndDate string) {
	res, err := c.callPurchasingSourceListSrvAPIRequirementList("/A_PurchasingSource", Material, Plant, VaridityEndDate)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) PurchasingSourceListSupplier(Material, Plant, Supplier, VaridityEndDate string) {
	res, err := c.callPurchasingSourceListSrvAPIRequirementSupplier("/A_PurchasingSource", Material, Plant, Supplier, VaridityEndDate)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) PurchasingSourceListSupplyingPlant(Material, Plant, SupplyingPlant, VaridityEndDate string) {
	res, err := c.callPurchasingSourceListSrvAPIRequirementSupplyingPlant("/A_PurchasingSource", Material, Plant, SupplyingPlant, VaridityEndDate)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) callPurchasingSourceListSrvAPIRequirementList(api, Material, Plant, VaridityEndDate, string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASING_SOURCE_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "Material, Plant, VaridityEndDate")
	params.Add("$filter", fmt.Sprintf("Material eq '%s' and Plant eq '%s' and VaridityEndDate eq '%s'", Material, Plant, VaridityEndDate))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callPurchasingSourceListSrvAPIRequirementSupplier(api, Material, Plant, Supplier, VaridityEndDate string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASING_SOURCE_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "Material, Plant, Supplier, VaridityEndDate")
	params.Add("$filter", fmt.Sprintf("Material eq '%s' and Plant eq '%s' and Supplier eq '%s' and VaridityEndDate eq '%s'", Material, Plant, Supplier, VaridityEndDate))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callPurchasingSourceListSrvAPIRequirementSupplyingPlant(api, Material, Plant, SupplyingPlant, VaridityEndDate string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASING_SOURCE_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "Material, Plant, SupplyingPlant, VaridityEndDate")
	params.Add("$filter", fmt.Sprintf("Material eq '%s' and Plant eq '%s' and SupplyingPlant eq '%s' and VaridityEndDate eq '%s'", Material, Plant, SupplyingPlant, VaridityEndDate))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}