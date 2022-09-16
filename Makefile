
build-mocks:
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	@~/go/bin/mockgen -source=domain/bet/bet.interfaces.go -destination=domain/bet/mock/bet.go -package=mock
	@~/go/bin/mockgen -source=domain/challenged/challenged.interfaces.go -destination=domain/challenged/mock/challenged.go -package=mock
	@~/go/bin/mockgen -source=domain/gambler/gambler.interfaces.go -destination=domain/gambler/mock/gambler.go -package=mock

test:
	go test ./...  -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html