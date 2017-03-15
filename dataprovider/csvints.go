package dataprovider

import (
	"encoding/csv"
	"os"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

type csvInts struct {
	r *csv.Reader
}

func NewCsvInts(path string) *csvInts {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	ci := &csvInts{}
	ci.r = csv.NewReader(f)

	return ci
}

func (ci *csvInts) Fetch() (int, error) {
	record, err := ci.r.Read()
	if err != nil {
		return -1, errors.Wrap(err, "failed to read value")
	}
	if len(record) != 1 {
		return -1, errors.New("record contains more than one int")
	}
	nr, err := strconv.ParseInt(strings.TrimRight(record[0], "\r"), 10, 0)
	if err != nil {
		return -1, errors.Wrap(err, "failed to parse int")
	}
	return int(nr), nil
}
