package oracle

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

func handleErr(err error, errMsg string) {
	if err != nil {
		log.Panic(errors.Wrap(err, errOpenDB))
	}
}

func connect() *sql.DB {
	env()
	return open()
}

func open() *sql.DB {
	db, err := sql.Open("oracle", getConnectionInfo())
	handleErr(err, errOpenDB)
	return db
}

func execute(sql string) {
	db := connect()
	defer func() {
		err := db.Close()
		handleErr(err, errCloseDB)
	}()
	_, err := db.Exec(sql)
	handleErr(err, errExecStmt)
}

func prepareAndExcute(org, mnemonic, passphrase string) int64 {
	db := connect()
	defer func() {
		err := db.Close()
		handleErr(err, errCloseDB)
	}()
	sql := fmt.Sprintf(insertRow, os.Getenv("ORACLE_TABLE_NAME"))
	stmt, err := db.Prepare(sql)
	handleErr(err, errPrepare)

	execResult, err := stmt.Exec(org, mnemonic, passphrase)
	handleErr(err, errExecStmt)

	row, err := execResult.RowsAffected()
	handleErr(err, errRowUpdate)
	return row
}

func getRow(org string) (string, string) {
	db := connect()
	defer func() {
		err := db.Close()
		handleErr(err, errCloseDB)
	}()
	var (
		mnemonic   string
		passphrase string
	)

	sql := fmt.Sprintf(selectRow, os.Getenv("ORACLE_TABLE_NAME"), org)
	row := db.QueryRow(sql)
	err := row.Scan(&mnemonic, &passphrase)
	handleErr(err, errScanRow)
	return mnemonic, passphrase
}

func env() {
	err := godotenv.Load(".env")
	if err != nil {
		handleErr(err, errLoadEnv)
	}
}

func getConnectionInfo() string {
	return "oracle://" +
		os.Getenv("ORACLE_USER_NAME") + ":" + os.Getenv("ORACLE_USER_PWD") +
		"@" + os.Getenv("ORACLE_SVR_IP") + ":" + os.Getenv("ORACLE_SVC_PORT") +
		"/" + os.Getenv("ORACLE_SVC_NAME")
}

// Error Response
const (
	errLoadEnv   string = ".env load failed"
	errOpenDB    string = "database open failed"
	errCloseDB   string = "database close failed"
	errExecStmt  string = "execute statement failed"
	errPrepare   string = "prepare statement falied"
	errRowUpdate string = "row update failed"
	errQueryRow  string = "query row failed"
	errScanRow   string = "next row in multiple rows"
)
