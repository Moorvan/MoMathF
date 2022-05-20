# MoMathF
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

## Run

Usage:

```shell
Usage of MoMathF:
  -config string
        config file path (default "./config.yaml")
  -image string
        image file path
```

cmd:

```shell
./build/MoMathF -config ./config_git.yaml -image ./testPics/1.png
```

The following content will be copied to your clipboard:

```txt
\varphi^{k} \stackrel{\operatorname{def}}{=} I\left(V^{0}\right) \wedge \operatorname{path}^{0, k} \wedge\left(\neg P\left(V^{k}\right)\right)
```

