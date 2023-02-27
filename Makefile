.PHONY : mac windows linux all mkdir
mac: prepare
	 CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o ./build/BiliTicketSeller_Mac/BiliTicketSeller main.go

windows: prepare
	go generate
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o ./build/BiliTicketSeller_Windows/BiliTicketSeller.exe ./

linux: prepare
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o ./build/BiliTicketSeller_Linux/BiliTicketSeller main.go

all: mac windows linux

prepare:
	mkdir -p ./build/
	mkdir -p ./build/BiliTicketSeller_Mac
	mkdir -p ./build/BiliTicketSeller_Windows
	mkdir -p ./build/BiliTicketSeller_Linux