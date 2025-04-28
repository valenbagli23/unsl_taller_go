package database

import (
	"fmt"
	"log"
)

func Migrate() {
	// Primero creamos la base de datos si no existe
	createDBCmd := `CREATE DATABASE IF NOT EXISTS db_go`
	_, err := DB.Exec(createDBCmd)
	if err != nil {
		log.Fatal("Error creando la base de datos:", err)
	}

	// Seleccionamos la base de datos
	useDBCmd := `USE db_go`
	_, err = DB.Exec(useDBCmd)
	if err != nil {
		log.Fatal("Error seleccionando la base de datos:", err)
	}

	//-----------------
	if _, err := DB.Exec("DROP TABLE IF EXISTS users"); err != nil {
		log.Fatal("No se puede borrar la tabla", err)
	}

	//---------------------
	// Ahora creamos las tablas
	query := queryUser()
	_, err = DB.Exec(query)
	if err != nil {
		log.Fatal("Error creando tabla users:", err)
	}

	fmt.Println("Migración completada: Base de datos y tablas creadas correctamente.")
}

func queryUser() string {
	return `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT  PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		address VARCHAR(255) NOT NULL,
		nickname VARCHAR(100),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		version INT DEFAULT 1
	);	
	`
}
