# Go REST API on Kubernetes

Demo REST API using Go. Easily deployable to Kubernetes.

## Running Locally

First, in `main.go`, change `DbHost` to `localhost`

### Via Docker Compose

1. `docker-compose up`

### Manually

#### Database

1. `docker run -d --name posts-api-go -e POSTGRES_USER=postgres-dev -e POSTGRES_PASSWORD=not-for-prod -e POSTGRES_DB=dev -v posts-api-go-data:/var/lib/postgresql/data -p 5432:5432 postgres:latest`

#### API

1. `go run main.go`
1. Starts @ http://localhost:8080

## Kubernetes

1. Create a Kubernetes cluster, e.g. GKE
1. Run `kompose convert`
1. Create a zip of the entire project folder
1. Upload zip via Cloud Shell
1. Unzip
1. `chmod +x kubectl.sh`
1. `./kubectl.sh`