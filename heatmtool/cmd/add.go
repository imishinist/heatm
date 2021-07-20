package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/imishinist/heatmtool"
	"github.com/spf13/cobra"
)

var (
	addFiles []string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add add heatmap counts",
	Run: func(cmd *cobra.Command, args []string) {
		heatmaps := make([]*heatmtool.HeatMap, 0, len(addFiles))
		for _, file := range addFiles {
			f, err := os.Open(file)
			if err != nil {
				log.Fatal(err)
			}

			var hm heatmtool.HeatMap
			if err := json.NewDecoder(f).Decode(&hm); err != nil {
				log.Fatal(err)
			}
			heatmaps = append(heatmaps, &hm)
		}

		dist := heatmaps[0]
		for i := 1; i < len(heatmaps); i++ {
			if err := dist.Add(heatmaps[i]); err != nil {
				log.Fatal(err)
			}
		}

		json.NewEncoder(os.Stdout).Encode(dist)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringArrayVar(&addFiles, "file", addFiles, "")
}
