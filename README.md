# Git Branch Batch Delete

## About

Delete more than one git branch at a time.

### Example

```
$ batchdelete      
? Select branches to delete:  [Use arrows to move, space to select, <right> to all, <left> to none, type to filter]
> [ ]  branch-1
  [ - ]  branch-2
  [ - ]  branch-3
  [ ]  main
$ Are you sure you want to delete the selected branches? (y/N) Y
$ Deleted branch 'branch-2'
$ Deleted branch 'branch-3'
$
```

## Install

`curl -L -O https://github.com/htluff/gitbatchdelete/install.sh`
