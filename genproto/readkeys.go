package main
import "github.com/h4ck3rm1k3/taginfo/database/taginfo"
import "github.com/h4ck3rm1k3/capnproto-gen"

func main() {
	var prototype taginfo_model.Keys	
	captnprotogen.Generate(prototype,"@0xbbfb7b1b11adeed6")
}
