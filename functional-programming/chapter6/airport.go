// chapter6
//
// Task: Write a function that returns the total hours of delays for the Seattle aiport (aiport code: SEA)
package chapter6

type Entry struct {
	Airport struct {
		Code string
		Name string
	}
	Statistics struct {
		Flights struct {
			Cancelled int
			Delayed   int
			OnTime    int `json:"On Time"`
			Total     int
		}
		MinutesDelayed struct {
			Carrier      int
			LateAircraft int `json:"Late Aircraft"`
			Security     int
			Weather      int
			Total        int
		} `json:"Minutes Delayed"`
	}
}

func Task(input []Entry) float64 {
	var (
		bySeattleAirport = func(input Entry) bool {
			return input.Airport.Code == "SEA"
		}

		toMinutesDelayed = func(input Entry) int {
			return input.Statistics.MinutesDelayed.Total
		}

		minutesAddition = func(a, b int) int {
			return a + b
		}
	)

	totalMinutes := float64(Reduce(
		Map( // Get minutes of delayed flights
			Filter(input, bySeattleAirport), // Get Seattle airport
			toMinutesDelayed,
		),
		minutesAddition,
	))

	return totalMinutes / 60.0
}
