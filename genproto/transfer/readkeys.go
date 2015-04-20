package main
import "github.com/h4ck3rm1k3/taginfo/database/taginfo"
import "github.com/h4ck3rm1k3/capnproto-gen/transfer"

func main() {
	var prototype taginfo_model.Keys	
	capnprotogen_transfer.Generate(prototype)
}
