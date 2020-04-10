package store

import (
	"database/sql"
	"fmt"

	"github.com/tedux/timing/pkg/model"
)

const (
	actionInsertSQL        = "INSERT INTO action(name,group_id) VALUES(?,?)"
	actionGetSQL           = "SELECT id,name,group_id FROM action WHERE id=?"
	actionListSQL          = "SELECT id,name,group_id FROM action"
	actionListByGroupIdSQL = "SELECT id,name,group_id FROM action WHERE group_id=?"
	actionDeleteSQL        = "DELETE FROM action WHERE id=?"
	actionUpdateGroupIdSQL = "UPDATE action SET group_id=? WHERE id=?"
)

type Action interface {
	Insert(action *model.Action) (id int64, err error)
	Get(id int64) (action *model.Action, err error)
	List() (actions []*model.Action, err error)
	ListByGroupId(groupId int64) (actions []*model.Action, err error)
	Delete(id int64) (err error)
	UpdateGroupId(id, groupId int64) (err error)
}

type action struct {
	db *sql.DB
}

func NewAction(db *sql.DB) Action {
	return action{db: db}
}

func (a action) Insert(action *model.Action) (id int64, err error) {
	rs, err := a.db.Exec(actionInsertSQL, action.Name, action.GroupId)
	if err != nil {
		return id, err
	}
	id, err = rs.LastInsertId()
	return
}

func (a action) Get(id int64) (*model.Action, error) {
	var action model.Action
	row := a.db.QueryRow(actionGetSQL, id)
	err := row.Scan(&action.Id, &action.Name, &action.GroupId)
	if err != nil {
		return nil, err
	}
	return &action, nil
}

func (a action) List() (actions []*model.Action, err error) {
	rows, err := a.db.Query(actionListSQL)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var action model.Action
		if err = rows.Scan(&action.Id, &action.Name, &action.GroupId); err != nil {
			return nil, err
		}
		actions = append(actions, &action)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return
}

func (a action) ListByGroupId(groupId int64) (actions []*model.Action, err error) {
	rows, err := a.db.Query(actionListByGroupIdSQL, groupId)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var action model.Action
		if err = rows.Scan(&action.Id, &action.Name, &action.GroupId); err != nil {
			return nil, err
		}
		actions = append(actions, &action)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return
}

func (a action) Delete(id int64) (err error) {
	_, err = a.db.Exec(actionDeleteSQL, id)
	return
}

func (a action) UpdateGroupId(id, groupId int64) (err error) {
	rs, err := a.db.Exec(actionUpdateGroupIdSQL, groupId, id)
	if err != nil {
		return err
	}
	num, err := rs.RowsAffected()
	if err != nil {
		return err
	} else if num != 1 {
		return fmt.Errorf("fail to update action group_id by id[%v]", id)
	}
	return
}
