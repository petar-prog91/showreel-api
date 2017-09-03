package models

type ClassGroup struct {
	ID                int    `db:"id, omitempty"`
	ClassName         string `db:"classname, omitempty"`
	ResponsiblePerson int    `db:"responsible_person, omitempty"`
	CreationDate      string `db:"creation_date, omitempty"`
}
