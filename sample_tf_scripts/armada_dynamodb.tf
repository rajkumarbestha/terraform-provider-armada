provider "armada" {
    access_key = "" //required, Can be given at run time also.
    secret_key = "" //required, Can be given at run time also.
}

resource "armada_dynamodb" "dev_dynamo" {
    applicationCode = "DOC" //required
    environment = "Dev" //required
    tableName = "test" //required
    readCapacity = "1" //required
    writeCapacity = "1" //required
    hashKey = "id" //required
    hashKeyType = "Number" //required
    purpose = "Do Not Approve" //required
    requestFrom = "1" //required   
    requestRaisedBy = "vgaddam@aetvn.com" //required
    endpoint = "" //optional
}

output "dynamodb_request_ID" {
	value = ["${armada_dynamodb.dev_dynamo.*.service_request_id_out}"]
	// service_request_id_out is a computed variable
}

// Every value you provide is a string and must be enclosed in a "".
