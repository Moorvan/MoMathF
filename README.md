# MoMathF
Call MathPix API to convert image to Latex formula

## API

We call the API from MathPix, you can apply for it from https://accounts.mathpix.com/account. You can fill id and key in the config_git.yaml and rename the file to config.yaml:

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

```shell
./build/MoMathF -path ./testPics/1.png
```

The following content will be copied to your clipboard:

```txt
\varphi^{k} \stackrel{\operatorname{def}}{=} I\left(V^{0}\right) \wedge \operatorname{path}^{0, k} \wedge\left(\neg P\left(V^{k}\right)\right)
```

