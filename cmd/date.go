package cmd

import (
	"fmt"

	"github.com/cloudingcity/gool/pkg/date"
	"github.com/spf13/cobra"
)

var dateToTimestampCmd = &cobra.Command{
	Use:                   "date-to-timestamp [date]",
	Aliases:               []string{"d2t"},
	Short:                 "Covert date to unix timestamp",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		timestamp, err := date.ToTimestamp(args[0])
		if err != nil {
			fmt.Println("Invalid date!")
			return
		}
		fmt.Println(timestamp)
	},
}
