package median

import (
	"github.com/g-leon/go-running-median/dataprovider"
	"github.com/ancientlore/go-avltree"
	"github.com/pkg/errors"
	"github.com/zfjagann/golang-ring"
	"io"
	"strconv"
	"log"
)

// stringSliceWriter writes a slice of strings
type stringSliceWriter interface{
	Write([]string) error
}

// int comparator for the tree based multiset
func compareInt(a interface{}, b interface{}) int {
	if a.(int) < b.(int) {
		return -1
	} else if a.(int) > b.(int) {
		return 1
	}
	return 0
}

// Get uses a dataprovider.IntFetcher to read one value
// at each step and write the median of the current data set
// in a csv file located at the given path.
// At each step no more than windowSize elements will be kept
// in memory.
func Get(dp dataprovider.IntFetcher, w stringSliceWriter, windowSize int) error {
	set := avltree.New(compareInt, avltree.AllowDuplicates)
	_ring := &ring.Ring{}
	_ring.SetCapacity(windowSize)

	log.Println("Running median...")

	for {
		val, err := dp.Fetch()
		if errors.Cause(err) == io.EOF {
			log.Println("Finished running median.")
			return nil
		}
		if err != nil {
			return errors.Wrap(err,"unable to fetch value")
		}

		set.Add(val)
		_ring.Enqueue(val)

		if set.Len() == 1 {
			w.Write([]string{"-1"})
		} else {
			switch set.Len() % 2 {
			case 1:
				err = w.Write([]string{strconv.Itoa(set.At(set.Len()/2).(int))})
			case 0:
				err = w.Write([]string{strconv.FormatFloat(float64(set.At(set.Len()/2-1).(int) +
					set.At(set.Len()/2).(int))/2, 'f', -1, 64)})
			}

			if err != nil {
				return errors.Wrap(err, "unable to write value")
			}
		}

		// If window is full then remove last inserted
		// element in order to make room for the next one.
		if len(_ring.Values()) == _ring.Capacity() {
			last := _ring.Dequeue()
			set.Remove(last)
		}
	}

	return nil
}
