package config

import (
	"os"
	"sync"

	logger "github.com/tcero76/marketplace/config/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	onceRead  sync.Once
	onceWrite sync.Once
	errRead   error
	errWrite  error
	dbRead    *gorm.DB
	dbWrite   *gorm.DB
)

func GetPostgresWrite(log *logger.LoggerLogstash) (*gorm.DB, error) {
	log.Info("Entra a Config Postgres WRITE...")
	onceWrite.Do(func() {
		dsn := os.Getenv("DNS_WRITE")
		log.Info("DNS_WRITE: ", dsn)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Error("Error al conectar a la base de datos (write):", err)
			errWrite = err
			return
		}
		log.Info("Conexión exitosa a PostgreSQL (write) con GORM")
		dbWrite = db
	})
	return dbWrite, errWrite
}

func GetPostgresRead(log *logger.LoggerLogstash) (*gorm.DB, error) {
	log.Info("Entra a Config Postgres READ...")
	onceRead.Do(func() {
		dsn := os.Getenv("DNS_READ")
		log.Info("DNS_READ: ", dsn)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Error("Error al conectar a PostgreSQL (read):", err)
			errRead = err
			return
		}

		log.Info("Conexión exitosa a PostgreSQL (read) con GORM")
		dbRead = db
	})

	return dbRead, errRead
}
