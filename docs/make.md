# Make

Aliases:
- `make`
- `set`
- `add`

## Making Hosts, Groups, Scripts, and Files
Your inventory is empty until you start adding items.<br />
These are the types you can add to your inventory with `make`:
- [`Hosts`](hosts.md)
- [`Groups`](groups.md)
- [`Scripts`](scripts.md)
- [`Files`](files.md)

Additionally, you can append [`labels`](labels.md) to any of these.

```bash
$ # For Hosts and Groups
$ gsh make <type> <name> [key=value...]
```
```bash
$ # For Scripts and Files
$ gsh make <type> <name> <path/to/file> [key=value...]
```
## Examples:

```bash
$ # setting host labels
$ gsh make host dev0 address=dev0.local develop=true project=foo
╔══════╕ Set!
║ Host └──────────╮
╿ Name  :  Labels
└ dev0 : address=dev0.local develop=true project=foo
```
```bash
$ # making a group
$ gsh make group develop develop=true
╔═══════╕ Set!
║ Group └─────────╮
╿ Name  :  Labels
└ develop : develop=true
```
```bash
$ # adding a script will calculate the sha256sum of the script
$ gsh make script hostname ./scripts/hostname.sh some=label another=label
╔════════╕ Set!
║ Script └────────╮
╿ Name  :  Labels
└ hostname : lang=bash revision=da5070f9 some=label another=label
```
```bash
$ # adding a file will also calculate the sha256sum
$ gsh make file nginx.conf ./configs/nginx.cfg class=webserver
╔══════╕ Set!
║ File └──────────╮
╿ Name  :  Labels
└ nginx.conf : filename=nginx.cfg revision=6fea02c9 class=webserver
```

^^ Here are four examples, one for each type of item you can add to your inventory.
