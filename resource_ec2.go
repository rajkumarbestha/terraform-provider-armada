package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"log"
	"os"
	"encoding/json"
	"net/http"
	"bytes"
	"time"
)

func dbaasEC2() *schema.Resource {
	return &schema.Resource {
		SchemaVersion: 1,
		Create:        createEC2,
		Read:          readEC2,
		Update:        updateEC2,
		Delete:        deleteEC2,
		Schema: map[string]*schema.Schema {
			"aws_region": &schema.Schema {
				Type:     schema.TypeString,
				Required: true,
			},
			"action": &schema.Schema {
				Type:     schema.TypeString,
				Required: true,
			},
			"application_code": &schema.Schema {
				Type:     schema.TypeString,
				Required: true,
			},
			"environment": &schema.Schema {
				Type:     schema.TypeString,
				Required: true,
			},
			"request_type_id": &schema.Schema {
				Type:     schema.TypeString,
				Required: true,
			},
			"request_from": &schema.Schema {
				Type:     schema.TypeString,
				Required: true,
			},
			"version": &schema.Schema {
				Type:     schema.TypeString,
				Required: true,
			},
			"source_server_name": &schema.Schema {
				Type:     schema.TypeString,
				Required: true,
			},
			"destination_server": &schema.Schema {
				Type:     schema.TypeString,
				Required: true,
			},
			"source_database_name": &schema.Schema {
				Type:     schema.TypeString,
				Required: true,
			},
			"destination_database_name": &schema.Schema {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_account": &schema.Schema {
				Type:     schema.TypeString,
				Required: true,
			},
			"dbs_dbnameto_restore": &schema.Schema {
				Type:     schema.TypeString,
				Required: true,
			},
			"server_type": &schema.Schema {
				Type:     schema.TypeString,
				Required: true,
			},
			"edrive_size": &schema.Schema {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_request_id": &schema.Schema {
				Type:     schema.TypeString,
				Optional: true,
			},
			"purpose": &schema.Schema {
				Type:     schema.TypeString,
				Required: true,
			},
			"request_raised_by": &schema.Schema {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_instance_type": &schema.Schema {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_request_id_out": &schema.Schema {
				Type: schema.TypeString,
				Computed: true,
			},
		}, // Resource Schema End.
	} // Resource End.
}


func createEC2(d *schema.ResourceData, meta interface{}) error {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND, 0666)
	 if err != nil {
		log.Fatal(err)
	 }
	defer file.Close()
	log.SetOutput(file)
	log.Printf("Terraform apply command given and control is in resource create function \r\n")
	request := make(map[string]interface{})
	request["aws_region"] = d.Get("aws_region").(string)
	request["version"] = d.Get("version").(string)
	request["environment"] = d.Get("environment").(string)
	request["action"] = d.Get("action").(string)
	request["dbInstanceType"] =d.Get("db_instance_type").(string)
	request["applicationCode"] = d.Get("application_code").(string)
	request["requestFrom"]=  d.Get("request_from").(string)
	request["requestTypeID"]=  d.Get("request_type_id").(string)
	request["requestRaisedBy"]=d.Get("request_raised_by").(string)
	request["sourceServerName"]= d.Get("source_server_name").(string)
	request["sourceDatabaseName"]= d.Get("source_database_name").(string)
	request["destinationServer"] = d.Get("destination_server").(string)
	request["destinationDatabaseName"] = d.Get("destination_database_name").(string)
	request["eDriveSize"] = d.Get("edrive_size").(string)
	request["serviceAccount"] =  d.Get("service_account").(string)
	request["purpose"] = d.Get("purpose").(string)
	request["dbsDbnameToRestore"] = d.Get("dbs_dbnameto_restore").(string)
	log.Printf("Response received from Num Verify API. \r\n")
	//ec2create := EC2CreateRequest{aws_region, version, environment, action, dbInstanceType, applicationCode, requestFrom, requestRaisedBy, requestTypeID, sourceServerName, sourceDatabaseName, destinationServer,/* eDriveSize,destinationDatabaseName, serviceAccount, purpose, dbsDbnameToRestore*/}
	ec2createjson, err := json.Marshal(request)
	log.Printf("Response received from Num Verify API. \r\n")
  req, err := http.NewRequest("POST", "https://1w3zqo2l9i.execute-api.us-east-1.amazonaws.com/default/python_lambda", bytes.NewBuffer(ec2createjson))
  req.Header.Set("X-Custom-Header", "myvalue")
  req.Header.Set("Content-Type", "application/json")

	//log.Println(reflect.TypeOf(v).String())
	log.Println(d.Id()) 
	awsAuth := v4.NewSigner(credentials.NewStaticCredentials(
	    meta.(*dbaasCreds).access_key,
	    meta.(*dbaasCreds).secret_key,
			"",))
	awsAuth.Sign(req, nil, "execute-api", "us-east-1", time.Now())
	client := &http.Client{}
	log.Println(meta.(*dbaasCreds).access_key)
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
	var record Response

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	log.Printf(record.RequestIDOut)

	//log.Printf("Mobile Number = %s \r\n", record.ID)
	d.SetId(record.RequestIDOut)
	d.Set("service_request_id_out", record.RequestIDOut)
	return nil
}

func readEC2(d *schema.ResourceData, meta interface{}) error {
	//hi := "heyy"

	return nil
}

func updateEC2(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteEC2(d *schema.ResourceData, meta interface{}) error {
	return nil
}
