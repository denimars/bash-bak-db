package main

import (
	"bytes"
	"dbbak/util"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func db(config util.Config, status bool) {
	if status {
		for _, db := range config.Database {
			timestamp := time.Now().Format("200601021504")
			backupFilename := fmt.Sprintf("%s%s.sql", db.DbName, timestamp)
			os.Setenv("MYSQL_PWD", db.DbPassword)
			cmd := exec.Command("mysqldump", "-u"+db.DbUser, db.DbName)
			backupFile, err := os.Create(backupFilename)
			if err != nil {
				fmt.Println("Error creating backup file:", err)
				return
			}
			defer backupFile.Close()
			cmd.Stdout = backupFile
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
			dbbakFolder := config.DbBakFolder
			err = os.MkdirAll(dbbakFolder, os.ModePerm)
			if err != nil {
				fmt.Println("Error creating dbbak folder:", err)
				return
			}
			destPath := filepath.Join(dbbakFolder, backupFilename)
			err = os.Rename(backupFilename, destPath)
			if err != nil {
				fmt.Println("Error moving file to dbbak folder:", err)
				return
			}
			fmt.Printf("Backup moved to %s successfully!\n", destPath)
		}
	}

}

func openFile(location string) (*os.File, bool) {
	sourceFile, err := os.Open(location)
	if err != nil {

		return sourceFile, false
	}
	return sourceFile, true
}

func fileBak(config util.Config, status bool) {
	if status {
		to, statusTo := openFile(config.DbBakFolder)
		for _, d := range config.Data {
			from, statusFrom := openFile(d.Location)
			if statusFrom && statusTo {
				_, err := io.Copy(to, from)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Folder not found!")
			}
			defer from.Close()
		}
		defer to.Close()
	}
}

func main() {
	config, status := util.LoadEnv()
	db(config, status)
	fileBak(config, status)
}
