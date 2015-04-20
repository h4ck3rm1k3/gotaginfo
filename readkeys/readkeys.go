package main
import "github.com/h4ck3rm1k3/taginfo/database/taginfo"
import "upper.io/db/sqlite"
import "upper.io/db"
import "log"
import "fmt"
import "github.com/h4ck3rm1k3/captnproto-gen"
import model "github.com/h4ck3rm1k3/taginfo/captnproto"
import capn "github.com/glycerine/go-capnproto"

func main() {

	var prototype taginfo_model.Keys
	
	captnprotogen.Generate(prototype)
	
	var records []taginfo_model.Keys
	
	var settings = sqlite.ConnectionURL{
		Database: taginfo_model.DBName,
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

	// Printing to stdout.
	for i, record := range records {
		fmt.Printf("record %n\n", i)
		fmt.Printf("Record %v\n", record)
		seg := capn.NewBuffer(nil)
		church := models.NewChurch(seg)
		
		
	}
	
}
