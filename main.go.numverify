package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
	"fmt"
	"log"
	"os"
	"encoding/json"
	"net/http"
)

var country string

type Numverify struct {
	Valid               bool   `json:"valid"`
	Number              string `json:"number"`
	LocalFormat         string `json:"local_format"`
	InternationalFormat string `json:"international_format"`
	CountryPrefix       string `json:"country_prefix"`
	CountryCode         string `json:"country_code"`
	CountryName         string `json:"country_name"`
	Location            string `json:"location"`
	Carrier             string `json:"carrier"`
	LineType            string `json:"line_type"`
}

func main() {
	opts := plugin.ServeOpts{
		ProviderFunc: Provider,
	}
	plugin.Serve(&opts)
}

func Provider() terraform.ResourceProvider {
	return &schema.Provider{ // Source https://github.com/hashicorp/terraform/blob/v0.6.6/helper/schema/provider.go#L20-L43
		Schema:        providerSchema(),
		ResourcesMap:  providerResources(),
	}
}

// List of supported configuration fields for your provider.
// Here we define a linked list of all the fields that we want to
// support in our provider (api_key, endpoint, timeout & max_retries).
// More info in https://github.com/hashicorp/terraform/blob/v0.6.6/helper/schema/schema.go#L29-L142
func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
	}
}

// List of supported resources and their configuration fields.
// Here we define da linked list of all the resources that we want to
// support in our provider. As an example, if you were to write an AWS provider
// which supported resources like ec2 instances, elastic balancers and things of that sort
// then this would be the place to declare them.
// More info here https://github.com/hashicorp/terraform/blob/v0.6.6/helper/schema/resource.go#L17-L81
func providerResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"numverify_API": &schema.Resource{
			SchemaVersion: 1,
			Create:        createFunc,
			Read:          readFunc,
			Update:        updateFunc,
			Delete:        deleteFunc,
			Schema: map[string]*schema.Schema{ // List of supported configuration fields for your resource
				"mobile_number": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"access_key": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"country": &schema.Schema{
					Type: schema.TypeString,
					Computed: true,
				},
			},
		},
	}
}


// The methods defined below will get called for each resource that needs to
// get created (createFunc), read (readFunc), updated (updateFunc) and deleted (deleteFunc).
// For example, if 10 resources need to be created then `createFunc`
// will get called 10 times every time with the information for the proper
// resource that is being mapped.
//
// If at some point any of these functions returns an error, Terraform will
// imply that something went wrong with the modification of the resource and it
// will prevent the execution of further calls that depend on that resource
// that failed to be created/updated/deleted.

func createFunc(d *schema.ResourceData, meta interface{}) error {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND, 0666)
	 if err != nil {
		log.Fatal(err)
	 }
	defer file.Close()
	log.SetOutput(file)
	log.Printf("Terraform apply command given and control is in resource create function \r\n")
	mobile_number := d.Get("mobile_number").(string)
	access_key := d.Get("access_key").(string)

	log.Printf("Mobile number from TF: %s \r\n", mobile_number)
	log.Printf("Access Key from TF: %s \r\n", access_key)
	// QueryEscape escapes the phone string so
	// it can be safely placed inside a URL query
	//safePhone := url.QueryEscape(phone)

	url := fmt.Sprintf("http://apilayer.net/api/validate?access_key=%s&number=%s", access_key, mobile_number)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		//return
	}
	log.Printf("Num Verify request created \r\n")
	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	log.Printf("Hitting Num Verify API : %s \r\n", url)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		//return
	}
	log.Printf("Response received from Num Verify API. \r\n")

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record Numverify

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	log.Printf("Printing the Response received. \r\n")

	log.Printf("Mobile Number = %s \r\n", record.InternationalFormat)
	log.Printf("Country   = %s \r\n", record.CountryName)
	log.Printf("Location  = %s \r\n", record.Location)
	log.Printf("Carrier   = %s \r\n", record.Carrier)
	log.Printf("LineType  = %s \r\n", record.LineType)

	country = record.CountryName
	d.SetId("country")
	//return nil;
	return readFunc(d, meta)
}

func readFunc(d *schema.ResourceData, meta interface{}) error {
	//hi := "heyy"
	d.Set("country", country)
	return nil
}

func updateFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}
