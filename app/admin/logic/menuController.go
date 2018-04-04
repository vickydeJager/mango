package logic

import (
	"strings"

	"github.com/louisevanderlith/mango/util/control"
)

type MenuController struct {
	control.UIController
}

func (ctrl *MenuController) Prepare() {
	ctrl.Data["menu"] = getMenu("/")
	ctrl.UIController.Prepare()
}

func (ctrl *MenuController) Setup(name string) {
	ctrl.UIController.Setup(name)
}

type Menu map[string]*menuItem

type menuItem struct {
	State     bool
	Text      string
	Path      string
	IconClass string
	Children  Menu
}

var _path string

func getMenu(path string) Menu {
	_path = path

	return getItems()
}

func getItems() Menu {
	result := make(Menu)

	homeKey, homeItem := newItem("Home", "/", "fa-home")
	result[homeKey] = homeItem
	result[homeKey].setActive()

	artifactKey, artifactItem := artifactMenu()
	result[artifactKey] = artifactItem

	commsKey, commsItem := commsMenu()
	result[commsKey] = commsItem

	folioKey, folioItem := folioMenu()
	result[folioKey] = folioItem

	userKey, userItem := userMenu()
	result[userKey] = userItem

	return result
}

func artifactMenu() (shortName string, result *menuItem) {
	shortName, result = newItem("Artifact API", "#", "fa-ban")

	uplKey, uplItem := newItem("Uploads", "/uploads", "fa-ban")
	result.Children[uplKey] = uplItem

	result.setActive()

	return shortName, result
}

func commsMenu() (shortName string, result *menuItem) {
	shortName, result = newItem("Comms API", "#", "fa-mail")

	msgKey, msgItem := newItem("Messages", "/comms", "fa-mail")
	result.Children[msgKey] = msgItem

	result.setActive()

	return shortName, result
}

func folioMenu() (shortName string, result *menuItem) {
	shortName, result = newItem("Folio API", "/site", "fa-web")

	result.setActive()

	return shortName, result
}

func userMenu() (shortName string, result *menuItem) {
	shortName, result = newItem("Secure API", "#", "fa-lock")

	usrKey, usrItem := newItem("Users", "/user", "fa-user")
	result.Children[usrKey] = usrItem

	result.setActive()

	return shortName, result
}

func newItem(text, path, iconClass string) (shortName string, result *menuItem) {
	shortName = getUniqueName(text)
	result = &menuItem{
		Text:      text,
		Path:      path,
		IconClass: iconClass,
		State:     false,
		Children:  make(Menu),
	}

	return shortName, result
}

func (item *menuItem) setActive() {
	if item.Path == _path {
		item.State = true
	}

	for _, v := range item.Children {
		if v.Path == _path {
			v.State = true
			item.State = true
		}
	}
}

func getUniqueName(raw string) string {
	return strings.ToLower(strings.Replace(raw, " ", "", -1))
}
