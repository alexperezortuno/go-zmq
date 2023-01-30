# go-zmq

### REQUIREMENTS
Ubuntu:

```shell
sudo apt install libczmq-dev --fix-missing
```

### List of packages

```shell
go list -m -u all
```

### Install packages

```shell
go mod tidy
```

### USAGE

Test 

```shell
go run main.go -s
```

Worker

```shell
go run main.go -w
```
