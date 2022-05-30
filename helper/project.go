package helper

import (
	"context"

	"github.com/guibedin/mongodbatlas-helper/model"
	"go.mongodb.org/atlas/mongodbatlas"
)

func getProjectClusters(client *mongodbatlas.Client, groupID string) []model.SimpleCluster {
	clusters, _, err := client.Clusters.List(context.Background(), groupID, nil)
	if err != nil {
		panic(err)
	}

	var simpleClusters []model.SimpleCluster
	for _, cluster := range clusters {
		simpleClusters = append(simpleClusters, model.SimpleCluster{
			Name:     cluster.Name,
			Hostname: cluster.MongoURI,
			ConnectionStrings: model.SimpleConnectionStrings{
				Standard:    cluster.ConnectionStrings.Standard,
				StandardSrv: cluster.ConnectionStrings.StandardSrv},
		})
	}
	return simpleClusters
}

func getProjectUsers(client *mongodbatlas.Client, groupID string) []model.SimpleUser {
	users, _, err := client.DatabaseUsers.List(context.Background(), groupID, nil)
	if err != nil {
		panic(err)
	}

	var simpleUsers []model.SimpleUser
	for _, user := range users {
		var simpleRoles []model.SimpleRole
		for _, role := range user.Roles {
			simpleRoles = append(simpleRoles, model.SimpleRole{
				CollectionName: role.CollectionName,
				DatabaseName:   role.DatabaseName,
				RoleName:       role.RoleName,
			})
		}

		simpleUsers = append(simpleUsers, model.SimpleUser{
			Username:     user.Username,
			DatabaseName: user.DatabaseName,
			Roles:        simpleRoles,
		})
	}

	return simpleUsers
}

func GetAllProjects(client *mongodbatlas.Client) []model.SimpleProject {
	projects, _, err := client.Projects.GetAllProjects(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	var simpleProjects []model.SimpleProject
	for _, project := range projects.Results {
		simpleProjects = append(simpleProjects, model.SimpleProject{
			GroupID:      project.ID,
			Name:         project.Name,
			ClusterCount: project.ClusterCount,
			Clusters:     getProjectClusters(client, project.ID),
			Users:        getProjectUsers(client, project.ID),
		})
	}

	return simpleProjects
}
