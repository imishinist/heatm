package cmd

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"

	"github.com/imishinist/heatm"
)

var (
	genimgFile     = "-"
	genimgFileName = "output.png"
)

// genimgCmd represents the genimg command
var genimgCmd = &cobra.Command{
	Use:   "genimg",
	Short: "generate image from heatmap json file",
	Run: func(cmd *cobra.Command, args []string) {
		in, err := input(genimgFile)
		if err != nil {
			cmd.PrintErrf("input error: %v", err)
			return
		}

		var heatmap heatm.HeatMap
		if err := json.NewDecoder(in).Decode(&heatmap); err != nil {
			cmd.PrintErrf("json decode error: %v", err)
			return
		}
		f, err := os.OpenFile(genimgFileName, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			cmd.PrintErr("open file error: %v", err)
			return
		}

		if err := output(f, "png", heatmap); err != nil {
			cmd.PrintErrf("generate output error: %v", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(genimgCmd)

	genimgCmd.Flags().StringVar(&genimgFileName, "output-file", genimgFileName, "")
	genimgCmd.Flags().StringVar(&genimgFile, "file", genimgFile, "")
}
