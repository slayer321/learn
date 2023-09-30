package server

import (
	"errors"
	"time"
)

var ErrIDNotFound = errors.New("ID not found")

type Activity struct {
	Time        time.Time `json:"time"`
	Description string    `json:"description"`
	ID          uint64    `json:"id"`
}

type Activities struct {
	Activities []Activity
}

func (a *Activities) Insert(activity Activity) uint64 {
	activity.ID = uint64(len(a.Activities))
	a.Activities = append(a.Activities, activity)
	return activity.ID
}

func (a *Activities) Retrieve(id uint64) (Activity, error) {
	if id >= uint64(len(a.Activities)) {
		return Activity{}, ErrIDNotFound
	}
	return a.Activities[id], nil
}

func (a *Activities) ListActivities() []Activity {
	return a.Activities
}
