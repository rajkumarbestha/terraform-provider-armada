provider "fakeserver"{
}

resource "fakeserver_API" "API" {
  name = "Rajkumar"
  job = "Associate Engineer"
}

output "ID" {
	value = ["${fakeserver_API.API.*.req_id}"]
}