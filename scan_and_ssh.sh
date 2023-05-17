#!/bin/bash

# Perform IP scanning with masscan
masscan_output=$(masscan -p 22 54.149.62.0/24 --rate=1000 2>/dev/null)

# Extract the IP address from the masscan output
ip_address=$(echo "$masscan_output" | grep -oP 'Discovered open port 22/tcp on \K[\d.]+')

# Check if an IP address was found
if [[ -z "$ip_address" ]]; then
    echo "No open SSH ports found in the specified IP range."
    exit 1
fi

echo "Attempting SSH connection to $ip_address"
ssh ubuntu@54.149.62.87
# ssh ubuntu@ec2-54-149-62-87.us-west-2.compute.amazonaws.com
