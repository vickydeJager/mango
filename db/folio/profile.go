package folio

import (
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Profile struct {
	db.Record
	Title          string      `orm:"size(128)" json:",omitempty"`
	Description    string      `orm:"size(512)" json:",omitempty"`
	ContactEmail   string      `orm:"size(128)" json:",omitempty"`
	ContactPhone   string      `orm:"size(20)" json:",omitempty"`
	URL            string      `orm:"size(128)" json:",omitempty"`
	ImageID        int64       `orm:"null"`
	StyleSheet     string      `orm:"size(50)"`
	SocialLinks    SocialLinks `orm:"reverse(many)" json:",omitempty"`
	PortfolioItems Portfolios  `orm:"reverse(many)" json:",omitempty"`
	AboutSections  Abouts      `orm:"reverse(many)" json:",omitempty"`
	Headers        Headers     `orm:"reverse(many)" json:",omitempty"`
}

func (p Profile) Validate() (bool, error) {
	return util.ValidateStruct(&p)
}
