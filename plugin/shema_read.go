package plugin

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type TableColumn struct {
	ColumnName    string `gorm:"column:COLUMN_NAME"`
	DataType      string `gorm:"column:DATA_TYPE"`
	IsNullable    string `gorm:"column:IS_NULLABLE"`
	ColumnDefault string `gorm:"column:COLUMN_DEFAULT"`
}

func Generate(dbName string) {
	// Replace with your MySQL database connection string
	dsn := fmt.Sprintf("root:rootroot@tcp(localhost:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbName)

	// Initialize Gorm MySQL driver
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Logger configuration
	})
	if err != nil {
		fmt.Println("Failed to connect to database")
		os.Exit(1)
	}

	// Example of retrieving table schema
	var tables []string
	result := db.Raw("SHOW TABLES").Scan(&tables)
	if result.Error != nil {
		fmt.Println("Failed to retrieve tables:", result.Error)
		return
	}

	fmt.Println("Tables in the database:")
	content := ""
	for _, table := range tables {
		fmt.Println(table)
		content = content + printTableSchema(db, dbName, table)
	}

	GenerateSchemaFile(content, fmt.Sprintf("graph/schema.%s.graphqls", dbName))
}

func printTableSchema(db *gorm.DB, dbName string, tableName string) string {
	// Example of retrieving detailed table schema details
	var columns []TableColumn
	result := db.Raw("SELECT COLUMN_NAME, DATA_TYPE, IS_NULLABLE, COLUMN_DEFAULT FROM information_schema.columns WHERE table_schema = ? AND table_name = ?", dbName, tableName).Scan(&columns)
	if result.Error != nil {
		fmt.Println("Failed to retrieve schema for table", tableName, ":", result.Error)
		return ""
	}

	fmt.Println("Columns for table", tableName, ":")
	return GenerateContentSchema(tableName, columns)
}
