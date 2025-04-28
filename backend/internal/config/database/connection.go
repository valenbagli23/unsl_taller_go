package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	// import aninomo: _
	// Se ejecuta solamente su funciono init
)

var DB *sql.DB

func Connect() {
	var err error

	dsn := "root:@tcp(127.0.0.1:3306)/?parseTime=true"

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error al abrir la conexión:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos:", err)
	}

	fmt.Println("Conexión a la base de datos exitosa!")
}
