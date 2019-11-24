package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ValidateNumber(currentNumber string) PhoneNumber {
	var finalNumber PhoneNumber
	if len(currentNumber) == 10 {
		//formatedNumber := "+52" + currentNumber
		//number = "+521234567890"

		client := &http.Client{}
		requestURL := configuration.HachiUrl + currentNumber
		//fmt.Printf("UQL Req: %v\n", requestURL)
		req, err := http.NewRequest("GET", requestURL, nil)
		//req.SetBasicAuth(configuration.TwilioAccountSid, configuration.TwilioAuthToken)
		q := req.URL.Query()

		q.Add("Type", "carrier")
		req.URL.RawQuery = q.Encode()
		resp, err := client.Do(req)
		if err != nil {
			//log.Fatal(err)
			//fmt.Printf("LeadData: %+v\n", err, resp.StatusCode)
		}
		bodyText, err := ioutil.ReadAll(resp.Body)
		s := string(bodyText)

		var number Number
		json.Unmarshal([]byte(s), &number)
		//fmt.Printf("LeadData: %+v\n", err, resp.StatusCode)
		//fmt.Printf("LeadData: %+v\n", phoneNumber)

		if resp.StatusCode == 404 {
			return finalNumber
		} else {
			if number.NumberType.Description == "MOVIL" {
				finalNumber.Type = "mobile"
			}
			if number.NumberType.Description == "FIJO" {
				finalNumber.Type = "landline"
			}
			finalNumber.Number = currentNumber
			finalNumber.Carrier = number.Provider.Description
		}

	} else {
		return finalNumber
	}
	return finalNumber
}
