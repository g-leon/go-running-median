package median

import (
	"testing"
	"os"
	"encoding/csv"
	"github.com/g-leon/go-running-median/dataprovider"
	"strconv"
)

func TestGet(t *testing.T) {
	// create input file
	fi, err := os.Create("input.csv")
	defer os.Remove(fi.Name())
	if err != nil {
		t.Error(err)
	}
	wi := csv.NewWriter(fi)
	if err != nil {
		t.Error(err)
	}
	wi.Write([]string{"3"})
	wi.Write([]string{"2"})
	wi.Write([]string{"1"})
	wi.Flush()

	// feed input file to data provider
	dp := dataprovider.NewCsvInts(fi.Name())

	// create output file
	fo, err := os.Create("output.csv")
	defer fo.Close()
	defer os.Remove(fo.Name())
	if err != nil {
		t.Error(err)
	}
	wo := csv.NewWriter(fo)
	if err != nil {
		t.Error(err)
	}

	// get medians
	err = Get(dp, wo, 3)
	if err != nil {
		t.Error(err)
	}
	wo.Flush()

	// read output file and check results
	fo.Seek(0, 0)
	r := csv.NewReader(fo)

	// return -1 if there is only one element in the window
	rec, err := r.Read()
	if err != nil {
		t.Error(err)
	}
	nr, err := strconv.ParseFloat(rec[0], 0)
	if err != nil {
		t.Error(err)
	}
	if nr != -1 {
		t.Errorf("Expected -1 got %v", rec)
	}

	// return average if nr of elements in the window is even
	rec, err = r.Read()
	if err != nil {
		t.Error(err)
	}
	nr, err = strconv.ParseFloat(rec[0],0)
	if err != nil {
		t.Error(err)
	}
	if nr != 2.5 {
		t.Errorf("Expected 2.5 got %v", rec)
	}

	// return median if nr of elements in the window is odd
	rec, err = r.Read()
	if err != nil {
		t.Error(err)
	}
	nr, err = strconv.ParseFloat(rec[0],0)
	if err != nil {
		t.Error(err)
	}
	if nr != 2 {
		t.Errorf("Expected 2 got %v", rec)
	}
}
