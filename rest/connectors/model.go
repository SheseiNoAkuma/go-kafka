package connectors

type Configuration interface {
	Name() string
	BaseUrl() string
	Auth() Authentication
}

type Authentication interface {
	UserName() string
	Password() string
}
