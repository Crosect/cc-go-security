package ccsecurity

import (
	"github.com/crosect/cc-go"
	secHttpClient "github.com/crosect/cc-go-security/web/client"
	"github.com/crosect/cc-go/web/client"
	"go.uber.org/fx"
)

func SecuredHttpClientOpt() fx.Option {
	return fx.Options(
		ccgo.ProvideProps(secHttpClient.NewSecurityProperties),
		fx.Provide(fx.Annotated{
			Group:  "contextual_http_client_wrapper",
			Target: NewSecuredHttpClient,
		}),
	)
}

func NewSecuredHttpClient(props *secHttpClient.SecurityProperties) ccgo.ContextualHttpClientWrapper {
	return func(client client.ContextualHttpClient) (client.ContextualHttpClient, error) {
		return secHttpClient.NewSecuredHttpClient(client, props), nil
	}
}
