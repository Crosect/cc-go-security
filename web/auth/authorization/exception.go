package authorization

import (
	"github.com/crosect/cc-go/exception"
	"net/http"
)

var (
	AccessDenied = exception.New(http.StatusForbidden, "Access is denied")
)
