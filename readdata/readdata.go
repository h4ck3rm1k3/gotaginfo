package main
import "github.com/h4ck3rm1k3/taginfo/database/taginfo"
//import "log"
import "os"
import "io"
import "bytes"
import "fmt"
import "github.com/h4ck3rm1k3/taginfo/capnproto"
import capn "github.com/glycerine/go-capnproto"
//import "encoding/hex"

func main() {
	file, err := os.Open("key.data")
	if err != nil {fmt.Println(err); os.Exit(-1); }
	defer file.Close()
	var record taginfo_model.Keys
	var buf bytes.Buffer
	for {
		seg, err := capn.ReadFromStream(file, &buf)
//		fmt.Println(hex.EncodeToString(buf.Bytes()))

		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err); os.Exit(-1); }
		}
		
		binrecord := model.ReadRootKeys(seg)
//		fmt.Println(binrecord)
//		fmt.Printf("BinRecord key %v\n", binrecord.Key())
//		fmt.Printf("Record count all %v\n", binrecord.Countall())
		model.Get(&binrecord, &record)
		fmt.Printf("Record %v\n", record)
		//record.Key
		//fmt.Printf("Record key %v\n", record.Key)
	}
}
