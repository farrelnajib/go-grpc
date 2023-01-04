package main

import (
	"fmt"
	"github.com/farrelnajib/go-rpc/product"
	"github.com/farrelnajib/go-rpc/product_impl/accessor"
	"github.com/farrelnajib/go-rpc/product_impl/impl"
	"github.com/farrelnajib/go-rpc/product_impl/nonrpchandler"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net"
	"os"
)

func dbSetup() (*gorm.DB, error) {
	username := os.Getenv("PSQL_USER")
	password := os.Getenv("PSQL_PASSWORD")
	dbName := os.Getenv("PSQL_DBNAME")
	dbHost := os.Getenv("PSQL_HOST")
	dbPort := os.Getenv("PSQL_PORT")
	envMode := os.Getenv("ENV_MODE")
	ssl := os.Getenv("SSL_MODE")

	dbUri := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=Asia/Jakarta",
		dbHost,
		username,
		password,
		dbName,
		dbPort,
		ssl,
	)

	enableLogging := envMode != "prod"

	conn, err := gorm.Open(
		postgres.Open(dbUri),
		&gorm.Config{},
	)

	if err != nil {
		return nil, err
	}

	if enableLogging {
		conn.Config.Logger = logger.Default.LogMode(logger.Info)
	}

	conn.AutoMigrate(&product.ProductORM{}, &product.ProductVariantORM{})

	return conn, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Could not load env: %v", err)
	}

	port := os.Getenv("RPC_PORT")
	nonRPCPort := os.Getenv("NON_RPC_PORT")
	db, err := dbSetup()
	if err != nil {
		log.Fatalf("Failed to connect db %v", err.Error())
	}

	productAccessor := accessor.NewAccessor(db)

	productService := impl.NewProductService(
		productAccessor,
	)

	server := grpc.NewServer()
	product.RegisterProductServiceServer(server, productService)
	log.Println("Starting RPC server at", port)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Could not listen to %v", port)
	}

	router := fiber.New()
	router.Get("/test-handler", nonrpchandler.SampleHandler())

	errs := make(chan error, 2)
	go func() {
		errs <- server.Serve(listener)
	}()
	go func() {
		errs <- router.Listen(nonRPCPort)
	}()

	for err := range errs {
		if err != nil {
			log.Fatalf("Error serving: %v:", err.Error())
		}
	}
}
