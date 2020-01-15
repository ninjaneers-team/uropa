# WORK IN PROGRESS!

# urOpa: Declarative configuration for Opa

urOpa provides declarative configuration and drift detection for Open Policy Agent.

[![Build Status](https://travis-ci.com/ninjaneers-team/uropa.svg?branch=master)](https://travis-ci.com/ninjaneers-team/uropa)

## Table of Content

- [**Features**](#features)
- [**Compatibility**](#compatibility)
- [**Installation**](#installation)
- [**Documentation**](#documentation)
- [**License**](#license)

## Features

- **Export**
  Existing Opa policies to a YAML configuration file
  This can be used to backup Opa's policies.
- **Import**  
  Opa's database can be populated using the exported or a hand written config
  file.
- **Diff and sync capabilities**  
  urOpa can diff the policies in the config file and
  the configuration in Open Policy Agent and then sync it as well.
  This can be used to detect config drifts or manual interventions.
- **Reverse sync**  
  urOpa supports a sync the other way as well, meaning if an
  entity is created in Opa and doesn't add it to the config file,
  urOpa will detect the change.
- **Validation**  
  urOpa can validate a YAML file that you backup or modify to catch errors
  early on.
- **Reset**  
  This can be used to drops all entities in Opa's DB.
- **Parallel operations**  
  All Admin API calls to Opa are executed in parallel using multiple
  threads to speed up the sync process.
- **Authentication with Opa**
  Custom HTTP headers can be injected in requests to Opa's Admin API
  for authentication/authorization purposes.
- **Manage Opa's config with multiple config file**  
  Split your Opa's configuration into multiple logical files based on a shared
  set of tags amongst entities.
- **Designed to automate configuration management**  
  urOpa is designed to be part of your CI pipeline and can be used to not only
  push configuration to Opa but also detect drifts in configuration.

## Compatibility

urOpa is compatible with Opa 1.x and Opa Enterprise >= 0.35.

## Installation

### macOS

If you are on macOS, install urOpa using brew:

```shell
$ brew tap ninjaneers-team/uropa
$ brew install uropa
```

### Linux

If you are Linux, you can either use the Debian or RPM archive from
the Github [release page](https://github.com/ninjaneers-team/uropa/releases)
or install by downloading the binary:

```shel
$ curl -sL https://github.com/ninjaneers-team/uropa/releases/download/v0.7.0/deck_0.7.0_linux_amd64.tar.gz -o uropa.tar.gz
$ tar -xf uropa.tar.gz -C /tmp
$ sudo cp /tmp/uropa /usr/local/bin/
```

### Docker image

Docker image is hosted on [Docker Hub](https://hub.docker.com/r/ninjaneers-team/uropa).

You can get the image with the command:

```
docker pull ninjaneers-team/uropa
```

## Documentation

You can use `--help` flag once you've urOpa installed on your system
to get help in the terminal itself.

The project's documentation site is
[https://uropa.github.com](https://uropa.github.com).

## Changelog

Changelog can be found in the [CHANGELOG.md](CHANGELOG.md) file.

## License

urOpa is licensed with Apache License Version 2.0.
Please read the [LICENSE](LICENSE) file for more details.

## Special Thanks

This project based on [hbagdi/deck](https://github.com/hbagdi/deck). Thanks to [Harry](https://github.com/hbagdi) for maintaining decK! 
