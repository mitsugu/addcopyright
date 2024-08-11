# addcopyright
## what is this?
  Adds copyright information to Exif of image data taken with a digital camera, etc.  
  Use this if your camera does not have the ability to write copyright information.

## Requirements
  * [ExifTool](https://exiftool.org/)

## Setup
### install addcopyright
```
go install github.com/mitsugu/addcopyright@<tag name>
```
### edit addcopyright.json
```
// addcopyright.json example
{
  "copyright": "MITSUGU OYAMA (online nick name OrzBruford)",
  "exiftool_path": "/usr/bin/exiftool"
}
```
  Note: Place addcopyright.json in **the current directory** where you run addcopyright.

### usage
```
addcopyright <input file path> <output file path>
```
### License
Apache License 2.0


## comment
enjoy!! 😀
