package main

import (
	"github.com/wnjoon/go-orm-db/maria"
)

func main() {
	oracleTest()
	db := maria.Connect("admin", "adminpw", "localhost", "3306", "blockchain")
	// maria.Insert("0x10319230812", "1", db)
	maria.Query("0x10319230812", db)

}

func oracleTest() {
	// oracle.CreateUser("admin", "adminpw")
	// oracle.CreateTable("MNEMONIC_TABLE")
	// oracle.Insert("COMPANY",
	// "leaf senior buddy flower caught faint season come laundry tobacco elder coast seat pupil manual whip ability category gloom crumble aware water because topic",
	// "passphrase")
	// bip := oracle.Select("COMPANY")
	// fmt.Println("Mnemonic: ", bip.Mnemonic, "\nPassphrase: ", bip.Passphrase)
}
