package sql

import "github.com/kbudke/Go-APIs/models"

var (
	// 	createRolesTable = `
	// 		CREATE TABLE IF NOT EXISTS Roles (
	// 			id VARCHAR(255),
	// 			name VARCHAR(255),
	// 			created_at DATETIME,
	// 			PRIMARY KEY (ID)
	// 		)`

	insertPerson = `
		INSERT INTO person (Name, Gender, homeworld, starships) VALUES (?, ?, ?, ?)
    `

// 	getRole = `
// 		SELECT name FROM Roles WHERE name=?
// 	`

// 	getAllRoles = `
// 		SELECT * FROM Roles
// 	`

// 	deleteRole = `
// 		DELETE FROM Roles WHERE id=?
// 	`
)

func (c client) AddPerson(person models.Person) (models.Person, error) {
	_, err := c.DB.Exec(insertPerson,
		person.name,
		person.gender,
		person.homeworld,
		person.starships)
	return roles, err
}

// func (c client) VerifyRole(name string) error {
// 	var role models.Role
// 	var err error
// 	err = c.DB.Get(&role, getRole, name)
// 	if err != nil {
// 		return err
// 	}

// 	return err
// }

// func (c client) GetAllRoles() ([]models.Role, error) {
// 	var roles []models.Role
// 	var err error
// 	err = c.DB.Select(&roles, getAllRoles)
// 	if err != nil {
// 		return roles, err
// 	}

// 	return roles, nil
// }

// func (c client) DeleteRole(roleID string) error {
// 	_, err := c.DB.Exec(deleteRole, roleID)

// 	return err
// }
