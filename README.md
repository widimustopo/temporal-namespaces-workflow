# temporal-namespaces-workflow

in this project , contain example temporal workflow with API framework go Echo 

how to run 

install docker if not yet installed
or read [Temporal Server is running locally](https://docs.temporal.io/docs/server/quick-install).

```
docker-compose up 
```

run API temporal-echo-namespaces
```
go run main.go
```

run it's workflow, temporal-namespaces-workflow 

```
go run main.go
```

create your own namespaces using a preconfigured CLI (tctl).
create an alias for `tctl`:

```bash
alias tctl="docker exec temporal-admin-tools tctl"
```

The following is an example of how to register a new namespace `test-namespace` with 1 day of retention:
```bash
tctl --ns test-namespace namespace register -rd 1
```