// package base_model

// import (
// 	"database/sql"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
//   )

// postgresDB, err := sql.Open("postgres", "mydb_dsn")
// gormDB, err := gorm.Open(postgres.New(postgres.Config{
// 	Conn: postgresDB,
// }), &gorm.Config{})