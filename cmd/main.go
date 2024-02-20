package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/k6mil6/test-app/internal/config"
	"github.com/k6mil6/test-app/internal/printer"
	"github.com/k6mil6/test-app/internal/storage/postgres"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	cfg := config.Get()

	db, err := sqlx.Connect("postgres", cfg.DatabaseDSN)
	if err != nil {
		log.Println("failed to connect to database: ", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	//if err := initialisation.Init(ctx, db); err != nil {
	//	log.Println("failed to initialisation: ", err)
	//}

	arg := os.Args[1]
	orderIDs := strings.Split(arg, ",")

	orderStorage := postgres.NewOrderStorage(db)
	shelfStorage := postgres.NewShelfStorage(db)

	page, err := printer.PrintPage(ctx, orderStorage, shelfStorage, orderIDs)
	if err != nil {
		fmt.Println("Error generating assembly page:", err)
		return
	}

	fmt.Println(page)

}
