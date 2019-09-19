package armada

type DummyResponse struct {
	RequestID string
}

type CreateResponse struct {
	Message         string `json:"message"`
	ServiceRequesID int `json:"serviceRequesID"`
	StatusCode      int         `json:"statusCode"`
	Data            interface{} `json:"data"`
}

type DeleteResponse struct {
	Message         string `json:"message"`
	StatusCode      int         `json:"statusCode"`
	Data            interface{} `json:"data"`
}

type armadaCreds struct {
	access_key string
	secret_key string
	endpoint string
}


type EC2CreateRequest struct {
	aws_region         string
	version            string
	environment        string
	action             string
	dbInstanceType     string
	applicationCode    string
	requestFrom        string
	requestRaisedBy    string
	requestTypeID      string
	sourceServerName   string
	sourceDatabaseName string
	destinationServer  string
	eDriveSize string
	destinationDatabaseName string
	serviceAccount string
	purpose string
	dbsDbnameToRestore string
}

type EC2DestroyRequest struct {
	aws_region              string
	version                 string
	environment             string
	action                  string
	applicationCode         string
	dbInstanceType          string
	requestFrom             string
	requestRaisedBy         string
	requestTypeID           string
	sourceServerName        string
	sourceDatabaseName      string
	destinationServer       string
	eDriveSize              string
	destinationDatabaseName string
	serviceAccount          string
	purpose                 string
	dbsDbnameToRestore      string
	serviceRequestID        string
}
