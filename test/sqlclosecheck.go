package main

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

type Dept struct {
	DeptNo int    `json:"deptno"`
	Dname  string `json:"dname"`
	Loc    string `json:"loc"`
}

func sqlclosecheck(ctx context.Context, sqlxDB *sqlx.DB) {
	rows, err := sqlxDB.QueryxContext(ctx, "SELECT * FROM dept WHERE `loc` = ?", "BOSTON")
	if err != nil {
		log.Fatal(err)
	}
	//defer rows.Close()
	{
		// 模拟发生错误，没有关闭就直接返回
		return
	}
	var list = make([]Dept, 0, 30)
	for rows.Next() {
		var dept Dept
		err := rows.StructScan(&dept)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, dept)
	}
	log.Printf("list: %+v", list)
}
