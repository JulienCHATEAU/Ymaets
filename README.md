## Ymaets

Ymaets is a 2D game where you drive a small ship and defend yourself against wild enemies !

## Requirements

- golang : `sudo apt install golang-go`

## Installation

I) Start by update your Linux packages

- `sudo apt update`
- `sudo apt updgrade`

II) If you already have a Go workspace go to II), otherwise let's create one.   
Move to a directory where you want to create your Go workspace and execute the following commands :

- `mkdir -p go/src go/bin`
- `export GOPATH=$(pwd)/go`

III) Then to download this project and its dependancies execute :

- `git clone https://github.com/JulienCHATEAU/Ymaets.git $GOPATH/src/Ymaets`
- `apt-get install libgl1-mesa-dev libxi-dev libxcursor-dev libxrandr-dev libxinerama-dev`
- `apt-get install libwayland-dev libxkbcommon-dev `
- `go get -v -u github.com/gen2brain/raylib-go/raylib`

IV) Finally, to compile the project use :

- `go install Ymaets`

## Run 

You can simply run the project executing the following command :

- `go run Ymaets`