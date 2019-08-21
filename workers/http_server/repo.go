package main

import "fmt"

var currentId int

// Give us some seed data
func init() {

}

func RepoFindNumber(requestedNumber string) Number {

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

	numberType := getDbNumberType(resultedNumber.IDNumberNumberType)
	fmt.Printf("NumberType: %v\n", numberType)

	provider := getDbProvider(resultedNumber.IDNumberProvider)
	fmt.Printf("Provider: %v\n", provider)

	town := getDbTown(resultedNumber.IDNumberTown)
	fmt.Printf("Town: %v\n", town)

	township := getDbTownship(town.IDTownship)
	fmt.Printf("Township: %v\n", township)

	state := getDbState(township.IDState)
	fmt.Printf("State: %v\n", state)

	var number Number

	number.IDNumber = resultedNumber.IDNumber
	number.Prefix = resultedNumber.Prefix
	number.Series = resultedNumber.Series
	number.InitialNumeration = resultedNumber.InitialNumeration
	number.FinalNumeration = resultedNumber.FinalNumeration
	number.NumberProvider = provider
	number.NumberNumberType = numberType

	numberTwonship := Township{township.IDTownship, township.Description, state}
	numberTown := Town{town.IDTown, town.Description, numberTwonship}

	number.NumberTown = numberTown

	/*for _, number := range numbersDb {

		twoStringPrefix := str[:2]

		if twoStringPrefix != "33"

		if n.Id == id {
			return t
		}
	}

	resultedNumber*/

	// return empty Todo if not found
	return number
}
