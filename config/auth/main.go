package auth

import "github.com/caarlos0/env"

// Auth auth config struct
type Auth struct {
	AccessSecret      string `env:"ACCESS_SECRET" envDefault:"ACCESS_SECRET"`
	Accessexpiration  int    `env:"ACCESS_EXPIRATION" envDefault:"60"`
	RefreshSecret     string `env:"REFRESH_SECRET" envDefault:"REFRESH_SECRET"`
	Refreshexpiration int    `env:"REFRESH_EXPIRATION" envDefault:"60"`
}

// New create Auth cconfig instance
func New() (authConfig *Auth) {
	env.Parse(authConfig)
	return
}