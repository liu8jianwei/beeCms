package models

type MenuModel struct {
	Mid     int `orm:"pk;auto"`
	Parent  int
	Name    string `orm:"size(45)"`
	Seq     int
	Formatt string `orm:"size(2048);default({})"`
}

type MenuTree struct {
	MenuModel
	child []MenuModel
}

func (m *MenuModel) TableName() string {
	return "xcms_menu"
}
