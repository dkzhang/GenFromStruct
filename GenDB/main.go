package main

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

func main() {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sql, args, err := psql.Select("*").From("elephants").Where("name IN (?,?)", "Dumbo", "Verna").ToSql()

	//sql == "INSERT INTO users (name,age) VALUES (?,?),(?,? + 5)"
	fmt.Printf("\nsql: %s \nargs: %v \nerr: %v", sql, args, err)
}
