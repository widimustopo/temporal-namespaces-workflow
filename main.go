package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/widimustopo/temporal-namespaces-workflow/activities"
	database "github.com/widimustopo/temporal-namespaces-workflow/databases"
	"github.com/widimustopo/temporal-namespaces-workflow/libs"
	"github.com/widimustopo/temporal-namespaces-workflow/workflows"
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
	paymentClient := initPaymentTemporalClient(loadConfig)
	productClient := initProductTemporalClient(loadConfig)

	workerRegister := initRegisterWorker(registerClient, db)
	workerOrder := initOrderWorker(paymentClient, db)
	workerExpired := initExpiredWorker(paymentClient, db)
	workerPayment := initPaymentWorker(paymentClient, db)
	workerPaymentFail := initPaymentFailWorker(paymentClient, db)
	workerProduct := initProductWorker(productClient, db)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-signals

	workerRegister.Stop()
	workerOrder.Stop()
	workerPayment.Stop()
	workerPaymentFail.Stop()
	workerExpired.Stop()
	workerProduct.Stop()

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

func initPaymentTemporalClient(cfg *libs.Config) client.Client {
	c, err := client.NewClient(client.Options{
		HostPort:  client.DefaultHostPort,
		Namespace: cfg.PaymentNamespaces,
	})

	if err != nil {
		logrus.Fatalln("Unable to create client", err)
	}

	return c
}

func initProductTemporalClient(cfg *libs.Config) client.Client {
	c, err := client.NewClient(client.Options{
		HostPort:  client.DefaultHostPort,
		Namespace: cfg.ProductNamespaces,
	})

	if err != nil {
		logrus.Fatalln("Unable to create client", err)
	}

	return c
}

func initRegisterWorker(temporalClient client.Client, db *gorm.DB) worker.Worker {
	wo := worker.Options{
		MaxConcurrentActivityExecutionSize: libs.MaxConcurrentSquareActivitySize,
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

func initOrderWorker(temporalClient client.Client, db *gorm.DB) worker.Worker {
	wo := worker.Options{
		MaxConcurrentActivityExecutionSize: libs.MaxConcurrentSquareActivitySize,
	}

	worker := worker.New(temporalClient, libs.OrderWorkflow, wo)

	worker.RegisterWorkflow(workflows.OrderWorkflow)
	worker.RegisterWorkflow(workflows.CounterProductWorkflow)
	newHandler := activities.HandlerActivities(db, temporalClient)
	worker.RegisterActivity(newHandler.Order)
	worker.RegisterActivity(newHandler.Counter)

	err := worker.Start()
	if err != nil {
		logrus.Error(err.Error())
		return nil
	}

	return worker
}

func initExpiredWorker(temporalClient client.Client, db *gorm.DB) worker.Worker {
	wo := worker.Options{
		MaxConcurrentActivityExecutionSize: libs.MaxConcurrentSquareActivitySize,
	}

	worker := worker.New(temporalClient, libs.ExpiredWorkflow, wo)

	worker.RegisterWorkflow(workflows.ExpiredWorkflow)

	newHandler := activities.HandlerActivities(db, temporalClient)
	worker.RegisterActivity(newHandler.Expired)

	err := worker.Start()
	if err != nil {
		logrus.Error(err.Error())
		return nil
	}

	return worker
}

func initPaymentWorker(temporalClient client.Client, db *gorm.DB) worker.Worker {
	wo := worker.Options{
		MaxConcurrentActivityExecutionSize: libs.MaxConcurrentSquareActivitySize,
	}

	worker := worker.New(temporalClient, libs.PaymentWorkflow, wo)

	worker.RegisterWorkflow(workflows.PaymentWorkflow)

	newHandler := activities.HandlerActivities(db, temporalClient)
	worker.RegisterActivity(newHandler.Payment)

	err := worker.Start()
	if err != nil {
		logrus.Error(err.Error())
		return nil
	}

	return worker
}

func initPaymentFailWorker(temporalClient client.Client, db *gorm.DB) worker.Worker {
	wo := worker.Options{
		MaxConcurrentActivityExecutionSize: libs.MaxConcurrentSquareActivitySize,
	}

	worker := worker.New(temporalClient, libs.PaymentFailWorkflow, wo)

	worker.RegisterWorkflow(workflows.PaymentFailWorkflow)

	newHandler := activities.HandlerActivities(db, temporalClient)
	worker.RegisterActivity(newHandler.PaymentFail)

	err := worker.Start()
	if err != nil {
		logrus.Error(err.Error())
		return nil
	}

	return worker
}

func initProductWorker(temporalClient client.Client, db *gorm.DB) worker.Worker {
	wo := worker.Options{
		MaxConcurrentActivityExecutionSize: libs.MaxConcurrentSquareActivitySize,
	}

	worker := worker.New(temporalClient, libs.AddProductWorkflow, wo)

	worker.RegisterWorkflow(workflows.AddProductWorkflow)

	newHandler := activities.HandlerActivities(db, temporalClient)
	worker.RegisterActivity(newHandler.AddProduct)

	err := worker.Start()
	if err != nil {
		logrus.Error(err.Error())
		return nil
	}

	return worker
}
