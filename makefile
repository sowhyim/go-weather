build:
	go build -o ./bin/weather ./main.go

install:
	sudo go build -o /usr/bin/weather ./main.go
