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
func getCalledNumbers() []grackleCall {

	var call = grackleCall{}

	result := []grackleCall{}

	db, err := sql.Open("mysql", configuration.GrackleDbUsername+":"+configuration.GrackleDbPassword+"@tcp("+configuration.GrackleDbServer+":"+configuration.GrackleDbPort+")/"+configuration.GrackleDbSchema)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT id, queue, server, number, date, status, uniqueid FROM grackle.callback where status < 3")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&call.id, &call.queue, &call.server, &call.number, &call.date, &call.status, &call.uniqueid)
		if err != nil {
			log.Fatal(err)
		}
		//log.Println(&cdrId, &calldate, &src, &dst)
		//fmt.Println(strconv.Itoa(call.id) + "," +call.queue + "," +call.server + "," +call.number + "," +call.date + "," +call.status + "," +call.uniqueid)
		result = append(result, grackleCall{call.id, call.queue, call.server, call.number, call.date, call.status, call.uniqueid})
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func getCallDetailRecord(calledNumber string, date string) []cdr {
	var c = cdr{}

	result := []cdr{}

	layout := "2006-01-02 15:04:05"

	originalDate, _ := time.Parse(layout, date)

	fmt.Println("OldDate:" + originalDate.Format(layout))

	newDate := originalDate.Add(time.Minute * 2)

	fmt.Println("NewDate:" + newDate.Format(layout))

	db, err := sql.Open("mysql", configuration.PbxDbUsername+":"+configuration.PbxDbPassword+"@tcp("+configuration.PbxDbServer+":"+configuration.PbxDbPort+")/"+configuration.PbxDbSchema)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sqlQuery := "SELECT cdrId, calldate,src, dst, disposition, billsec FROM asteriskcdrdb.cdr where calldate between '" + date + "' and '" + newDate.Format(layout) + "' and disposition = 'ANSWERED' and dcontext = 'from-internal-xfer'  and dst like '%" + calledNumber + "'"

	//fmt.Println("SQL:" + sqlQuery)

	rows, err := db.Query(sqlQuery)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&c.cdrID, &c.calldate, &c.src, &c.dst, &c.disposition, &c.billsec)
		if err != nil {
			log.Fatal(err)
		}
		//log.Println(&cdrId, &calldate, &src, &dst)
		//fmt.Println(strconv.Itoa(c.cdrId) + "," + c.calldate + "," + c.src + "," + c.dst + "," + c.disposition)
		result = append(result, cdr{c.cdrID, c.calldate, c.src, c.dst, c.disposition, c.billsec})
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return result
}

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
