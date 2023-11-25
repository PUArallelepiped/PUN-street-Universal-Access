# !/bin/bash

# This script is used to generate the go server code from the swagger.yaml file
# It is used to update the code when the swagger.yaml file is updated
# put this script in /backend and run it with ./codegen.sh

docker run --rm -v ${PWD}:/local parsertongue/swagger-codegen-cli:latest generate \
      -i /local/api/swagger.yaml -l go-server -o /local/codegen-temp

sudo mv codegen-temp/api/swagger.yaml api/swagger.yaml
sudo mv codegen-temp/go/model* swagger/
sudo rm -rf codegen-temp
USER=$(whoami)
sudo chown $USER:$USER api/swagger.yaml swagger/model*
gofmt -w .
sed -i '/style: simple/d' api/swagger.yaml 
sed -i '/explode: false/d' api/swagger.yaml