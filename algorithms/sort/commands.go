package sort

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"log"
)

func GetSortCommands() cli.Commands {
	return cli.Commands{
		{
			Name: "selection_sort",
			Usage: "Sort array with selection sort algorithm",
			Action: func(c *cli.Context) error {
				var slice []int

				err := json.Unmarshal([]byte(c.Args().Get(0)), &slice)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(selectionSort(slice))

				return nil
			},
		},
	}
}