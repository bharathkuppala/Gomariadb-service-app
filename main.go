package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	kafkaservice "github.com/mariaDB/module/kafka-service"
	utility "github.com/mariaDB/module/utilities"

	"github.com/mariaDB/module/student"

	"github.com/mariaDB/module/config"
	"github.com/mariaDB/module/database/gormdb"
)

var (
	err error
)

func main() {
	utility.Logger = log.New(os.Stdout, "", log.Lmicroseconds|log.Lshortfile)

	conf := config.ReadEnv()
	if conf == nil {
		log.Println("Failed to get env variables from .yml file")
		return
	}

	utility.GormDatabase, err = gormdb.NewGormConfig(conf)
	if err != nil {
		utility.Logger.Println("problem in establishing connection: "+err.Error(), err)
		return
	}
	fmt.Println("After gorm in gorm????????????????????????")
	// Create table in db ie mariadb_test
	utility.GormDatabase.AutoMigrate(&student.Student{})

	// create producer
	utility.Producer, err = kafkaservice.InitProducer()
	if err != nil {
		utility.Logger.Println("Error producer: ", err.Error())
		return
	}

	router := http.NewServeMux()

	// router implementation
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/api/v1/register", student.Register)
	router.HandleFunc("/api/v1/login", student.Login)

	// Configuring server ...
	server := &http.Server{
		Addr:    ":" + fmt.Sprint(conf.Server.Port),
		Handler: router,
	}

	flag.String("Listening on port", fmt.Sprint(conf.Server.Port), "port to listen on")
	flag.VisitAll(func(flag *flag.Flag) {
		log.Println(flag.Name, "->", flag.Value)
	})

	// Calls serve to handle incoming requst ...
	if err := server.ListenAndServe(); err != nil {
		utility.Logger.Println("error in spinning up the server: " + err.Error())
		return
	}

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, welcome to mariadb web app"))
}
