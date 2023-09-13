package gplus

type OrderCondition struct {
	//排序的字段名
	Columns []string
	//排序(orderBy[constants.Asc|constants.Desc]
	Order string
	//排序顺序
	OrderIndex int
}
