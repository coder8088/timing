package app

import (
	"github.com/tedux/timing/pkg/db"
	"github.com/tedux/timing/pkg/store"
)

type App interface {
	ActionGroup
	Action
	Timing
}

type app struct {
	actionGroup
	action
	timing
}

func New() App {
	return app{
		actionGroup: actionGroup{store.NewActionGroup(db.SQLite)},
		action:      action{store.NewAction(db.SQLite)},
		timing:      timing{store.NewTiming(db.SQLite)},
	}
}
