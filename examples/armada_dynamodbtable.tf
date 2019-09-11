
resource "armada_dynamodb" "dev_dynamo" {
    application_code = "DOC"
    environment = "Dev" //required
    table_name = "test" //required
    read_capacity = "1" //required
    write_capacity = "1" //required
    hashkey = "id" //required
    hashkey_type = "Number" //required
    purpose = "Terraform Plugin Testing" //required
    request_from = "1" //required
    request_raised_by = "rajkumarb@virtusa.com" //required
    endpoint = "" //optional
}

output "dynamodb_request_ID" {
	value = ["${armada_dynamodb.dev_dynamo.*.service_request_id_out}"]
	// service_request_id_out is a computed variable
}