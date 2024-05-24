all: local linux_x64 mac_intel mac_silicon win64
linux_x64: FORCE
	GOOS=linux GOARCH=amd64 go build -o check_scanner_status *.go
mac_intel: FORCE
	GOOS=darwin GOARCH=amd64 go build -o check_scanner_status.osx_intel *.go
mac_silicon: FORCE
	GOOS=darwin GOARCH=arm64 go build -o check_scanner_status.osx_m1 *.go
win64: FORCE
	GOOS=windows GOARCH=amd64 go build -o check_scanner_status.exe *.go
local: FORCE
	go build

FORCE: ;
