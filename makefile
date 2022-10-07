source:=api meta object

gen:
	$(foreach n, $(source), cd $(n)server; go generate ./..; cd ..)

build-all:
	$(foreach n, $(source), go build -o bin/$(n) $(n)server/main.go;)

start: build run

build:
	go build -o bin/$(n) $(n)server/main.go

run:
	./bin/$(n) app test_conf/$(n)-server-$(i).yaml

clear:
	clear
	rm -r /workspaces/temp/*
	go test -v metaserver/test/api_test.go -test.run TestClearEtcd