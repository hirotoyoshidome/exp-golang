## Build
* exec

```
go run main.go
```

* build

```
go build -o Hello hello.go
./Hello
```

## Docker

```
docker build -t mygolang .
docker run --rm -it -d -v $(pwd):/var/golang --name mygolang mygolang:latest
docker exec -it mygolang sh
```
