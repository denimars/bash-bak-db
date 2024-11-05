#!/bin/bash

# Set database variables
DB_NAME="your_database_name"      # Replace with your database name
DB_USER="your_username"           # Replace with your MySQL username
DB_PASSWORD="your_password"       # Replace with your MySQL password

# Get the current date and time for the filename in the format sipaham<year><month><date><hour><minute>
TIMESTAMP=$(date +"%Y%m%d%H%M")
BACKUP_FILENAME="sipaham${TIMESTAMP}.sql"

# Perform the MySQL dump
mysqldump -u $DB_USER -p$DB_PASSWORD $DB_NAME > $BACKUP_FILENAME

# Verify if the dump was successful
if [ $? -eq 0 ]; then
    echo "Backup successful! File created: $BACKUP_FILENAME"
else
    echo "Backup failed!"
fi