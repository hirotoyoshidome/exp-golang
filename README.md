## install

```
brew install go
```

## exec
* exec

```
go run hello.go
```

* build

```
go build hello.go
./hello
```

## Docker

* build

```
docker build -t my_golang_div:latest .
```
* exec

```
docker run --rm -it my_golang_dev:latest /bin/sh
go run hello.go
exit
```
