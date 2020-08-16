# Delete

Aliases:
- `rm`
- `remove`
- `del`
- `delete`

## Deleting Hosts, Groups, Scripts, and Files
You can remove or delete items from your inventory after adding them.<br />
These are the types you can remove from your inventory with `delete`:
- [`Hosts`](hosts.md)
- [`Groups`](groups.md)
- [`Scripts`](scripts.md)
- [`Files`](files.md)

The caveat here is that you can only supply a `<type>` and a `<name>`

```bash
$ # For Hosts, Groups, Scripts, and Files
$ gsh delete <type> <name>
```
## Examples:

```bash
$ # deleting a host
$ gsh del host demo1
╔══════╕ Deleted!
║ Host └──────────╮
╿ Name  :  Labels
└ demo1 : address=demo1.local demo=true
```
```bash
$ # deleting a group
$ gsh delete group demo
╔═══════╕ Deleted!
║ Group └─────────╮
╿ Name  :  Labels
└ demo : demo=true
```
```bash
$ # deleting a script
$ gsh delete script cpuinfo
╔════════╕ Deleted!
║ Script └────────╮
╿ Name  :  Labels
└ cpuinfo : lang=bash revision=8f0aaf02
```
```bash
$ # deleting a file
$ gsh delete file nginx.conf
╔══════╕ Deleted!
║ File └──────────╮
╿ Name  :  Labels
└ nginx.conf : filename=nginx.cfg revision=6fea02c9 class=webserver
```

^^ Here are four examples, one for each type of item you can delete from your inventory.
