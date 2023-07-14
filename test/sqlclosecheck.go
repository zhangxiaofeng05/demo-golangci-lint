package main

import (
	"context"
	"github.com/jmoiron/sqlx"
	"log"
)

func sqlclosecheck(sqlxDB *sqlx.DB) {
	ctx := context.Background()
	rows, err := sqlxDB.QueryxContext(ctx, "SELECT deptno FROM dept WHERE `dname` = ?", "SALES")
	if err != nil {
		log.Fatal(err)
	}
	//defer rows.Close()
	for rows.Next() {
		var deptno int
		err := rows.Scan(&deptno)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("deptno: %v", deptno)
	}
}
