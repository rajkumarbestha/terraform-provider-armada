provider "fakeserver"{
}

resource "fakeserver_API" "API" {
  service = "RDS"
  access_key = "2346abfccbbf1f028b7db6afecbc14c5"
}

output "request_ID_from_API" {
	value = ["${fakeserver_API.API.*.req_id}"]
}