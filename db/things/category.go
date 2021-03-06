package things

import (
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Category struct {
	db.Record
	Name          string        `orm:"size(50)"`
	Description   string        `orm:"size(255)"`
	SubCategories Subcategories `orm:"reverse(many)"`
}

func (o Category) Validate() (bool, error) {
	return util.ValidateStruct(&o)
}
