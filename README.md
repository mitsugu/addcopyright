# addcopyright
## what is this?
  Adds copyright information to Exif of image data taken with a digital camera, etc.  
  Use this if your camera does not have the ability to write copyright information.

## Requirements
  * [ExifTool](https://exiftool.org/)
  * [ImageMagick](https://imagemagick.org/)

## Setup
### install addcopyright
#### use go install
```
go install github.com/mitsugu/addcopyright@<tag name>
```
#### download release zip file
1. download [release zip file](https://github.com/mitsugu/addcopyright/releases)
2. extract the zip file
3. Please put it in any directory.
4. Pass the path to the addcopyright directory. Or put the addcopyright executable file in a directory that is in your path.

### edit addcopyright.json
```
// addcopyright.json example
{
  "copyright": "MITSUGU OYAMA (online nick name OrzBruford)",
  "exiftool_path": "/usr/bin/exiftool",
  "imagemagick_path": "./imagemagick/magick"
}
```
  Note: Place addcopyright.json in **the current directory** where you run addcopyright.

### usage
```
addcopyright [--config <configuration json file path] --input <input file path> --output <output file path>
or
addcopyright [-c <configuration json file path] -i <input file path> -o <output file path>
```
### License
[Apache License 2.0](LICENSE.en.txt)


## comment
enjoy!! 😀
