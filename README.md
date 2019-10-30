# practice_golang

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

### 将来的にBlockChainのEthereumの開発で利用するかも（BitCoinで利用されているC++は学習コスト高すぎた）
→当面はSolidityとJSをやるかも

