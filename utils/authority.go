package utils

import (
	"github.com/crosect/cc-go-security/web/auth/authorization"
	"github.com/crosect/cc-go-security/web/auth/authorization/authority"
)

func ConvertRolesToSimpleAuthorities(roles []string) []authority.GrantedAuthority {
	authorities := make([]authority.GrantedAuthority, 0)
	if roles == nil {
		return authorities
	}
	for _, role := range roles {
		authorities = append(authorities, authority.NewSimpleGrantedAuthority(authorization.RolePrefix+role))
	}
	return authorities
}
