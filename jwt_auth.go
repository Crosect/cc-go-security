package ccsecurity

import (
	"errors"
	"fmt"
	"github.com/crosect/cc-go-security/web/auth/authen"
	"github.com/crosect/cc-go-security/web/config"
	"github.com/crosect/cc-go-security/web/filter"
	"go.uber.org/fx"
)

func JwtAuthFilterOpt() fx.Option {
	return fx.Provide(fx.Annotated{
		Group:  "authentication_filter",
		Target: NewJwtAuthFilter,
	})
}

type JwtAuthFilterIn struct {
	fx.In
	SecurityProperties  *config.HttpSecurityProperties
	AuthProviderManager *authen.ProviderManager
}

func NewJwtAuthFilter(in JwtAuthFilterIn) (filter.AuthenticationFilter, error) {
	if in.SecurityProperties.Jwt == nil {
		return nil, errors.New("missing JWT Auth config")
	}
	in.AuthProviderManager.AddProvider(authen.NewJwtAuthProvider())
	jwtFilter, err := filter.JwtAuthSecurityFilter(in.SecurityProperties.Jwt)
	if err != nil {
		return nil, fmt.Errorf("cannot init JWT Security Filter: [%v]", err)
	}
	return jwtFilter, nil
}
