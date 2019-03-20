build_win:
	env GOOS=windows GOARCH=amd64 go build ./server
build:
	go build  ./server
