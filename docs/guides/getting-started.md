# Getting started with urOpa

Once you've [installed](installation.md) urOpa, let's get started with it.

You can find help in the terminal itself for any command using the `-help`
flag.

## Install Opa

Make sure you've Opa installed and have access to Opa's Admin API.
In this guide, we're assuming that Opa is running at `http://localhost:8181`.
Please change it to the network address where Opa is running in your case.

## Create the configuration

Let's create the `Opa.yaml` file now. We're going to make the following changes:

```shell
# your Opa.yaml file should look like:
$ cat Opa.yaml
_format_version: "1.1"
policies:
  - id: example
    raw: |-
       package example
       default allow = false
```

## diff and sync the configuration to Opa

```
# let's perform a diff
uropa diff
# you should see urOpa reporting that the properties you had changed
# in the file are going to be changed by urOpa in Opa's database.

# let's apply the changes
uropa sync

# curl Opa's Admin API to see the updated route and service in Opa.

# you can also run the diff command, which will report no changes
uropa diff
```

## Drift detection using urOpa

Go ahead and now create a consumer in Opa.

```shell
$ curl -s -XPOST http://localhost:8001/consumers -d 'username=dodo' | jq
{
  "custom_id": null,
  "created_at": 1573162649,
  "id": "ed32faa1-9105-488e-8722-242e9d266717",
  "tags": null,
  "username": "dodo"
}
```

Note that we have created this consumer in Opa but the consumer doesn't exist
in `Opa.yaml` file we've saved on disk.

Let's see what urOpa reports on a diff now.

```shell
$ uropa diff
deleting consumer dodo
```

Since the file does not contain the consumer definition, urOpa reports that
a `sync` run will delete the consumer from Opa's database.

Let's go ahead and run the sync process.

```shell
$ uropa sync
```

Now, looking up curl http://localhost:8001/consumers/dodo
{"message":"Not found"}the consumer in Opa's database will return a `404`:

```shell
$ curl http://localhost:8001/consumers/dodo
{"message":"Not found"}
```

This shows how urOpa can detect changes done directly using Opa's Admin API
can be detected by urOpa. You can configure your CI or run a `cronjob` in which
urOpa detects if any changes exist in Opa that are not part of your configuration
file, and alert your teams if such a discrepancy is present.


## Reset your configuration

Finally, you can reset the configuration of Opa using urOpa.
The changes performed by this command are irreversible(unless you've created a
backup using `uropa dump`) so please be careful.


```shell
$ uropa reset
This will delete all configuration from Opa's database.
> Are you sure? y
```

And that's it.
Start using urOpa to declaratively configure your Opa installation today!

