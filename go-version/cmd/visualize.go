/*
Copyright © 2022 Asabs18 & davegi
This file is part of the `nantucket` project.
*/
package cmd

import (
	"log"
	"os"

	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-gota/gota/dataframe"
	"github.com/spf13/cobra"
)

func readData(filename string) [][]opts.LineData {
	items := make([][]opts.LineData, 0)

	df := readCsv(filename)

	names := df.Names()

	//sel1 := df.Select([]int{1})

	for i := 1; i <= 4; i++ {
		series := df.Col(names[i])
		for j := 1; j < series.Len(); j++ {
			val, _ := series.Elem(j).Int()
			items[i] = append(items[i], opts.LineData{Value: val})
		}
	}
	return items
}

func readCsv(filename string) dataframe.DataFrame {
	file, err := os.Open(filename)

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	df := dataframe.ReadCSV(file)

	return df
}

// visualizeCmd represents the visualize command
var visualizeCmd = &cobra.Command{
	Use:   "visualize",
	Short: "Create a histogram from the input file",
	Long: `
_visualize_ creates a histogram of deaths, bounded by dmin and dmax.
 from the supplied input file, which must have been created by _analyze_'
 
 The resultant histogram is written to the supplied output file.`,
	Run: func(cmd *cobra.Command, args []string) {
		readData("cmd/simulations/simulation#2023-03-02-17-00-31.csv")
	},
}

func init() {
	rootCmd.AddCommand(visualizeCmd)

	// HACK: where to put config vars?
	var dmin float32
	var dmax float32
	var output_file string
	// # HACK: where should these strings live?
	// # BUGBUG: aren't hidden?
	visualizeCmd.Flags().MarkHidden("tprob")
	visualizeCmd.Flags().MarkHidden("vprob")
	visualizeCmd.Flags().Float32P("dmin", "n", dmin, "dmin")
	visualizeCmd.Flags().Float32P("dmax", "x", dmax, "dmax")
	visualizeCmd.Flags().String("output-file", output_file, "output_file")
}
