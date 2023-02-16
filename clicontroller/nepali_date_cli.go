package clicontroller

import (
	"fmt"
	"os"

	"github.com/fatih/color"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                   "nepalidate ",
	DisableFlagsInUseLine: true,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},

	Short: "Nepali Tithi",
	Long:  `CLi app to display nepalidate, time, festivals, panchang, thithi and english date.`,
	Run: func(cmd *cobra.Command, args []string) {

		aaja := Scrape()

		green := color.New(color.FgGreen).SprintFunc()
		red := color.New(color.FgRed).SprintFunc()
		fmt.Printf("\n")
		fmt.Println(red("\t\tDate(B.S):  "), green(aaja.Date))
		fmt.Println(red("\t\tDate(A.D):  "), green(aaja.EnglishDate))
		fmt.Println(red("\t\tTime:      "), green(aaja.Time))
		fmt.Println(red("\t\tEvents:     "), green(aaja.Event))
		fmt.Println(red("\t\tPanchang:   "), green(aaja.Panchang))
		fmt.Println(red("\t\tThithi:     "), green(aaja.Tithi))
		fmt.Printf("\n")

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
