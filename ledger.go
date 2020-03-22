package ledger

import (
	"fmt"
	"strings"
	"time"
)

type Ledger struct {
	Header    []string
	Documents []Document
}

type Document struct {
	date     string
	no       int
	acc      []string
	sub      []string
	div      []string
	taxin    []bool
	taxClass []string
	//taxRate  []taxRate
	amt  []float64
	note []string
}

func (l *Ledger) InputDate() string {
	var d string
	layout := "2006-01-02"
	for {
		fmt.Print("Input [date] (YYYY-MM-DD): ")
		_, err := fmt.Scan(&d)
		if err != nil {
			fmt.Println("Invalid input.")
			continue
		}
		// Validate
		t, err := time.Parse(layout, d)
		if err != nil {
			fmt.Println("Invalid. Do not forget -(hyphen).")
			continue
		}
		tString := t.String()
		tSplit := strings.Split(tString, " ")
		return tSplit[0]
	}
}

func (l *Ledger) no() int {
	return len(l.Documents)
}

func (l *Ledger) InputAcc(a *AccMaster) string {
	var idx int
	isEmpty := false
	for {
		if len(a.accName) == 0 {
			fmt.Println("Empty master.")
			isEmpty = true
			break
		}
		for i, v := range a.accName {
			fmt.Printf("%v.%v ", i, v)
		}
		fmt.Printf("\nSelect [acc]: ")
		_, err := fmt.Scan(&idx)
		if err != nil {
			fmt.Println("Invalid.")
			continue
		}
		if idx > len(a.accName) {
			fmt.Println("Select from list above.")
			continue
		}
		break
	}
	if !isEmpty {
		return a.accName[idx]
	}
	return ""
}

type Master interface {
	ColumnName() string
	List() []string
	Len() int
}

func (l *Ledger) InputMaster(m Master) string {
	var idx int
	isEmpty := false
	list := m.List()
	for {
		if m.Len() == 0 {
			fmt.Println("Empty master.")
			isEmpty = true
			break
		}
		for i, v := range list {
			fmt.Printf("%v.%v ", i, v)
		}
		fmt.Printf("\nSelect [%v]: ", m.ColumnName())
		_, err := fmt.Scan(&idx)
		if err != nil {
			fmt.Println("Invalid.")
			continue
		}
		if idx > m.Len() {
			fmt.Println("Select from list above.")
			continue
		}
		break
	}
	if !isEmpty {
		return list[idx]
	}
	return ""
}
