package db

import (
	"fmt"

	"github.com/marveldo/gogin/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
func Setup(c *config.Config) (*gorm.DB , error) {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		c.Host, c.DBUser, c.DBPassword, c.DBname, c.Port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		 TranslateError: true,
	})
}

func Get_db_models() []interface{} {
	var models []interface{}
    models = append(models, &TestModel{}, &UserModel{})
	return models
}