package testUtil

import "github.com/crosect/cc-go/config"

type JwtTestProperties struct {
	PrivateKey string
}

func NewJwtTestProperties(loader config.Loader) (*JwtTestProperties, error) {
	props := &JwtTestProperties{}
	if err := loader.Bind(props); err != nil {
		return nil, err
	}
	return props, nil
}

func (j *JwtTestProperties) Prefix() string {
	return "app.security.http.jwt"
}
