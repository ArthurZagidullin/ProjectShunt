build_win:
	env GOOS=windows GOARCH=amd64 go build -o ./cmd/winserver.exe ./server
build:
	go build -o ./cmd/linserver ./server
