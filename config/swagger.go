package config

import (
	"github.com/JOIN-M-Y/server/docs"
)

// Swagger swagger config struct
type Swagger struct{}

// NewSwagger create swagger configuration instance
func NewSwagger() *Swagger {
	docs.SwaggerInfo.Title = "JOIN api"
	docs.SwaggerInfo.Description = "JOIN Rest api"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = "localhost:5000"
	return &Swagger{}
}
