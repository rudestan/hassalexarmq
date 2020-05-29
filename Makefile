.PHONY: build_lambda_full
build_lambda_full: build_lambda zip_lambda
	rm ./build/lambda_rmq_publisher

.PHONY: build_consumer_arm_full
build_consumer_arm_full: build_consumer_arm tar_consumer_arm
	rm ./build/arm_alexa_rmq_consumer

.PHONY: build_lambda
build_lambda:
	GOOS=linux GOARCH=amd64 go build -o ./build/lambda_rmq_publisher ./cmd/lambda-publisher/lambda_rmq_publisher.go

.PHONY: build_consumer_arm
build_consumer_arm:
	GOOS=linux GOARCH=arm go build -o ./build/arm_alexa_rmq_consumer ./cmd/alexareq-consumer/alexareq_consumer.go

.PHONY: zip_lambda
zip_lambda:
	zip -j ./build/lambda_rmq_publisher.zip ./build/lambda_rmq_publisher

.PHONY: tar_consumer_arm
tar_consumer_arm:
	tar -czvf ./build/arm_alexa_rmq_consumer.tar.gz ./build/arm_alexa_rmq_consumer

.PHONY: build
build:
	go build -o ./build/lambda_rmq_publisher ./cmd/lambda-publisher/lambda_rmq_publisher.go
	go build -o ./build/alexa_rmq_consumer ./cmd/alexareq-consumer/alexareq_consumer.go