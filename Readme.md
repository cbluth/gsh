# gsh : The Missing Host Manager

`gsh` intends to offer a new way to manage sets of hosts and a collection of accompanying scripts and files.
You can organize your inventory of hosts with groups and labels. You can add scripts and files (like configs) to your inventory. Then, `gsh` works over ssh, and you can execute scripts and/or copy files en masse to your groups of servers.
See TODO section below for future plans.

### Requirements

This should work on any linux host, like your workstation. Mac os builds are being tested.
Access to some ssh servers is needed to use all the features of `gsh`. Currently only bash is supported, with plans for more languages.

### Installation
You can download the latest release here: https://github.com/cbluth/gsh/releases <br />
Or use these options:

#### A) Use Go Install
To build with go install:
```bash
$ go install github.com/cbluth/gsh
```

#### B) Clone via git, and build in docker

To build in docker, golang not need to be installed:
```bash
$ git clone https://github.com/cbluth/gsh && cd gsh
$ ./build.sh    # this will make a `release.tgz` file
$ tar -xzvf release.tgz # makes linux and mac builds
$ sudo cp ./build/gsh-linux /usr/local/bin/gsh
$ sudo chmod a+rx /usr/local/bin/gsh
```

#### C) Clone via git, and build locally


To build with go build:
```bash
$ git clone https://github.com/cbluth/gsh && cd gsh
$ go build . -o gsh
$ sudo cp gsh /usr/local/bin/gsh
$ sudo chmod a+rx /usr/local/bin/gsh
```

## How To Use It??

You can use `gsh` to execute scripts over ssh, and manage an inventory of servers.<br />
For usage, see [the docs](docs/)

## Hack or Contribute
This project intentionally avoids importing 3rd party libraries, and tries to stick to the basic libraries provided by golang.
All contributions and bug reports are welcome, please just open an issue on the github issues page.
This is a new project, with lots of issues, if you see any bugs, please report them! :)

To hack on this project, git clone to anywhere, and edit the code directly.
After applying edits to the codebase, you can use `./gsh.sh` to compile and run `gsh` via go, with the same arguments as running the program normally; eg: `./gsh.sh make host dev0 address=dev0.local`

## TODO
- Implement `copy` feature
- Cleanup orphaned files/scripts in ~/.gsh
- Implement `sudo=true` feature for scripts
- Support other languages like python, nodejs, and ruby

## License

This project is licensed under the MIT License, see the [LICENSE](LICENSE) file for details.
