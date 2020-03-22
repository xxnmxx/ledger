// ToDo Make other headers.

package ledger

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type class string

type column string

var columnName = []column{
	0: "acc",
	1: "sub",
}

// *****Acc*****

// AccClass contains classes of the acc.
var AccClass = []class{
	0:  "currentAsset",
	1:  "fixedAsset",
	2:  "currentLiability",
	3:  "fixedLiability",
	4:  "equit",
	5:  "sales",
	6:  "cogs",
	7:  "sga",
	8:  "nopinc",
	9:  "nopexp",
	10: "specinc",
	11: "specexp",
	12: "tax",
}

// AccMaster contains names and classes of accs.
type AccMaster struct {
	columnName column
	accName    []string
	accClass   []class
}

// LoadAccMaster loads data from a csv file.
func LoadAccMaster(n string) *AccMaster {
	a := AccMaster{
		columnName: columnName[0],
		accName:    []string{},
		accClass:   []class{},
	}
	f, err := os.Open(n)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		a.accName = append(a.accName, rec[0])
		a.accClass = append(a.accClass, class(rec[1]))
	}
	return &a
}

// WriteAccMaster writes current AccMaster to csv a file.
func (a *AccMaster) WriteAccMaster(n string) {
	f, err := os.Create(n)
	if err != nil {
		log.Fatal(err)
	}
	w := csv.NewWriter(f)
	prep := a.transformDim()
	for _, record := range prep {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv", err)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func (a *AccMaster) transformDim() [][]string {
	tfm := make([][]string, 0)
	//temp := make([]string, 0)
	for i, n := range a.accName {
		temp := []string{}
		temp = append(temp, n, string(a.accClass[i]))
		tfm = append(tfm, temp)
	}
	return tfm
}

// CreateAccMaster returns accMaster.
func CreateAccMaster() *AccMaster {
	return &AccMaster{
		columnName: columnName[0],
		accName:    []string{},
		accClass:   []class{},
	}
}

// AddAccMaster adds name and accClass as master data.
func (a *AccMaster) AddAccMaster(n string, c class) {
	uniq := true
	for _, v := range a.accName {
		if n == v {
			uniq = false
		}
	}
	if uniq {
		a.accName = append(a.accName, n)
		a.accClass = append(a.accClass, c)
	} else {
		fmt.Printf("%v has already existed in the master.\n", n)
	}
}

// CheckShape returns whether the shape is ok or not.
func (a *AccMaster) CheckShape() bool {
	return len(a.accName) == len(a.accClass)
}

// ColumnName returns the name of the column.
func (a *AccMaster) ColumnName() column {
	return a.columnName
}

// List returns the list of the item.
func (a *AccMaster) List() []string {
	list := make([]string, 0)
	for _, v := range a.accName {
		list = append(list, v)
	}
	return list
}

// *****Sub*****
type SubMaster struct {
	columnName column
	subName    []string
	atrbAcc    []string
}

// LoadSubMaster loads data from a csv file.
func LoadSubMaster(n string) *SubMaster {
	s := SubMaster{
		columnName: columnName[1],
		subName:    []string{},
		atrbAcc:    []string{},
	}
	f, err := os.Open(n)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		s.subName = append(s.subName, rec[0])
		s.atrbAcc = append(s.atrbAcc, rec[1])
	}
	return &s
}

// WriteSubMaster writes current SubMaster to csv a file.
func (s *SubMaster) WriteSubMaster(n string) {
	f, err := os.Create(n)
	if err != nil {
		log.Fatal(err)
	}
	w := csv.NewWriter(f)
	prep := s.transformDim()
	for _, record := range prep {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv", err)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func (s *SubMaster) transformDim() [][]string {
	tfm := make([][]string, 0)
	for i, n := range s.subName {
		temp := []string{}
		temp = append(temp, n, string(s.atrbAcc[i]))
		tfm = append(tfm, temp)
	}
	return tfm
}

// CreateSubMaster returns SubMaster.
func CreateSubMaster() *SubMaster {
	return &SubMaster{
		columnName: columnName[1],
		subName:    []string{},
		atrbAcc:    []string{},
	}
}

// AddSubMaster adds data.
func (s *SubMaster) AddSubMaster(a *AccMaster, n string) {
	uniq := true
	for _, v := range s.subName {
		if n == v {
			uniq = false
		}
	}
	if uniq {
		s.subName = append(s.subName, n)
		for i, v := range a.accName {
			fmt.Printf("%v.%v ", i, v)
		}
		var idx int
		fmt.Printf("\nSelect attrAcc: ")
		fmt.Scan(&idx)
		fmt.Print("\n")
		s.atrbAcc = append(s.atrbAcc, a.accName[idx])
	} else {
		fmt.Printf("%v has already existed in the master.\n", n)
	}
}

// CheckShape returns whether the shape is ok or not.
func (s *SubMaster) CheckShape() bool {
	return len(s.subName) == len(s.atrbAcc)
}

// ColumnName returns the name of the column.
func (s *SubMaster) ColumnName() column {
	return s.columnName
}

// List returns the list of the item.
func (s *SubMaster) List() []string {
	list := make([]string, 0)
	for _, v := range s.subName {
		list = append(list, v)
	}
	return list
}

// *****Div*****

//type rec struct {
//date     string
//no       uint
//acc      acc
//sub      sub
//div      div
//taxin    bool
//taxclass taxclass
//taxrate  taxrate
//amt      float64
//note     string
//}
