.PHONY: build clean deploy deploy-prod test

build:
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w -X 'main.Stage=serverless'" -o bin/blogbackend/main main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean test build
	AWS_PROFILE=jonas sls deploy --verbose --stage dev --region eu-north-1

deploy-prod: clean test build
	AWS_PROFILE=jonas sls deploy --verbose --stage prod --region eu-west-1

test: 
	go test ./...
