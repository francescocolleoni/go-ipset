# go-ipset
`go-ipset` is a Go wrapper of `iptables`' `ipset` utility.

## Introduction
`go-ipset` exposes core CRUD operations like `create`, `update`, `view/list` and `delete` of IP sets.

This wrapper requires `ipset` to be available on the target system (typically a Linux machine).

For more informations, see:
- [ipset](https://ipset.netfilter.org)
- [netfilter](https://www.netfilter.org)
- [ipset manual](https://ipset.netfilter.org/ipset.man.html)

## Prerequisites and testing
This library should work wherever `ipset` is available and executable by the user running the Go program which is using it.

The project includes a preconfigured [Vagrant](https://www.vagrantup.com) machine that can be used as a sandboxed environment for running tests and, eventually, run your own Go programs that link `go-ipset`.

The Vagrant machine definition file is available under folder `/vm`.

### Run the Vagrant VM
- copy of `/vm/Vagrantfile.example` as `/vm/Vagrantfile`
- edit `/vm/Vagrantfile` according to your needs
  - if needed, you can append your own setup code at the end of this file
  - last lines of `Vagrantfile` contain references to `setup.sh`, which will download Go 1.18 and install `ipset` utility
  - line `config.vm.synced_folder "/<path/to/go/src>/go-ipset", "/sandbox"` must be configured appropriately
    - once up, the VM will contain `go-ipset` code and executables in directory `/sandbox`
- from `/vm`, run the VM with command `vagrant up`
  - after the initial setup, you should be able to connect to the VM with command `vagrant ssh`
  - use `vagrant halt` to stop the VM, `vagrant destroy` to destroy it

For more details about Vagrant, [check the official documentation](https://www.vagrantup.com/docs/installation).
