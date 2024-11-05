#!/bin/bash

# Define variables for ease of use and readability
REMOTE_USER="username"
REMOTE_SERVER="remote_server_ip"
REMOTE_PATH="/path/to/remote/directory"
LOCAL_PATH="/path/to/local/directory"

# Execute rsync command with SSH
rsync -avz -e ssh "${REMOTE_USER}@${REMOTE_SERVER}:${REMOTE_PATH}" "${LOCAL_PATH}"
