## filtering-camera

Camera stream processing using Javascript and WebAssembley

Engines: Node **v18**, Golang **v1.19**

**DEV**: https://localhost:3000

**PRODUCTION**: https://camera.dyum.in (https://filtering-camera-wasm.vercel.app)

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

Binary can be compiled using [TinyGo](https://tinygo.org), but it causes WASM module to crash randomly on iOS devices

Install TinyGo (MacOS): https://tinygo.org/getting-started/install/macos

Compile the binary using TinyGo

```shell script
tinygo build -o ../public/bin.wasm -target wasm main.go
```

You would need to replace [/public/wasm_exec.js](/public/wasm_exec.js) file with a proper one if WASM binary was compiled with TinyGo

### Linting

Using [ESLint](https://eslint.org)

### Vercel deployment

Application is automatically deployed to [Vercel](https://vercel.com) and is available at https://filtering-camera-wasm.vercel.app

### License

[MIT](LICENSE.md)
