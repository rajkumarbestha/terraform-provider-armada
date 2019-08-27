provider "armada"{
  access_key = "" //required, Can be given at run time also.
  secret_key = "" //required, Can be given at run time also.
}

resource "armada_EC2" "dev_ec2" {
  version = "" //required
  application_code = "" //required
  environment = "" //required (dev,QA,UAT)
  request_type_id = "" //required 
  db_instance_type = "" //required 
  request_from = "1" //required (1 for API)
  source_server_name = "" //required 
  destination_server = "" //required 
  source_database_name = "" //required 
  destination_database_name = "" //required 
  dbs_dbnameto_restore = "" //required 
  service_account = "" //required 
  server_type = "" //required 
  edrive_size = "" //required 
  request_raised_by ="" //required (email)
  purpose = "" //required
  endpoint = "" //optional
}

output "request_ID" {
	value = ["${dbaas_EC2.dev_ec2.*.service_request_id_out}"]
	// service_request_id_out is a computed variable
}

// Every value you provide is a string and must be enclosed in a "".
