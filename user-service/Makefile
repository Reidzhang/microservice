build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/microservice/user-service \
		proto/user/user.proto
	docker build -t user-service .