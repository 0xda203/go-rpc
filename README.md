
# go-rpc

A basic information system that can be distributed across multiple servers, each of which implements a repository of information about parts or components (parts), built using Remote Procedure Call (RPC) in Go.

## Usage/Examples

Starting local nameserver on port 9000

```javascript
go run cmd/naming/main.go -host 127.0.0.1:9000
```

Starting local part repository server on port 9001 with name 'server1' 
that connects and registers itself into to a nameserver at 127.0.0.1:9000

```javascript
go run cmd/service/main.go -host 127.0.0.1 -port 9001 -name server1 -ns 127.0.0.1:9000
```

Starting a local client connecting to a nameserver (for naming server lookup) listening on 9000

```javascript
go run cmd/client/main.go -ns 127.0.0.1:9000
```


## Installation

Must have go 1.18 installed on machine.

```bash
  cd go-rpc
  go get github.com/google/uuid v1.3.0
```
    