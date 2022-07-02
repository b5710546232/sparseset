bench:
	go test -bench=.  -benchmem ./benchmark/compare_sparse_map_test.go

.PHONY: test
test:
	go test -v -count=1 ./...	