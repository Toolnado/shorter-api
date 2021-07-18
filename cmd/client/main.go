package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Toolnado/shorter-api/internal/api"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	var (
		url    string
		comand string
		ending bool
	)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("port not found")
	}

	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := api.NewLinkShortenerClient(conn)

	for !ending {
		fmt.Printf("To shorten the url, enter 'CREATE'\nTo get the original url, enter 'GET'\nTo exit the application, enter 'END'\n")
		fmt.Scanf("%v\n", &comand)
		switch comand {
		case "CREATE":
			fmt.Printf("Please enter url\n")
			fmt.Scanf("%v\n", &url)

			if url != "" {
				res, err := client.Create(context.Background(), &api.CreateRequest{Url: url})
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Abbreviated url for %s = %s\n", url, res.GetShortUrl())
			} else {
				fmt.Printf("Empty url\n")
			}

		case "GET":
			fmt.Printf("Please enter shorturl\n")
			fmt.Scanf("%v\n", &url)

			if url != "" {
				res, err := client.Get(context.Background(), &api.GetRequest{ShortUrl: url})
				if err != nil {
					log.Fatal(err)
				}
				if res.Url == "" {
					fmt.Printf("url not found\n")
				} else {
					fmt.Printf("Original url for %s = %s\n", url, res.GetUrl())
				}

			} else {
				fmt.Printf("Empty url\n")
			}

		case "END":
			fmt.Printf("Bye\n")
			ending = true

		default:
			fmt.Printf("Command %s not recognized\n", comand)
		}

	}
	os.Exit(0)
}
