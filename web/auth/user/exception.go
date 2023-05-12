package user

import (
	"github.com/crosect/cc-go/exception"
	"net/http"
)

var (
	NotFound = exception.New(http.StatusNotFound, "User not found")
)
