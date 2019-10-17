package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	datadir = flag.String("datadir", "", "Data directory for LevelDB")
)

func main() {
	flag.Parse()
	if *datadir == "" {
		log.Fatal("datadir empty")
	}

	db, err := leveldb.OpenFile(*datadir, nil)
	if err != nil {
		log.Fatal("failed to open db: ", err)
	}
	defer db.Close()

	iter := db.NewIterator(nil, nil)
	if err != nil {
		log.Fatal("failed to get iterator: ", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	defer w.Flush()

	for iter.Next() {
		fmt.Fprintf(w, "%s \t %s\n", string(iter.Key()), string(iter.Value()))
	}
	iter.Release()
	if err = iter.Error(); err != nil {
		log.Fatal("failed to release iterator: ", err)
	}
}
