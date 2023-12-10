package logic

import (
	"task11/structs"
	"time"
)

type EventsModel interface {
	Create(eNew structs.EventNoId) (structs.Event, error)
	SelectById(id structs.EventID) (structs.Event, error)
	SelectBetweenDates(start, end time.Time) ([]structs.Event, error)
	Update(e structs.Event) (structs.Event, error)
	Delete(id structs.EventID) error
}

type EventAPI struct {
	m EventsModel
}

func NewEventAPI(m EventsModel) *EventAPI {
	return &EventAPI{
		m: m,
	}
}

func (api *EventAPI) Create(eNew structs.EventNoId) (structs.Event, error) {
	return api.m.Create(eNew)
}

func (api *EventAPI) Update(e structs.Event) (structs.Event, error) {
	return api.m.Update(e)
}

func (api *EventAPI) Delete(id structs.EventID) error {
	return api.m.Delete(id)
}

func (api *EventAPI) ForDay(date time.Time) ([]structs.Event, error) {
	return api.m.SelectBetweenDates(date, date)
}

func (api *EventAPI) ForWeek(date time.Time) ([]structs.Event, error) {
	const week = 7 * 24 * time.Hour
	return api.m.SelectBetweenDates(date.Add(-week), date)
}

func (api *EventAPI) ForMonth(date time.Time) ([]structs.Event, error) {
	const month = 30 * 24 * time.Hour
	return api.m.SelectBetweenDates(date.Add(-month), date)
}
