servers:
	make -j 2 api-server client-server
api-server:
	go run cmd/api/main.go
client-server:
	go run cmd/client/main.go