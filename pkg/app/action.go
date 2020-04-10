package app

import (
	"github.com/tedux/timing/pkg/db"
	"github.com/tedux/timing/pkg/model"
	"github.com/tedux/timing/pkg/store"
)

type Action interface {
	AddAction(action *model.Action) (id int64, err error)
	GetAction(id int64) (*model.Action, error)
	ListActions() (actions []*model.Action, err error)
	ListActionsByGroupId(groupId int64) (actions []*model.Action, err error)
	UpdateGroupId(id, groupId int64) error
	DeleteAction(id int64) error
}

type action struct {
	store store.Action
}

func NewAction() Action {
	return action{store: store.NewAction(db.SQLite)}
}

func (a action) AddAction(action *model.Action) (id int64, err error) {
	return a.store.Insert(action)
}

func (a action) GetAction(id int64) (*model.Action, error) {
	return a.store.Get(id)
}

func (a action) ListActions() (actions []*model.Action, err error) {
	return a.store.List()
}

func (a action) ListActionsByGroupId(groupId int64) (actions []*model.Action, err error) {
	return a.store.ListByGroupId(groupId)
}

func (a action) UpdateGroupId(id, groupId int64) error {
	return a.store.UpdateGroupId(id, groupId)
}

func (a action) DeleteAction(id int64) error {
	return a.store.Delete(id)
}
