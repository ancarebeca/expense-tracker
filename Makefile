all: build run

install:
	glide install

build:
	go build -v -race -o expense-tracker

rebuild:
	go build -v -race -a -o expense-tracker

run:
	./expense-tracker

clean:
	rm ./expense-tracker

functional-test:
	./services ginkgo -r; 

unit-test:	
	go test

