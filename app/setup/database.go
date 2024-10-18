package setup

import (
	"database/sql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"pismo/data/models"
	"time"
)

func ConnectDatabase() *gorm.DB {

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: getConnection(),
	}), &gorm.Config{})

	if err != nil {
		log.Fatalln("[gorm.Open]: ", err)
	}

	_ = db.AutoMigrate(&models.Account{})
	_ = db.AutoMigrate(&models.OperationType{})
	_ = db.AutoMigrate(&models.Transaction{})

	return db
}

func getConnection() *sql.DB {

	postgresDsn := os.Getenv("POSTGRES_DSN")
	var counts int
	counts = 0

	var conn *sql.DB
	var err error

	for {
		conn, err = openDB(postgresDsn)
		if err != nil {
			Logger.Info("Postgres not yet ready...")
			counts++
		} else {
			Logger.Info("Connected to Postgres")
			break
		}

		if counts > 10 {
			Logger.Error("Can't connect to database!", err)
			panic("Can't connect to database!")
		}

		Logger.Info("Backing off for two seconds")
		time.Sleep(2 * time.Second)
	}

	return conn
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
