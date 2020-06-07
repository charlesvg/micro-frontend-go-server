## Simple http server in Go

##### Goals
* Serve as a basis for a containerized micro-frontend

##### Features

* Copies files from the `./web` directory to memory
* Serves files to port `8080`
* Supports configuration through `configs/application.yml`

##### Installing
* Install dependencies `go get ./...`

##### Running
* Build executable `go build cmd/mifrogo.go`
* Run the executable