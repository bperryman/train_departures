package traindeps

import (
	"fmt"
	"time"
)

// Go departure board...

// TrainStatus defines that status that the departing train is in.
// The values are limited to: OnTime, Delayed, Cancelled, Late and Unknown
type TrainStatus int

const (
	// OnTime - the train is expect at it's scheduled departure time
	OnTime TrainStatus = iota + 1

	// Delayed - there is a delay in the departure, no further details about the delay
	Delayed

	// Cancelled - the train has been cancelled
	Cancelled

	// Late - the train is delayed and is running late by a known amout of time
	Late

	// Unknown - No know information about the train exists
	Unknown
)

func (ts TrainStatus) String() string {
	switch ts {
	case OnTime:
		return "On Time"
	case Delayed:
		return "Delayed"
	case Cancelled:
		return "Cancelled"
	case Late:
		return "Late"
	default:
		return "Unknown"
	}
}

// Board represents a snapshot in time of the departure board at a station.
// Station is the statation the departure board is for.
// AsAt is the time the snapshot was taken.
// Departures holds the departing trains.
type Board struct {
	Station    string
	AsAt       time.Time
	Departures []Departure
}

// Departure represents a single train departure.
// Destination is the destination of the train
// Platform is the platform the train will be departing from - will be blank until a platform is assigned.
// DepartureTime is the scheduled departure time for the train.
// ExpectedTime is the actual time the trains is expected to leave.
// LateByMins is the delay in minutes of the train when the status is late
// Status is the status of the train.
type Departure struct {
	Destination   string
	Platform      string
	DepartureTime time.Time
	ExpectedTime  time.Time
	LateByMins    int
	Status        TrainStatus
}

func (d Departure) String() string {
	return fmt.Sprintf("%s - %s - %s - %02d:%02d - %02d:%02d(%d mins late)",
		d.Destination, d.Status, d.Platform,
		d.DepartureTime.Hour(), d.DepartureTime.Minute(),
		d.ExpectedTime.Hour(), d.ExpectedTime.Minute(),
		d.LateByMins)
}
