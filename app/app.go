package app

import (
	"fmt"
	"github.com/Dontunee/banking/domain"
	"github.com/Dontunee/banking/logger"
	"github.com/Dontunee/banking/service"
	"github.com/jmoiron/sqlx"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func sanityCheck() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbAddress := os.Getenv("DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	serverAddress := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")

	if dbUser == "" {
		logger.Fatal("DB_USER Environment variable is not set")
	}
	if dbPassword == "" {
		logger.Fatal("DB_PASSWORD Environment variable is not set")
	}
	if dbAddress == "" {
		logger.Fatal("DB_ADDRESS Environment variable is not set")
	}
	if dbPort == "" {
		logger.Fatal("DB_PORT Environment variable is not set")
	}
	if dbName == "" {
		logger.Fatal("DB_NAME Environment variable is not set")
	}
	if serverAddress == "" {
		logger.Fatal("SERVER_ADDRESS Environment variable is not set")
	}
	if serverPort == "" {
		logger.Fatal("SERVER_PORT Environment variable is not set")
	}
}
func Start() {

	sanityCheck()
	router := mux.NewRouter()

	//wiring
	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	transactionRepositoryDb := domain.NewTransactionRepositoryDb(dbClient)
	customerHandler := CustomerHandlers{service: service.NewCustomerService(customerRepositoryDb)}
	accountHandler := AccountHandlers{service: service.NewAccountService(accountRepositoryDb)}
	transactionHandler := TransactionHandlers{service: service.NewTransactionService(accountRepositoryDb, transactionRepositoryDb)}

	//define the route and handler which provides response and brings request from client

	router.
		HandleFunc("/customers", customerHandler.getAllCustomers).
		Methods(http.MethodGet).
		Name("GetAllCustomers")

	router.
		HandleFunc("/customers/{customer_id:[0-9]+}", customerHandler.getCustomerById).
		Methods(http.MethodGet).
		Name("GetCustomer")

	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/account", accountHandler.createAccount).
		Methods(http.MethodPost).
		Name("NewAccount")

	router.
		HandleFunc("/transactions/create/{account_id:[0-9]+}", transactionHandler.createTransaction).
		Methods(http.MethodPost).
		Name("NewTransaction")

	authMiddleware := AuthMiddleware{domain.NewAuthRepository()}
	router.Use(authMiddleware.authorizationHandler())

	// starts the server on asn ip and port and use default multiplexer

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router)
	if err != nil {
		logger.Fatal("unable to start server with error" + err.Error())
	}
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbAddress := os.Getenv("DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	client, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddress, dbPort, dbName))

	if err != nil {
		panic(err)
	}

	//see "Important settings"  section
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
