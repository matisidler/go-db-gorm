package storage

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Creamos la conexión a la BD.
//Utilizamos el Patrón Singleton para que solo se ejecute una vez.

//Creamos dos variables que van a poder ser usadas por todos los archivos del paquete storage.
//Con once hacemos el patrón Singleton para que se ejecute una sola vez.
var (
	db   *gorm.DB
	once sync.Once
)

type Driver string

const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

func NewConnection(d Driver) {
	switch d {
	case MySQL:
		newMySqlDB()
	case Postgres:
		newPqDB()
	default:
		log.Fatal("Parametro no valido")
	}

}

func newPqDB() *gorm.DB {
	once.Do(func() {
		var err error
		db, err = gorm.Open(postgres.Open("postgres://postgres:password@localhost:5432/gocrud?sslmode=disable"), &gorm.Config{})
		if err != nil {
			log.Fatalf("can't open DB %v", err)
		}

		fmt.Println("Conectado a Postgres.")
	})
	return db
}

func newMySqlDB() *gorm.DB {
	once.Do(func() {
		var err error
		db, err = gorm.Open(mysql.Open("root:password@tcp(localhost:3306)/gocrud?parseTime=true"), &gorm.Config{})
		if err != nil {
			log.Fatalf("can't open DB %v", err)
		}

		fmt.Println("Conectado a MySQL.")
	})
	return db
}
