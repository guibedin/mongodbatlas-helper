package model

import "fmt"

type SimpleCluster struct {
	Name              string
	Hostname          string
	ConnectionStrings SimpleConnectionStrings
}

func (sc SimpleCluster) String() string {
	return fmt.Sprintf("[ Name: %s - Hostname: %s\nConnection Strings: %s ]\n", sc.Name, sc.Hostname, sc.ConnectionStrings)
}
