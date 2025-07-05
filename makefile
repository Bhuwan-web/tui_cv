build:
	go build -o cv main.go

build_win:
	GOOS=windows GOARCH=amd64 go build -o cv.exe main.go

build_linux:
	GOOS=linux GOARCH=amd64 go build -o cv_linux main.go

run: build
	cv
