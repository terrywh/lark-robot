.PHONY: all

PACKAGE=github.com/terrywh/lark-robot/v1

CLIENT_SOURCE=client/client.go
TOOLKIT_SOURCE=toolkit/hash.go toolkit/toolkit.go
SERVER_SOURCE=server/init.go server/init-debug.go server/main.go

all: server-debug server-lark-robot

server-debug: ${SERVER_SOURCE} ${CLIENT_SOURCE} ${TOOLKIT_SOURCE}
	go build -o $@ -tags debug ${PACKAGE}/server

server-lark-robot: ${SERVER_SOURCE} ${CLIENT_SOURCE} ${TOOLKIT_SOURCE}
	go build -o $@ ${PACKAGE}/server
