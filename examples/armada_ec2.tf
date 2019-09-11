
resource "armada_ec2" "dev_ec2" {
  version = "2016" //required
  application_code = "DBT" //required
  environment = "UAT" //required (dev,QA,UAT)
  request_type_id = "6" //required 
  db_instance_type = "t2.small" //required 
  request_from = "1" //required (1 for API)
  source_server_name = "AZV-DBS4E-GBRD7" //required 
  destination_server = "null" //required 
  source_database_name = "BSM_DEV" //required 
  destination_database_name = "null" //required 
  dbs_dbnameto_restore = "BRD_IAM" //required 
  service_account = "SVCDBS6EUDBT" //required 
  server_type = "DBS" //required 
  edrive_size = "40" //required 
  request_raised_by ="rajkumarb@virtusa.com" //required (email)
  purpose = "Terraform Plugin Testing" //required
  endpoint = "" //optional
}

output "ec2_request_ID" {
	value = ["${armada_ec2.dev_ec2.*.service_request_id_out}"]
	// service_request_id_out is a computed variable
}

// Every value you provide is a string and must be enclosed in a "".
