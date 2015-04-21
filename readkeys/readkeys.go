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

	var record taginfo_model.Keys
//	var records []taginfo_model.Keys
	
	var settings = sqlite.ConnectionURL{
		Database: "./data/" + taginfo_model.DBName,
	}
	sess, err := db.Open(sqlite.Adapter, settings)
	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}
	defer sess.Close()

	file, err := os.OpenFile("key.data", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {fmt.Println(err); os.Exit(-1); }
	defer file.Close()
	
	dataCollection, err := sess.Collection(record.TableName())
	if err != nil {
		log.Fatalf("sess.Collection(): %q : %s\n", err, 
			taginfo_model.Keys.TableName)
	}
	var res db.Result

	res = dataCollection.Find()
	
	fmt.Printf("after find\n")
	
	total, err := res.Count()
	
	if err != nil {
		fmt.Printf("error %q",err)
	}
	
	fmt.Printf("Found Records %v\n", total)
	
	for {
		err = res.Next(&record)
		if err == nil {
			//fmt.Printf("Record %v\n", record)
			seg := capn.NewBuffer(nil)
			key := model.NewRootKeys(seg)
			model.Set(record,key)
			buf := bytes.Buffer{}
			seg.WriteTo(&buf)
			file.Write(buf.Bytes())
		}else if err == db.ErrNoMoreRows {
			break
		}else{
			log.Fatalf("res.Next(): %q\n", err)
		}
		
	}
	
}
