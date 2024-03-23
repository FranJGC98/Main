package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Contact structure

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}
	return nil
}

// Load contacts from a json

func loadContacsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return nil
	}
	return nil
}

func main() {

	//Contact slice

	var contacts []Contact

	//Load existing contacts from a file

	err := loadContacsFromFile(&contacts)
	if err != nil {
		fmt.Println("Error to download the contacts:", err)
	}

	//Create bufio's instance

	reader := bufio.NewReader(os.Stdin)
	for {

		//Show ths options to the user

		fmt.Print("==== CONTACT MANAGER ====\n",
			"1. Add new contact\n",
			"2. Show all the contacts\n",
			"3. Exit\n",
			"Choose one option:\n",
		)

		// Read user's options

		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error to read the option", err)
			return
		}
		// Manage the user's choice

		switch option {
		case 1:
			var c Contact

			// Add and create contact

			fmt.Print("Name: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Phone: ")
			c.Phone, _ = reader.ReadString('\n')

			//Add contact to Slice

			contacts = append(contacts, c)
			//Save in changes in json
			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error to save the option", err)

			}
		case 2:

			//Show all the contacts

			fmt.Println("=====================================================================")
			for index, contact := range contacts {
				fmt.Printf("%d. Name: %s Email: %s Number: %s\n",
					index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("=====================================================================")

		case 3:
			//exir to the program
			return
		default:
			fmt.Println("Invalid option")

		}
	}
}
