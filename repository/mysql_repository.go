package repository

/*
This class stores the payouts into the database.
*/

import (
	"database/sql"
	"fmt"
	"strings"
	"takeHomeTest/errorHandlers"
	"takeHomeTest/models"
)

type Repository interface {
	StorePayouts(payouts []models.Payout) error
}

type repository struct {
	dbInstance   *sql.DB
	errorHandler errorHandlers.ErrorHandlers
}

func NewRepository(dbInstance *sql.DB, errorHandler errorHandlers.ErrorHandlers) Repository {
	return &repository{
		dbInstance:   dbInstance,
		errorHandler: errorHandler,
	}
}

/*
This method stores the created payouts to mysql database.
It prints the total affected rows which are equal to number of payout transactions
and returns err
*/
func (r *repository) StorePayouts(payouts []models.Payout) error {
	valueStrings := make([]string, 0, len(payouts))
	valueArgs := make([]interface{}, 0, len(payouts)*3)
	for _, post := range payouts {
		valueStrings = append(valueStrings, "(?, ?, ?)")
		valueArgs = append(valueArgs, post.SellerReference)
		valueArgs = append(valueArgs, post.Amount)
		valueArgs = append(valueArgs, post.Currency)
	}
	stmt := fmt.Sprintf("INSERT INTO Payout (SellerReference, Amount, Currency) VALUES %s",
		strings.Join(valueStrings, ","))

	tx, _ := r.dbInstance.Begin()
	result, err := r.dbInstance.Exec(stmt, valueArgs...)

	if err != nil {
		r.errorHandler.FailOnError(err, "Rolling back the database changes")
		tx.Rollback()
	} else {
		tx.Commit()
	}
	rows, _ := result.RowsAffected()
	fmt.Println("Affected rows:", rows)
	return err
}
