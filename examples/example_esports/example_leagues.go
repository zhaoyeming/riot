// Example script demonstrating how to list out available leagues.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yuhanfang/riot/esports"
)

func main() {
	for i := 1; i <= 15; i++ {
		ctx := context.Background()
		client := esports.NewClient(http.DefaultClient)

		// Select a tournament, and pick a match in that tournament.
		leagues, err := client.GetLeagues(ctx, esports.League(int64(i)))
		fmt.Println(err)
		if len(leagues.Leagues) > 0 {
			fmt.Println(leagues.Leagues[0].ID)
			fmt.Println(leagues.Leagues[0].Slug)
			fmt.Println(leagues.Leagues[0].Name)
			fmt.Println()
			fmt.Println()
		}
	}
}
