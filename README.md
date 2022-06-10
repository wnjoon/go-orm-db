# go-orm-db
Test ORM using golang

## Oracle

Wish to use ORM but no reference and libraries found.  
And There is **NO WAY** to use Oracle DB in **MAC M1**. -> Installed it to AWS Ubuntu 20.04.  

> Posted on my blog : [Handling oracle database using golang](https://wnjoon.github.io/go/oraclego/)

## MariaDB

### Docker

1. Get Docker image

```sh
$ docker pull mariadb
# If you want another verion (10.3)
$ docker pull mariadb:10.3
```

2. Start container

```sh
$ docker run --name testdbcontainer -e MYSQL_ROOT_PASSWORD=test -p 3306:3306 -d mariadb
# If you want to run another version (10.3)
$ docker run --name testdbcontainer -e MYSQL_ROOT_PASSWORD=test -p 3306:3306 -d mariadb:10.3
```

3. Execute

```sh
$ docker exec -it testdbcontainer mysql -u root -p 
```

### go.mod

```sh
# GORM
$ go get -u github.com/jinzhu/gorm
# MySQL driver for GORM
$ go get github.com/go-sql-driver/mysql
```

### Create database and user with grant

```sh
> create database bcstatus;
> create user 'user'@'%' identified by 'userpw';
> select host,user from mysql.user;
> grant all privileges on bcstatus.* to 'user'@'localhost';
> rename user 'user'@'localhost' to 'user'@'%';
```

### Create Table
```sh
> 
// CREATE TABLE statuses(  
//     id INT NOT NULL AUTO_INCREMENT,  
//     tx_id VARCHAR(100) NOT NULL,  
//     code VARCHAR(40) NOT NULL,      
//     PRIMARY KEY ( id )
// );
```


### Log

Set global query for tracking queries only generated from mysql.  

```sh
# Inside MariaDB
# -> docker exec -it testdbcontainer mysql -u root -p 

MariaDB [(none)]> SET GLOBAL general_log = 1;
Query OK, 0 rows affected (0.00 sec)
MariaDB [(none)]> SELECT @@log_output, @@general_log, @@general_log_file;
+--------------+---------------+--------------------+
| @@log_output | @@general_log | @@general_log_file |
+--------------+---------------+--------------------+
| FILE         |             1 | mach-W650EH.log    |
+--------------+---------------+--------------------+
1 row in set (0.00 sec)
MariaDB [(none)]>
```


## References
- [gorm.io](https://gorm.io/)
- [Connecting a Go application to Oracle Database - Developers Blog, Oracle](https://blogs.oracle.com/developers/post/connecting-a-go-application-to-oracle-database)
- [lucasjellema/go-oracle-database - github](https://github.com/lucasjellema/go-oracle-database)
- [mariadb - using in mac M1](https://two-track.tistory.com/21)
- [Go언어에서 ORM, gORM - Medium](https://medium.com/@amoebamach/go%EC%96%B8%EC%96%B4%EC%97%90%EC%84%9C-orm-gorm-83ab33ecdc98)