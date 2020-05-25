.PHONY: build_lambda
build_lambda:
	GOOS=linux GOARCH=amd64 go build -o ./build/rmq_lambda_publisher ./cmd/rmq-lambda-publisher/rmq_lambda_publisher.go

.PHONY: build
build:
	go build -o ./build/rmq_lambda_publisher ./cmd/rmq-lambda-publisher/rmq_lambda_publisher.go