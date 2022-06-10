package oracle

import (
	"fmt"
	"os"

	_ "github.com/sijms/go-ora/v2"
)

type BIP struct {
	Mnemonic   string
	Passphrase string
}

func CreateTable(name string) {
	// env()
	sql := fmt.Sprintf(createTable, name)
	execute(sql)
}

func CreateUser(name string, password string) {
	sql := fmt.Sprintf(createUser, name, password, os.Getenv("ORACLE_TABLE_TEMP_NAME"))
	fmt.Println(sql)
	execute(sql)
	authority(name)
}

func authority(name string) {
	execute(fmt.Sprintf("GRANT CONNECT TO %s", name))
	execute(fmt.Sprintf("GRANT RESOURCE TO %s", name))
	execute(fmt.Sprintf("GRANT DBA TO %s", name))
}

func Insert(org, mnemonic, passphrase string) {
	prepareAndExcute(org, mnemonic, passphrase)
}

func Select(org string) *BIP {
	mnemonic, passphrase := getRow(org)
	return &BIP{Mnemonic: mnemonic, Passphrase: passphrase}
}
