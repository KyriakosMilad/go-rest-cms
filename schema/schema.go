package schema

import (
	"github.com/KyriakosMilad/go-rest-cms/database"
	"log"
)

var schemas = []interface{ GetTable() string }{
	// register schemas here
}

func Migrate() {
	for _, v := range schemas {
		err := database.CreateTable(v)
		if err != nil {
			log.Fatalln("error can't migrate " + v.GetTable() + " : " + err.Error())
		}
	}
}

func Drop() {
	for _, v := range schemas {
		err := database.DropTable(v)
		if err != nil {
			log.Fatalln("error can't drop " + v.GetTable() + " : " + err.Error())
		}
	}
}
