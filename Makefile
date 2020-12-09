run:
	go run ./example/table/main.go
	go run ./example/chf/main.go
	go run ./example/csv/main.go
	go run ./example/gold/main.go
	go run ./example/pretty/main.go
	
test:
	go test -v .

testrandomint:
	go test -v -run TestRandomInteger .

cover:
	go test -cover .
