# Installation

Clone this repo with `git install https://github.com/jschiarizzi/storj-ipfs-gateway.git` to the folder in your GOPATH `$GOPATH/src/github.com`. By default this should be `~/go/bin`

Move into the folder:
`cd $GOPATH/src/github.com/ipfs-s3c-storj-plugin`

Install locally
`gx install --local`
Here we use gx to publish 3 ipfs packages:go-ipfs-config,go-ipfs and iptb.  Maybe in your site it is hard to get these three packages
just check https://github.com/zyfrank/go-ipfs, https://github.com/zyfrank/go-ipfs-config and https://github.com/zyfrank/iptb, clone to your local env., switch to branch storj-s3c-plugin, then use `gx publish -f` to your local IPFS node, so `gx install --local` can find these three packages.

Install dependencies and initialize S3 gateway.
```
make install
./build/ipfs init --profile s3c-storjds
```

Change $IPFS_PATH/config (commonly it is ~/.ipfs/config),  input your "accessKey" and "secretKey" which are used to access storj s3 gateway

Start your test storj env. by using storj-sim network run

Now start the ipfs daemon .`/build/ipfs daemon`.

When you `./build/ipfs` add `*`, the file will be fed to storj.

### Usage of installer scripts

Scripts for installing can be found under ./shell. Install Storj first, and then the IPFS gateway

Assume golang (>=1.11) has been setup.

Before do following experiment, backup your local ipfs env.

First, you should run ```./shell/setupStorj.sh```

After that, you should start up storj test network on a new termial by using command  ```storj-sim network run```

Now ```./shell/setupIpfsStorjPlugin.sh``` can be run to setup ipfs-s3c-storj-plugin env.

If all successfully, You should switch to ~/src/github.com/ipfs-s3c-storj-plugin/build and use new built ipfs command. Now you can try ```./ipfs add```, ```./ipfs cat```, ```./ipfs ls```, ```./ipfs get``` etc.  


I have recorded two tty sessions under ./record directory
To replay these two records, you should install ttyrec, in Ubuntu, run following command

```sudo apt-get install ttyrec```

Now
```ttyplay ./record/storjInstallRec``` can replay the storj setup process, during replay, ctrl+'f' can be used to increase playback speed.

```ttyplay ./record/ipfs-s3c-storj-pluginInstallRec``` can replay the ipfs-s3c-storj plugin setup and some ipfs command test process.
