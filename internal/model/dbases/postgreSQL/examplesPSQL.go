package postgreSQL

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func PSQLExamples(db *sqlx.DB) {
	selectFromCountries1(db, "LARP Club")
}

func selectFromCountries1(db *sqlx.DB, title string) {
	rows, _ := db.Query(`SELECT c.country_name FROM public.countries c
					JOIN venues v on c.country_code = v.country_code
					JOIN events e on v.venue_id = e.venue_id
					WHERE e.title = $1;`, title)
	var contryNames []string
	for rows.Next() {
		var tempCN string
		if err := rows.Scan(&tempCN); err != nil {
			fmt.Println(err.Error())
		}
		contryNames = append(contryNames, tempCN)
	}
	fmt.Println(contryNames)
}

func insertCountries(db *sqlx.DB) {
	_, err := db.Exec(`INSERT INTO public.countries (country_code, country_name) 
			VALUES ('us','United States'), ('mx','Mexico'), ('au','Australia'),
			('gb','United Kingdom'), ('de','Germany'), ('ru','Russia'),('kz','Kazakhstan');`)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func insertCities(db *sqlx.DB) {
	_, err := db.Exec(`INSERT INTO public.cities (name, postal_code, country_code)
			VALUES ('Portland','97205','us');`)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func insertVenues(db *sqlx.DB) {
	_, err := db.Exec(`INSERT INTO public.venues (name, postal_code, country_code) 
			VALUES ('Crystal Ballroom','97205','us'),('VooDoo Donuts','97205','us');`)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func insertEvents(db *sqlx.DB) {
	_, err := db.Exec(`INSERT INTO events (title, starts, ends, venue_id)
			VALUES ('LARP Club', '2012-02-15 17:30', '2012-02-15 19:30', 2),
			       ('April Fools Day', '2012-04-01 00:00', '2012-04-01 23:59', NULL),
			       ('Christmas Day', '2012-12-25 00:00', '2012-12-25 23:59', NULL);`)
	if err != nil {
		fmt.Println(err.Error())
	}
}
