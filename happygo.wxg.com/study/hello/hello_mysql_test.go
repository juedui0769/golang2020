package hello

import (
	"database/sql"
	"fmt"
	"testing"
	//_ "github.com/Go-SQL-Driver/MySQL"
	_ "github.com/go-sql-driver/mysql"
)

func TestMysql(t *testing.T)  {
	fmt.Println()
	// username:password@protocol(address)/dbname?param=value
	db, err := sql.Open("mysql", "root:root@/books?charset=utf8")
	checkError(err)

	rows, err := db.Query("select * from t_book")
	checkError(err)

	for rows.Next() {
		var id int
		var name string
		var author string
		var category string
		var description string
		err = rows.Scan(&id, &name, &author, &category, &description)
		checkError(err)
		fmt.Printf("id: %d, name: %s, author: %s, category: %s, description: %s \r\n",
			id, name, author, category, description)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}


