# mongodbatlas-helper
Helper functions for using the MongoDB Atlas Go Client - https://github.com/mongodb/go-client-mongodb-atlas

# Motivation
In my current job, it is usefeul to get some information about all projects of a given Atlas organization.

Instead of using `go-client-mongodb-atlas` everytime, I decided to create this simple script. It uses the `go-client-mongodb-atlas` to get the all information that I need in a single function call.

# Essential Information
The information you get when using this helper is a list of `Projects` with this structure:
```go
type SimpleProject struct {
    GroupID      string
	Name         string
	ClusterCount int
	Clusters     []SimpleCluster
	Users        []SimpleUser
}
```

You get only two fields of information for each cluster:
```go
type SimpleCluster struct {
	Name     string
	Hostname string
}
```

And, for each database user in this project, you get:
```go
type SimpleUser struct {
	Username     string
	DatabaseName string
	Roles        []SimpleRole
}

type SimpleRole struct {
	CollectionName string
	DatabaseName   string
	RoleName       string
}
```

This **Essential Information** was defined by my needs, so I tried to keep it simple.

# How To Use

## Setup
You need to export two environment variables, which are the public and private key used to connect to your Atlas Organization:
```
export MONGODB_ATLAS_PUBLIC_KEY=<your_public_key>
export MONGODB_ATLAS_PRIVATE_KEY=<your_private_key>
```

## Usage
Setup your client and then call `GetAllProjects(client)`:
```golang
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
```