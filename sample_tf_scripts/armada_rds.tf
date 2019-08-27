provider "armada"{
  access_key = "" //required, Can be given at run time also.
  secret_key = "" //required, Can be given at run time also.
}

resource "armada_rds" "dev_rds" {
  database_engine = "oracle-ee" //required
  application_code = "DOC" //required
  environment = "UAT" //required
  request_type_id = "5" //required
  request_from = "1" //required
  database_engine_verison = "12.1" //required
  db_restore = "false" //required
  db_restore_path = "" //required
  database_name = "DOC" //required
  allocated_storage = "100" //required
  db_instance_class = "db.t2.medium" //required
  request_created_date = "2019-08-26T09:35:41.707Z" //required
  request_raised_by ="rajkumarb@aenetworks.com" //required
  purpose = "Terraform Plugin Testing" //required
  endpoint = "" //optional
}

output "rds_request_ID" {
	value = ["${armada_rds.dev_rds.*.service_request_id_out}"]
	// service_request_id_out is a computed variable
}
