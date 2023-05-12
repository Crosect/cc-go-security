package authen

import (
	"github.com/crosect/cc-go/exception"
	"net/http"
)

var BadCredentials = exception.New(http.StatusUnauthorized, "Bad credentials")
