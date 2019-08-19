package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func cleanDB() {
	db, err := sql.Open("mysql", configuration.DbUsername+":"+configuration.DbPassword+"@tcp("+configuration.DbServer+":"+configuration.DbPort+")/"+configuration.DbSchema)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sql := "call CleanDb"

	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("SQL: %v, Result: %v\n", sql, affect)

}

func insertNumberTypesIntoDb() {

	db, err := sql.Open("mysql", configuration.DbUsername+":"+configuration.DbPassword+"@tcp("+configuration.DbServer+":"+configuration.DbPort+")/"+configuration.DbSchema)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sql := "INSERT INTO `hachi`.`number_type` (`description`) VALUES ('FIJO'), ('MOVIL');"

	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("SQL: %v, Result: %v\n", sql, affect)
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

func insertProvidersIntoDb(providers []string) {
	for _, provider := range providers {
		db, err := sql.Open("mysql", configuration.DbUsername+":"+configuration.DbPassword+"@tcp("+configuration.DbServer+":"+configuration.DbPort+")/"+configuration.DbSchema)

		if err != nil {
			log.Fatal(err)
		}

		defer db.Close()

		sql := "INSERT INTO `hachi`.`provider` (`description`) VALUES ('" + provider + "');"

		stmt, err := db.Prepare(sql)
		if err != nil {
			log.Fatal(err)
		}

		res, err := stmt.Exec()
		if err != nil {
			log.Fatal(err)
		}

		affect, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("SQL: %v, Result: %v\n", sql, affect)
	}
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

//insertStatesIntoDb Insert States retreived from CSV file to Database
func insertStatesIntoDb(states []string) {
	for _, state := range states {
		db, err := sql.Open("mysql", configuration.DbUsername+":"+configuration.DbPassword+"@tcp("+configuration.DbServer+":"+configuration.DbPort+")/"+configuration.DbSchema)

		if err != nil {
			log.Fatal(err)
		}

		defer db.Close()

		sql := "INSERT INTO `hachi`.`state` (`description`) VALUES ('" + state + "');"

		stmt, err := db.Prepare(sql)
		if err != nil {
			log.Fatal(err)
		}

		res, err := stmt.Exec()
		if err != nil {
			log.Fatal(err)
		}

		affect, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("SQL: %v, Result: %v\n", sql, affect)
	}
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

//insertTownshipsIntoDb Insert States retreived from CSV file to Database
func insertTownshipsIntoDb(townships []string, statesDb []State) {
	var stateID int
	db, err := sql.Open("mysql", configuration.DbUsername+":"+configuration.DbPassword+"@tcp("+configuration.DbServer+":"+configuration.DbPort+")/"+configuration.DbSchema)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sql := "INSERT INTO `hachi`.`township` (`description`, `id_state`) VALUES "

	statesProcessed := 0

	for _, township := range townships {

		townshipState := strings.Split(township, "-")
		stateID = 0
		_ = stateID
		for _, state := range statesDb {
			if state.Description == townshipState[0] {
				stateID = state.IDState
			}
		}

		if stateID != 0 {
			//fmt.Printf("Township: %v, State: %v, StateID: %v\n", townshipState[1], townshipState[0], stateID)
			//INSERT INTO `hachi`.`township` (`description`, `id_state`) VALUES ('Test', '1');

			stateIDStr := strconv.Itoa(stateID)

			if statesProcessed > 0 {
				sql = sql + ","
			}

			sql = sql + " ('" + townshipState[1] + "','" + stateIDStr + "')"

			//_ = sql
			statesProcessed++
		}
	}

	if statesProcessed > 0 {

		stmt, err := db.Prepare(sql)
		if err != nil {
			log.Fatal(err)
		}

		res, err := stmt.Exec()
		if err != nil {
			log.Fatal(err)
		}

		affect, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("SQL: %v, Result: %v\n", sql, affect)
	}

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

//insertTownsIntoDb Insert States retreived from CSV file to Database
func insertTownsIntoDb(towns []string, townshipsDb []TownshipDb, statesDb []State) {
	var townshipID int

	db, err := sql.Open("mysql", configuration.DbUsername+":"+configuration.DbPassword+"@tcp("+configuration.DbServer+":"+configuration.DbPort+")/"+configuration.DbSchema)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sql := "INSERT INTO `hachi`.`town` (`description`, `id_township`) VALUES "

	townsProcessed := 0

	for _, town := range towns {

		townTownship := strings.Split(town, "-")

		if townTownship[0] == "CANCUN" {
			fmt.Printf("CANCUN\n")
		}

		townshipID = 0
		_ = townshipID

		for _, state := range statesDb {
			_ = state

			if townTownship[2] == state.Description {
				for _, township := range townshipsDb {
					if township.Description == townTownship[1] && township.IDState == state.IDState {
						townshipID = township.IDTownship
					}
				}
			}
		}

		if townshipID != 0 {
			//fmt.Printf("Township: %v, State: %v, StateID: %v\n", townshipState[1], townshipState[0], stateID)
			//INSERT INTO `hachi`.`township` (`description`, `id_state`) VALUES ('Test', '1');

			townshipIDStr := strconv.Itoa(townshipID)

			if townsProcessed > 0 {
				sql = sql + ","
			}

			sql = sql + " ('" + townTownship[0] + "','" + townshipIDStr + "')"

			//_ = sql
			townsProcessed++
		}
	}

	if townsProcessed > 0 {

		stmt, err := db.Prepare(sql)
		if err != nil {
			log.Fatal(err)
		}

		res, err := stmt.Exec()
		if err != nil {
			log.Fatal(err)
		}

		affect, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("SQL: %v, Result: %v\n", sql, affect)
	}

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

/*
func insertNumbersIntoDb(towns []string, townshipsDb []TownshipDb, statesDb []State, providers []Provider) {
	var townshipID int

	db, err := sql.Open("mysql", configuration.DbUsername+":"+configuration.DbPassword+"@tcp("+configuration.DbServer+":"+configuration.DbPort+")/"+configuration.DbSchema)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sql := "INSERT INTO `hachi`.`town` (`description`, `id_township`) VALUES "

	townsProcessed := 0

	for _, town := range towns {

		townTownship := strings.Split(town, "-")

		if townTownship[0] == "CANCUN" {
			fmt.Printf("CANCUN\n")
		}

		townshipID = 0
		_ = townshipID

		for _, state := range statesDb {
			_ = state

			if townTownship[2] == state.Description {
				for _, township := range townshipsDb {
					if township.Description == townTownship[1] && township.IDState == state.IDState {
						townshipID = township.IDTownship
					}
				}
			}
		}

		if townshipID != 0 {
			//fmt.Printf("Township: %v, State: %v, StateID: %v\n", townshipState[1], townshipState[0], stateID)
			//INSERT INTO `hachi`.`township` (`description`, `id_state`) VALUES ('Test', '1');

			townshipIDStr := strconv.Itoa(townshipID)

			if townsProcessed > 0 {
				sql = sql + ","
			}

			sql = sql + " ('" + townTownship[0] + "','" + townshipIDStr + "')"

			//_ = sql
			townsProcessed++
		}
	}

	if townsProcessed > 0 {

		stmt, err := db.Prepare(sql)
		if err != nil {
			log.Fatal(err)
		}

		res, err := stmt.Exec()
		if err != nil {
			log.Fatal(err)
		}

		affect, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("SQL: %v, Result: %v\n", sql, affect)
	}

}

/*
func updateGrackleRegistry(grackleCallID int, status int) {
	db, err := sql.Open("mysql", configuration.GrackleDbUsername+":"+configuration.GrackleDbPassword+"@tcp("+configuration.GrackleDbServer+":"+configuration.GrackleDbPort+")/"+configuration.GrackleDbSchema)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	stmt, err := db.Prepare("update callback set status =? where id=?")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(status, grackleCallID)
	if err != nil {
		log.Fatal(err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(affect)

}
*/
