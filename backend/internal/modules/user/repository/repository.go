package repository

import (
	"database/sql"
	"log"

	"backend/internal/modules/user/domain"
)

// Estructura que encapsula la base de datos
type UserRepository struct {
	DB *sql.DB
}

// Constructor
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Devuelve Error, si es null entonces la operacion se cumplió con éxito
func (r *UserRepository) Create(user *domain.User) error {
	query := `INSERT INTO users (id, name, address, nickname, created_at, updated_at, version) 
			  VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := r.DB.Exec(query, user.ID, user.Name, user.Address, user.NickName, user.CreatedAt, user.UpdatedAt, user.Version)
	if err != nil {
		log.Println("Error al crear el usuario:", err)
		return err
	}
	return nil
}

// Devuelve todos los usuarios
func (r *UserRepository) FindAll() ([]domain.User, error) {
	var users []domain.User
	query := `SELECT id, name, address, nickname, created_at, updated_at, version FROM users`

	rows, err := r.DB.Query(query)
	if err != nil {
		log.Println("Error al obtener usuarios:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Address, &user.NickName, &user.CreatedAt, &user.UpdatedAt, &user.Version); err != nil {
			log.Println("Error al escanear usuario:", err)
			return nil, err
		}
		users = append(users, user)
	}

	// Verificar si hubo error al iterar sobre las filas
	if err := rows.Err(); err != nil {
		log.Println("Error durante la iteración de los resultados:", err)
		return nil, err
	}

	return users, nil
}
