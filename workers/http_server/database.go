package main

import (
	"database/sql"
	"fmt"
	"log"
)

func getTwoPrefixNumbers() []Prefix {
	var prefix = Prefix{}
	result := []Prefix{}

	db, err := sql.Open("mysql", configuration.DbUsername+":"+configuration.DbPassword+"@tcp("+configuration.DbServer+":"+configuration.DbPort+")/"+configuration.DbSchema)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("select prefix from number where prefix <= 99 group by prefix")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&prefix.Numeration)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, Prefix{prefix.Numeration})
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func getDbNumber(prefix string, series string, numeration string) NumberDb {
	var number = NumberDb{}
	result := NumberDb{}

	db, err := sql.Open("mysql", configuration.DbUsername+":"+configuration.DbPassword+"@tcp("+configuration.DbServer+":"+configuration.DbPort+")/"+configuration.DbSchema)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	row := db.QueryRow("SELECT * FROM number where prefix=? and series=? and initial_numeration<=? and final_numeration>=?", prefix, series, numeration, numeration)

	error := row.Scan(&number.IDNumber, &number.Prefix, &number.Series, &number.InitialNumeration, &number.FinalNumeration, &number.IDNumberProvider, &number.IDNumberNumberType, &number.IDNumberTown)

	switch error {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return result
	case nil:
		return number
	default:
		panic(error)
	}
}

func getDbNumbers() []NumberDb {
	var number = NumberDb{}
	result := []NumberDb{}

	db, err := sql.Open("mysql", configuration.DbUsername+":"+configuration.DbPassword+"@tcp("+configuration.DbServer+":"+configuration.DbPort+")/"+configuration.DbSchema)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT id_number, prefix, series, initial_numeration, final_numeration, id_provider, id_number_type, id_town FROM town")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&number.IDNumber, &number.Prefix, &number.Series, &number.InitialNumeration, &number.FinalNumeration, &number.IDNumberProvider, &number.IDNumberNumberType, &number.IDNumberTown)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, NumberDb{number.IDNumber, number.Prefix, number.Series, number.InitialNumeration, number.FinalNumeration, number.IDNumberProvider, number.IDNumberNumberType, number.IDNumberTown})
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func getDbNumberTypes() []NumberType {
	var numberType = NumberType{}
	result := []NumberType{}

	db, err := sql.Open("mysql", configuration.DbUsername+":"+configuration.DbPassword+"@tcp("+configuration.DbServer+":"+configuration.DbPort+")/"+configuration.DbSchema)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT id_number_type, description FROM number_type")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&numberType.IDNumberType, &numberType.Description)
		if err != nil {
			log.Fatal(err)
		}
		//log.Println(&cdrId, &calldate, &src, &dst)
		//fmt.Println(strconv.Itoa(call.id) + "," +call.queue + "," +call.server + "," +call.number + "," +call.date + "," +call.status + "," +call.uniqueid)
		result = append(result, NumberType{numberType.IDNumberType, numberType.Description})
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func getDbProviders() []Provider {
	var provider = Provider{}
	result := []Provider{}

	db, err := sql.Open("mysql", configuration.DbUsername+":"+configuration.DbPassword+"@tcp("+configuration.DbServer+":"+configuration.DbPort+")/"+configuration.DbSchema)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT id_provider, description FROM provider")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&provider.IDProvider, &provider.Description)
		if err != nil {
			log.Fatal(err)
		}
		//log.Println(&cdrId, &calldate, &src, &dst)
		//fmt.Println(strconv.Itoa(call.id) + "," +call.queue + "," +call.server + "," +call.number + "," +call.date + "," +call.status + "," +call.uniqueid)
		result = append(result, Provider{provider.IDProvider, provider.Description})
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func getDbStates() []State {
	var state = State{}
	result := []State{}

	db, err := sql.Open("mysql", configuration.DbUsername+":"+configuration.DbPassword+"@tcp("+configuration.DbServer+":"+configuration.DbPort+")/"+configuration.DbSchema)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT id_state, description FROM state")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&state.IDState, &state.Description)
		if err != nil {
			log.Fatal(err)
		}
		//log.Println(&cdrId, &calldate, &src, &dst)
		//fmt.Println(strconv.Itoa(call.id) + "," +call.queue + "," +call.server + "," +call.number + "," +call.date + "," +call.status + "," +call.uniqueid)
		result = append(result, State{state.IDState, state.Description})
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func getDbTownships() []TownshipDb {
	var township = TownshipDb{}
	result := []TownshipDb{}

	db, err := sql.Open("mysql", configuration.DbUsername+":"+configuration.DbPassword+"@tcp("+configuration.DbServer+":"+configuration.DbPort+")/"+configuration.DbSchema)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT id_township, description, id_state FROM township")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&township.IDTownship, &township.Description, &township.IDState)
		if err != nil {
			log.Fatal(err)
		}
		//log.Println(&cdrId, &calldate, &src, &dst)
		//fmt.Println(strconv.Itoa(call.id) + "," +call.queue + "," +call.server + "," +call.number + "," +call.date + "," +call.status + "," +call.uniqueid)
		result = append(result, TownshipDb{township.IDTownship, township.Description, township.IDState})
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func getDbTowns() []TownDb {
	var town = TownDb{}
	result := []TownDb{}

	db, err := sql.Open("mysql", configuration.DbUsername+":"+configuration.DbPassword+"@tcp("+configuration.DbServer+":"+configuration.DbPort+")/"+configuration.DbSchema)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT id_town, description, id_township FROM town")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&town.IDTown, &town.Description, &town.IDTownship)
		if err != nil {
			log.Fatal(err)
		}
		//log.Println(&cdrId, &calldate, &src, &dst)
		//fmt.Println(strconv.Itoa(call.id) + "," +call.queue + "," +call.server + "," +call.number + "," +call.date + "," +call.status + "," +call.uniqueid)
		result = append(result, TownDb{town.IDTown, town.Description, town.IDTownship})
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return result
}
