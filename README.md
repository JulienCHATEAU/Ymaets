## Ymaets

Ymaets is a 2D game where you drive a small ship and defend yourself against wild enemies !

## Requirements

- golang : `sudo apt install golang-go`

## Installation

I) Start by updating your Linux packages

- `sudo apt update`
- `sudo apt updgrade`

II) If you already have a Go workspace go to III), otherwise let's create one.   
Move to a directory where you want to create your Go workspace and execute the following commands :

- `mkdir -p go/src go/bin`
- `export GOPATH=$(pwd)/go`

III) Then to download this project and its dependancies execute these following commands :

- `git clone https://github.com/JulienCHATEAU/Ymaets.git $GOPATH/src/Ymaets`
- `apt-get install libgl1-mesa-dev libxi-dev libxcursor-dev libxrandr-dev libxinerama-dev`
- `apt-get install libwayland-dev libxkbcommon-dev `
- `go get -v -u github.com/gen2brain/raylib-go/raylib`
- `go get github.com/nickdavies/go-astar/astar`

IV) Finally, to compile the project use :

- `go install Ymaets`

## Advice

Note that in II), your GOPATH variable will be unset when you quit your terminal.  
Add `export GOPATH=<your_go_workspace_path>` to your .bashrc file in order to make it persistent.

## Run 

You can simply run the project executing the following command :

- `go run Ymaets`