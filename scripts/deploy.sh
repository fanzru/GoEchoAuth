#!/bin/sh
set -e

cd ..
echo "Deploying application ..."

# Update codebase
git fetch origin main
echo "failed 1"
git reset --hard origin/main
echo "failed 2"

echo "Formating Code 🛠"
go fmt

echo "Installing dependency 🛠"
go mod tidy
            
echo "Restart pm2 service 🔥"
pm2 restart deployment.json
