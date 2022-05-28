package main

import (
	"fmt"
	"log"
	"os"

	"github.com/guibedin/mongodbatlas-helper/helper"
	"github.com/mongodb-forks/digest"
	"go.mongodb.org/atlas/mongodbatlas"
)

func main() {
	// Setup mongodb client
	t := digest.NewTransport(os.Getenv("MONGODB_ATLAS_PUBLIC_KEY"), os.Getenv("MONGODB_ATLAS_PRIVATE_KEY"))
	tc, err := t.Client()
	if err != nil {
		log.Fatalf(err.Error())
	}
	client := mongodbatlas.NewClient(tc)

	// Call helper passing client as a parameter
	projects := helper.GetAllProjects(client)
	fmt.Println(projects)
}
