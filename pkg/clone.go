package pkg

import (
	"context"
	"fmt"
	"github.com/JuanRulliansyah/pgklone/utils"
	"github.com/jackc/pgx/v4/pgxpool"
)

func CloneDatabase(sourceURL, targetURL string) error {
	ctx := context.Background()

	sourcePool, err := pgxpool.Connect(ctx, sourceURL)
	if err != nil {
		return fmt.Errorf("failed to connect to source database: %w", err)
	}
	defer sourcePool.Close()
	fmt.Println("Connected to source database.")

	targetPool, err := pgxpool.Connect(ctx, targetURL)
	if err != nil {
		return fmt.Errorf("failed to connect to target database: %w", err)
	}
	defer targetPool.Close()
	fmt.Println("Connected to target database.")

	schemaSQL := `SELECT table_name FROM information_schema.tables WHERE table_schema = 'public';`
	rows, err := sourcePool.Query(ctx, schemaSQL)
	if err != nil {
		return fmt.Errorf("failed to fetch schema from source database: %w", err)
	}
	defer rows.Close()
	fmt.Println("Fetched schema from source database.")

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return fmt.Errorf("failed to scan row: %w", err)
		}

		fmt.Printf("Processing table: %s\n", tableName)

		createTableSQL, err := utils.GetCreateTableSQL(ctx, sourcePool, tableName)
		if err != nil {
			return fmt.Errorf("failed to get table schema: %w", err)
		}
		fmt.Printf("Create table SQL: %s\n", createTableSQL)
		if _, err := targetPool.Exec(ctx, createTableSQL); err != nil {
			return fmt.Errorf("failed to create table in target database: %w", err)
		}
		fmt.Printf("Table %s created. Starting data copy...\n", tableName)

		if err := utils.CopyData(ctx, sourcePool, targetPool, tableName); err != nil {
			return fmt.Errorf("failed to copy data for table %s: %w", tableName, err)
		}
	}

	fmt.Println("Database cloned successfully.")
	return nil
}
