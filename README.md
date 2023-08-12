# BusLive's API

## Prerequisites

- Go (>= 1.20)
- Terraform

## Development

```shell
make server.run  # run local dev server
```

## Deployment

There are several Terraform workspaces by each deployment stages.

- `default` (stand for production environment)
- `dev`

```shell
terraform -chdir=deploy/terraform workspace select <STAGE>

make lambda.build
make lambda.deploy
```
