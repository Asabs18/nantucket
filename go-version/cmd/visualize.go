/*
Copyright © 2022 Asabs18 & davegi
This file is part of the `nantucket` project.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// visualizeCmd represents the visualize command
var visualizeCmd = &cobra.Command{
	Use:   "visualize",
	Short: "Create a histogram from the input file",
	Long: `
_visualize_ creates a histogram of deaths, bounded by dmin and dmax.
 from the supplied input file, which must have been created by _analyze_'
 
 The resultant histogram is written to the supplied output file.`,
	Run: func(cmd *cobra.Command, args []string) {
		//func (c *Command) Flag(name string) (flag *flag.Flag)
		fmt.Println("visualize called")

		tprob := cmd.Flag("tProb")
		fmt.Println(tprob.Value)
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
