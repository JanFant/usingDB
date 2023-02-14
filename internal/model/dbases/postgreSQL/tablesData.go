package postgreSQL

import "time"

type cities struct {
	name        string
	postCode    string
	countryCode string
}

type countries struct {
	countryName string
	countryCode string
}

type events struct {
	eventID int
	title   string
	starts  time.Time
	ends    time.Time
	venueID int
}

type venues struct {
	venueID       int
	name          string
	streetAddress string
	typeV         string
	postalCode    string
	countryCode   string
}

type month struct {
	jan int
	feb int
	mar int
	apr int
	may int
	jun int
	jul int
	aug int
	sep int
	oct int
	nov int
	dec int
}
