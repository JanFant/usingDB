package postgreSQL

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"io"
	"os"
)

func PSQLExamples(db *sqlx.DB) {
	//selectFromCountries1(db, "LARP Club")
	selectYearAndMonth(db)
}

func selectYearAndMonth(db *sqlx.DB) {
	rows, _ := db.Query(`SELECT year, coalesce(jan,0) as jan, coalesce(feb,0) as feb, coalesce(mar,0) as mar,
             								coalesce(apr,0) as apr, coalesce(may,0) as may, coalesce(jun,0) as jun,
             								coalesce(jul,0) as jul, coalesce(aug,0) as aug, coalesce(sep,0) as sep,
             								coalesce(oct,0) as oct, coalesce(nov,0) as nov, coalesce(dec,0) as dec
								FROM crosstab(
                      				'SELECT extract(year from starts) as year,
                              				extract(month from starts) as month, count(*)
                      				 FROM events GROUP BY year, month',
                      				'SELECT generate_series as month FROM generate_series(1,12)'
								) AS (
                        			year int,
                        			jan int, feb int, mar int, apr int, may int, jun int,
                        			jul int, aug int, sep int, oct int, nov int, dec int
								) ORDER BY YEAR;`)
	var (
		mapYear = make(map[int]month)
	)

	for rows.Next() {
		var (
			tYear  int
			tMonth month
		)
		if err := rows.Scan(&tYear, &tMonth.jan, &tMonth.feb, &tMonth.mar, &tMonth.apr, &tMonth.may, &tMonth.jun,
			&tMonth.jul, &tMonth.aug, &tMonth.sep, &tMonth.oct, &tMonth.nov, &tMonth.dec); err != nil {
			fmt.Println(err.Error())
		}
		mapYear[tYear] = tMonth
	}
	fmt.Println(mapYear)

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

func readAddEventFile(path string) (str string, err error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(`Can't open file add_event.sql -err:"'`, err.Error())
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("close file err", err.Error())
		}
	}(file)
	data := make([]byte, 64)
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		str = str + string(data[:n])
	}
	return str, nil
}
