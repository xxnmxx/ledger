package main

import (
	"fmt"

	"github.com/xxnmxx/ledger"
)

func main() {
	a := ledger.CreateAccMaster()
	a.AddAccMaster("cash", ledger.AccClass[0])
	a.AddAccMaster("land", ledger.AccClass[1])
	a.AddAccMaster("ap", ledger.AccClass[1])
	fmt.Printf("%+v\n", a)
	//a.WriteAccMaster("AccMaster")
	//m2 := ledger.LoadAccMaster("AccMaster")
	//fmt.Printf("%+v", m2)
	m := ledger.CreateMenu()
	s := ledger.CreateSubMaster()
	s.AddSubMaster(a, "mizuho")
	m.AddItem(a, s)
	m.SelectMenu()
	fmt.Println(m.Selected)
}
