package folio

import (
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type SocialLink struct {
	db.Record
	Icon    string   `orm:"size(25)"`
	URL     string   `orm:"size(128)"`
	Profile *Profile `orm:"rel(fk)" json:",omitempty"`
}

func (o SocialLink) Validate() (bool, error) {
	return util.ValidateStruct(&o)
}
