package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println(`
	╦ ╦╔═╗╦  ╔═╗╔═╗╔╦╗╔═╗  ╔╦╗╔═╗  ╔═╗╦╦  ╔═╗  ╔╦╗╔═╗╔╗╔╔═╗╔═╗╔═╗╦═╗
	║║║║╣ ║  ║  ║ ║║║║║╣    ║ ║ ║  ╠╣ ║║  ║╣   ║║║╠═╣║║║╠═╣║ ╦║╣ ╠╦╝
	╚╩╝╚═╝╩═╝╚═╝╚═╝╩ ╩╚═╝   ╩ ╚═╝  ╚  ╩╩═╝╚═╝  ╩ ╩╩ ╩╝╚╝╩ ╩╚═╝╚═╝╩╚═
	`)
	fmt.Println(`-----------------------------------------------------`)
	fmt.Println("1. View Current Directory Content")
	fmt.Println("2. Create a New File")
	fmt.Println("3. Create a New Folder")
	fmt.Println("4. View File Content")
	fmt.Println("5. Navigate to a Different Folder")
	fmt.Println("6. Rename a File / Folder")
	fmt.Println("7. Delete a File / Folder")
	fmt.Println("8. Exit")
	fmt.Println(`-----------------------------------------------------`)
	for {
		// Asking user to input choice
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// View current directory contents
			listCurrentDirectoryContents()
		case 2:
			// Create a new file
			createFile()
		case 3:
			// Create a new folder
			createFolder()
		case 4:
			// View file contents
			viewFileContents()
		case 5:
			// Navigate to a different folder
			navigateToFolder()
		case 6:
			// Rename a file or folder
			renameFileOrFolder()
		case 7:
			// Delete a file or folder
			deleteFileOrFolder()
		case 8:
			// Exit the program
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please enter a valid option.")
		}
	}
}

// Function to list files and directories in the current directory
func listCurrentDirectoryContents() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Contents of current directory:")
	for _, file := range files {
		fmt.Println(file.Name())
	}
	fmt.Println(`-----------------------------------------------------`)
}

// Function to create a new file
func createFile() {
	var fileName string
	fmt.Print("Enter file name to create: ")
	fmt.Scanln(&fileName)

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	fmt.Println("File created successfully.")
	fmt.Println(`-----------------------------------------------------`)
}

// Function to create a new folder
func createFolder() {
	var folderName string
	fmt.Print("Enter folder name to create: ")
	fmt.Scanln(&folderName)

	err := os.Mkdir(folderName, 0755) // 0755 is the permission mode
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Folder created successfully.")
	fmt.Println(`-----------------------------------------------------`)
}

// Function to view contents of a file
func viewFileContents() {
	var fileName string
	fmt.Print("Enter file name to view contents: ")
	fmt.Scanln(&fileName)

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File contents:")
	fmt.Println(string(data))
	fmt.Println(`-----------------------------------------------------`)
}

// Function to navigate to a different folder
func navigateToFolder() {
	var folderName string
	fmt.Print("Enter folder name to navigate (or '..' to go up): ")
	fmt.Scanln(&folderName)

	err := os.Chdir(folderName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Changed directory to:", folderName)
	fmt.Println(`-----------------------------------------------------`)
}

// Function to rename a file or folder
func renameFileOrFolder() {
	var oldName, newName string
	fmt.Print("Enter current file/folder name: ")
	fmt.Scanln(&oldName)
	fmt.Print("Enter new file/folder name: ")
	fmt.Scanln(&newName)

	err := os.Rename(oldName, newName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File/folder renamed successfully.")
	fmt.Println(`-----------------------------------------------------`)
}

// Function to delete a file or folder
func deleteFileOrFolder() {
	var name string
	fmt.Print("Enter file/folder name to delete: ")
	fmt.Scanln(&name)

	err := os.RemoveAll(name)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File/folder deleted successfully.")
	fmt.Println(`-----------------------------------------------------`)
}
