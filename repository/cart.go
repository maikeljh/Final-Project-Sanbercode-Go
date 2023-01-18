package repository

import (
	"database/sql"
	"final-project-sanbercode-go-batch-41/structs"
	"time"
)

func GetAllCart(db *sql.DB) (err error, results []structs.Cart) {
	sql := `SELECT * FROM cart`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var cart = structs.Cart{}
		err = rows.Scan(&cart.ID, &cart.ProductID, &cart.Count, &cart.UserID, &cart.CreatedAt, &cart.UpdatedAt)
		if err != nil {
			panic(err)
		}

		sql = `SELECT * FROM product WHERE id = $1`
		var product structs.Product
		err1 := db.QueryRow(sql, cart.ProductID).Scan(&product.ID, &product.Name, &product.CategoryID, &product.Price, &product.Description, &product.CreatedAt, &product.UpdatedAt)
		if err1 != nil {
			panic(err1)
		}
		cart.Product = product

		results = append(results, cart)
	}

	return
}

func InsertCart(db *sql.DB, cart structs.Cart) (err error) {
	sql := "INSERT INTO cart (product_id, count, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"

	time := time.Now()
	errs := db.QueryRow(sql, cart.ProductID, cart.Count, cart.UserID, time, time)

	return errs.Err()
}

func UpdateCart(db *sql.DB, cart structs.Cart) (err error) {
	sqlQuery := "UPDATE user SET product_id = $1, count = $2, user_id = $3, updated_at = $4 WHERE id = $5"

	time := time.Now()
	res, errs := db.Exec(sqlQuery, cart.ProductID, cart.Count, cart.UserID, time, cart.ID)

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

func DeleteCart(db *sql.DB, cart structs.Cart) (err error) {
	sqlQuery := "DELETE FROM cart WHERE id = $1"

	res, errs := db.Exec(sqlQuery, cart.ID)
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

func GetCartByUserId(db *sql.DB, id int64) (err error, results []structs.Cart) {
	sql := "SELECT * from cart WHERE user_id = $1"

	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var cart = structs.Cart{}

		err = rows.Scan(&cart.ID, &cart.ProductID, &cart.Count, &cart.UserID, &cart.CreatedAt, &cart.UpdatedAt)
		if err != nil {
			panic(err)
		}

		var product structs.Product
		err1 := db.QueryRow(sql, cart.ProductID).Scan(&product.ID, &product.Name, &product.CategoryID, &product.Price, &product.Description, &product.CreatedAt, &product.UpdatedAt)
		if err1 != nil {
			panic(err1)
		}
		cart.Product = product

		results = append(results, cart)
	}

	return
}
