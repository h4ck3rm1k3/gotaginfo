package main
import "github.com/h4ck3rm1k3/taginfo/database/taginfo"
import "upper.io/db/sqlite"
import "upper.io/db"
import "log"
import "os"
import "bytes"
import "fmt"
import "github.com/h4ck3rm1k3/taginfo/capnproto"
import capn "github.com/glycerine/go-capnproto"

func main() {

	var prototype taginfo_model.Keys
	var records []taginfo_model.Keys
	
	var settings = sqlite.ConnectionURL{
		Database: "/home/h4ck3rm1k3/goroot/" + taginfo_model.DBName,
	}
	sess, err := db.Open(sqlite.Adapter, settings)
	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}
	defer sess.Close()
	
	dataCollection, err := sess.Collection(prototype.TableName())
	if err != nil {
		log.Fatalf("sess.Collection(): %q : %s\n", err, 
			taginfo_model.Keys.TableName)
	}
	var res db.Result

	res = dataCollection.Find()
	
	fmt.Printf("after find\n")
	// Query all results and fill the birthday variable with them.
	err = res.All(&records)

	fmt.Printf("after all\n")
	if err != nil {
		log.Fatalf("res.All(): %q\n", err)
	}

	file, err := os.OpenFile("key.data", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {fmt.Println(err); os.Exit(-1); }
	defer file.Close()
	

	for i, record := range records {
		fmt.Printf("record %n\n", i)
		fmt.Printf("Record %v\n", record)
		seg := capn.NewBuffer(nil)
		key := model.NewKeys(seg)
		model.Get(key,record)
		buf := bytes.Buffer{}
		seg.WriteTo(&buf)
		file.Write(buf.Bytes())
		
	}
	
}
