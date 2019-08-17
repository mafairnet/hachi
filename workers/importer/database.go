package main

type grackleCall struct {
	id       int
	queue    int
	server   int
	number   string
	date     string
	status   int
	uniqueid string
}

func getUser() {

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
