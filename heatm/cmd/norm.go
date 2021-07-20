package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"image/png"
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/imishinist/heatm"
)

var (
	normFile = "-"
)

// normCmd represents the norm command
var normCmd = &cobra.Command{
	Use:   "norm",
	Short: "norm normalize heatmap counts",
	Run: func(cmd *cobra.Command, args []string) {
		in, err := input(normFile)
		if err != nil {
			cmd.PrintErrf("input error: %v", err)
			return
		}

		var heatmap heatm.HeatMap
		if err := json.NewDecoder(in).Decode(&heatmap); err != nil {
			cmd.PrintErrf("json decode error: %v", err)
			return
		}

		heatmap.Normalize()
		json.NewEncoder(os.Stdout).Encode(heatmap)
	},
}

func init() {
	rootCmd.AddCommand(normCmd)

	normCmd.Flags().StringVar(&normFile, "file", normFile, "")
}

func input(file string) (*os.File, error) {
	if file == "-" {
		return os.Stdin, nil
	}
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func output(w io.Writer, output string, hm heatm.HeatMap) error {
	if output == "json" {
		if err := json.NewEncoder(w).Encode(hm); err != nil {
			return fmt.Errorf("json decode error: %w", err)
		}
		return nil
	} else if output == "png" {
		img := hm.Img()

		if err := png.Encode(w, img); err != nil {
			return fmt.Errorf("png encode error: %w", err)
		}
		return nil
	} else {
		return errors.New("unsupported output type")
	}
}
