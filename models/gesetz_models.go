package models

func GetGesetze() []Gesetz {
	var gesetze []Gesetz

	rows, err := db.Query("SELECT * FROM Gesetz")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var gesetz Gesetz
		err := rows.Scan(&gesetz.ID, &gesetz.Date, &gesetz.Titel, &gesetz.Topic, &gesetz.Description, &gesetz.Details)
		if err != nil {
			panic(err)
		}

		gesetze = append(gesetze, gesetz)
	}
	return gesetze
}

func GesetzById(id int64) Gesetz {
	var gesetz Gesetz
	rows, err := db.Query("SELECT * FROM Gesetz WHERE id = ?", id)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&gesetz.ID, &gesetz.Date, &gesetz.Titel, &gesetz.Topic, &gesetz.Description, &gesetz.Details)

	if err != nil {
		panic(err)
	}

	return gesetz
}

func PutGesetz(gesetz Gesetz_put) {

	rows, err := db.Query("INSERT INTO Gesetz (titel,date,topic,description,details) VALUES (?,?,?,?,?)", gesetz.Titel, gesetz.Date, gesetz.Topic, gesetz.Description, gesetz.Details)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

}
