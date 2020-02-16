package search

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"strconv"
)

func GetSearchCommands() cli.Commands {
	return cli.Commands{
		{
			Name: "binary_search",
			Usage: "Search value in sorted array",
			Action: func(c *cli.Context) error {
				var slice []int
				var value, err = strconv.Atoi(c.Args().Get(1))
				if err != nil {
					log.Fatal(err)
				}

				err = json.Unmarshal([]byte(c.Args().Get(0)), &slice)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(binarySearch(slice, value))

				return nil
			},
		},
	}
}