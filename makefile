source:=api meta object admin

define build-ui
	cd src/adminserver/ui; yarn
endef

gen:
	$(foreach n, $(source), cd src/$(n)server; go generate ./..; cd ..)

build-all: $(build-ui)
	$(foreach n, $(source), go build -o bin/$(n) src/$(n)server/main.go;)

start: build run

build:
ifeq ($(n),admin)
	$(build-ui)
endif
	go build -o bin/$(n) src/$(n)server/main.go

run:
	./bin/$(n) app test_conf/$(n)-server-$(i).yaml

clear:
	clear
	rm -rf /workspaces/temp/*
	go test -v src/metaserver/test/api_test.go -test.run TestClearEtcd