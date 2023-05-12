package service

import (
	"github.com/crosect/cc-go-security/web/auth/authen"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

type JwtService interface {
	GetAuthentication(token *jwt.Token, request *http.Request) (authen.Authentication, error)
}
