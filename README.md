# Train Departures

## Overview
A simple project that uses web scraping to capture UK Train departures from the Nation Rail website.

This is currently using the [OJP board](https://ojp.nationalrail.co.uk/service/ldbboard/dep/KGX) which as of August 2023
looks like it will be replaced with an improved board.

## Current API
To fetch a train departure board it is as easy as referencing the project and calling:

```go
package main

import (
	"fmt"
	"github.com/bperryman/traindeps"
	"os"
)

func main() {
	board, err := traindeps.LoadFromInternet("STATION")
	if err != nil {
		fmt.Printf("Error loading the departure board")
		os.Exit(-1)
	}
	for _, dep := range board.Departures {
		fmt.Printf("To: %s, leaving from: %s\n", dep.Destination, dep.Platform)
    }
}
```

Here `STATION` is the 3-letter code, for example `KGX` is the code for London Kings Cross. A list of these can be found on
[Wikipedia](https://en.wikipedia.org/wiki/UK_railway_stations_–_A)

## History
During a particularly bad spell of train lateness I got so fed up that I wrote a scraper to pull down all the delays on
the journeys that I would make. Just for kicks I ended up writing this in Common Lisp, Elixir, Python and the go version
here.

Then the pandemic happened, along with a lack off journeys into work, and this project just sort of languished.
Now I want to use this for a notice board, based around the
[Pimoroni Galactic Unicorn](https://shop.pimoroni.com/products/space-unicorns?variant=40842033561683)
- yes, yes it looked cool and I bought it, so now I have a solution looking for a problem (again).

## TODO
- [X] Rename the package to something more go friendly.
- [ ] Update the scraper to retrieve the departure boards from the new presentation (currently beta?)
- [ ] Update to include any special notifiers that are displayed - for example delays due to bad weather or special
services being run
- [ ] Allow for a specific destination - so your journey into work for example.
- [ ] Allow for intermediate stations on the journey to be retrieved.
- [ ] Update any testing.
