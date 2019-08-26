package main

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

func armadaEc2() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createEC2,
		Read:          readEC2,
		Update:        updateEC2,
		Delete:        deleteEC2,
		Schema: map[string]*schema.Schema{
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
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"source_server_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"destination_server": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"source_database_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"destination_database_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"service_account": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"dbs_dbnameto_restore": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"edrive_size": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"service_request_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"purpose": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"request_raised_by": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"db_instance_type": &schema.Schema{
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
		}, // Resource Schema End.
	} // Resource End.
}

func createEC2(d *schema.ResourceData, meta interface{}) error {
	// Logging.
	file, err := os.OpenFile("logs/ec2_create.log", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	log.Printf("Terraform apply command given and control is in resource create function. \r\n")

	// Make a request object.
	ec2CreateReq := make(map[string]interface{})
	ec2CreateReq["version"] = d.Get("version").(string)
	ec2CreateReq["environment"] = d.Get("environment").(string)
	ec2CreateReq["dbInstancestype"] = d.Get("db_instance_type").(string)
	ec2CreateReq["applicationCode"] = d.Get("application_code").(string)
	ec2CreateReq["requestFrom"] = d.Get("request_from").(string)
	ec2CreateReq["requestTypeID"] = d.Get("request_type_id").(string)
	ec2CreateReq["requestRaisedBy"] = d.Get("request_raised_by").(string)
	ec2CreateReq["sourceServerName"] = d.Get("source_server_name").(string)
	ec2CreateReq["sourceDatabaseName"] = d.Get("source_database_name").(string)
	ec2CreateReq["destinationServer"] = d.Get("destination_server").(string)
	ec2CreateReq["destinationDatabaseName"] = d.Get("destination_database_name").(string)
	ec2CreateReq["eDriveSize"] = d.Get("edrive_size").(string)
	ec2CreateReq["serviceAccount"] = d.Get("service_account").(string)
	ec2CreateReq["purpose"] = d.Get("purpose").(string)
	ec2CreateReq["dbsDbnameToRestore"] = d.Get("dbs_dbnameto_restore").(string)
	ec2CreateReq["serverType"] = d.Get("server_type").(string)

	// Convert to JSON.
	ec2CreateJson, err := json.Marshal(ec2CreateReq)
	log.Printf("Json Obtained %s", string(ec2CreateJson))

	endpoint := "https://dv2-api.dbaas.aenetworks.com/aws/ec2/create"
	endpoint_from_tf := d.Get("endpoint").(string)
	if len(endpoint_from_tf) > 0 {
		endpoint = endpoint_from_tf
	}
	log.Printf("endpoint %s", endpoint)
	
	// Create Request.
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(ec2CreateJson))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	log.Printf("Request Created")

	// AWS IAM AUTH.
	awsAuth := v4.NewSigner(credentials.NewStaticCredentials(
		meta.(*dbaasCreds).access_key,
		meta.(*dbaasCreds).secret_key,
		""))
	awsAuth.Sign(req, bytes.NewReader(ec2CreateJson), "execute-api", "us-east-1", time.Now())
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
		log.Printf("Decoded. %v %d", data.Message, data.ServiceRequesID)

		// Setting the ResourceID.
		d.SetId(strconv.Itoa(data.ServiceRequesID))

		// Setting the RequestID Out.
		d.Set("service_request_id_out", strconv.Itoa(data.ServiceRequesID))
		return nil
	}
	return errors.New(data.Message)

}

func readEC2(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func updateEC2(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteEC2(d *schema.ResourceData, meta interface{}) error {
	// Logging.
	file, err := os.OpenFile("logs/ec2_destroy.log", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	// Create Request.
	endpoint := fmt.Sprintf("https://dv2-api.dbaas.aenetworks.com/aws/ec2/delete?serviceRequestId=%s", d.Id())
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
