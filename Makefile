CURRDIR := $(shell pwd)
GOPATH := ${CURRDIR}/.build
PYENV := ${CURRDIR}/.pyenv

godeps:
	GOPATH=${GOPATH} go get google.golang.org/grpc
	GOPATH=${GOPATH} go get github.com/golang/protobuf/protoc-gen-go

pydeps: pyvenv
	${PYENV}/bin/pip install -r pysrc/requirements.txt

pyvenv:
	virtualenv ${PYENV}

gen-go: godeps dirs
	PATH=${GOPATH}/bin:${PATH} protoc -I proto proto/srv.proto --go_out=plugins=grpc:.build/src/protobuf

gen-py: pydeps
	${PYENV}/bin/python -m grpc_tools.protoc -I./proto --python_out=./pysrc --grpc_python_out=./pysrc ./proto/srv.proto

go-server: gen-go
	GOPATH=${GOPATH} go build ./gosrc/main.go

py-client: gen-py
	${PYENV}/bin/python pysrc/main.py

dirs:
	mkdir -p ${GOPATH}/src/protobuf/
