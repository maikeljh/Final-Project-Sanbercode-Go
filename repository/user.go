package repository

import (
	"database/sql"
	"final-project-sanbercode-go-batch-41/structs"
	"time"
)

func GetAllUsers(db *sql.DB) (err error, results []structs.User) {
	sql := `SELECT * FROM user`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var user = structs.User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Address, &user.PhoneNumber, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, user)
	}

	return
}

func InsertUser(db *sql.DB, user structs.User) (err error) {
	sql := "INSERT INTO user (name, address, phone_number, username, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	time := time.Now()
	errs := db.QueryRow(sql, user.Name, user.Address, user.PhoneNumber, user.Username, user.Password, time, time)

	return errs.Err()
}

func UpdateUser(db *sql.DB, user structs.User) (err error) {
	sqlQuery := "UPDATE user SET name = $1, address = $2, phone_number = $3, updated_at = $4 WHERE id = $5"

	time := time.Now()
	res, errs := db.Exec(sqlQuery, user.Name, user.Address, user.PhoneNumber, time, user.ID)

	if errs != nil {
		panic(errs)
	}

	n, _ := res.RowsAffected()

	if n == 0 {
		err = sql.ErrNoRows
	} else {
		err = nil
	}
	return err
}

func DeleteUser(db *sql.DB, user structs.User) (err error) {
	sqlQuery := "DELETE FROM user WHERE id = $1"

	res, errs := db.Exec(sqlQuery, user.ID)
	n, _ := res.RowsAffected()

	if errs != nil {
		panic(errs)
	}

	if n == 0 {
		err = sql.ErrNoRows
	} else {
		err = nil
	}
	return err
}
