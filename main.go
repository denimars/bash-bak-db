package main

import (
	"bytes"
	"dbbak/util"
	"fmt"
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
			finalLocation := fmt.Sprintf("%v/%v", util.Location(), dbbakFolder)
			err = os.MkdirAll(finalLocation, os.ModePerm)
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

func splitDate(fileName string) (time.Time, string) {
	isCollect := false
	dateOnString := ""
	for i := len(fileName) - 1; i >= 0; i-- {
		if isCollect {
			if util.IsNumeric(string(fileName[i])) {
				dateOnString = string(fileName[i]) + dateOnString
			} else {
				isCollect = false
			}

		}
		if fileName[i] == '.' && !isCollect {
			isCollect = true
		}

	}
	fmt.Println(dateOnString)
	return time.Now(), dateOnString
}

func deleteBak(config util.Config) {
	var err error
	var files []os.DirEntry
	path := fmt.Sprintf("%v/%v", util.Location(), config.DbBakFolder)
	if files, err = os.ReadDir(path); err == nil {
		for _, file := range files {
			if !file.IsDir() {
				splitDate(file.Name())
				// fmt.Println(dateOnString)
				fmt.Println("----")
			}
		}
	} else {
		fmt.Println(err)
	}

}

func main() {
	config, _ := util.LoadEnv()
	// config, status := util.LoadEnv()
	// db(config, status)
	deleteBak(config)
}
