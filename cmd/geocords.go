package cmd

import (
	"github/johnwesonga/gowapi/api"

	"github.com/spf13/cobra"
)

var geocordsCmd = &cobra.Command{
	Use:   "geocords",
	Short: "returns lat,lon given a city",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		api.GetGeoCords(city, countryCode)

	},
}

var city, countryCode string

func init() {
	rootCmd.AddCommand(geocordsCmd)
	geocordsCmd.Flags().StringVarP(&city, "city", "c", "", "city name (required)")
	geocordsCmd.MarkFlagRequired("city")
	geocordsCmd.Flags().StringVarP(&countryCode, "code", "", "", "country code e.g. GB (optional)")
}
