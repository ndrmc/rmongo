package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ndrmc/analytics/pkg/database"

	_ "github.com/lib/pq"
	"github.com/ndrmc/analytics/pkg/common"
	"github.com/ndrmc/analytics/pkg/models"
)

func main() {
	conf := common.LoadConfiguration("/Users/yared/src/gospace/src/github.com/ndrmc/analytics/config.json")
	initDB(conf)
	loadOperation(1)
}

func countOperations() {
	count := models.HowManyOperations()
	fmt.Printf("Total number of Operations is: %d", count)
}

func loadOperation(id int64) {
	result := models.GetOperation(id)
	fmt.Printf("Found operation %s and %d dispatches", result.Name.String, len(result.Dispatches))

	// buf, err := json.MarshalIndent(result, "", "\t")
	// if err != nil {
	// 	common.LogError(err)
	// }
	// fmt.Println(string(buf))
}

func loadOperations() {
	total := models.GetAllOperations()
	ops, err := json.MarshalIndent(total, "", "\t")
	if err != nil {
		common.LogError(err)
	}
	fmt.Println("===========================")
	fmt.Println(string(ops))
}

func initDB(conf common.Config) {
	var err error

	pgInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.PgHost, conf.PgPort, conf.PgUser, conf.PgPass, conf.DbName)

	database.Con, err = sql.Open("postgres", pgInfo)
	if err != nil {
		log.Panicf("Error making connection to database. Detail: %s", err)
	}

	err = database.Con.Ping()
	if err != nil {
		log.Panicf("Error making connection to database. Detail: %s", err)
	}

	common.LogInfo("Successfuly connected to database")

}