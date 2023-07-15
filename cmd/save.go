package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save device configuration",
	Long:  `Save your device to a local config file by specifying a name and the MAC address of the device.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Create new type to save JSON data to config file
		type Device struct {
			Name       string `json:"name"`
			MACAddress string `json:"macAddress"`
		}

		// Check if two arguments have been given in the command, else cancel
		if len(args) != 2 {
			fmt.Println("Provide a name and a MAC address and try again")
			return
		}
		// Check if MAC address is valid
		mac := args[1]
		match, _ := regexp.MatchString("^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$", mac)
		if !match {
			fmt.Println("Invalid MAC address format, use format XX:XX:XX:XX:XX:XX")
			return
		}

		// Assign given arguments to Device type
		newDevice := Device{
			Name:       args[0],
			MACAddress: mac,
		}

		// Read existing JSON file
		filePath := "config.json"
		configData, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println("Failed to read JSON file:", err)
			return
		}

		// Unmarshal existing JSON data into slice of type Device
		var deviceList []Device
		err = json.Unmarshal(configData, &deviceList)
		if err != nil {
			fmt.Println("Failed to Unmarshall JSON data:", err)
		}

		deviceList = append(deviceList, newDevice)

		// Convert device data to JSON
		newDeviceList, err := json.MarshalIndent(deviceList, "", "    ")
		if err != nil {
			fmt.Println("Failed to convert data to JSON:", err)
			return
		}

		// Write JSON data to the configuration file
		err = ioutil.WriteFile("config.json", newDeviceList, 0644)
		if err != nil {
			fmt.Println("Failed to write configuration file:", err)
			return
		}

		fmt.Println("Configuration data saved successfully!")

		fmt.Println("Name:", newDevice.Name)
		fmt.Println("MAC:", newDevice.MACAddress)
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)

	// Add new flags to the `save` command
	saveCmd.Flags().StringP("name", "n", "", "Name of the device")       // -n, --name flags
	saveCmd.Flags().StringP("mac", "m", "", "MAC address of the device") // -m, --mac flags

	// Customizing the "usage" display
	saveCmd.SetUsageTemplate(`Usage:
	woled save [flags]

	Flags:
	-n, --name string   Name of the device
	-m, --mac string    MAC address of the device

	Examples:
	gowol save -n "My Device" -m "00:11:22:33:44:55"
	`)
}