## January 1, 2020

This CRUD was originally built to test kickwasm version 5.0.0 and it's tools.

This current version was built to test kickwasm version 15 and it's tools. The CRUD WIKI has been updated to this version of the CRUD.

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

## If you want to build CRUD on linux, windows, darwin

### Download kickwasm and it's tools and dependencies

``` shell
~$ go get -u github.com/josephbudd/kickwasm
~$ cd ~/go/src/github.com/josephbudd/kickwasm
~/go/src/github.com/josephbudd/kickwasm$ make install
~/go/src/github.com/josephbudd/kickwasm$ make test
~/go/src/github.com/josephbudd/kickwasm$ make dependencies
~/go/src/github.com/josephbudd/kickwasm$ make proofs
```

### Build and run the CRUD

``` shell
$ cd ~/go/src/github.com/josephbudd/crud
$ kickbuild -rp -mp -run
```

[![building and running this crud](https://i.vimeocdn.com/video/803693464.webp?mw=550&amp;mh=310&amp;q=70)](https://vimeo.com/351949802)
