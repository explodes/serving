
deps:
	go get github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen

proto:
	for x in **/*.proto; do protoc --go_out=paths=source_relative,plugins=grpc:. $$x; done

clean:
	rm **/**.pb.go

mocks:
	mockgen -destination expz/mock_service_pb_test.go -package expz github.com/explodes/serving/expz ExpzServiceClient

.PHONY: proto