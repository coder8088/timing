package store

import (
	"database/sql"

	"github.com/tedux/timing/pkg/model"
)

const (
	actionGroupInsertSQL = "INSERT INTO action_group(name) VALUES(?)"
	actionGroupGetSQL    = "SELECT id,name FROM action_group WHERE id=?"
	actionGroupListSQL   = "SELECT id,name FROM action_group"
	actionGroupDeleteSQL = "DELETE FROM action_group WHERE id=?"
)

type ActionGroup interface {
	Insert(actionGroup *model.ActionGroup) (id int64, err error)
	Get(id int64) (*model.ActionGroup, error)
	List() (actionGroups []*model.ActionGroup, err error)
	Delete(id int64) error
}

type actionGroup struct {
	db *sql.DB
}

func NewActionGroup(db *sql.DB) ActionGroup {
	return actionGroup{db: db}
}

func (a actionGroup) Insert(actionGroup *model.ActionGroup) (id int64, err error) {
	rs, err := a.db.Exec(actionInsertSQL, actionGroup.Name)
	if err != nil {
		return id, err
	}
	id, err = rs.LastInsertId()
	return
}

func (a actionGroup) Get(id int64) (*model.ActionGroup, error) {
	var actionGroup model.ActionGroup
	row := a.db.QueryRow(actionGroupGetSQL, id)
	err := row.Scan(&actionGroup.Id, &actionGroup.Name)
	if err != nil {
		return nil, err
	}
	return &actionGroup, nil
}

func (a actionGroup) List() (actionGroups []*model.ActionGroup, err error) {
	rows, err := a.db.Query(actionGroupListSQL)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var actionGroup model.ActionGroup
		if err = rows.Scan(&actionGroup.Id, &actionGroup.Name); err != nil {
			return
		}
		actionGroups = append(actionGroups, &actionGroup)
	}

	err = rows.Err()
	return
}

func (a actionGroup) Delete(id int64) error {
	_, err := a.db.Exec(actionGroupDeleteSQL, id)
	return err
}
