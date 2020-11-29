run:
	go run ./example/table/main.go

test:
	go test -v .

testcheck:
	go test -v -run TestCheckArg .

cover:
	go test -cover .
