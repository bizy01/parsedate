## yacc build
yacc:
	goyacc -o parser/parser.y.go parser/parser.y

## test
test:
	go test ./... -v