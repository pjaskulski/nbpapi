run:
	go run ./example/main.go

test:
	go test -v .

testcheck:
	go test -v -run TestCheckArg .

cover:
	go test -cover .
