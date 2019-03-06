
deps:
	go get github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen

proto:
	for x in **/*.proto; do protoc --go_out=paths=source_relative,plugins=grpc:. $$x; done

clean:
	rm **/**.pb.go

mocks:
	mockgen -destination expz/mock_expz/mock_service.pb.go github.com/explodes/serving/expz ExpzServiceClient,Client
	mockgen -destination logz/mock_logz/mock_service.pb.go github.com/explodes/serving/logz LogzServiceClient,Client
	mockgen -destination userz/mock_userz/mock_service.pb.go github.com/explodes/serving/userz UserzServiceClient,Client

test:
	./runtests.sh

gen: proto mocks

.PHONY: proto