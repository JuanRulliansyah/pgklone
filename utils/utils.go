package utils

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"strings"
)

func GetCreateTableSQL(ctx context.Context, pool *pgxpool.Pool, tableName string) (string, error) {
	schemaSQL := fmt.Sprintf(`
        SELECT 'CREATE TABLE "%s" ('
        || array_to_string(
            array_agg(column_name || ' ' || column_type || 
            CASE WHEN is_nullable = 'NO' THEN ' NOT NULL' ELSE '' END),
            ', '
        )
        || ');'
        FROM (
            SELECT 
                column_name,
                CASE 
                    WHEN data_type = 'character varying' THEN 'VARCHAR(' || character_maximum_length || ')'
                    WHEN data_type = 'character' THEN 'CHAR(' || character_maximum_length || ')'
                    ELSE data_type 
                END AS column_type,
                is_nullable
            FROM information_schema.columns
            WHERE table_name = $1
        ) AS columns
    `, tableName)

	row := pool.QueryRow(ctx, schemaSQL, tableName)
	var createTableSQL string
	if err := row.Scan(&createTableSQL); err != nil {
		return "", fmt.Errorf("could not generate CREATE TABLE SQL for %s: %w", tableName, err)
	}

	return createTableSQL, nil
}

func CopyData(ctx context.Context, sourcePool, targetPool *pgxpool.Pool, tableName string) error {
	fmt.Printf("Copying data for table: %s\n", tableName)

	dataSQL := fmt.Sprintf(`SELECT * FROM "%s";`, tableName)
	rows, err := sourcePool.Query(ctx, dataSQL)
	if err != nil {
		return fmt.Errorf("could not query source table %s: %w", tableName, err)
	}
	defer rows.Close()

	columns := GetColumnNames(rows)
	if len(columns) == 0 {
		return fmt.Errorf("no columns found for table %s", tableName)
	}
	placeholders := GetPlaceholders(len(columns))
	insertSQL := fmt.Sprintf(`INSERT INTO "%s" (%s) VALUES (%s);`, tableName, strings.Join(columns, ", "), placeholders)

	fmt.Printf("Insert SQL: %s\n", insertSQL)

	tx, err := targetPool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	rowCount := 0
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return fmt.Errorf("could not retrieve row values: %w", err)
		}

		if _, err = tx.Exec(ctx, insertSQL, values...); err != nil {
			return fmt.Errorf("could not insert data into %s: %w", tableName, err)
		}

		rowCount++
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("error iterating rows: %w", err)
	}

	fmt.Printf("Copied %d rows for table %s\n", rowCount, tableName)
	return nil
}

func GetColumnNames(rows pgx.Rows) []string {
	columns := rows.FieldDescriptions()
	columnNames := make([]string, len(columns))
	for i, col := range columns {
		columnNames[i] = string(col.Name)
	}
	return columnNames
}

func GetPlaceholders(count int) string {
	placeholders := make([]string, count)
	for i := 0; i < count; i++ {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}
	return strings.Join(placeholders, ", ")
}
