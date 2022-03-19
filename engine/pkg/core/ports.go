package core

type HttpClient interface {
	Client(baseurl string) WebserviceClient
}

type WebserviceClient interface {
	Post(data interface{}) (interface{}, error)
	Get(data interface{}) (interface{}, error)
}
