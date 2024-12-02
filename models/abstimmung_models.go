package models

func GetAbstimmungen() []Abstimmung {
	var abstimmungen []Abstimmung

	rows, err := db.Query("SELECT * FROM Abstimmung")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var abstimmung Abstimmung
		err := rows.Scan(&abstimmung.ID, &abstimmung.GesetzId, &abstimmung.AbgeordneteId, &abstimmung.Vote)
		if err != nil {
			panic(err)
		}

		abstimmungen = append(abstimmungen, abstimmung)
	}
	return abstimmungen
}

func AbstimmungById(id int64) Abstimmung {
	var abstimmung Abstimmung
	rows, err := db.Query("SELECT * FROM Abstimmung WHERE id = ?", id)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	rows.Next()
	err = rows.Scan(&abstimmung.ID, &abstimmung.GesetzId, &abstimmung.AbgeordneteId, &abstimmung.Vote)

	if err != nil {
		panic(err)
	}

	return abstimmung
}

func PutAbstimmung(abstimmung Abstimmung_put) {

	rows, err := db.Query("INSERT INTO Abstimmung (gesetzId,abgeorneteId, vote) VALUES (?,?,?)", abstimmung.GesetzId, abstimmung.AbgeordneteId, abstimmung.Vote)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

}
