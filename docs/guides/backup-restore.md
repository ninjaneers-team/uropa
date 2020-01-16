# Backup and restore of Opa's configuration

You can use urOpa to backup and restore the entire or only a subset of
Opa's configuration.

To back up Opa's configuration, use the `dump` command:

```shell
$ uropa dump
# this generates a Opa.yaml file with the entire configuration of Opa
```

And then restore this file back to Opa using `sync` command:

```shell
$ uropa diff # a dry-run where urOpa shows the changes it will perform
$ uropa sync # actually re-creates the entities in Opa
```
