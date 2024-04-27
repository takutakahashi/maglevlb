/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"math/rand"

	maglev "github.com/dgryski/go-maglev"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// simulateCmd represents the simulate command
var simulateCmd = &cobra.Command{
	Use:   "simulate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		u := uint64(12345)
		m := map[string]int{
			"a": 0,
			"b": 0,
			"c": 0,
		}
		names := []string{"a", "b", "c"}
		table := maglev.New(names, 3)
		for i := 0; i < 10000; i++ {
			n := rand.Uint64()
			m[names[table.Lookup(uint64(n))]]++
		}
		logrus.Info(m)
		logrus.Info(table.Lookup(u))
		table.Rebuild([]int{2})
		m = map[string]int{
			"a": 0,
			"b": 0,
			"c": 0,
		}
		for i := 0; i < 10000; i++ {
			n := rand.Uint64()
			m[names[table.Lookup(uint64(n))]]++
		}
		logrus.Info(m)
		logrus.Info(table.Lookup(u))
		table.Rebuild([]int{})
		m = map[string]int{
			"a": 0,
			"b": 0,
			"c": 0,
		}
		for i := 0; i < 10000; i++ {
			n := rand.Uint64()
			m[names[table.Lookup(uint64(n))]]++
		}
		logrus.Info(m)
		logrus.Info(table.Lookup(u))
	},
}

func init() {
	rootCmd.AddCommand(simulateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// simulateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// simulateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
