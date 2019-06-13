# storj-ipfs-gateway
An IPFS gateway with a [Storj](http://storj.io) implementation as a backend. This allows for easily viewing files stored in the Storj network directly in a web browser.

## Attribution
This experimental ipfs plugin combines code from and is heavily inspired by [RTrade's storj ipfs plugin](https://github.com/RTradeLtd/storj-ipfs-ds-plugin). It also uses code from https://github.com/ipfs/go-ds-s3 and https://github.com/storj/storj.

Special thanks to [RTrade](https://www.rtradetechnologies.com/) for supporting this project and inspiring the development direction it has taken so far.

## Running this Gateway
Clone this project and run `main.go`. Running and entire S3 Gateway and more instructions/requirements can be found in [Installation.md](https://github.com/jschiarizzi/storj-ipfs-gateway/blob/master/Installation.md).

## Usage of a Storj Gateway
Storj gives us the decentralized cloud storage, compatible with S3 buckets, through their service [Tardigrade](https://tardigrade.io/). This gateway acts as a web-based explorer for the files that are stored in Storj services like Tardigrade. The purpose of keeping this gateway running is to demonstrate how fast and easy storage is on IPFS systems like Storj.  We also provide instructions on using this project to setup your own gateway.

### Example Uses
After setting up your own gateway you could use it to create an online ebook store, with ebook files stored with Tardigrade. Or you could use a gateway like this to run a social art platform with persistent storage. 
