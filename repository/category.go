package repository

import (
	"database/sql"
	"final-project-sanbercode-go-batch-41/structs"
	"time"
)

func GetAllCategory(db *sql.DB) (err error, results []structs.Category) {
	sql := "SELECT * FROM category"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var category = structs.Category{}

		err = rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, category)
	}

	return
}

func InsertCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "INSERT INTO category (name, created_at, updated_at) VALUES ($1, $2, $3)"

	time := time.Now()
	errs := db.QueryRow(sql, category.Name, time, time)

	return errs.Err()
}

func UpdateCategory(db *sql.DB, category structs.Category) (err error) {
	sqlQuery := "UPDATE category SET name = $1, updated_at = $2 WHERE id = $3"

	time := time.Now()
	res, errs := db.Exec(sqlQuery, category.Name, time, category.ID)

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

func DeleteCategory(db *sql.DB, category structs.Category) (err error) {
	sqlQuery := "DELETE FROM category WHERE id = $1"

	res, errs := db.Exec(sqlQuery, category.ID)
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
