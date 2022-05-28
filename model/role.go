package model

import "fmt"

type SimpleRole struct {
	CollectionName string
	DatabaseName   string
	RoleName       string
}

func (sr SimpleRole) String() string {
	return fmt.Sprintf("[ Collection: %s - Database: %s - Role: %s ]\n", sr.CollectionName, sr.DatabaseName, sr.RoleName)
}
