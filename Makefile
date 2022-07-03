bench:
	go test -bench=. -benchtime=1s -benchmem ./benchmark/compare_test.go

.PHONY: test
test:
	go test -v -count=1 ./...	