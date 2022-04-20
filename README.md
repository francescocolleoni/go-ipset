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

### Testing
- assuming that `make` is available on the target development environment, run `make` from the project root to run tests

## Supported options
The following list illustrates options supported by `go-ipset` in various scenarios; sets of alternative options are enclosed in `{}`, where options are separated by operator `|`, while `[<term>]` indicates that `<term>` is optional:
- `bitmap:ip`
  - create
    - **`range { fromip-toip | ip/cidr }`** `[ netmask cidr ] [ timeout value ] [ counters ] [ comment ] [ skbinfo ]`
  - add, delete
    - `{ ip | fromip-toip | ip/cidr }`
  - test
    - `ip`
- `bitmap:ip,mac`
  - create
    - **`range { fromip-toip | ip/cidr }`** `[ timeout value ] [ counters ] [ comment ] [ skbinfo ]`
  - add, delete, test
    - `ip[,macaddr]`
- `bitmap:port`
  - create 
    - **`range fromport-toport`** `[ timeout value ] [ counters ] [ comment ] [ skbinfo ]`
  - add, delete
    - `{ [proto:]port | [proto:]fromport-toport }`
  - test
    - `[proto:]port`
- `hash:ip`
  - create
    - `[ family { inet | inet6 } ]` **or**
    - `[ hashsize value ] [ maxelem value ] [ netmask cidr ] [ timeout value ] [ counters ] [ comment ] [ skbinfo ]`
  - add, delete, test
    - `ip`
- `hash:mac`
  - create
    - `[ hashsize value ] [ maxelem value ] [ timeout value ] [ counters ] [ comment ] [ skbinfo ]`
  - add, delete, test
    - `macaddr`
- `hash:ip,mac`
  - create
    - `[ family { inet | inet6 } ]` **or**
    - `[ hashsize value ] [ maxelem value ] [ timeout value ] [ counters ] [ comment ] [ skbinfo ]`
  - add, delete, test
    - `ip,macaddr`
- `hash:net`
  - create
    - `[ family { inet | inet6 } ]` **or**
    - `[ hashsize value ] [ maxelem value ] [ timeout value ] [ counters ] [ comment ] [ skbinfo ]`
  - add, delete, test
    - ip[/cidr]
- `hash:net,net`
  - create
    - `[ family { inet | inet6 } ]` **or**
    - `[ hashsize value ] [ maxelem value ] [ timeout value ] [ counters ] [ comment ] [ skbinfo ]`
  - add, delete, test
    - `ip[/cidr],ip[/cidr]`
- `hash:ip,port`
  - create
    - `[ family { inet | inet6 } ]` **or**
    - `[ hashsize value ] [ maxelem value ] [ timeout value ] [ counters ] [ comment ] [ skbinfo ]`
  - add, delete, test
    - `ip,[proto:]port`
- `hash:net,port`
  - create
    - `[ family { inet | inet6 } ]` **or**
    - `[ hashsize value ] [ maxelem value ] [ timeout value ] [ counters ] [ comment ] [ skbinfo ]`
  - add, delete, test
    - `ip[/cidr],[proto:]port`
- `hash:ip,port,ip`
  - create
    - `[ family { inet | inet6 } ]` **or**
- `[ hashsize value ] [ maxelem value ] [ timeout value ] [ counters ] [ comment ] [ skbinfo ]`
  - add, delete, test
    - `ip,[proto:]port,ip`
- `hash:ip,port,net`
  - create
    - `[ family { inet | inet6 } ]` **or**
    - `[ hashsize value ] [ maxelem value ] [ timeout value ] [ counters ] [ comment ] [ skbinfo ]`
  - add, delete, test
    - `ip,[proto:]port,ip[/cidr]`
- `hash:ip,mark`
  - create 
    - `[ family { inet | inet6 } ]` **or**
    - `[ markmask value ] [ hashsize value ] [ maxelem value ] [ timeout value ] [ counters ] [ comment ] [ skbinfo ]`
  - add, delete, test
    - ip,mark
- `hash:net,port,net`
  - create
    - `[ family { inet | inet6 } ]` **or**
    - `[ hashsize value ] [ maxelem value ] [ timeout value ] [ counters ] [ comment ] [ skbinfo ]`
  - add, delete, test
    - `ip[/cidr],[proto:]port,ip[/cidr]`
- `hash:net,iface`
  - create
    - `[ family { inet | inet6 } ]` **or**
    - `[ hashsize value ] [ maxelem value ] [ timeout value ] [ counters ] [ comment ] [ skbinfo ]`
  - add, delete, test
    - `ip[/cidr],[physdev:]iface`
- `list:set`
  - create
    - `[ size value ] [ timeout value ] [ counters ] [ comment ] [ skbinfo ]`
  - add, delete, test
    - `setname [ { before | after } setname ]`

## Supported sets
- bitmaps
  - `bitmap:ip`
  - `bitmap:ip,mac`
  - `bitmap:port`
- hashes
  - `hash:ip`
  - `hash:mac`
  - `hash:ip,mac`
  - `hash:net`
  - `hash:net,net`
  - `hash:ip,port`
  - `hash:net,port`
  - `hash:ip,port,ip`
  - `hash:ip,port,net`
  - `hash:ip,mark`
  - `hash:net,port,net`
  - `hash:net,iface`
- lists
  - `list:set`

## Supported IP, port and network interface formats
`go-ipset` supports different formats when defining arguments for create set, add/delete entries to/from sets and testing entry containment for a given set; such formats depend on set type, as listed below (ref.: [ipset manual](https://ipset.netfilter.org/ipset.man.html)).

IP addresses are always IPv4, while IP addresses expressed as ip/cidr strings must be formatted as four pairs of numbers, separated by `.` followed by `/<cidr>`; this means that addresses like `1.1.1.1/16` are acceptable, while those like `1.1.0/16` are not.
- IP addresses and intervals
  - `ip`
    - ex.: `1.1.1.1`
    - **not supported**: `1.1.0`
  - `ip[/cidr]`
    - ex.: `1.1.1.1`
    - ex.: `1.1.1.1/16`
    - **not supported**: `1.1.0/16`
  - `fromip-toip`
    - ex.: `1.1.1.1-2.2.2.2`
- port numbers and intervals
  - `port`:
    - ex.: `12345`
  - `[proto:]port`:
    - ex.: `12345`
    - ex.: `tcp:12345`
  - `fromport-toport`
    - ex.: `12345-56789`
  - `[proto:]fromport-toport`
    - ex.: `12345-56789`
    - ex.: `tcp:12345-56789`
- other
  - `macaddr`
    - ex.: `A1:2B:C3:4D:E5:6F`
  - `mark` 
    - any value in closed (inclusive) interval `[0, 4294967295]`
  - `iface`
    - ex.: `eth0`
  - `[physdev:]iface`
    - ex.: `eth0`
    - ex.: `physdev:eth0`
  - `setname`
    - any word that is the name of a set 
