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
	// database credentials load on env
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	// create a timestamp in the format sipaham<year><month><date><hour><minute>
	timestamp := time.Now().Format("200601021504") // YYYYMMDDHHMM
	backupFilename := fmt.Sprintf("sipaham%s.sql", timestamp)

	// set the MYSQL_PWD environment variable to securely pass the password
	os.Setenv("MYSQL_PWD", dbPassword)

	// construct the mysqldump command without the password flag
	cmd := exec.Command("mysqldump", "-u"+dbUser, dbName)

	// create the backup file
	backupFile, err := os.Create(backupFilename)
	if err != nil {
		fmt.Println("Error creating backup file:", err)
		return
	}
	defer backupFile.Close()

	// set the command's output to the backup file
	cmd.Stdout = backupFile

	// capture only stderr
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	// run the command
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error creating backup:", err)
		fmt.Printf("Details: %s\n", stderr.String())
		return
	}

	fmt.Printf("Backup successful! File created: %s\n", backupFilename)
}
