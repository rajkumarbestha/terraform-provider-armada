provider "armada"{
}

resource "armada_EC2" "dev_ec2" {
  aws_region = "us-east-1"
  version = "1.0"
  action= "create"
  application_code = "DBA"
  environment = "dev"
  request_type_id = "6"
  db_instance_type = "t2.medium"
  request_from = "1"
  source_server_name = "sourceserver"
  destination_server = "destinationserver"
  source_database_name = "sourcedatabase"
  destination_database_name = "destinationdatabase"
  dbs_dbnameto_restore = "db"
  service_account = "iybfayw64yfb"
  server_type = "dbs"
  edrive_size = "100"
  request_raised_by ="rajkumar"
  purpose = "testing"
}

output "request_ID" {
	value = ["${dbaas_EC2.dev_ec2.*.service_request_id_out}"]
}
