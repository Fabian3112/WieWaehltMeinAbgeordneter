package models

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type User_put struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Abgeordnete struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Partei string `json:"patei"`
	Bio    string `json:"bio"`
}

type Abgeordnete_put struct {
	Name   string `json:"name"`
	Partei string `json:"partei"`
	Bio    string `json:"bio"`
}

type AbgeordneteVote struct {
	ID           int64        `json:"id"`
	Name         string       `json:"name"`
	Partei       string       `json:"patei"`
	Bio          string       `json:"bio"`
	Abstimmungen []Abstimmung `json:abstimmungen`
}

type Gesetz struct {
	ID          int64  `json:"id"`
	Titel       string `json:"titel"`
	Date        string `json:"date"`
	Topic       string `json:"topic"`
	Description string `json:"description"`
	Details     string `json:"details"`
}

type Gesetz_put struct {
	Titel       string `json:"titel"`
	Date        string `json:"date"`
	Topic       string `json:"topic"`
	Description string `json:"description"`
	Details     string `json:"details"`
}

type Abstimmung struct {
	ID            int64  `json:"id"`
	GesetzId      int64  `json:"gesetzId"`
	AbgeordneteId int64  `json:"abgeordneteId"`
	Vote          string `json:"vote"`
}

type Abstimmung_put struct {
	GesetzId      int64  `json:"gesetzId"`
	AbgeordneteId int64  `json:"abgeordneteId"`
	Vote          string `json:"vote"`
}
