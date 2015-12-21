# CI Registerer

## What's this?

This is a library for building client to register a Go REST API server to Submit Agent.


## How to Use

* Requirements
  * Go 1.5.x and later
  * A running instance of [elsa](https://github.com/cytoscape-ci/elsa)

```go

// Create registration object
reg := NewRegistration("myservicename", "192.168.99.100", 3000, 4, "v1")

// Register it
err := RegisterService("http://my-elsa-instance:8080/registration", reg, RetrySetting{RetryInterval:10, RetryMax:5})

```

For testing, you can use official [Docker image](https://hub.docker.com/r/cytoscape/elsa/) to run local instance of elsa.
