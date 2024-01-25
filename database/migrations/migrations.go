package migrations

import (
	"fmt"
	"os"

	"github.com/RajaSunrise/crud-fiber/database"
	"github.com/RajaSunrise/crud-fiber/models/entity"
)

func Migrate() {
	err := database.DB.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stdout, []any{"succes to migrate"}...)
}
