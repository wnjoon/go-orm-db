package oracle

const (
	// 인덱스, 암호화된 니모닉, 암호화된 패스프레이즈
	createTable = "CREATE TABLE %s ( ORG VARCHAR2(100), MNEMONIC VARCHAR2(1000), PASSPHRASE VARCHAR2(100))"
	// dropTable   = "DROP TABLE MNEMONIC_TABLE PURGE"
	createUser = "CREATE USER %s IDENTIFIED BY %s DEFAULT TABLESPACE USERS TEMPORARY TABLESPACE %s "
	insertRow  = "INSERT INTO %s ( ORG, MNEMONIC, PASSPHRASE) VALUES (:org, :mnemonic, :passphrase)"
	selectRow  = "SELECT MNEMONIC, PASSPHRASE FROM %s WHERE ORG = '%s'"
)
