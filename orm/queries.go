// Package orm provide DB connection and custom queries
// this file is used to define the custom queries requiered
package orm

var (

	//GetUserQuery ...
	GetUserQuery = `
	SELECT u.id, co.name, r.name, u.name, u.lastname, u.email, u.is_active
	FROM users as u
	LEFT JOIN companies AS co ON  u.company_id = co.id
	LEFT JOIN roles AS r ON u.role_id = r.id
	WHERE u.id = $1;`
)
