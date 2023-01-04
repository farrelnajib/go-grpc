package main

import (
	"fmt"
	"github.com/farrelnajib/go-rpc/discount"
	"github.com/farrelnajib/go-rpc/discount_impl/accessor"
	"github.com/farrelnajib/go-rpc/discount_impl/impl"
	"github.com/farrelnajib/go-rpc/product"
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

	conn.AutoMigrate(&discount.DiscountORM{}, &discount.RowORM{})

	return conn, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Could not load env: %v", err)
	}

	port := os.Getenv("RPC_PORT")
	db, err := dbSetup()
	if err != nil {
		log.Fatalf("Failed to connect db %v", err.Error())
	}

	discountAccessor := accessor.NewAccessor(db)
	productService := product.NewClientService(os.Getenv("PRODUCT_PORT"))

	discountService := impl.NewDiscountService(
		discountAccessor,
		productService,
	)

	server := grpc.NewServer()
	discount.RegisterDiscountServiceServer(server, discountService)
	log.Println("Starting RPC server at", port)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Could not listen to %v", port)
	}

	log.Fatal(server.Serve(listener))
}
