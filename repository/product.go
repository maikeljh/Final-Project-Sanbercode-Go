package repository

import (
	"database/sql"
	"final-project-sanbercode-go-batch-41/structs"
	"time"
)

func GetAllProducts(db *sql.DB) (err error, results []structs.Product) {
	sql := `SELECT * FROM product`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var product = structs.Product{}
		err = rows.Scan(&product.ID, &product.Name, &product.CategoryID, &product.Price, &product.Description, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, product)
	}

	return
}

func InsertProduct(db *sql.DB, product structs.Product) (err error) {
	sql := "INSERT INTO product (name, category_id, price, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"

	time := time.Now()
	errs := db.QueryRow(sql, product.Name, product.CategoryID, product.Price, product.Description, time, time)

	return errs.Err()
}

func UpdateProduct(db *sql.DB, product structs.Product) (err error) {
	sqlQuery := "UPDATE user SET name = $1, category_id = $2, price = $3, description = $4, updated_at = $5 WHERE id = $6"

	time := time.Now()
	res, errs := db.Exec(sqlQuery, product.Name, product.CategoryID, product.Price, product.Description, time, product.ID)

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

func DeleteProduct(db *sql.DB, product structs.Product) (err error) {
	sqlQuery := "DELETE FROM rpdocut WHERE id = $1"

	res, errs := db.Exec(sqlQuery, product.ID)
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

func GetProductByCategoryID(db *sql.DB, id int) (err error, results []structs.Product) {
	sql := "SELECT * FROM product WHERE category_id = $1"

	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var product = structs.Product{}

		err = rows.Scan(&product.ID, &product.Name, &product.CategoryID, &product.Price, &product.Description, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, product)
	}

	return
}
