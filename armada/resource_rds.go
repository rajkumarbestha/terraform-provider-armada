package armada

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/hashicorp/terraform/helper/schema"
)

func armadaRds() *schema.Resource {
	return &schema.Resource {
		SchemaVersion: 1,
		Create:        createRDS,
		Read:          readRDS,
		Update:        updateRDS,
		Delete:        deleteRDS,
		Schema: map[string]*schema.Schema {
			"application_code": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"environment": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"request_type_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"request_from": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"database_engine": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"database_engine_verison": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"allocated_storage": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"db_instance_class": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"db_restore": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"db_restore_path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"database_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"purpose": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"request_raised_by": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"request_created_date": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"service_request_id_out": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func createRDS(d *schema.ResourceData, meta interface{}) error {
	// Logging.
	file, err := os.OpenFile("logs/rds_create.log", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	log.Printf("Terraform apply command given and control is in resource create function. \r\n")

	// Make a request object.
	rdsCreateReq := make(map[string]interface{})
	rdsCreateReq["applicationCode"] = d.Get("application_code").(string)
	rdsCreateReq["databaseEngine"] = d.Get("database_engine").(string)
	rdsCreateReq["environment"] = d.Get("environment").(string)
	rdsCreateReq["requestFrom"] = d.Get("request_from").(string)
	rdsCreateReq["requestTypeID"] = d.Get("request_type_id").(string)
	rdsCreateReq["requestRaisedBy"] = d.Get("request_raised_by").(string)
	rdsCreateReq["databaseEngineVerison"] = d.Get("database_engine_verison").(string)
	rdsCreateReq["dbRestore"] = d.Get("db_restore").(string)
	rdsCreateReq["dbRestorePath"] = d.Get("db_restore_path").(string)
	var db_array[1] map[string]interface{}
	db_map := make(map[string]interface{})
	db_map["key"] = d.Get("database_name").(string)
	db_array[0] = db_map
	rdsCreateReq["databaseName"] = db_array
	rdsCreateReq["allocatedStorage"] = d.Get("allocated_storage").(string)
	rdsCreateReq["purpose"] = d.Get("purpose").(string)
	rdsCreateReq["dbInstanceClass"] = d.Get("db_instance_class").(string)
	rdsCreateReq["requestCreatedDate"] = d.Get("request_created_date").(string)

	// Convert to JSON.
	rdsCreateJson, err := json.Marshal(rdsCreateReq)
	log.Printf("Json Obtained %s", string(rdsCreateJson))

	endpoint := "https://dv2-api.dbaas.aenetworks.com/aws/rds/create"
	endpoint_from_tf := d.Get("endpoint").(string)
	if len(endpoint_from_tf) > 0 {
		endpoint = endpoint_from_tf
	}
	log.Printf("endpoint %s", endpoint)

	// Create Request
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(rdsCreateJson))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	log.Printf("Request Created")

	// AWS IAM AUTH.
	awsAuth := v4.NewSigner(credentials.NewStaticCredentials(
		meta.(*dbaasCreds).access_key,
		meta.(*dbaasCreds).secret_key,
		""))
	awsAuth.Sign(req, bytes.NewReader(rdsCreateJson), "execute-api", "us-east-1", time.Now())
	log.Printf("Signed.")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Printf("Response Received %s", bodyString)

	var data CreateResponse

	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if data.StatusCode == 200 {
		log.Printf("Checing Decode JSON. %v %d", data.Message, data.ServiceRequesID)

		// Setting the ResourceID.
		d.SetId(strconv.Itoa(data.ServiceRequesID))

		// Setting the RequestID Out.
		d.Set("service_request_id_out", strconv.Itoa(data.ServiceRequesID))
		return nil
	}
	return errors.New(data.Message)
	return nil
}

func readRDS(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func updateRDS(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteRDS(d *schema.ResourceData, meta interface{}) error {
	// Logging.
	file, err := os.OpenFile("logs/rds_destroy.log", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	// Create Request.
	endpoint := fmt.Sprintf("https://dv2-api.dbaas.aenetworks.com/aws/rds/delete?serviceRequestId=%s", d.Id())
	req, err := http.NewRequest("DELETE", endpoint, nil)
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	log.Printf("Request Created %s", endpoint)

	// AWS IAM AUTH.
	awsAuth := v4.NewSigner(credentials.NewStaticCredentials(
		meta.(*dbaasCreds).access_key,
		meta.(*dbaasCreds).secret_key,
		""))
	awsAuth.Sign(req, nil, "execute-api", "us-east-1", time.Now())
	log.Printf("Signed.")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Printf("Response Received %s", bodyString)

	var data DeleteResponse

	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if data.StatusCode == 200 {
		log.Printf("Checking Decode JSON. %v", data.Message)
		log.Printf("Service Request ID to delete: %s", d.Id())

		// Deleting the resource.
		d.SetId(" ")
		return nil
	}
	return errors.New(data.Message)
}