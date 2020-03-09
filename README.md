## Go development

### bash configure
```
export GO111MODULE=on
export PATH='/usr/local/go/bin':$PATH
export GOPATH=$HOME/Study/go
export PATH=$PATH:$GOPATH/bin 
# export GOPROXY="https://athens.azurefd.net"
# export GOPROXY="https://goproxy.io"
# export GOPROXY="https://goproxyus.herokuapp.com"
export GOPROXY="https://goproxy.cn"
```

### iris install
```
go get -u github.com/kataras/iris/v12@latest 
```

### init go mod for new project (import v12 suffix)
```
go mod init
go build
```

### run project
```
go run main.go
```

### build without debug info
```
go build -ldflags "-s -w"
```

### install
```
go install -ldflags '-s'
```