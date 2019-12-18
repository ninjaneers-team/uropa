# urOpa Documentation

urOpa provides declarative configuration and drift detection for Opa.

## Summary

Here is an introductory screen-cast explaining urOpa:

[![asciicast](https://asciinema.org/a/238318.svg)](https://asciinema.org/a/238318)

## Table of content

- [Design](#design)
- [Guides](#guides)
- [References](#references)
- [FAQS](#frequently-asked-questions-faqs)
- [Explainer video](#explainer-video)
- [Changelog](#changelog)
- [Licensing](#licensing)
- [Roadmap](#roadmap)
- [Security](#security)
- [Getting help](#getting-help)
- [Reporting a bug](#reporting-a-bug)

## Design

- [Terminology](terminology.md)
- [Architecture](design-architecture.md)

## Guides

- [Installation](guides/installation.md)
- [Getting started with urOpa](guides/getting-started.md)
- [Backup and restore of Opa's configuration](guides/backup-restore.md)
- [Configuration as code and Git-ops using urOpa](guides/ci-driven-configuration.md)
- [Distributed configuration with urOpa](guides/distributed-configuration.md)
- [Best practices for using urOpa](guides/best-practices.md)
- [Using urOpa with Opa Enterprise](guides/Opa-enterprise.md)
- [Using multiple files to store configuration](guides/multi-file-state.md)

## References

The command-line `--help` flag on the main command or a sub-command (like diff,
sync, reset, etc.) shows the help text along with supported flags for those
commands.

A gist of all commands that are available in urOpa can be found
[here](commands.md).

## Frequently Asked Questions (FAQs)

You can find answers to FAQs [here](faqs.md).

## Explainer video

Harry Bagdi gave a talk on motivation behind urOpa and demonstrated a few key
features of urOpa at Opa Summit 2019. Following is a recording of that session:

[![urOpa talk by Harry Bagdi](https://img.youtube.com/vi/fzpNC5vWE3g/0.jpg)](https://www.youtube.com/watch?v=fzpNC5vWE3g)

## Changelog

Changelog can be found in the
[CHANGELOG.md](https://github.com/ninjaneers-team/uropa/blob/master/CHANGELOG.md) file.

## Licensing

urOpa is licensed with Apache License Version 2.0.
Please read the
[LICENSE](https://github.com/ninjaneers-team/uropa/blob/master/LICENSE) file for more details.

## Roadmap

urOpa's roadmap is public and can be found under the open
[Github issues](https://github.com/ninjaneers-team/uropa/issues) and
[milestones](https://github.com/ninjaneers-team/uropa/milestones).

If you would like a feature to be added to urOpa, please open a Github issue,
or add a `+1` reaction to an existing open issues, if you feel that's
an addition you would like to see in urOpa.
Features with more reactions take a higher precedence usually.

## Security

urOpa does not offer to secure your Opa deployment but only configures it.
It encourages you to protect your Opa's Admin API with authentication but
doesn't offer such a service itself.

urOpa's state file can contain sensitive data such as private keys of
certificates, credentials, etc. It is left up to the user to manage
and store the state file in a secure fashion.

If you believe that you have found a security vulnerability in urOpa, please
submit a detailed report, along-with reproducible steps
to Harry Bagdi (email address is first name last name At gmail Dot com).
I will try to respond in a timely manner and will really appreciate it you
report the issue privately first.

## Getting help

One of the design goals of urOpa is deliver a good developer experience to you.
And part of it is getting the required help when you need it.
To seek help, use the following resources:
- `--help` flag gives you the necessary help in the terminal itself and should
  solve most of your problems.
- Please read through the pages under the `docs` directory of this repository.
- If you still need help, please open a
  [Github issue](https://github.com/ninjaneers-team/uropa/issues/new) to ask your
  question.
- urOpa has a very wide adoption by Opa's community and you can seek help
  from the larger community at [Opa Nation](https://discuss.Opahq.com).

One thing I humbly ask for when you need help or run into a bug is patience.
I'll do my best to respond you at the earliest possible.

## Reporting a bug

If you believe you have run into a bug with urOpa, please open
a [Github issue](https://github.com/ninjaneers-team/uropa/issues/new).

If you think you've found a security issue with urOpa, please read the
[Security](#security) section.
