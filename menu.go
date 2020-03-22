package ledger

import (
	"fmt"
	"log"
)

// Menu contains items for select and selected items.
type Menu struct {
	Items    []Item
	Selected []int
}

type Item interface {
	ColumnName() string
	List() []string
}

// CreateMenu returns menu struct.
func CreateMenu() *Menu {
	m := new(Menu)
	m.Items = make([]Item, 0)
	m.Selected = make([]int, 0)
	return m
}

// SelectMenu returns input form.
func (m *Menu) SelectMenu() {
	for i, item := range m.Items {
		var sel int
		temp := item.List()
		for j, v := range temp {
			fmt.Printf("%v.%v ", j, v)
		}
		fmt.Printf("\nSelect [%v]: ", m.Items[i].ColumnName())
		if _, err := fmt.Scan(&sel); err != nil {
			log.Fatal(err)
		}
		m.Selected = append(m.Selected, sel)
	}
}

// Register regist the data to the ledger.
//func (l *Ledger) Register()

// AddItem adds a item to the menu.
func (m *Menu) AddItem(i ...Item) {
	for _, v := range i {
		m.Items = append(m.Items, v)
	}
}
