# Copy (WIP)

Aliases:
- `copy`
- `cp`

## Copying Files and Scripts to Hosts and Groups
You can copy files or scripts to sets of hosts.<br />

```bash
$ # For copying scripts/files to sets of hosts
$ gsh cp <type> <name> <-n|-g> <host|group> [dst=/target/destination] [chmod=ug+rw,o-a]
```

## Examples:

```bash
$ # copying a file to host demo1
$ gsh copy file sshrc -n demo1 dst=/etc/sshrc chmod=a+rx
╔══════╕ File Copied!
║ File └──────────╮
╿ Name  :  Labels
└ sshrc : scope=host target=demo1 dst=/etc/sshrc chmod=a+rx
```
```bash
$ # copying a file to group
$ gsh copy file nginx.conf -g group1 dst=/etc/nginx/nginx.conf
╔══════╕ File Copied!
║ File └──────────╮
╿ Name  :  Labels
└ nginx.conf : scope=group target=group1 dst=/etc/nginx/nginx.conf
```
