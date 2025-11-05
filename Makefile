build:
	go build -o bin/lumen cmd/main.go

single_sample:
	go run cmd/main.go -samples=1 -width=800

low_quality:
	go run cmd/main.go -samples=10 -width=600

medium_quality:
	go run cmd/main.go -samples=25 -width=800

high_quality:
	go run cmd/main.go -samples=100 -width=1280

ultra_quality:
	go run cmd/main.go -samples=200 -width=1920

prof:
	go run cmd/main.go -cpuprofile=prof/cpu.prof -heapprofile=prof/heap.prof
