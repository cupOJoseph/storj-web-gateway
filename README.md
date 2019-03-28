# storj-ipfs-gateway
An IPFS gateway with a storj implementation backend

## ipfs-s3c-storj-plugin
This experimental ipfs plugin combines code and is heavily inspired by https://github.com/RTradeLtd/storj-ipfs-ds-plugin and https://github.com/ipfs/go-ds-s3.

## Usage 

clone this repo to your $GOPATH/src/github.com 

cd $GOPATH/src/github.com/ipfs-s3c-storj-plugin 

gx install --local 
(here I personnally gx publish 3 ipfs packages:go-ipfs-config,go-ipfs and iptb.  maybe in your site it is hard to get these three packages
just check https://github.com/zyfrank/go-ipfs, https://github.com/zyfrank/go-ipfs-config and https://github.com/zyfrank/iptb, clone to your local env. ,switch to branch storj-s3c-plugin, then use gx publish -f to your local IPFS node, so gx install --local can find these three packages)

make install

./build/ipfs init --profile s3c-storjds

change $IPFS_PATH/config (commonly it is ~/.ipfs/config),  input your "accessKey" and "secretKey" which are used to access storj s3 gateway

start your test storj env. by using storj-sim network run

now start ipfs daemon  ./build/ipfs daemon

when you ./build/ipfs add *, the file will be feed to storj




# Usage of shell script under ./shell

Assume golang (>=1.11) has been setup.

Before do following experiement, backup your local ipfs env.

First, you should run ```./shell/setupStorj.sh```

After that, you should start up storj test network on a new termial by using command  ```storj-sim network run```

Now ```./shell/setupIpfsStorjPlugin.sh``` can be run to setup ipfs-s3c-storj-plugin env.

If all successfully, You should switch to ~/src/github.com/ipfs-s3c-storj-plugin/build and use new built ipfs command. Now you can try ```./ipfs add```, ```./ipfs cat```, ```./ipfs ls```, ```./ipfs get``` etc.  


I have recorded two tty sessiones under ./record directory
To replay these two records, you should install ttyrec, in unbuntu, run following command

```sudo apt-get install ttyrec```

Now
```ttyplay ./record/storjInstallRec``` can replay the storj setup process, during replay, ctrl+'f' can be used to increase playback speed. 

```ttyplay ./record/ipfs-s3c-storj-pluginInstallRec``` can replay the ipfs-s3c-storj plugin setup and some ipfs command test process.
