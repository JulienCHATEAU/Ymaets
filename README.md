## Ymaets

Ymaets is a 2D game where you drive a small ship and defend yourself against wild enemies !

## Installation

I) If you already have a Go workspace go to II), otherwise let's create one.   
Move to a directory where you want to create your Go workspace and execute the following commands :

- `mkdir -p go/src go/bin`
- `export GOPATH=$(pwd)/go`

II) Then to download this project and its dependancies execute :

- `git clone https://github.com/JulienCHATEAU/Ymaets.git $GOPATH/src/Ymaets`
- `go get -v -u $GOPATH/src/github.com/gen2brain/raylib-go/raylib`

III) Finally, to compile the project use :

- `go install Ymaets`

## Run 

You can simply run the project executing the following command :

- `go run Ymaets`