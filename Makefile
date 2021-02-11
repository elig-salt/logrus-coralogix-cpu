.PHONY: build
build:
	go build -o bin/logrus-coralogix-cpu

.PHONY: run
run:
	./bin/logrus-coralogix-cpu

.PHONY: profile
profile:
	go tool pprof -http ":9000" -seconds 60 logrus-coralogix-cpu http://localhost:6060/debug/pprof/profile
