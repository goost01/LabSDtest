NAME="default"
gen:
	protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     helloworld/helloworld.proto

clean:
	rm helloworld/*.go

player:
	go run greeter_client/main.go $(NAME)
	
server:
	go run greeter_server/main.go