package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	database "github.com/widimustopo/temporal-namespaces-module/databases"
	"github.com/widimustopo/temporal-namespaces-module/libs"
	"github.com/widimustopo/temporal-namespaces-module/temporal/activities"
	"github.com/widimustopo/temporal-namespaces-module/temporal/workflows"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"gorm.io/gorm"
)

var loadConfig *libs.Config

func main() {
	//load config from .env
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal(err)
	}

	//load config from config.yml
	loadConfig, err = libs.NewEnvConfig()
	if err != nil {
		panic(err)
	}

	//initiate database connection
	db := database.OpenDB(loadConfig)

	registerClient := initRegisterTemporalClient(loadConfig)

	workerRegister := initRegisterWorker(registerClient, db)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-signals

	workerRegister.Stop()

}
func initRegisterTemporalClient(cfg *libs.Config) client.Client {
	c, err := client.NewClient(client.Options{
		HostPort:  client.DefaultHostPort,
		Namespace: cfg.MemberNamespaces,
	})

	if err != nil {
		logrus.Fatalln("Unable to create client", err)
	}

	return c
}

func initRegisterWorker(temporalClient client.Client, db *gorm.DB) worker.Worker {
	wo := worker.Options{
		//	MaxConcurrentActivityExecutionSize: libs.MaxConcurrentSquareActivitySize,
	}

	worker := worker.New(temporalClient, libs.RegisterWorkflow, wo)

	worker.RegisterWorkflow(workflows.RegisterWorkflow)

	newHandler := activities.HandlerActivities(db, temporalClient)
	worker.RegisterActivity(newHandler.Register)

	err := worker.Start()
	if err != nil {
		logrus.Error(err.Error())
		return nil
	}

	return worker
}
