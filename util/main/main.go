package main

import (
	"flag"
	"fmt"
	"os"
	util "simplebank/util"
	"strings"
)

func main() {
	// Define command-line flags
	updateKey := flag.String("key", "", "The configuration key to update")
	configFile := flag.String("config", "../temp_config.enc", "The path to the encrypted config file")
	key := flag.String("encryptionkey", "ruEanqdOKvgzoN9n", "The encryption key")
	flag.Parse()

	// Load the encrypted config file
	encryptedData, err := os.ReadFile(*configFile)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		os.Exit(1)
	}

	// Decrypt the encrypted config data
	decryptedData, err := util.DecryptData(encryptedData, []byte(*key))
	if err != nil {
		fmt.Println("Error decrypting config:", err)
		os.Exit(1)
	}

	// Convert decrypted data to a map
	configMap := make(map[string]string)
	lines := strings.Split(string(decryptedData), "\n")
	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			configMap[parts[0]] = parts[1]
		}
	}

	// Prompt the user for the new value
	if *updateKey != "" {
		newInput := util.GetUserInput(fmt.Sprintf("Enter a new value for key '%s': ", *updateKey))
		configMap[*updateKey] = newInput
	} else {
		fmt.Println("No key specified for update.")
		os.Exit(1)
	}

	// Convert the map back to a string
	var updatedData []string
	for key, value := range configMap {
		updatedData = append(updatedData, key+"="+value)
	}
	updatedConfig := []byte(strings.Join(updatedData, "\n"))

	// Encrypt the updated configuration
	encryptedConfig, err := util.EncryptData(updatedConfig, []byte(*key))
	if err != nil {
		fmt.Println("Error encrypting config:", err)
		os.Exit(1)
	}

	// Write the encrypted configuration back to the file
	err = os.WriteFile(*configFile, encryptedConfig, 0644)
	if err != nil {
		fmt.Println("Error writing updated config file:", err)
		os.Exit(1)
	}

	fmt.Println("Configuration key updated successfully.")
}


=UNIQUE(SORT(CHOOSECOLS(FILTER(Inventory,ISNUMBER(SEARCH(W2,Inventory[Item])*SEARCH(W3,Inventory[Item])*SEARCH(W4,Inventory[Item])*SEARCH(W5,Inventory[Item])*SEARCH(W6,Inventory[Item])*SEARCH(W7,Inventory[Item])*SEARCH(W8,Inventory[Item])*SEARCH(W9,Inventory[Item])*SEARCH(W10,Inventory[Item])*SEARCH(W11,Inventory[Item])*SEARCH(W12,Inventory[Item])*SEARCH(W13,Inventory[Item])),"No results"),{1,2},{2}))
