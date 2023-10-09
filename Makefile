#  sqlStr := postgresql://root:secret@localhost:5434/simple_bank?sslmode=disable
pass := nghjKO0hr\$$237!.C
sqlStr := postgresql://esystem_user:$(pass)@34.143.228.170:5432/Customer001?sslmode=disable
pr: 
	echo $(sqlStr)
env: 
	source ~/.bash_profile
	
gen: 
	protoc 	--go-grpc_out=. --proto_path=proto proto/*.proto \
	-I. \
	-I${GOPATH}/src \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=. --grpc-gateway_out=:. \
  --openapiv2_out=:swagger

clean:
	rm pb/*.go 

server1:
	go run cmd/server/main.go -port 50051

server2:
	go run cmd/server/main.go -port 50052

server1-tls:
	go run cmd/server/main.go -port 50051 -tls

server2-tls:
	go run cmd/server/main.go -port 50052 -tls

server-cmd:
	go run cmd/server/main.go -port 8080

server-tls:
	go run cmd/server/main.go -port 8080 -tls

rest:
	go run cmd/server/main.go -port 8081 -type rest -endpoint gopherface.local:8080 -tls

client:
	go run cmd/client/main.go -address 0.0.0.0:8080

client-tls:
	go run cmd/client/main.go -address gopherface.local:8080 -tls

test1:
	go test -cover -race ./...

cert:
	cd cert; ./gen.sh; cd ..

# .PHONY: gen clean server client test cert 

postgres:
	docker run --name postgres12 --network bank-network -p 5434:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.2-alpine

postgresstart:
	docker container restart postgres12

mysql:
	docker run --name mysql8 -p 3306:3306  -e MYSQL_ROOT_PASSWORD=secret -d mysql:8

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

psql:
	docker exec -it postgres12 psql -U root

migrateup:
	migrate -path db/migration -database $(sqlStr) -verbose up

migrateup1:
	migrate -path db/migration -database $(sqlStr) -verbose up 1

migrategoto:
	migrate -path db/migration -database $(sqlStr) -verbose goto $(v)

migrateforce:
	migrate -path db/migration -database $(sqlStr) -verbose force $(v)

migratedown:
	migrate -path db/migration -database $(sqlStr) -verbose down

migratedown1:
	migrate -path db/migration -database $(sqlStr) -verbose down 1

sqlc:
	exec /Users/rhickmercado/Documents/Programming/go/src/sqlc/sqlc/main generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/mock.go simplebank/db/datastore Store

setdb2:
	source setdb2
build:
	GOOS=linux GOARCH=amd64 go build -o cesystem
	scp -P 2222 cesystem rhickmercado@34.142.152.56:/home/rhickmercado
    scp cesystem rhickmercado@34.143.228.170:/home/rhickmercado/cmd/server
	ssh rhickmercado@34.143.228.170

	GOOS=windows GOARCH=amd64 go build -o lesystem

	GOOS=linux GOARCH=386 go build
	scp jobs  root@example.com:/var/www/go
	scp esystem.service root@example.com:/var/www/go
	
.PHONY: psql postgresstart migrategoto migrateforce env gen clean server-cmd client test1 cert env postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mockall action setdb2, build


