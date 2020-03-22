package main

import (
	"github.com/xxnmxx/ledger"
)

func main() {
	l := new(ledger.Ledger)
	a := ledger.LoadAccMaster("AccMaster")
	s := ledger.LoadSubMaster("SubMaster")
	d := ledger.LoadDivMaster("DivMaster")
	l.InputMaster(a)
	l.InputMaster(s)
	l.InputMaster(d)
}
