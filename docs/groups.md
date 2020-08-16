# Groups
A group represents collection of [`hosts`](hosts.md) on the network.

Aliases:
- `group`
- `groups`

## Making, Updating, Deleting, and Listing Groups
Your group inventory is empty until you have added a group.<br />
A group will only contain hosts if the hosts also have matching labels.
These are the types operations you can do with a group:
- [`make`](make.md)         # Make or Update a group in inventory
- [`delete`](delete.md)     # Delete a group from inventory
- [`show`](show.md)         # Show or List group
- [`execute`](execute.md)   # Execute scripts on groups of hosts
- [`copy`](copy.md)         # Copy files/scripts to groups of hosts

## Special Labels

Groups use their assigned labels to determine which hosts are in the group.
For example a group made as: `gsh make group baz project=baz`, this will create a group called `baz` and all hosts with `project=baz` as a label will be in the group.

## Examples:

```bash
$ # making a group with labels
$ gsh make group develop develop=true
╔═══════╕ Set!
║ Group └─────────╮
╿ Name  :  Labels
└ develop : develop=true
```
```bash
$ # updating the same group with different labels
$ gsh make group develop environment=develop
╔═══════╕ Set!
║ Group └─────────╮
╿ Name  :  Labels
└ develop : environment=develop
```
```bash
$ # deleting a group
$ gsh del group demo1
╔═══════╕ Deleted!
║ Group └─────────╮
╿ Name  :  Labels
└ demo1 : demo=true
```