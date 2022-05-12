package vface

type IRow interface {
	IModel
	Cells() []IModel
}

type Row struct {
	Model
	CellModels []IModel
}

var _ IRow = &Row{}

func (row *Row) Cells() []IModel {
	return row.CellModels
}

// Optional interface for models denoting cells
type ICell interface {
	IModel
	IsHeader() bool
}

type ITable interface {
	IModel
	Rows() []IRow
}

type Table struct {
	Model
	RowModels []IRow
}

var _ ITable = &Table{}

func (tabke *Table) Rows() []IRow {
	return tabke.RowModels
}

type RowView struct{ View[IRow] }

func (c *RowView) isHeader(model IModel) bool {
	if m2, ok := model.(ICell); ok {
		return m2.IsHeader()
	}
	return false
}

type TableView struct{ View[ITable] }
