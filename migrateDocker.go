package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func migrateDocker() {

	// var user = "root"
	// var password = "secret"
	// var ip = "172.17.0.2:3306"
	// var dbname = "testdb"

	db, err := sql.Open("mysql", user+":"+password+"@tcp("+ip+")/")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

	_, err = db.Exec("CREATE DATABASE if not exists " + dbname)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE " + dbname)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("CREATE TABLE if not exists person (id int NOT NULL AUTO_INCREMENT, name varchar(40), PRIMARY KEY (id));")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Printf("Users Table successfully migrated....\n")
	}

	stmt, err = db.Prepare("insert ignore into person (id, name) values(?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(nil, "Joao Silva")
	_, err = stmt.Exec(nil, "Jose Souza")
	_, err = stmt.Exec(nil, "Junior Joao")
	_, err = stmt.Exec(nil, "Jose Junior")

	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()

	stmt, err = db.Prepare("CREATE TABLE if not exists widgets (id int NOT NULL AUTO_INCREMENT, name varchar(40), color varchar(40), price varchar(10), melts varchar(10), inventory int, PRIMARY KEY (id));")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Printf("widgets Table successfully migrated....\n")
	}

	stmt, err = db.Prepare("insert ignore into widgets (id, name, color, price, melts, inventory) values(?,?,?,?,?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(nil, "Gold", "red", "4.22", "yes", 23)
	_, err = stmt.Exec(nil, "Silver", "magenta", "2.22", "no", 232)
	_, err = stmt.Exec(nil, "Bronze", "white", "5.22", "yes", 2)

	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()
}

// insert into `widgets` (`id`, `name`, `color`,`price`,`melts`,`inventory`) values (null,"Bastiao Jose","red","5.23", true, 24);
//insert into `widgets` (`id`, `name`, `color`,`price`,`melts`,`inventory`) values (null,"Testando Teste","green","10.27", false, 11);
