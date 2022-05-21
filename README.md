# MoMathF Server
Call MathPix API to convert image to Latex formula

## API

We call the API from MathPix, you can apply for it from https://accounts.mathpix.com/account. You can fill id and key in the config_git.yaml.

```yaml
api:
  api-id: xxxxxxx
  api-key: xxxxxxxxxxxxxxxxx
```

## Build

```shell
go build -tags dev/prod -o ./build
```

