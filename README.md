## Sept 24, 2019

This CRUD was originally built to test kickwasm version 5.0.0 and it's tools.

This current version was built to test kickwasm version 10.0.0 and it's tools. The CRUD WIKI has been updated to this version of the CRUD.

Read the CRUD WIKI to learn how kickwasm and it's tools are used.

Below is an older video of the crud app with the cookie crumbs in the GUI. They only added complexity to a simple application so they are removed in the source code. I'll update the video when I get a chance.

To build crud you will need it's source code and it's dependencies. However do not use the **-u** flag with go get. Instead manually download the dependencies as shown below.

## If you want to run CRUD on ubuntu linux

### Get and run CRUD

``` shell

$ go get github.com/josephbudd/crud
$ cd ~/go/src/github.com/josephbudd/crud
$ ./crud

```

## If you want to build CRUD on linux

### Get CRUD's other dependencies

* [the boltdb package.](https://github.com/boltdb/bolt)
* [the yaml package.](https://gopkg.in/yaml.v2)
* [the gorilla websocket package.](https://github.com/gorilla/websocket)

``` shell

$ go get github.com/boltdb/bolt/...
$ go get gopkg.in/yaml.v2
$ go get github.com/gorilla/websocket

```

### Get the kickpack tool

The new renderer build scripts which are **renderer/build.sh** and **renderer/buildPack.sh** use kickpack. So you will need to download, build and install kickpack.

``` shell

$ go get -u github.com/josephbudd/kickpack
$ cd ~/go/src/github.com/josephbudd/kickpack
$ go install

```

### Build the renderer process with renderer/build.sh

``` shell

nil@NIL:~/go/src/github.com/josephbudd/crud$ cd renderer/
nil@NIL:~/go/src/github.com/josephbudd/crud/renderer$ ./build.sh

STEP 1:
REMOVE YOUR PREVIOUS BUILD OF /home/nil/go/src/github.com/josephbudd/crud/renderer/spawnpack
rm -r /home/nil/go/src/github.com/josephbudd/crud/renderer/spawnpack

STEP 2:
WRITE THE SOURCE CODE FOR YOUR NEW spawnpack PACKAGE.
 * The spawnpack package is your renderer's spawn html templates.
cd /home/nil/go/src/github.com/josephbudd/crud/site
kickpack -nu -o=/home/nil/go/src/github.com/josephbudd/crud/renderer/spawnpack ./spawnTemplates
cd /home/nil/go/src/github.com/josephbudd/crud/renderer
 * Success. The source code for your new spawnpack package is written.

STEP 3:
BUILD THE RENDERER GO CODE INTO WEB ASSEMBLY CODE AT ../site/app.wasm
GOARCH=wasm GOOS=js go build -o ../site/app.wasm main.go panels.go
 * Success. The renderer go code is compiled into web assembly code at ../site/app.wasm

STEP 4:
REMOVE YOUR PREVIOUS BUILD OF crudsitepack
rm -r /home/nil/go/src/github.com/josephbudd/crudsitepack

STEP 5:
WRITE THE crudsitepack PACKAGE SOURCE CODE.
 * The crudsitepack package will pretend to be a file store
     but it will actually just read the files
     in the site folder and return their contents.
 * See func serveFileStore in Serve.go.
cd /home/nil/go/src/github.com/josephbudd/crud
kickpack -nu -nopack -o=/home/nil/go/src/github.com/josephbudd/crudsitepack ./site ./Http.yaml
 * Success. The crudsitepack package source code has been written.

STEP 6:
BUILD THE crudsitepack PACKAGE.
cd /home/nil/go/src/github.com/josephbudd/crudsitepack
go build
 * Success.
 * You have successfully compiled the crudsitepack package object code.

STEP 7:
BUILD THE MAIN PROCESS GO CODE INTO THE EXECUTABLE crud.
   You will do so with the following 2 commands...
   cd ..
   go build

```

### Build the main process

``` shell

$ cd ..
$ go build
$ ./crud

```

[![building and running this crud](https://i.vimeocdn.com/video/803693464.webp?mw=550&amp;mh=310&amp;q=70)](https://vimeo.com/351949802)
