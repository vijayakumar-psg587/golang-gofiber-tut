****Startup repo with GoLang****

***
This repo servers as the kickstarter for golang projects
Also servers as a starting point for gofiber/gprc/microservice development



****Steps to run in local****
- Download the repo
- Run `go install` to install all dependencies
- Now run `go run main.go` to get started

****Additional features to Note****
- To download/install a particular version of dependecy - ` go get github.com/lib/pq@v1.8.0`
- To download latest use `go get github.com/lib/pq@latest`
- TO download and update all transitive dependecies of a particular module `
  go get -u github.com/lib/pq`
- To list dependecies of main module `  go list -m all`
- To list versions of a package `go list -m -versions github.com/lib/pq`
- To remove unused dep `
  go mod tidy`


  Go tool let's you download remote packages either by passing their import path to go get or calling go get ./... inside your project's directory to recursively get all remote dependencies.
