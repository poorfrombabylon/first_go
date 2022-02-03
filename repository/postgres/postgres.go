package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"test_task/service"
)

type d struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
	sslmode  string
}

var databaseUrl = d{
	host:     "localhost",
	port:     5432,
	user:     "postgres",
	password: "228",
	dbname:   "restapi",
	sslmode:  "disable",
}

//var databaseUrl = "host=postgres dbname=restapi sslmode=disable"

type PostgresDB struct {
	db *sql.DB
}

func NewPostgresDB() *PostgresDB {
	db, err := sql.Open("postgres", fmt.Sprintf(" user=%s dbname=%s password=%s sslmode=%s",
		databaseUrl.user, databaseUrl.dbname, databaseUrl.password, databaseUrl.sslmode))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("connected")
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("postgres database does not exist")
	}

	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS restapi (LongUrls varchar(255) PRIMARY KEY, ShortUrls varchar(255) not null unique);")
	_, err = db.Exec(query)
	if err != nil {
		fmt.Println(err)
	}

	return &PostgresDB{db: db}
}

func (db *PostgresDB) GetUrl(longUrl string) (string, error) {

	var shortUrl string
	fmt.Println(longUrl)
	//statement := fmt.Sprintf()
	row := db.db.QueryRow("select ShortUrls from restapi where LongUrls = $1", longUrl)

	err := row.Scan(&shortUrl)

	if err != nil {
		return "", errors.New(fmt.Sprintf("error %v\n", err))
	}

	return shortUrl, nil
}


func (db *PostgresDB) PostUrl(longUrl string) (string, error) {
	shortUrl := service.Shorten(longUrl)
	fmt.Println(longUrl)
	//var kek string
	statement := fmt.Sprintf("INSERT INTO restapi (LongUrls, ShortUrls) VALUES ($1, $2)")
	_, err := db.db.Exec(statement, longUrl, shortUrl)
	if err != nil {
		fmt.Println("LongUrl is already in db")
		shortUrl, _ = db.GetUrl(longUrl)
		//return "", errors.New(fmt.Sprintf("error  %v", err))
	}
	//if err := db.db.QueryRow("INSERT INTO restapi (LongUrls, ShortUrls) VALUES ($1, $2) RETURNING ShortUrl", longUrl, shortUrl).Scan(&kek); err != nil {
	//	return "", errors.New(fmt.Sprintf("error %v", err))
	//}
	return shortUrl, nil
}

