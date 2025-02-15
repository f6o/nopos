
proto: hands/hands.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./hands/hands.proto

prisma:
	go run github.com/steebchen/prisma-client-go db push