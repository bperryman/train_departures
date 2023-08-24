package traindeps

import (
	"fmt"
	"testing"
	"time"
)

func TestCanImportValidDepartureBoard(tst *testing.T) {
	// Load up the board
	captured := time.Now()
	board, err := LoadFromFile("./test-files/simple.html", "CHX", captured)
	if err != nil {
		tst.Errorf("Received an error loading from file")
		tst.Errorf("%s", err)
	}
	if captured != board.AsAt {
		tst.Errorf("Expected AsAt date to match supplied date")
	}
	if board.Station != "CHX" {
		tst.Errorf("Expected board Station to be CHX")
	}
	if board.Departures == nil {
		tst.Errorf("No departures picked up from file")
	}
	if len(board.Departures) != 27 {
		tst.Errorf("Found %d departures, expected 27", len(board.Departures))
	}

	// Test that first departure is for Dartford via Bexleyheath
	if board.Departures[0].Destination != "Dartford via Bexleyheath" {
		for i, c := range board.Departures[0].Destination {
			fmt.Printf("%2d => %c => %x\n", i, c, c)
		}
		tst.Errorf("Expected first departure to be for Dartford via Bexleyheath - received %s", board.Departures[0].Destination)
	}
	if board.Departures[0].Platform != "6" {
		tst.Errorf("Expected first departure platform to be 6")
	}
	if board.Departures[0].DepartureTime.Hour() != 22 || board.Departures[0].DepartureTime.Minute() != 24 {
		tst.Errorf("Expected first departure time to be at 22:24")
	}
	if board.Departures[0].ExpectedTime.Hour() != 22 || board.Departures[0].ExpectedTime.Minute() != 24 {
		tst.Errorf("Expected first departure expected time to be at 22:24")
	}
	if board.Departures[0].LateByMins != 0 {
		tst.Errorf("Expected first departure to be late by 0 minutes")
	}
	if board.Departures[0].Status != OnTime {
		tst.Errorf("Expoected first departure status to be OnTime got %s", board.Departures[0].Status)
	}
}

func TestCanImportLateDepartureBoard(tst *testing.T) {
	// Load up the board
	captured := time.Now()
	board, err := LoadFromFile("./test-files/delayed-with-time.html", "CHX", captured)
	if err != nil {
		tst.Errorf("Received an error loading from file")
		tst.Errorf("%s", err)
	}
	if captured != board.AsAt {
		tst.Errorf("Expected AsAt date to match supplied date")
	}
	if board.Station != "CHX" {
		tst.Errorf("Expected board Station to be CHX")
	}
	if board.Departures == nil {
		tst.Errorf("No departures picked up from file")
	}
	if len(board.Departures) != 41 {
		tst.Errorf("Found %d departures, expected 41", len(board.Departures))
	}

	// Test that first departure is for Dartford via Bexleyheath
	if board.Departures[0].Destination != "Slade Green" {
		for i, c := range board.Departures[0].Destination {
			fmt.Printf("%2d => %c => %x\n", i, c, c)
		}
		tst.Errorf("Expected first departure to be for Slade Green - received %s", board.Departures[0].Destination)
	}
	if board.Departures[0].Platform != "" {
		tst.Errorf("Expected first departure platform to be empty")
	}
	if board.Departures[0].DepartureTime.Hour() != 16 || board.Departures[0].DepartureTime.Minute() != 15 {
		tst.Errorf("Expected first departure time to be at 16:15")
	}
	if board.Departures[0].ExpectedTime.Hour() != 16 || board.Departures[0].ExpectedTime.Minute() != 28 {
		tst.Errorf("Expected first departure expected time to be at 16:28")
	}
	if board.Departures[0].LateByMins != 13 {
		tst.Errorf("Expected first departure to be late by 13 minutes")
	}
	if board.Departures[0].Status != Late {
		tst.Errorf("Expoected first departure status to be Late got %s", board.Departures[0].Status)
	}
}

