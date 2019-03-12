
deps:
	go get github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen

proto:
	for x in **/*.proto; do protoc --go_out=paths=source_relative,plugins=grpc:. $$x; done

clean:
	rm **/**.pb.go

mocks:
	mockgen -destination expz/mock_expz/mock_service_client.pb.gen.go github.com/explodes/serving/expz ExpzServiceClient
	mockgen -destination logz/mock_logz/mock_service_client.pb.gen.go github.com/explodes/serving/logz LogzServiceClient
	mockgen -destination logz/mock_logz/mock_client.gen.go github.com/explodes/serving/logz Client
	mockgen -destination userz/mock_userz/mock_service_client.pb.gen.go github.com/explodes/serving/userz UserzServiceClient
	mockgen -destination userz/mock_userz/mock_client.gen.go github.com/explodes/serving/userz Client
	mockgen -destination userz/mock_userz/mock_storage.gen.go github.com/explodes/serving/userz Storage

test:
	./runtests.sh

gen: proto mocks

.PHONY: proto