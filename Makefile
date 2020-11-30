run:
	go run ./example/table/main.go
	go run ./example/chf/main.go
	go run ./example/csv/main.go
	go run ./example/gold/main.go
	go run ./example/pretty/main.go
	
test:
	go test -v .

testcheck:
	go test -v -run TestCheckArg .

cover:
	go test -cover .
