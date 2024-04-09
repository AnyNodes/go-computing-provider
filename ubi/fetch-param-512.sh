#!/bin/bash

#Download the ubi-task parameter
if [ -z "$PARENT_PATH" ]; then
    echo "PARENT_PATH is empty. Skipping ubi-task parameter download."
else
    if [[ "$(lscpu | grep "Vendor ID" | awk '{print $3}')" == "AuthenticAMD" ]]; then
        docker run --rm -v "$PARENT_PATH":/var/tmp/filecoin-proof-parameters sxk1633/lotus-shed-amd:latest lotus-shed fetch-params --proving-params 512MiB
    else
        docker run --rm -v "$PARENT_PATH":/var/tmp/filecoin-proof-parameters sxk1633/lotus-shed-intel:latest lotus-shed fetch-params --proving-params 512MiB
    fi
fi

