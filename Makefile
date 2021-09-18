build:
	go build -o bin/webauthncf cmd/cli/main.go

run-server:
	go run cmd/cli/main.go server --port 8888

# compile-all:
# 	echo "Compiling for every OS and Platform"
# 	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 cmd/cli/main.go 
# 	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 cmd/cli/main.go
# 	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go

