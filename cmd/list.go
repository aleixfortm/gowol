package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display a list of saved devices",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Device list:")

		type Device struct {
			Name       string `json:"name"`
			MACAddress string `json:"macAddress"`
		}

		// Read existing JSON file
		filePath := "data.json"
		fileData, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println("Failed to read JSON file:", err)
			return
		}
		// Stop function if there is no data to show
		if len(fileData) == 0 {
			return
		}
		// Unmarshal existing JSON data into slice of type Device
		var deviceList []Device
		err = json.Unmarshal(fileData, &deviceList)
		if err != nil {
			fmt.Println("Failed to Unmarshall JSON data:", err)
		}

		for i, deviceData := range deviceList {
			fmt.Println(" ", i, deviceData.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Customizing the "usage" display
	listCmd.SetUsageTemplate(`Usage:
	woled list

Arguments:
	None

Examples:
	woled list
	`)
}