func TestCanImportDelayedDepartureBoard(tst *testing.T) {
	// Load up the board
	captured := time.Now()
	board, err := LoadFromFile("./test-files/delayed.html", "CHX", captured)
	if err != nil {
		tst.Errorf("Received an error loading from file")
		tst.Errorf("%s", err)
	}
	if captured != board.AsAt {
		tst.Errorf("Expected AsAt date to match supplied date")
	}
	if board.Station != "CHX" {
		tst.Errorf("Expected board Station to be CHX")
	}
	if board.Departures == nil {
		tst.Errorf("No departures picked up from file")
	}
	if len(board.Departures) != 42 {
		tst.Errorf("Found %d departures, expected 42", len(board.Departures))
	}

	// Test that first departure is for Dartford via Bexleyheath
	if board.Departures[0].Destination != "Slade Green" {
		for i, c := range board.Departures[0].Destination {
			fmt.Printf("%2d => %c => %x\n", i, c, c)
		}
		tst.Errorf("Expected first departure to be for Slade Green - received %s", board.Departures[0].Destination)
	}
	if board.Departures[0].Platform != "" {
		tst.Errorf("Expected first departure platform to be empty")
	}
	if board.Departures[0].DepartureTime.Hour() != 16 || board.Departures[0].DepartureTime.Minute() != 15 {
		tst.Errorf("Expected first departure time to be at 16:15")
	}
	if board.Departures[0].ExpectedTime.Hour() != 0 || board.Departures[0].ExpectedTime.Minute() != 0 {
		tst.Errorf("Expected first departure expected time to be at 00:00 got %02d:%02d",
			board.Departures[0].ExpectedTime.Hour(), board.Departures[0].ExpectedTime.Minute())
	}
	if board.Departures[0].LateByMins != 0 {
		tst.Errorf("Expected first departure to be late by 0 minutes")
	}
	if board.Departures[0].Status != Delayed {
		tst.Errorf("Expoected first departure status to be Delayed got %s", board.Departures[0].Status)
	}
}

func TestCanImportCancelledDepartureBoard(tst *testing.T) {
	// Load up the board
	captured := time.Now()
	board, err := LoadFromFile("./test-files/cancelled.html", "CHX", captured)
	if err != nil {
		tst.Errorf("Received an error loading from file")
		tst.Errorf("%s", err)
	}
	if captured != board.AsAt {
		tst.Errorf("Expected AsAt date to match supplied date")
	}
	if board.Station != "CHX" {
		tst.Errorf("Expected board Station to be CHX")
	}
	if board.Departures == nil {
		tst.Errorf("No departures picked up from file")
	}
	if len(board.Departures) != 36 {
		tst.Errorf("Found %d departures, expected 36", len(board.Departures))
	}

	// Test that first departure is for Dartford via Bexleyheath
	if board.Departures[4].Destination != "Hayes (Kent)" {
		for i, c := range board.Departures[4].Destination {
			fmt.Printf("%2d => %c => %x\n", i, c, c)
		}
		tst.Errorf("Expected first departure to be for Slade Green - received %s", board.Departures[4].Destination)
	}
	if board.Departures[4].Platform != "" {
		tst.Errorf("Expected first departure platform to be empty")
	}
	if board.Departures[4].DepartureTime.Hour() != 17 || board.Departures[4].DepartureTime.Minute() != 55 {
		tst.Errorf("Expected first departure time to be at 17:55")
	}
	if board.Departures[4].ExpectedTime.Hour() != 0 || board.Departures[4].ExpectedTime.Minute() != 0 {
		tst.Errorf("Expected first departure expected time to be at 00:00 got %02d:%02d",
			board.Departures[4].ExpectedTime.Hour(), board.Departures[4].ExpectedTime.Minute())
	}
	if board.Departures[4].LateByMins != 0 {
		tst.Errorf("Expected first departure to be late by 0 minutes")
	}
	if board.Departures[4].Status != Cancelled {
		tst.Errorf("Expoected first departure status to be Delayed got %s", board.Departures[4].Status)
	}
}

func TestCanProcessStatusText(tst *testing.T) {
	expectations := []struct {
		txt    string
		status TrainStatus
	}{
		{"On time", OnTime},
		{"Cancelled", Cancelled},
		{"Delayed", Delayed},
		{"17:02 5 mins late", Late},
		{"17:02 1 min late", Late},
	}

	for _, e := range expectations {
		if processStatusText(e.txt) != e.status {
			tst.Errorf("Failed to match %s to %d", e.txt, e.status)
		}
	}
}

func TestCanFetchCorrectDelay(tst *testing.T) {
	expectations := []struct {
		txt                string
		hours, mins, delay int
	}{
		{"On time", 0, 0, 0},
		{"Cancelled", 0, 0, 0},
		{"Delayed", 0, 0, 0},
		{"17:02 5 mins late", 17, 2, 5},
		{"17:02 1 min late", 17, 2, 1},
	}

	for _, e := range expectations {
		expected, delay := expectedAndDelayed(e.txt)
		if delay != e.delay {
			tst.Errorf("Failed to match %s to %d delay", e.txt, e.delay)
		}
		if expected.Hour() != e.hours {
			tst.Errorf("Failed to match %d hours to %d", expected.Hour(), e.hours)
		}
		if expected.Minute() != e.mins {
			tst.Errorf("Failed to match %d minutes to %d", expected.Minute(), e.mins)
		}
	}
}
