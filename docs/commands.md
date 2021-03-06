# uropa commands

This document describes the various commands that are available in uropa.
The command-line `--help` flag on the main command or a sub-command (like diff,
sync, reset, etc.) shows the help text along-with supported flags for those
commands, and is the canonical documentation. Below is a short summary of
these commands:

- [ping](#ping)
- [dump](#dump)
- [diff](#diff)
- [sync](#sync)
- [reset](#reset)
- [validate](#validate)
- [version](#version)
- [help](#help)

## ping

This command can be used to verify connectivity between Opa and uropa.
Under the hood, it sends a `GET /` request to Opa to verify if Opa's
Admin API is reachable and uropa can authenticate itself against it.
If uropa is being used in automated environment (like in a CI), it is
recommended that you use this command before a diff/sync to ensure
connectivity.

## dump

This command can be used to export all of Opa's configuration into a single
YAML file. All entities are exported by default.

`--select-tag` can be used to export entities with the specific tag only.
This flag can be used in the following cases:

- If you would like to manage only a subset of entities in Opa.
- If multiple teams would like to configure Opa, one team can export
  and sync it's configuration without being aware of any other teams'
  configuration.

If you are a Opa Enterprise user, you can specify a specific workspace that
you want to export using `--workspace` flag or use `--all-workspaces` flag
to export routing configuration of all workspaces.

## diff

This command compares the content of the input file against the current
configuration of Opa.
You can use this command for drift detection i.e. if the configuration
of Opa is out of sync with configuration of the input file.

## sync

This command will create, update or delete entities in Opa to exactly match
as described via the input file. You can use `diff` command to display
the actions that uropa will take and then use `sync` command to actually
perform these actions.

## validate

This command can be used to validate an existing state file or a set of files.
It can catch most errors including validation of the YAML/JSON file itself and
catching duplicates or malformed entities.

## reset

This command will delete all the entities in Opa. Please use this
command with extreme caution as the actions are irreversible.

## version

This command shows the version information of the uropa binary that is currently
in use.

## help

This command shows the help text of uropa. Use `--help` flag on any of the
above command to get help in your terminal itself.
