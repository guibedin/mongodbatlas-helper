package model

import (
	"fmt"
	"strings"
)

type SimpleUser struct {
	Username     string
	DatabaseName string
	Roles        []SimpleRole
}

func (su SimpleUser) String() string {
	var rolesStr []string
	for _, role := range su.Roles {
		rolesStr = append(rolesStr, role.String())
	}
	roles := strings.Trim(strings.Join(rolesStr, ""), "\n")

	return fmt.Sprintf("[ Username: %s - Database: %s - Roles: %s ]\n", su.Username, su.DatabaseName, roles)
}
