## Simple http server in Go

##### Goals
* Serve as a basis for a containerized micro-frontend

##### Features

* Copies files from the `./web` directory to memory
* Supports configuration through `configs/app.yml`

##### Installing
* Install dependencies `go get ./...`

##### Running
* Build executable `go build cmd/app.go`
* Run the executable