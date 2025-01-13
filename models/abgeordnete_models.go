package models

func GetAbgeordnete() []Abgeordnete {
	var abgeordnete []Abgeordnete

	rows, err := db.Query("SELECT * FROM Abgeordnete")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var abgeordneter Abgeordnete
		err := rows.Scan(&abgeordneter.ID, &abgeordneter.Name, &abgeordneter.Partei, &abgeordneter.Bio)
		if err != nil {
			panic(err)
		}

		abgeordnete = append(abgeordnete, abgeordneter)
	}
	return abgeordnete
}

func AbgeordneterById(id int64) Abgeordnete {
	var abgeordnete Abgeordnete
	rows, err := db.Query("SELECT * FROM Abgeordnete WHERE id = ?", id)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	rows.Next()
	err = rows.Scan(&abgeordnete.ID, &abgeordnete.Name, &abgeordnete.Partei, &abgeordnete.Bio)

	if err != nil {
		panic(err)
	}

	return abgeordnete
}

func PutAbgeordneter(abgeordneter Abgeordnete_put) {

	rows, err := db.Query("INSERT INTO Abgeordnete (name,partei,bio) VALUES (?,?,?)", abgeordneter.Name, abgeordneter.Partei, abgeordneter.Bio)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

}

func GetAbgeordneteVote(id int64) AbgeordneteVote {
	var abgeordnete AbgeordneteVote
	rows, err := db.Query("SELECT * FROM Abgeordnete WHERE id = ?", id)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	rows.Next()
	err = rows.Scan(&abgeordnete.ID, &abgeordnete.Name, &abgeordnete.Partei, &abgeordnete.Bio)

	if err != nil {
		panic(err)
	}

	var abstimmungen []Abstimmung

	rows, err = db.Query("SELECT * FROM Abstimmung Where abgeorneteId = ?", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var absimmung Abstimmung
		err := rows.Scan(&absimmung.ID, &absimmung.GesetzId, &absimmung.AbgeordneteId, &absimmung.Vote)
		if err != nil {
			panic(err)
		}

		abstimmungen = append(abstimmungen, absimmung)
	}
	abgeordnete.Abstimmungen = abstimmungen

	return abgeordnete
}
