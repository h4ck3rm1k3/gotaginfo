package main
import "github.com/h4ck3rm1k3/taginfo/database/taginfo"
//import "upper.io/db/sqlite"
//import "upper.io/db"
//import "log"
//import "fmt"
import "github.com/h4ck3rm1k3/captnproto-gen"
//import "github.com/glycerine/go-capnproto"

func main() {
	var prototype taginfo_model.Keys	
	captnprotogen.Generate(prototype,"@0xbbfb7b1b11adeed6")
}
