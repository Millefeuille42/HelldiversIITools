package redisEvent

import (
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"encoding/json"
)

type EventType string

const (
	NewMessageEventType      EventType = "new_message"
	NewOrderEventType        EventType = "new_order"
	PlanetLostEventType      EventType = "planet_lost"
	PlanetLiberatedEventType EventType = "planet_liberated"
)

type Event interface {
	GetType() EventType
	ToJson() ([]byte, error)
}

type Generic struct {
	Type EventType
	Data string
}

func SendEvent(event Event) error {
	data, err := event.ToJson()
	if err != nil {
		return err
	}

	msg, err := json.Marshal(Generic{
		Type: event.GetType(),
		Data: string(data),
	})
	if err != nil {
		return err
	}

	return Client.Publish(Context, "events", msg).Err()
}

type NewMessageEvent struct {
	lib.NewsMessage
}

func (n NewMessageEvent) ToJson() ([]byte, error) {
	return json.Marshal(n)
}

func (n NewMessageEvent) GetType() EventType {
	return NewMessageEventType
}

type NewOrderEvent struct {
	lib.MajorOrder
}

func (n NewOrderEvent) ToJson() ([]byte, error) {
	return json.Marshal(n)
}

func (n NewOrderEvent) GetType() EventType {
	return NewOrderEventType
}

type NewPlanetEvent struct {
	lib.Planet
}

func (n NewPlanetEvent) ToJson() ([]byte, error) {
	return json.Marshal(n)
}

func (n NewPlanetEvent) GetType() EventType {
	if n.Planet.Owner != lib.Humans {
		return PlanetLostEventType
	}
	return PlanetLiberatedEventType
}
