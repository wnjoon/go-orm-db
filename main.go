package main

import (
	"fmt"

	"github.com/wnjoon/go-orm-db/oracle"
)

func main() {
	// oracle.CreateUser("admin", "adminpw")
	// oracle.CreateTable("MNEMONIC_TABLE")
	// oracle.Insert("COMPANY",
	// "leaf senior buddy flower caught faint season come laundry tobacco elder coast seat pupil manual whip ability category gloom crumble aware water because topic",
	// "passphrase")
	bip := oracle.Select("COMPANY")
	fmt.Println("Mnemonic: ", bip.Mnemonic, "\nPassphrase: ", bip.Passphrase)
}
