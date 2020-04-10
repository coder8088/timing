package app

import (
	"time"

	"github.com/tedux/timing/pkg/db"
	"github.com/tedux/timing/pkg/model"
	"github.com/tedux/timing/pkg/store"
)

type Timing interface {
	StartTiming(actionName string) (id int64, err error)
	StopTiming(id int64) error
	GetTiming(id int64) (action *model.Timing, err error)
	ListTimings() (actions []*model.Timing, err error)
	SearchTimingsByActionAndDt(actionName, dt string) (actions []*model.Timing, err error)
}

type timing struct {
	store store.Timing
}

func NewTiming() Timing {
	return timing{store: store.NewTiming(db.SQLite)}
}

func (a timing) StartTiming(actionName string) (id int64, err error) {
	now := time.Now()
	action := &model.Timing{Name: actionName, StartTime: now, Dt: now.Format("2006-01-02")}
	return a.store.Insert(action)
}

func (a timing) StopTiming(id int64) error {
	stopTime := time.Now()
	action, err := a.store.Get(id)
	if err != nil {
		return err
	}
	durationSeconds := int(stopTime.Sub(action.StartTime).Seconds())
	return a.store.UpdateStopTimeAndDuration(id, stopTime, durationSeconds)
}

func (a timing) ListTimings() (actions []*model.Timing, err error) {
	return a.store.List()
}

func (a timing) GetTiming(id int64) (action *model.Timing, err error) {
	return a.store.Get(id)
}

func (a timing) SearchTimingsByActionAndDt(actionName, dt string) (actions []*model.Timing, err error) {
	if actionName == "" && dt == "" {
		return a.ListTimings()
	}
	if actionName != "" && dt == "" {
		return a.store.ListByActionName(actionName)
	}
	if actionName == "" && dt != "" {
		return a.listTimingsByDt(dt)
	}

	return a.store.ListByActionNameAndDt(actionName, dt)
}

func (a timing) listTimingsByDt(dt string) (actions []*model.Timing, err error) {
	return a.store.ListByDt(dt)
}
