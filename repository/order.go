package repository

import (
	"database/sql"
	"final-project-sanbercode-go-batch-41/structs"
	"time"
)

func GetAllOrder(db *sql.DB) (err error, results []structs.Order) {
	sql := "SELECT * FROM user_order"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var order = structs.Order{}

		err = rows.Scan(&order.ID, &order.CartID, &order.UserID, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			panic(err)
		}

		sql = `SELECT * FROM cart WHERE id = $1`
		var cart structs.Cart
		err1 := db.QueryRow(sql, order.CartID).Scan(&cart.ID, &cart.ProductID, &cart.Count, &cart.UserID, &cart.CreatedAt, &cart.UpdatedAt)
		if err1 != nil {
			panic(err1)
		}

		sql = `SELECT * FROM product WHERE id = $1`
		var product structs.Product
		err2 := db.QueryRow(sql, cart.ProductID).Scan(&product.ID, &product.Name, &product.CategoryID, &product.Price, &product.Description, &product.CreatedAt, &product.UpdatedAt)
		if err2 != nil {
			panic(err2)
		}

		cart.Product = product
		order.Cart = cart

		results = append(results, order)
	}

	return
}

func InsertOrder(db *sql.DB, order structs.Order) (err error) {
	sql := "INSERT INTO user_order (cart_id, user_id , created_at, updated_at) VALUES ($1, $2, $3, $4)"

	time := time.Now()
	errs := db.QueryRow(sql, order.CartID, order.UserID, time, time)

	return errs.Err()
}

func UpdateOrder(db *sql.DB, order structs.Order) (err error) {
	sqlQuery := "UPDATE user_order SET cart_id = $1, user_id = $2, updated_at = $3 WHERE id = $4"

	time := time.Now()
	res, errs := db.Exec(sqlQuery, order.CartID, order.UserID, time, order.ID)

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

func DeleteOrder(db *sql.DB, order structs.Order) (err error) {
	sqlQuery := "DELETE FROM user_order WHERE id = $1"

	res, errs := db.Exec(sqlQuery, order.ID)
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

func GetOrderByUserId(db *sql.DB, id int64) (err error, results []structs.Order) {
	sql := "SELECT * from user_order WHERE user_id = $1"

	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var order = structs.Order{}

		err = rows.Scan(&order.ID, &order.CartID, &order.UserID, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			panic(err)
		}

		sql = `SELECT * FROM cart WHERE id = $1`
		var cart structs.Cart
		err1 := db.QueryRow(sql, order.CartID).Scan(&cart.ID, &cart.ProductID, &cart.Count, &cart.UserID, &cart.CreatedAt, &cart.UpdatedAt)
		if err1 != nil {
			panic(err1)
		}

		sql = `SELECT * FROM product WHERE id = $1`
		var product structs.Product
		err2 := db.QueryRow(sql, cart.ProductID).Scan(&product.ID, &product.Name, &product.CategoryID, &product.Price, &product.Description, &product.CreatedAt, &product.UpdatedAt)
		if err2 != nil {
			panic(err2)
		}

		cart.Product = product
		order.Cart = cart

		results = append(results, order)
	}

	return
}
