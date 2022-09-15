package u_sql

type Orm struct {
	Table           string
	Columns         []Column
	PrimaryKeyCount int
}

type Column struct {
	Field         string
	PrimaryKey    bool
	AutoIncrement bool
	NotNull       bool
	Unique        bool
	Type          string
	Value         interface{}
	ValueDefault  interface{}
}

func (c *Column) isPKandAI() bool {
	return c.PrimaryKey && c.AutoIncrement
}
