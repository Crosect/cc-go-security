package middleware

import (
	"github.com/crosect/cc-go-security/utils"
	"github.com/crosect/cc-go-security/web/auth/authen"
	"github.com/crosect/cc-go-security/web/auth/authorization"
	"github.com/crosect/cc-go-security/web/constant"
	secContext "github.com/crosect/cc-go-security/web/context"
	"github.com/crosect/cc-go-security/web/filter"
	"github.com/crosect/cc-go/exception"
	"github.com/crosect/cc-go/web/log"
	"github.com/crosect/cc-go/web/response"
	"net/http"
)

func Auth(
	authManager authen.AuthenticationManager,
	accessDecisionManager authorization.AccessDecisionManager,
	filters []filter.AuthenticationFilter,
) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			matchedUrl := secContext.GetMatchedUrlProtection(r)
			if matchedUrl == nil {
				log.Debug(r.Context(), "Authentication is not required for this request, skip")
				next.ServeHTTP(w, r)
				return
			}

			// Create and run the security filter chain
			filterChainHandler := filter.CreateAuthenticationHandler(filters, filter.UnauthorizedHandler)
			authentication := filterChainHandler(w, r)
			if authentication == nil {
				log.Info(r.Context(), "Authentication is required to access this resource")
				writeAuthenticationDirective(w, matchedUrl.UnauthorizedWwwAuthenticateHeaders)
				response.WriteError(w, exception.Unauthorized)
				return
			}

			authentication, err := authManager.Authenticate(authentication)
			if err != nil {
				log.Info(r.Context(), "Authentication failed, error [%s]", err.Error())
				response.WriteError(w, exception.Unauthorized)
				return
			}

			if !authentication.Authenticated() {
				log.Info(r.Context(), "Authentication failed, the request is unauthenticated")
				response.WriteError(w, exception.Unauthorized)
				return
			}

			restrictedAuthorities := utils.ConvertRolesToSimpleAuthorities(matchedUrl.Roles)
			if err := accessDecisionManager.Decide(authentication, restrictedAuthorities); err != nil {
				if _, ok := err.(exception.Exception); ok {
					log.Info(r.Context(), "Authorization failed, error [%s]", err.Error())
				} else {
					log.Error(r.Context(), "Error when trying to authorize request, error [%v]", err)
				}
				response.WriteError(w, err)
				return
			}

			// Continues to the next step
			next.ServeHTTP(w, secContext.AttachAuthentication(r, authentication))
		})
	}
}

func writeAuthenticationDirective(w http.ResponseWriter, unauthorizedWwwAuthenticateHeaders []string) {
	for _, val := range unauthorizedWwwAuthenticateHeaders {
		w.Header().Add(constant.HeaderWWWAuthenticate, val)
	}
}
