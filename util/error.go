package util

import (
	"log"

	"github.com/pkg/errors"
)

// Error Response
const (
	ErrLoadEnv   string = ".env load failed"
	ErrOpenDB    string = "database open failed"
	ErrSqlDB     string = "sql database open failed"
	ErrCloseDB   string = "database close failed"
	ErrExecStmt  string = "execute statement failed"
	ErrPrepare   string = "prepare statement falied"
	ErrRowUpdate string = "row update failed"
	ErrQueryRow  string = "query row failed"
	ErrScanRow   string = "next row in multiple rows"
)

func HandleErr(err error, errMsg string) {
	if err != nil {
		log.Panic(errors.Wrap(err, errMsg))
	}
}
