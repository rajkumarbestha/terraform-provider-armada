package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/hashicorp/terraform/helper/schema"
)


func armadaDynamodb() *schema.Resource {
	return &schema.Resource {
		Create:        createDynamoDB,
		Read:          readDynamoDB,
		Update:        updateDynamoDB,
		Delete:        deleteDynamoDB,
		Schema: map[string]*schema.Schema {
			"application_code": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"environment": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"request_from": &schema.Schema{
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
			"hashkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"hashkey_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"table_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"read_capacity": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"write_capacity": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"service_request_id_out": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDynamoDB(d *schema.ResourceData, meta interface{}) error {
	// Logging.
	CreateDirIfNotExist("logs")
	file, err := os.OpenFile("logs/create_resource.log", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetFlags(log.Lshortfile)
	log.SetOutput(file)

	log.Printf("Terraform apply command given and control is in resource create function. \r\n")

	// Make a request object.
	dynamoCreateReq := make(map[string]interface{})
	dynamoCreateReq["environment"] = d.Get("environment").(string)
	dynamoCreateReq["hashKey"] = d.Get("hashkey").(string)
	dynamoCreateReq["applicationCode"] = d.Get("application_code").(string)
	dynamoCreateReq["requestFrom"] = d.Get("request_from").(string)
	dynamoCreateReq["hashKeyType"] = d.Get("hashkey_type").(string)
	dynamoCreateReq["tableName"] = d.Get("table_name").(string)
	dynamoCreateReq["readCapacity"] = d.Get("read_capacity").(string)
	dynamoCreateReq["writeCapacity"] = d.Get("write_capacity").(string)
	dynamoCreateReq["purpose"] = d.Get("purpose").(string)
	dynamoCreateReq["requestRaisedBy"] = d.Get("request_raised_by").(string)

	// Convert to JSON.
	dynamoCreateJson, err := json.Marshal(dynamoCreateReq)
	log.Printf("Json Obtained %s", string(dynamoCreateJson))

	//endpoint := "https://1w3zqo2l9i.execute-api.us-east-1.amazonaws.com/default/python_lambda" for testing.
	endpoint := "/dynamodbtable/createtable"
	endpoint_from_tf := meta.(*armadaCreds).endpoint
	if len(endpoint_from_tf) > 0 {
		endpoint = endpoint_from_tf+endpoint
	}
	log.Printf("endpoint %s", endpoint)
	req, err := http.NewRequest("POST", endpoint, nil)
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	log.Printf("Request Created")

	// AWS IAM AUTH.
	awsAuth := v4.NewSigner(credentials.NewStaticCredentials(
		meta.(*armadaCreds).access_key,
		meta.(*armadaCreds).secret_key,
		""))
	awsAuth.Sign(req, bytes.NewReader(dynamoCreateJson), "execute-api", "us-east-1", time.Now())
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
		log.Printf("Checking Decode JSON. %v %d", data.Message, data.ServiceRequesID)

		// Setting the ResourceID.
		d.SetId(strconv.Itoa(data.ServiceRequesID))

		// Setting the RequestID Out.
		d.Set("service_request_id_out", strconv.Itoa(data.ServiceRequesID))
		return nil
	}
	return errors.New(data.Message)
}

func readDynamoDB(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func updateDynamoDB(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteDynamoDB(d *schema.ResourceData, meta interface{}) error {
	// Logging.
	CreateDirIfNotExist("logs")
	file, err := os.OpenFile("logs/destroy_resource.log", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetFlags(log.Lshortfile)
	log.SetOutput(file)
	endpoint :=fmt.Sprintf("/dynamodbtable/DeleteTable/%s", d.Id())
	endpoint_from_tf := meta.(*armadaCreds).endpoint
	if len(endpoint_from_tf) > 0 {
		endpoint = endpoint_from_tf+endpoint
	}
	//endpoint := fmt.Sprintf("https://dv2-api.dbaas.aenetworks.com/aws/dynamodbtable/DeleteTable/%s", d.Id())
	req, err := http.NewRequest("POST", endpoint, nil)
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	log.Printf("Request Created %s as", endpoint)

	// AWS IAM AUTH.
	awsAuth := v4.NewSigner(credentials.NewStaticCredentials(
		meta.(*armadaCreds).access_key,
		meta.(*armadaCreds).secret_key,
		""))
	awsAuth.Sign(req, nil, "execute-api", "us-east-1", time.Now())
	log.Printf("Signed the URL.")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	// Capture and log response.
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Printf("Response Received is %s", bodyString)

	var data DeleteResponse

	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if data.StatusCode == 200 {
		log.Printf("Checking Decode JSON. %v", data.Message)
		log.Printf("Service Request ID %s deleted", d.Id())

		// Deleting Resource.
		d.SetId(" ")
		return nil
	}
	return errors.New(data.Message)
}