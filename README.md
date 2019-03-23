# storj-ipfs-gateway
An IPFS gateway with a storj implementation backend


# Usage of shell script under ./shell

Assume golang (>=1.11) has been setup.

Before do following experiement, backup your local ipfs env.

First, you should run ./shell/setupStorj.sh

After that, you should start up storj test network on a new termial by using command  storj-sim network run

Now ./shell/setupIpfsStorjPlugin.sh can be run to setup ipfs-s3c-storj-plugin env.

If all successfully, You should switch to ~/src/github.com/ipfs-s3c-storj-plugin/build and use new built ipfs command. Now you can try ipfs add, ipfs cat, ipfs ls, ipfs get etc.  