package oracle

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/wnjoon/go-orm-db/util"
)

func connect() *sql.DB {
	util.Env()
	return open()
}

func open() *sql.DB {
	db, err := sql.Open("oracle", getConnectionInfo())
	util.HandleErr(err, util.ErrOpenDB)
	return db
}

func execute(sql string) {
	db := connect()
	defer func() {
		err := db.Close()
		util.HandleErr(err, util.ErrCloseDB)
	}()
	_, err := db.Exec(sql)
	util.HandleErr(err, util.ErrExecStmt)
}

func prepareAndExcute(org, mnemonic, passphrase string) int64 {
	db := connect()
	defer func() {
		err := db.Close()
		util.HandleErr(err, util.ErrCloseDB)
	}()
	sql := fmt.Sprintf(insertRow, os.Getenv("ORACLE_TABLE_NAME"))
	stmt, err := db.Prepare(sql)
	util.HandleErr(err, util.ErrPrepare)

	execResult, err := stmt.Exec(org, mnemonic, passphrase)
	util.HandleErr(err, util.ErrExecStmt)

	row, err := execResult.RowsAffected()
	util.HandleErr(err, util.ErrRowUpdate)

	return row
}

func getRow(org string) (string, string) {
	db := connect()
	defer func() {
		err := db.Close()
		util.HandleErr(err, util.ErrCloseDB)
	}()
	var (
		mnemonic   string
		passphrase string
	)

	sql := fmt.Sprintf(selectRow, os.Getenv("ORACLE_TABLE_NAME"), org)
	row := db.QueryRow(sql)
	err := row.Scan(&mnemonic, &passphrase)
	util.HandleErr(err, util.ErrScanRow)
	return mnemonic, passphrase
}

func getConnectionInfo() string {
	return "oracle://" +
		os.Getenv("ORACLE_USER_NAME") + ":" + os.Getenv("ORACLE_USER_PWD") +
		"@" + os.Getenv("ORACLE_SVR_IP") + ":" + os.Getenv("ORACLE_SVC_PORT") +
		"/" + os.Getenv("ORACLE_SVC_NAME")
}
