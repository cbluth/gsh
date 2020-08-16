# Hosts
A host represents an ssh server running on a host on the network.
[`Groups`](groups.md) can be created to represent groups of hosts.

Aliases:
- `host`
- `hosts`
- `node`
- `nodes`

## Making, Updating, Deleting, and Listing Hosts
Your hosts inventory is empty until you have added a host.<br />
These are the types operations you can do with a host:
- [`make`](make.md)         # Make or Update a host in inventory
- [`delete`](delete.md)     # Delete a host from inventory
- [`show`](show.md)         # Show or List hosts
- [`execute`](execute.md)   # Execute scripts on hosts
- [`copy`](copy.md)         # Copy files/scripts to hosts

## Special Labels

Hosts use special [`labels`](labels.md), please provide an `address` label when adding a host to your inventory.
Supplying an `address` label allows `gsh` to connect to the servers address via ssh.

## Examples:

```bash
$ # making a host with labels
$ gsh make host dev0 address=dev0.local develop=true project=foo
╔══════╕ Set!
║ Host └──────────╮
╿ Name  :  Labels
└ dev0 : address=dev0.local develop=true project=foo
```
```bash
$ # updating the same host with different labels
$ gsh make host dev0 address=dev0.local develop=true project=bar
╔══════╕ Set!
║ Host └──────────╮
╿ Name  :  Labels
└ dev0 : address=dev0.local develop=true project=bar
```
```bash
$ # showing hosts with a label
$ gsh show hosts demo=true
╔═══════╕ Total: 2
║ Hosts └─────────╮
╿ Name  :  Labels
├ demo2 : address=demo2.local demo=true
└ demo3 : address=demo3.local demo=true
```
```bash
$ # deleting a host
$ gsh del host demo1
╔══════╕ Deleted!
║ Host └──────────╮
╿ Name  :  Labels
└ demo1 : address=demo1.local demo=true
```