# glimpse

A tool to quickly catch a "glimpse" of your images, textures, etc...

# Build
Just run `make` in the directory you clone the repo.

This will:
- run go mod init 
- run go mod tidy to install dependencies
- run go build to generate the binary

**NOTE:** Go must be installed on your system for this build process to work.

# Tech Stack
I built this tool using Go w/ Raylib Go Bindings

## Checklist

- [x] Load single file asset and display it
- [ ] Load multiple file assets and swap between them
- [ ] display file assets in tabs to swap on
