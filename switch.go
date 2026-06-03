package main

import "fmt"

func addContact(m map[string]string) {
	m["mahi"] = "123456789"
	m["Raksh"] = "987345490"
	fmt.Println("Contacts Added")
}

func search(m map[string]string) {
	p, found := m["mahi"]

	if found {
		fmt.Println("Phone Number:", p)
	} else {
		fmt.Println("Contact not found")
	}
}

func view(m map[string]string) {
	fmt.Println("Contacts List:")
	for n, p := range m {
		fmt.Println(n, "->", p)
	}
}

func deleteContact(m map[string]string) {
	delete(m, "mahi")
	fmt.Println("Contact Deleted")
}

func main() {
	m := map[string]string{}

	var choice int

	for {
		fmt.Println("\n1. Add Contact")
		fmt.Println("2. Search Contact")
		fmt.Println("3. View Contacts")
		fmt.Println("4. Delete Contact")
		fmt.Println("5. Exit")

		fmt.Print("Enter Choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addContact(m)

		case 2:
			search(m)

		case 3:
			view(m)

		case 4:
			deleteContact(m)

		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid Choice")
		}
	}
}
