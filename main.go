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

func CopyDir(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("source directory does not exist: %v", err)
	}
	err = os.MkdirAll(dst, srcInfo.Mode())
	if err != nil {
		return fmt.Errorf("could not create destination directory: %v", err)
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return fmt.Errorf("could not read source directory: %v", err)
	}
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := CopyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := CopyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("could not open source file: %v", err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("could not create destination file: %v", err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return fmt.Errorf("error copying file: %v", err)
	}
	sourceInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("could not retrieve source file info: %v", err)
	}
	err = os.Chmod(dst, sourceInfo.Mode())
	if err != nil {
		return fmt.Errorf("could not set permissions on destination file: %v", err)
	}

	return nil
}

func fileBak(config util.Config, status bool) {
	if status {

		for _, d := range config.Data {
			err := CopyDir(d.From, d.To)
			fmt.Println(d.From)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func main() {
	config, status := util.LoadEnv()
	db(config, status)
	fileBak(config, status)
}
