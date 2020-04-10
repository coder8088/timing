package store

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/tedux/timing/pkg/model"
)

const (
	timingInsertSQL                    = "INSERT INTO timing(action,start_time,dt) VALUES(?,?,?)"
	timingGetSQL                       = "SELECT id,action,duration_seconds,start_time,stop_time,dt FROM timing WHERE id=?"
	timingListSql                      = "SELECT id,action,duration_seconds,start_time,stop_time,dt FROM timing WHERE stop_time IS NOT NULL "
	timingListByActionNameSQL          = "SELECT id,action,duration_seconds,start_time,stop_time,dt FROM timing WHERE stop_time IS NOT NULL AND action=?"
	timingListByDtSQL                  = "SELECT id,action,duration_seconds,start_time,stop_time,dt FROM timing WHERE stop_time IS NOT NULL AND dt=?"
	timingListByActionNameAndDtSQL     = "SELECT id,action,duration_seconds,start_time,stop_time,dt FROM timing WHERE stop_time IS NOT NULL AND action=? AND dt=?"
	timingUpdateStopTimeAndDurationSQL = "UPDATE timing SET stop_time=?,duration_seconds=? WHERE id=?"
	timingDeleteSQL                    = "DELETE FROM timing WHERE id=?"
)

type Timing interface {
	Insert(action *model.Timing) (id int64, err error)
	Get(id int64) (action *model.Timing, err error)
	List() (list []*model.Timing, err error)
	ListByActionName(actionName string) (list []*model.Timing, err error)
	ListByDt(dt string) (list []*model.Timing, err error)
	ListByActionNameAndDt(actionName, dt string) (list []*model.Timing, err error)
	UpdateStopTimeAndDuration(id int64, stopTime time.Time, duration int) (err error)
	Delete(id int64) (err error)
}

type timing struct {
	db *sql.DB
}

func NewTiming(db *sql.DB) Timing {
	return timing{db: db}
}

func (t timing) Insert(action *model.Timing) (id int64, err error) {
	stmt, err := t.db.Prepare(timingInsertSQL)
	if err != nil {
		return id, err
	}

	res, err := stmt.Exec(action.Name, action.StartTime, action.Dt)
	if err != nil {
		return id, err
	}

	return res.LastInsertId()
}

func (t timing) Get(id int64) (*model.Timing, error) {
	var action model.Timing
	row := t.db.QueryRow(timingGetSQL, id)
	if err := row.Scan(&action.Id, &action.Name, &action.DurationSeconds, &action.StartTime, &action.StopTime, &action.Dt); err != nil {
		return nil, err
	}
	return &action, nil
}

func (t timing) List() (list []*model.Timing, err error) {
	rows, err := t.db.Query(timingListSql)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var action model.Timing
		if err = rows.Scan(&action.Id, &action.Name, &action.DurationSeconds, &action.StartTime, &action.StopTime, &action.Dt); err != nil {
			return nil, err
		}
		list = append(list, &action)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return
}

func (t timing) ListByActionName(actionName string) (list []*model.Timing, err error) {
	rows, err := t.db.Query(timingListByActionNameSQL, actionName)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var action model.Timing
		if err = rows.Scan(&action.Id, &action.Name, &action.DurationSeconds, &action.StartTime, &action.StopTime, &action.Dt); err != nil {
			return nil, err
		}
		list = append(list, &action)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return
}

func (t timing) ListByDt(dt string) (list []*model.Timing, err error) {
	rows, err := t.db.Query(timingListByDtSQL, dt)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var action model.Timing
		if err = rows.Scan(&action.Id, &action.Name, &action.DurationSeconds, &action.StartTime, &action.StopTime, &action.Dt); err != nil {
			return nil, err
		}
		list = append(list, &action)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return
}

func (t timing) ListByActionNameAndDt(actionName, dt string) (list []*model.Timing, err error) {
	rows, err := t.db.Query(timingListByActionNameAndDtSQL, actionName, dt)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var action model.Timing
		if err = rows.Scan(&action.Id, &action.Name, &action.DurationSeconds, &action.StartTime, &action.StopTime, &action.Dt); err != nil {
			return nil, err
		}
		list = append(list, &action)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return
}

func (t timing) UpdateStopTimeAndDuration(id int64, stopTime time.Time, duration int) (err error) {
	stmt, err := t.db.Prepare(timingUpdateStopTimeAndDurationSQL)
	if err != nil {
		return err
	}

	rs, err := stmt.Exec(stopTime, duration, id)
	if err != nil {
		return err
	}
	num, err := rs.RowsAffected()
	if err != nil {
		return err
	} else if num != 1 {
		return fmt.Errorf("fail to update action stop_time and duration by id %v", id)
	}
	return
}

func (t timing) Delete(id int64) (err error) {
	rs, err := t.db.Exec(timingDeleteSQL, id)
	if err != nil {
		return err
	}
	num, err := rs.RowsAffected()
	if err != nil {
		return err
	} else if num != 1 {
		return fmt.Errorf("fail to delete action by id %v", id)
	}
	return
}
