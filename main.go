package main

import (
	"log"
	"github.com/g-leon/go-running-median/dataprovider"
	"os"
	"encoding/csv"
	"github.com/g-leon/go-running-median/median"
)

func main() {
	ci := dataprovider.NewCsvInts("./test2.csv")
	f, err := os.Create("./test2_result.csv")
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()
	err = median.Get(ci, w, 100)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	ci = dataprovider.NewCsvInts("./test3.csv")
	f, err = os.Create("./test3_result.csv")
	defer f.Close()
	w = csv.NewWriter(f)
	defer w.Flush()
	err = median.Get(ci, w, 1000)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	ci = dataprovider.NewCsvInts("./test4.csv")
	f, err = os.Create("./test4_result.csv")
	defer f.Close()
	w = csv.NewWriter(f)
	defer w.Flush()
	err = median.Get(ci, w, 10000)
	if err != nil {
		log.Fatalf("%+v", err)
	}
}
