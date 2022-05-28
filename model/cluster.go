package model

import "fmt"

type SimpleCluster struct {
	Name     string
	Hostname string
}

func (sc SimpleCluster) String() string {
	return fmt.Sprintf("[ Name: %s - Hostname: %s ]\n", sc.Name, sc.Hostname)
}
