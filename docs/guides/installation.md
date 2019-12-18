# Installation

urOpa is entirely written in Go. The build process builds a single static binary,
which makes it easy and convenient to install urOpa.

You can follow along installation instructions based on your
Operating System(OS):

### macOS

If you are on macOS, install urOpa using brew:

```shell
$ brew tap ninjaneers-team/uropa
$ brew install uropa
```

### Linux

If you are Linux, you can either use the Debian or RPM archive from
the Github [release page](https://github.com/ninjaneers-team/uropa/releases)
or install by downloading a compressed archive, which contains the binary:

```shel
$ curl -sL https://github.com/ninjaneers-team/uropa/releases/download/v0.7.0/deck_0.7.0_linux_amd64.tar.gz -o uropa.tar.gz
$ tar -xf uropa.tar.gz -C /tmp
$ sudo cp /tmp/uropa /usr/local/bin/
```

### Docker image

If your workflow requires a Docker image, then you can use `ninjaneers-team/uropa` Docker
image from the official Docker hub:

```
docker pull ninjaneers-team/uropa
```

You will have to mount the state files into the container as volumes so that
urOpa can read the files during diff/sync procedures.

If you're integrating urOpa into your CI system, you can either install urOpa
into the system itself, use the Docker based environment or pull the binaries
hosted on Github in each job.

