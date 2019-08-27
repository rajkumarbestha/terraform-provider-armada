provider "armada"{
  access_key = var.access_key
  secret_key = var.secret_key
}

resource "armada_rds" "dev_rds" {
  database_engine = "oracle-ee"
  application_code = "DOC"
  environment = "UAT"
  request_type_id = "5"
  request_from = "1"
  database_engine_verison = "12.1"
  db_restore = "false"
  db_restore_path = ""
  database_name = "DOC"
  allocated_storage = "100"
  db_instance_class = "db.t2.medium"
  request_created_date = "2019-08-26T09:35:41.707Z"
  request_raised_by ="rajkumarb"
  purpose = "Terraform Plugin Testing"
}

output "rds_request_ID" {
	value = ["${armada_rds.dev_rds.*.service_request_id_out}"]
}
