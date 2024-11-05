package main

import (
	"bytes"
	"dbbak/util"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	util.LoadEnv()
	// Database credentials
	dbName := os.Getenv("DB_NAME")         // Replace with your database name
	dbUser := os.Getenv("DB_USER")         // Replace with your MySQL username
	dbPassword := os.Getenv("DB_PASSWORD") // Replace with your MySQL password

	// Create a timestamp in the format sipaham<year><month><date><hour><minute>
	timestamp := time.Now().Format("200601021504") // YYYYMMDDHHMM
	backupFilename := fmt.Sprintf("sipaham%s.sql", timestamp)

	// Set the MYSQL_PWD environment variable to securely pass the password
	os.Setenv("MYSQL_PWD", dbPassword)

	// Construct the mysqldump command without the password flag
	cmd := exec.Command("mysqldump", "-u"+dbUser, dbName)

	// Create the backup file
	backupFile, err := os.Create(backupFilename)
	if err != nil {
		fmt.Println("Error creating backup file:", err)
		return
	}
	defer backupFile.Close()

	// Set the command's output to the backup file
	cmd.Stdout = backupFile

	// Capture only stderr
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	// Run the command
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error creating backup:", err)
		fmt.Printf("Details: %s\n", stderr.String()) // Print stderr output if any
		return
	}

	fmt.Printf("Backup successful! File created: %s\n", backupFilename)
}
