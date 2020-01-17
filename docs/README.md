# uropa Documentation

uropa provides declarative configuration and drift detection for Opa.

## Table of content

- [References](#references)
- [FAQS](#frequently-asked-questions-faqs)
- [Changelog](#changelog)
- [Licensing](#licensing)
- [Security](#security)
- [Getting help](#getting-help)
- [Reporting a bug](#reporting-a-bug)

## Design

- [Terminology](terminology.md)
- [Architecture](design-architecture.md)

## Guides

- [Installation](guides/installation.md)
- [Getting started with uropa](guides/getting-started.md)
- [Backup and restore of Opa's configuration](guides/backup-restore.md)
- [Configuration as code and Git-ops using uropa](guides/ci-driven-configuration.md)
- [Best practices for using uropa](guides/best-practices.md)
- [Using multiple files to store configuration](guides/multi-file-state.md)

## References

The command-line `--help` flag on the main command or a sub-command (like diff,
sync, reset, etc.) shows the help text along with supported flags for those
commands.

A gist of all commands that are available in uropa can be found
[here](commands.md).

## Frequently Asked Questions (FAQs)

You can find answers to FAQs [here](faqs.md).

## Changelog

Changelog can be found in the
[CHANGELOG.md](https://github.com/ninjaneers-team/uropa/blob/master/CHANGELOG.md) file.

## Licensing

uropa is licensed with Apache License Version 2.0.
Please read the
[LICENSE](https://github.com/ninjaneers-team/uropa/blob/master/LICENSE) file for more details.

## Security

uropa does not offer to secure your Opa deployment but only configures it.
It encourages you to protect your Opa's Admin API with authentication but
doesn't offer such a service itself.

uropa's state file can contain sensitive data such as private keys of
certificates, credentials, etc. It is left up to the user to manage
and store the state file in a secure fashion.

If you believe that you have found a security vulnerability in uropa, please
submit a detailed report, along-with reproducible steps
to Harry Bagdi (email address is first name last name At gmail Dot com).
I will try to respond in a timely manner and will really appreciate it you
report the issue privately first.

## Getting help

One of the design goals of uropa is deliver a good developer experience to you.
And part of it is getting the required help when you need it.
To seek help, use the following resources:
- `--help` flag gives you the necessary help in the terminal itself and should
  solve most of your problems.
- Please read through the pages under the `docs` directory of this repository.
- If you still need help, please open a
  [Github issue](https://github.com/ninjaneers-team/uropa/issues/new) to ask your
  question.

One thing I humbly ask for when you need help or run into a bug is patience.
I'll do my best to respond you at the earliest possible.

## Reporting a bug

If you believe you have run into a bug with uropa, please open
a [Github issue](https://github.com/ninjaneers-team/uropa/issues/new).

If you think you've found a security issue with uropa, please read the
[Security](#security) section.
