package db_utils

import (
	"context"
	"fmt"
	"log"
)

func TruncateAll() {
	ctx := context.Background()
	rows, err := Conn.Query(ctx, `
        SELECT table_name, table_schema
        FROM information_schema.tables
        WHERE table_schema IN('public') AND table_type = 'BASE TABLE'
    `)
	if err != nil {
		log.Fatalf("Failed to fetch table names: %v\n", err)
	}

	// Truncate each table
	for rows.Next() {
		var tableName string
		var tableSchema string
		if err := rows.Scan(&tableName, &tableSchema); err != nil {
			log.Fatalf("Failed to scan table name: %v\n", err)
		}

		var table string

		if tableSchema != "public" {
			table = fmt.Sprintf("%s.%s", tableSchema, tableName)
		} else {
			table = tableName
		}

		_, err := Conn.Exec(ctx, fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", table))
		if err != nil {
			log.Fatalf("Failed to truncate table %s: %v\n", tableName, err)
		}
		// fmt.Printf("Truncated table: %s\n", tableName)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating rows: %v\n", err)
	}

	// fmt.Println("All tables truncated successfully.")
}
