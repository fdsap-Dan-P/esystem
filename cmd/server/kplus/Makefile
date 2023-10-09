
build:
	GOOS=linux GOARCH=amd64 go build
	scp server rhickmercado@34.143.228.170:/eSystem/server/cmd/server
	scp esystem.service rhickmercado@34.143.228.170:/eSystem/server/cmd/server

	sudo systemctl start esystem

.PHONY: build

