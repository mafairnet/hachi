package main

var currentId int

// Give us some seed data
func init() {

}

func RepoFindNumber(requestedNumber string) NumberDb {

	var numberPrefix string
	var numberSeries string
	var numberNumeration string

	resultedNumber := NumberDb{}

	twoPrefixNumbers := getTwoPrefixNumbers()

	isTwoPrefixNumber := false
	_ = isTwoPrefixNumber

	for _, twoPrefixNumber := range twoPrefixNumbers {
		requestedNumberTwoStringPrefix := requestedNumber[:2]
		if requestedNumberTwoStringPrefix == twoPrefixNumber.Numeration {
			isTwoPrefixNumber = true
		}
	}

	if isTwoPrefixNumber {
		numberPrefix = requestedNumber[:2]
		numberSeries = requestedNumber[:6]
		numberSeries = numberSeries[len(numberSeries)-4:]
		numberNumeration = requestedNumber[len(requestedNumber)-4:]
	} else {
		numberPrefix = requestedNumber[:3]
		numberSeries = requestedNumber[:6]
		numberSeries = numberSeries[len(numberSeries)-3:]
		numberNumeration = requestedNumber[len(requestedNumber)-4:]
	}

	resultedNumber = getDbNumber(numberPrefix, numberSeries, numberNumeration)

	/*for _, number := range numbersDb {

		twoStringPrefix := str[:2]

		if twoStringPrefix != "33"

		if n.Id == id {
			return t
		}
	}*/
	// return empty Todo if not found
	return resultedNumber
}
