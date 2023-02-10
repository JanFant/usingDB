package postgreSQL

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"usingDB/internal/model/config"
)

var (
	tableCountries = `
	CREATE TABLE countries (
		country_code char(2) PRIMARY KEY,
		country_name text UNIQUE
	);`

	tableCities = `
	CREATE TABLE cities (
		name text NOT NULL,
		postal_code varchar(9) CHECK (postal_code <> ''),
		country_code char(2) REFERENCES countries,
		PRIMARY KEY (country_code, postal_code)
	);`

	tableVenues = `
	CREATE TABLE venues (
		venue_id SERIAL PRIMARY KEY,
		name varchar(255),
		street_address text,
		type char(7) CHECK ( type in ('public','private') ) DEFAULT 'public',
		postal_code varchar(9),
		country_code char(2),
		FOREIGN KEY (country_code, postal_code)
		REFERENCES cities (country_code, postal_code) MATCH FULL
	);`

	tableEvents = `
	CREATE TABLE events (
	    event_id SERIAL PRIMARY KEY,
	    title text,
	    starts timestamp,
	    ends timestamp,
	    venue_id int,
	    FOREIGN KEY (venue_id) REFERENCES venues (venue_id)
	    );`
)

func ConnectPSQLD() (*sqlx.DB, error) {
	db, err := sqlx.Open(config.GlobalConfig.PSQLConfig.Type, config.GlobalConfig.PSQLConfig.GetPSQLUrl())
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(config.GlobalConfig.PSQLConfig.MaxOpenConst)
	db.SetMaxIdleConns(config.GlobalConfig.PSQLConfig.MaxIdleConst)

	//Countries
	_, err = db.Exec(`SELECT * FROM public.countries limit(1);`)
	if err != nil {
		fmt.Println("Create Countries table!")
		db.MustExec(tableCountries)
		insertCountries(db)
	}

	//Cities
	_, err = db.Exec(`SELECT * FROM public.cities limit(1);`)
	if err != nil {
		fmt.Println("Create Cities table!")
		db.MustExec(tableCities)
		insertCities(db)
	}

	//Venues
	_, err = db.Exec(`SELECT * FROM public.venues limit(1);`)
	if err != nil {
		fmt.Println("Create Venues table!")
		db.MustExec(tableVenues)
		insertVenues(db)
	}

	//Events
	_, err = db.Exec(`SELECT * FROM public.events limit(1);`)
	if err != nil {
		fmt.Println("Create Events table!")
		db.MustExec(tableEvents)
		insertEvents(db)
	}

	return db, nil
}
