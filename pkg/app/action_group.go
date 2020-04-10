package app

import (
	"github.com/tedux/timing/pkg/db"
	"github.com/tedux/timing/pkg/model"
	"github.com/tedux/timing/pkg/store"
)

type ActionGroup interface {
	AddActionGroup(group *model.ActionGroup) (id int64, err error)
	GetActionGroup(id int64) (*model.ActionGroup, error)
	ListActionGroups() (groups []*model.ActionGroup, err error)
	DeleteActionGroup(id int64) error
}

type actionGroup struct {
	store store.ActionGroup
}

func NewActionGroup() ActionGroup {
	return actionGroup{store: store.NewActionGroup(db.SQLite)}
}

func (a actionGroup) AddActionGroup(group *model.ActionGroup) (id int64, err error) {
	return a.store.Insert(group)
}

func (a actionGroup) GetActionGroup(id int64) (*model.ActionGroup, error) {
	return a.store.Get(id)
}

func (a actionGroup) ListActionGroups() (groups []*model.ActionGroup, err error) {
	return a.store.List()
}

func (a actionGroup) DeleteActionGroup(id int64) error {
	return a.store.Delete(id)
}
