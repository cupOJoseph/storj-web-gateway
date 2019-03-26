#!/bin/bash

#extract keys from test network configuration, so ipfs-s3c-storj-plugin can get access right
storj_config_dir=~/.local/share/storj/
access_key=$(grep access-key ${storj_config_dir}/local-network/gateway/0/config.yaml)
access_key=${access_key:14:28}
secret_key=$(grep secret-key ${storj_config_dir}/local-network/gateway/0/config.yaml)
secret_key=${secret_key:14:28}
export STORJ_ACCESS_KEY=${access_key}
export STORJ_SECRET_KEY=${secret_key}

#clone and build ipfs-s3c-storj-plugin
plugin_dir=${GOPATH}/src/github.com/ipfs-s3c-storj-plugin
if [ ! -d ${plugin_dir} ]; then
    echo "clone ipfs-s3c-storj-plugin"
    git clone https://github.com/zyfrank/ipfs-s3c-storj-plugin.git ${plugin_dir}
fi

#initialize ipfs-s3c-storj-plugin configuration profile
cd ${plugin_dir}
rm -rf ~/.ipfs
make install
cd build
./ipfs init --profile s3c-storjds
