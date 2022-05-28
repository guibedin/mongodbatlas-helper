package model

import (
	"fmt"
	"strings"
)

type SimpleProject struct {
	GroupID      string
	Name         string
	ClusterCount int
	Clusters     []SimpleCluster
	Users        []SimpleUser
}

func (sp SimpleProject) String() string {
	var usersStr []string
	for _, user := range sp.Users {
		usersStr = append(usersStr, user.String())
	}
	users := strings.Trim(strings.Join(usersStr, ""), "\n")

	var clustersStr []string
	for _, cluster := range sp.Clusters {
		clustersStr = append(clustersStr, cluster.String())
	}
	clusters := strings.Trim(strings.Join(clustersStr, ""), "\n")

	return fmt.Sprintf("GroupID: %s\nName: %s\nCluster Count: %d\nClusters:\n%s\nUsers:\n%s", sp.GroupID, sp.Name, sp.ClusterCount, clusters, users)
}
