package main
import (
	"compress/bzip2"
	"fmt"
	"io"
//	"io/ioutil"
	"os"
//	"path/filepath"
//	"sort"
)
// based on github.com/lhcb-org/lbpkr

func Uncompress(filename string, out_filename string) error{
	fmt.Printf("decompressing %s\n", filename)
	dbfile, err := os.Create(out_filename)
	if err != nil {
		return err
	}
	defer dbfile.Close()

	var infile *os.File;
	if infile, err = os.Open(filename); err != nil {
		return err
	}
	defer infile.Close()

	r := bzip2.NewReader(infile)
	_, err = io.Copy(dbfile, r)
	return err
	
}
