# Labels

A label it just a key value pair, and is typically desgniated by the `=` equals sign in the middle of a string, like so:
- `color=green`
- `develop=true`
- `environment=production`
- `project=foobar`

Labels are heavily used by `gsh`. You can add a label to any item in your inventory, and some inventory items take or use specific labels.


## Hosts
Hosts use the `address` label to connect via ssh client, eg: `address=myhost.example.org`

## Groups
Groups use assigned labels to select hosts with the same label. <br />
For example, doing `gsh make group green color=green` would create a group called "green", and all hosts that have `color=green` labels assigned to them, then the same hosts would be in the group.

## Scripts

