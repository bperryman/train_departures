package train_departures

import (
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// LoadFromFile processes a National Rail HTML snapshot, stored in a file,
// and returns the departure board that it represents.
// No attempt is made to identify the station or the time the information
// was captured so these must also be supplied as additional arguments.
func LoadFromFile(fileName, station string, asAt time.Time) (*Board, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	// Extract the departures from here.
	departures, err := extractDeparturesHTML(file)
	if err != nil {
		return nil, err
	}
	board := Board{Station: station, AsAt: asAt, Departures: departures}
	err = file.Close()
	return &board, err
}

// LoadFromInternet makes a call to fetch the departure board for the specified
// station and returns the associated departure board.
func LoadFromInternet(station string) (*Board, error) {
	resp, err := http.Get("https://ojp.nationalrail.co.uk/service/ldbboard/dep/" + station)
	if err != nil {
		return nil, err
	}
	departures, err := extractDeparturesHTML(resp.Body)
	if err != nil {
		return nil, err
	}
	board := Board{Station: station, AsAt: time.Now(), Departures: departures}
	err = resp.Body.Close()
	return &board, err
}

func isLate(statusText string) bool {
	pattern := regexp.MustCompile(`(\d{2}:\d{2}) (\d+) mins? late`)
	matches := pattern.FindStringSubmatch(statusText)
	if matches == nil {
		return false
	}
	return true
}

func expectedAndDelayed(statusText string) (time.Time, int) {
	pattern := regexp.MustCompile(`(\d{2}:\d{2}) (\d+) mins? late`)
	matches := pattern.FindStringSubmatch(statusText)
	var expected time.Time
	if matches == nil {
		return expected, 0
	}
	newTime, err := time.Parse("15:04", matches[1])
	if err != nil {
		return expected, 0
	}
	expected = newTime
	delayMins, err := strconv.Atoi(matches[2])
	if err != nil {
		return expected, 0
	}
	return expected, delayMins
}

func processStatusText(statusText string) (status TrainStatus) {
	switch statusText {
	case "On time":
		status = OnTime
	case "Cancelled":
		status = Cancelled
	case "Delayed":
		status = Delayed
	default:
		if isLate(statusText) {
			status = Late
		} else {
			status = Unknown
		}
	}
	return
}

// TODO: Workout just what sort of return value I need here.
func extractDeparturesHTML(htmlStream io.Reader) ([]Departure, error) {
	doc, err := goquery.NewDocumentFromReader(htmlStream)
	if err != nil {
		return nil, err
	}
	var processedDepartures []Departure
	doc.Find("div.tbl-cont tbody tr").Each(func(i int, s *goquery.Selection) {
		var dep Departure
		s.Find("td").Each(func(tdi int, tds *goquery.Selection) {
			switch tdi {
			case 0:
				dep.DepartureTime, _ = time.Parse("15:04", cleanText(tds.Text()))
			case 1:
				dep.Destination = cleanText(tds.Text())
			case 2:
				dep.Status = processStatusText(cleanText(tds.Text()))
				dep.ExpectedTime, dep.LateByMins = expectedAndDelayed(cleanText(tds.Text()))
			case 3:
				dep.Platform = cleanText(tds.Text())
			}
		})
		if dep.Status == OnTime {
			dep.ExpectedTime = dep.DepartureTime
		}
		processedDepartures = append(processedDepartures, dep)
	})
	return processedDepartures, nil
}
