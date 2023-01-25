## filtering-camera

Engines: Node **v18**, Golang **v1.19**

**DEV**: https://localhost:3000

**PRODUCTION**: *TBD*

### Deploy

```shell script
git clone https://github.com/peterdee/filtering-camera
cd ./filtering-camera
nvm use 18
npm ci
```

### Configure HTTPS

HTTPS is required for camera access on mobile devices

Install `mkcert`

```shell script
brew install mkcert
```

Generate certificate for `localhost`

```shell script
mkcert -key-file key.pem -cert-file cert.pem example.com *.example.com localhost
```

Copy generated `cert.pem` and `key.pem` to [/serve](/serve) directory

### Launch

Run locally

```shell script
npm run dev
```

Build application

```shell script
npm run build
```

Serve static files

```shell script
npm run serve
```

### Rebuild WASM

Use Go **v1.19**

```shell script
gvm use go1.19
```

Compile WASM binary using Node script

```shell script
npm run compile
```

Alternatively, compile WASM binary manually:

1. Navigate to [/processing-wasm](/processing-wasm) directory

```shell script
cd ./processing-wasm
```

2. Compile WASM binary

```shell script
GOOS=js GOARCH=wasm go build -o ../public/bin.wasm
```

### Linting

Using [ESLint](https://eslint.org)

### License

[MIT](LICENSE.md)
