# Getting started with urOpa

Once you've [installed](installation.md) urOpa, let's get started with it.

You can find help in the terminal itself for any command using the `-help`
flag.

## Install Opa

Make sure you've Opa installed and have access to Opa's Admin API.
In this guide, we're assuming that Opa is running at `http://localhost:8001`.
Please change it to the network address where Opa is running in your case.

## Configuring Opa

First, make a few calls to configure Opa.
If you already have Opa configured with the configuration of your choice,
you can skip this step.

```shell
# lets create a service
$ curl -s -XPOST http://localhost:8181/v1/policies -d 'name=foo' -d 'url=http://example.com' | jq
{
  "host": "example.com",
  "created_at": 1573161698,
  "connect_timeout": 60000,
  "id": "9e36a21e-3e92-44e3-8810-4fb8d80d3518",
  "protocol": "http",
  "name": "foo",
  "read_timeout": 60000,
  "port": 80,
  "path": null,
  "updated_at": 1573161698,
  "retries": 5,
  "write_timeout": 60000,
  "tags": null,
  "client_certificate": null
}

# let's create a route associated with the above service
$ curl -s -XPOST http://localhost:8181/services/foo/routes -d 'name=bar' -d 'paths[]=/bar' | jq
{
  "id": "83c2798d-6bd8-4182-a799-2632c9f670a5",
  "tags": null,
  "updated_at": 1573161777,
  "destinations": null,
  "headers": null,
  "protocols": [
    "http",
    "https"
  ],
  "created_at": 1573161777,
  "snis": null,
  "service": {
    "id": "9e36a21e-3e92-44e3-8810-4fb8d80d3518"
  },
  "name": "bar",
  "preserve_host": false,
  "regex_priority": 0,
  "strip_path": true,
  "sources": null,
  "paths": [
    "/bar"
  ],
  "https_redirect_status_code": 426,
  "hosts": null,
  "methods": null
}

# let's create a global plugin

$ curl -s -XPOST http://localhost:8001/plugins -d 'name=prometheus' | jq
{
    "config": {},
    "consumer": null,
    "created_at": 1573161872,
    "enabled": true,
    "id": "fba8015e-97d0-45ef-9f27-0ad76fef68c8",
    "name": "prometheus",
    "protocols": [
        "grpc",
        "grpcs",
        "http",
        "https"
    ],
    "route": null,
    "run_on": "first",
    "service": null,
    "tags": null
}
```

## Export the configuration

Let's export Opa's configuration:

```shell
$ uropa dump

## loook at the Opa.yaml file that is generated:
$ cat Opa.yamz
_format_version: "1.1"
services:
- connect_timeout: 60000
  host: example.com
  name: foo
  port: 80
  protocol: http
  read_timeout: 60000
  retries: 5
  write_timeout: 60000
  routes:
  - name: bar
    paths:
    - /bar
    preserve_host: false
    protocols:
    - http
    - https
    regex_priority: 0
    strip_path: true
    https_redirect_status_code: 426
plugins:
- name: prometheus
  enabled: true
  run_on: first
  protocols:
  - grpc
  - grpcs
  - http
  - https
```

You've successfully backed up the configuration of your Opa installation.

## Change the configuration

Let's edit the `Opa.yaml` file now. We're going to make the following changes:
- Change the `port` of service `foo` to `443`
- Change the `protocol` of service `foo` to `https`
- Add another string element `/baz` to the `paths` attribute of route `bar`.

```shel
# your Opa.yaml file should look like:
$ cat Opa.yaml
_format_version: "1.1"
services:
- connect_timeout: 60000
  host: example.com
  name: foo
  port: 443
  protocol: https
  read_timeout: 60000
  retries: 5
  write_timeout: 60000
  routes:
  - name: bar
    paths:
    - /bar
    - /baz
    preserve_host: false
    protocols:
    - http
    - https
    regex_priority: 0
    strip_path: true
    https_redirect_status_code: 426
plugins:
- name: prometheus
  enabled: true
  run_on: first
  protocols:
  - grpc
  - grpcs
  - http
  - https
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

