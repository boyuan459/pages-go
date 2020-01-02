package casbin

import (
	"pages/component"

	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
	"github.com/jinzhu/gorm"
)

type Casbin struct {
	gorm.Model
	Ptype    string `json:"ptype"`
	RoleName string `json:"rolename"`
	Path     string `json:"path"`
	Method   string `json:"method"`
}

func Enforcer() *casbin.Enforcer {
	adapter := gormadapter.NewAdapterByDB(component.DB)
	e := casbin.NewEnforcer("config/auth_model.conf", adapter)
	e.LoadPolicy()
	return e
}

func (c *Casbin) AddPolicy(cb Casbin) bool {
	e := Enforcer()
	return e.AddPolicy(cb.RoleName, cb.Path, cb.Method)
}
