test:
	go test -v .

bench:
	go test -run=none -bench=. .

prof:
	go test -bench=Echo4 -cpuprofile=cpu.log .
	go tool pprof -text -nodecount=20 ./cli_demo_20160825.test ./cpu.log

doc:
	@echo 'http://localhost:8080/pkg/github.com/mkocikowski/cli_demo_20160825/'
	godoc -http=:8080
