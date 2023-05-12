package testUtil

import (
	ccgo "github.com/crosect/cc-go"
	"go.uber.org/fx"
)

func JwtTestUtilOpt() fx.Option {
	return fx.Options(
		ccgo.ProvideProps(NewJwtTestProperties),
		fx.Provide(NewJwtTestUtil),
	)
}
