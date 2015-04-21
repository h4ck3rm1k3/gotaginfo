export LD_LIBRARY_PATH := $(HOME)/install/lib64
export PATH := $(PATH):$(HOME)/install/bin
export GOPATH := $(HOME)/goroot
compile:
	#~/install/bin/go build github.com/h4ck3rm1k3/taginfo
	~/install/bin/go install github.com/h4ck3rm1k3/taginfo
	./bin/taginfo

readtables:
	#~/install/bin/go build github.com/h4ck3rm1k3/taginfo/readkeys
	~/install/bin/go install github.com/h4ck3rm1k3/taginfo/readkeys
	$(GOPATH)/bin/readkeys

readdata2:
	~/install/bin/go install github.com/h4ck3rm1k3/taginfo/readdata
	$(GOPATH)/bin/readdata

genproto:
	~/install/bin/go install github.com/h4ck3rm1k3/taginfo/genproto
	$(GOPATH)/bin/genproto

genproto_transfer:
	~/install/bin/go install github.com/h4ck3rm1k3/taginfo/genproto/transfer
	$(GOPATH)/bin/transfer

deps:
	#go install github.com/glycerine/go-capnproto/capnpc-go
	#go get github.com/glycerine/go-capnproto
	#go get upper.io/db
	#go get upper.io/db/sqlite
	#go get github.com/lib/pq/hstore
	#go install github.com/h4ck3rm1k3/gorm
	#go get github.com/mattn/go-sqlite3
	#go get	github.com/h4ck3rm1k3/lbpkr
	#go install	github.com/h4ck3rm1k3/gographviz
	#go get -v golang.org/x/tools/cmd/gorename
	#go build github.com/h4ck3rm1k3/lbpkr
	#go get github.com/lhcb-org/lbpkr
	#go get github.com/lhcb-org/lbpkr
	#go get github.com/mattn/go-sqlite3

generate:
	bash generate.sh
