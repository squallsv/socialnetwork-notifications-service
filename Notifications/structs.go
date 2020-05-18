package Notifications

import (
	"encoding/json"
	"errors"
	"time"
)

type EventType string

type Notification struct {
	Id string
	UserId string
	FromUserId string
	EventType EventType
	Preview string
	CreatedOn time.Time
}

const (
	NewFollower EventType = "NewFollower"
	NewDirectMessage EventType = "NewDirectMessage"
)

func (eventType *EventType) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type ET EventType
	var r *ET = (*ET)(eventType)
	err := json.Unmarshal(b, &r)
	if err != nil{
		panic(err)
	}
	switch *eventType {
	case NewFollower, NewDirectMessage:
		return nil
	}
	return errors.New("invalid event type")
}