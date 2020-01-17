# uropa Helm Chart

This is an implementation of uropa batch job for kubernetes to migrate Open Policy Agent policies

## Pre Requisites
* Kubernetes 1.9+

## Chart Details
This chart will do the following:
* Start a job container as post-install step to migrate policies 

## Configuration

The chart can be customized using the following configurable parameters:

| Parameter                       | Description                                                     | Default                      |
| ------------------------------- | ----------------------------------------------------------------| -----------------------------|
| `image.repository`              | Container image name                                  | `ninjaneers/uropa` |
| `image.tag`                     | Container image tag                                   | `latest`                    |
| `image.pullPolicy`              | Container pull policy                                 | `Always`                     |
| `imagepullSecret`              | Pod pull secret                                       | ``                     |
| `config.opaHost`                  | Opa host url                                         | `http://localhost:8181`                  |
| `config.data`           | Opa.yaml configuration file        | `{}`                         |

Specify parameters using `--set key=value[,key=value]` argument to `helm install`

Alternatively a YAML file that specifies the values for the parameters can be provided like this: