#!/bin/bash

#if neccessary, clone storj, build and install it, if failed, reference https://github.com/storj/storj
type storj-sim >/dev/null 2>&1
if [ $? -ne 0 ]; then
    echo "build storj and install storj"
    if [ ! -d ~/storj ]; then
        git clone https://github.com/storj/storj.git ~/storj

        if [ $? -ne 0 ]; then
            rm -rf ~/storj
            echo "can not clone storj repo, please install storj manually "
            exit
        else
            echo "build and install storj to you env."
            cd ~/storj
            go install -v ./cmd/...
            if [ $? -ne 0 ]; then
                echo "can not build and install storj, please install storj manually "
                exit
            fi
        fi
    fi
fi

#setup storj local test network configuration
storj_config_dir=~/.local/share/storj/
if [ ! -d ${storj_config_dir}/local-network ]; then
    echo "setup storj network"
    storj-sim network setup
fi