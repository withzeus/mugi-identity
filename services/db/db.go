package db

import (
	"github.com/withzeus/mugi-identity/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ORM struct {
	*gorm.DB
}

type Result struct {
	Error        error
	RowsAffected int64
}

func NewORM() *ORM {
	dbUrl := config.AppEnv.DatabaseUrl
	if dbUrl == "" {
		panic("DATABASE URL not set")
	}

	DB, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	return &ORM{DB}
}

func (o *ORM) CreateResult(statement *gorm.DB) Result {
	return Result{
		Error:        statement.Error,
		RowsAffected: statement.RowsAffected,
	}
}
