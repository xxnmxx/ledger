package ledger

import "fmt"

// TaxClassMaster hold tax classification data.
type TaxClassMaster struct {
	columnName string
	taxClass   []string
}

var taxClassMaster = TaxClassMaster{
	columnName: columnName[3],
	taxClass: []string{
		"Sales-Taxed",
		"Sales-Reduced",
		"Sales-Exempt",
		"Cost-Taxed",
		"Cost-Reduced",
		"Cost-Spec",
		"Cost-ForTaxed",
		"Cost-ForGeneral",
		"NoTax",
	},
}

// ColumnName returns the name of the column.
func (tc *TaxClassMaster) ColumnName() string {
	return tc.columnName
}

// Len returns the length.
func (tc *TaxClassMaster) Len() int {
	return len(tc.taxClass)
}

// List returns the list of the item.
func (tc *TaxClassMaster) List() []string {
	list := make([]string, 0)
	for _, v := range tc.taxClass {
		list = append(list, v)
	}
	return list
}

// TaxRateMaster holds rate data.
type TaxRateMaster struct {
	columnName string
	taxRate    []float64
}

var taxRateMaster = TaxRateMaster{
	columnName: columnName[4],
	taxRate: []float64{
		0,
		0.1,
		0.08,
		0.05,
	},
}

// ColumnName returns the name of the column.
func (tr *TaxRateMaster) ColumnName() string {
	return tr.columnName
}

// Len returns the length.
func (tr *TaxRateMaster) Len() int {
	return len(tr.taxRate)
}

// List returns the list of the item.
func (tr *TaxRateMaster) List() []string {
	list := make([]string, 0)
	for _, v := range tr.taxRate {
		s := fmt.Sprint(v)
		list = append(list, s)
	}
	return list
}
