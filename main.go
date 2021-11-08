//code by vikash parashar 7/11/21 09:08pm
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type user struct {
	ID    int    `json:"id"`
	NAME  string `json:"name"`
	AGE   int    `json:"age"`
	EMAIL string `json:"email"`
}

var DB *sql.DB

func main() {
	var opts int
	fmt.Println("Let's Connect To Database And Create A Table In Our Database")
	GetDb()
	defer GetDb().Close() // after performing the operations lets close the db , when we done !

	var opt2 int

	for i := 1; i == opt2; i++ {
		fmt.Println("Choose From Below Operation Which You Want To Perform With Database . Options Are Mention Below : 1-5")
		fmt.Println("1. Create A New User In Table In Database")
		fmt.Println("2. Get All Users From Table in Database")
		fmt.Println("3. Get Single User From Table in Database")
		fmt.Println("4. Update User From Table in Database")
		fmt.Println("5. Delete User From Table in Database")
		fmt.Scanf("%d", &opts)

		switch opts {
		case 1:
			CreateUser(DB)
		case 2:
			GetAllUsers(DB)
		case 3:
			GetUserById(DB)
		case 4:
			UpdateUser(DB)
		case 5:
			DeleteUser(DB)

		}
		fmt.Println("Do You Wish To Continue ?")
		fmt.Println(" Press 1 For Continue , 0 To Exit !")
		fmt.Scanf("%d", &opt2)
	}

}

// connection config. to our database
func GetDb() (db *sql.DB) {
	// os.Remove("./test.db") // lets remove existing database
	// os.Create("./test.db") // lets create a new file for our database
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You Are Now Connected To Database")
	DB = db
	//defer db.Close()
	CreateTable(db)
	return db
}

// lets create a new table in our database in which we store our users infoo/data/entities
func CreateTable(db *sql.DB) {
	myschema := `
	CREATE TABLE IF NOT EXISTS USERS(
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		NAME TEXT NOT NULL,
		AGE INTEGER NOT NULL,
		EMAIL TEXT NOT NULL
	)
	;
	`
	stmt, err := db.Prepare(myschema)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table Is Created Successfullu .")
	n, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("No. Of Rows Affected Are : %d", n)

}

// lets get all the data from our users table in database
func GetAllUsers(db *sql.DB) {
	// myquery1 := `
	// SELECT * FROM USERS;
	// `

	// var user user
	// rows, err := db.Query(myquery1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// //var Users []user
	// for rows.Next() {
	// 	rows.Scan(&user.ID, &user.NAME, &user.AGE, &user.EMAIL)
	// }
	row, err := db.Query("SELECT * FROM USERS ORDER BY ID")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var Id int
		var Name string
		var Age int
		var Email string
		row.Scan(&Id, &Name, &Age, &Email)
		log.Println("User: ", Id, " ", Name, " ", Age, " ", Email)
	}

}

// lest get user from our database with his id
func GetUserById(db *sql.DB) {
	myquery2 := `
	SELECT * FROM USERS WHERE ID = (?);
	`
	var user user
	rows, err := db.Query(myquery2)
	if err != nil {
		log.Fatal(err)
	}
	//var Users []user
	for rows.Next() {
		rows.Scan(&user.ID, &user.NAME, &user.AGE, &user.EMAIL)
	}
	fmt.Println("Result As Per Your Choice")
	fmt.Println(user)
	defer rows.Close()

}

// lets create a new user/entry in our users table
func CreateUser(db *sql.DB) {
	mystatement := `
	INSERT INTO USERS(NAME , AGE , EMAIL) VALUES(?,?,?);
	`
	var user user
	// var Name string
	// var Age int
	// var Email string
	stmt, err := db.Prepare(mystatement)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Enter User Name :")
	fmt.Scanf("%s", &user.NAME)
	fmt.Println("Enter User Age :")
	fmt.Scanf("%d", &user.AGE)
	fmt.Println("Enter User Email :")
	fmt.Scanf("%s", &user.EMAIL)
	res, err := stmt.Exec(user.NAME, user.AGE, user.EMAIL)
	if err != nil {
		log.Fatal(err)
	}
	n, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("No.Of Rows Affected !", n)

}

// delete an entry our users table
func DeleteUser(db *sql.DB) {
	// myquery3 := `
	// DELETE FROM USERS WHERE ID = (?);
	// `
}

// UPDATING EXISTING USER IN DATABASE
func UpdateUser(db *sql.DB) {

	myquery4 := `
	UPDATE USERS
	SET NAME = ?, AGE = ?, EMAIL=?
	WHERE ID = (?);
	
	`
	var user user
	fmt.Println("Enter User ID For Which you Want To Update Data ")
	fmt.Scanf("%d", &user.ID)
	fmt.Println("Enter User Name :")
	fmt.Scanf("%s", &user.NAME)
	fmt.Println("Enter User Age :")
	fmt.Scanf("%d", &user.AGE)
	fmt.Println("Enter User Email :")
	fmt.Scanf("%s", &user.EMAIL)
	stmt, err := db.Prepare(myquery4)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(user.NAME, user.AGE, user.EMAIL, &user.ID)
	if err != nil {
		log.Fatal(err)
	}
	n, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("No.Of Rows Affected !", n)

}
