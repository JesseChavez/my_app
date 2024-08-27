BINARY_NAME=my_app

build:
	ENKI_ENV=production yarn build
	ENKI_ENV=production mv frontend/builds public/assets
	GOOS=linux GOARCH=amd64 go build -o dist/${BINARY_NAME}-linux-amd64 cmd/main.go

clean:
	go clean
	rm -rfv dist/*
	rm -rfv frontend/builds
	rm -rfv public/assets
