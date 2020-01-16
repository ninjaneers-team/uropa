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

