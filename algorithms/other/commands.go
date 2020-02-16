package other

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"log"
)

func GetOtherCommands() cli.Commands {
	return cli.Commands{
		{
			Name: "sum",
			Usage: "Sum all values in array",
			Action: func(c *cli.Context) error {
				var slice []int

				err := json.Unmarshal([]byte(c.Args().Get(0)), &slice)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(sum(slice))

				return nil
			},
		},
	}
}