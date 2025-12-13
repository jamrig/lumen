build:
	go build -o bin/lumen cmd/main.go

single_sample:
	go run cmd/main.go -samples=1 -width=1000

low_quality:
	go run cmd/main.go -samples=10 -width=1000

medium_quality:
	go run cmd/main.go -samples=25 -width=1000

high_quality:
	go run cmd/main.go -samples=50 -width=1280

ultra_quality:
	go run cmd/main.go -samples=100 -width=1920

high_samples:
	go run cmd/main.go -samples=100 -width=400

prof:
	go run cmd/main.go -cpuprofile=prof/cpu.prof -heapprofile=prof/heap.prof
