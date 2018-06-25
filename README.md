# Directory Loader

## Just a simple directory loader for ProximaX
This is basically a CLI tool that creates a NEM Hash of the root directory that can then be loaded from any ProximaX Node.

### Steps to run
+ Unpack the dist here ![windows](https://testnet1.gateway.proximax.io/xpxfs/a233dd90f6c0e0c9ffc39eb433040b7827463b01ce87de3202904802a8a19beb/go-xpx-loader.zip)
+ Make sure that all artifacts are on the same directory
+ Open terminal and run `go-xpx-loader.exe daemon`
+ Open another terminal and run `go-xpx-loader.exe add -r <file or dir> ---uploadparam=<jsonconfigpath>`

#### NOTE: jsonconfigpath sample - https://github.com/proximax-storage/xpx-dirloader/blob/master/params/uploadPath.json

You'll get a NEM Hash like the one below

```bash

Last File --: HexGL-master/textures.full/tracks/cityscape/start/start.jpg
Last Bytes --: %!s(int64=32273)
Last Name --:
Loading Directory....
"bb2c743bfe6738408208c87ea275b8804be2d367d69a41bdceb403d5102c8916"
NEM Blockchain Hash --: "bb2c743bfe6738408208c87ea275b8804be2d367d69a41bdceb403d5102c8916"
```

You can then use this hash to load the entire directory on ProximaX

```xml
https://testnet.gateway.proximax.io/xpxfs/bb2c743bfe6738408208c87ea275b8804be2d367d69a41bdceb403d5102c8916/
```
