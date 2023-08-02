/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/simpletextbr/fullcycle-ports-and-adapters/adapters/cli"
	"github.com/spf13/cobra"
)

var action string
var productId string
var productName string
var price float64

// pcliCmd represents the pcli command
var pcliCmd = &cobra.Command{
	Use:   "pcli",
	Short: "pcli is a cli to manipulate products",
	Long: `pcli is a cli created on the course Full Cycle - Ports and Adapters by Code Education.
	You can use it to manipulate products.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Run(&productService, action, productId, productName, price)
	},
}

func init() {
	rootCmd.AddCommand(pcliCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pcliCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pcliCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	pcliCmd.Flags().StringVarP(&action, "action", "a", "", "Action to be executed(create, update, delete, get)")
	pcliCmd.Flags().StringVarP(&productId, "id", "i", "", "Product ID")
	pcliCmd.Flags().StringVarP(&productName, "name", "n", "", "Product Name")
	pcliCmd.Flags().Float64VarP(&price, "price", "p", 0, "Product Price")
}
