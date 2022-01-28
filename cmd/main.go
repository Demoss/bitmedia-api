package main

import (
	"context"
	"encoding/json"
	"fmt"
	bitmedia_api "github.com/bitmedia-api"
	"github.com/bitmedia-api/config"
	"github.com/bitmedia-api/pkg/db/adaptor/mongodb"
	"github.com/bitmedia-api/pkg/handler"
	"github.com/bitmedia-api/pkg/repository"
	"github.com/bitmedia-api/pkg/service"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()

	cfg := config.GetConfig()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	conn := mongodb.ParamsToConnect{
		User:     cfg.MongoDB.Username,
		Host:     cfg.MongoDB.Host,
		Port:     cfg.MongoDB.Port,
		Password: cfg.MongoDB.Password,
		Database: cfg.MongoDB.Database,
		AuthDB:   cfg.MongoDB.AuthDB,
	}

	client, err := mongodb.NewMongo(ctx, conn)
	if err != nil {
		return err
	}

	repos := repository.NewRepository(client)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	res, err := http.Get(cfg.URL + os.Getenv("api-key"))
	if err != nil {
		return err
	}
	var block bitmedia_api.Block

	err = json.NewDecoder(res.Body).Decode(&block)
	if err != nil {
		return err
	}

	lastReceivedBlock, err := strconv.ParseInt(bitmedia_api.HexaNumberToInteger(block.Result), 16, 64)

	go func() {
		for {
			AddNewBlocks(ctx, lastReceivedBlock, cfg.URL, services)
		}
	}()

	Init1000Blocks(ctx, services, lastReceivedBlock)

	fmt.Println(handlers)

	srv := new(bitmedia_api.Server)
	go func() {
		if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("error occured on server shutting down: %s", err.Error())
	}
	return nil
}

func Init1000Blocks(ctx context.Context, service *service.Service, lastReceiveBlock int64) {

	var blockByNumber bitmedia_api.BlockByNumber

	for i := 0; i < 1; i++ {

		temp := lastReceiveBlock - int64(i)
		hexInt := strconv.FormatInt(temp, 16)
		hexInt = "0x" + hexInt
		fmt.Println(hexInt)

		url := fmt.Sprintf("https://api.etherscan.io/api?module=proxy&action=eth_getBlockByNumber&tag=%v&boolean=true&apikey=%s", hexInt, os.Getenv("api-key"))
		get, err := http.Get(url)
		if err != nil {
			return
		}
		err = json.NewDecoder(get.Body).Decode(&blockByNumber)
		if err != nil {
			return
		}
		err = service.SaveBlockByNumber(ctx, blockByNumber)
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func AddNewBlocks(ctx context.Context, lastReceivedBlock int64, url string, service *service.Service) {
	for {
		res, err := http.Get(url + os.Getenv("api-key"))
		if err != nil {
			return
		}
		var block bitmedia_api.Block

		err = json.NewDecoder(res.Body).Decode(&block)
		if err != nil {
			return
		}
		newBlock, err := strconv.ParseInt(bitmedia_api.HexaNumberToInteger(block.Result), 16, 64)
		if err != nil {
			return
		}

		if newBlock > lastReceivedBlock {
			fmt.Println("true")
			var blockByNumber bitmedia_api.BlockByNumber
			hexInt := strconv.FormatInt(newBlock, 16)
			hexInt = "0x" + hexInt
			fmt.Println("addNew")
			url := fmt.Sprintf("https://api.etherscan.io/api?module=proxy&action=eth_getBlockByNumber&tag=%v&boolean=true&apikey=%s", hexInt, os.Getenv("api-key"))
			get, err := http.Get(url)
			if err != nil {
				return
			}
			err = json.NewDecoder(get.Body).Decode(&blockByNumber)
			if err != nil {
				return
			}
			err = service.SaveBlockByNumber(ctx, blockByNumber)
			if err != nil {
				return
			}
			lastReceivedBlock = newBlock
		}
		time.Sleep(200 * time.Millisecond)
	}

}
