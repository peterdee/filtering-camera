## filtering-camera

Camera stream processing using Javascript and WebAssembley

Engines: Node **v22**, Golang **v1.24**

**DEV**: https://localhost:3000

**PRODUCTION**: https://camera.dyum.in / https://filtering-camera-wasm.vercel.app

### Deploy

Clone the repository and install necessary dependencies

```shell script
cd ./filtering-camera
nvm use 22
npm ci
```

### Configure HTTPS

HTTPS is required for camera access on mobile devices

Generate necessary files using OpenSSL

```shell script
# Generate key file
openssl genrsa -out key.pem 2048

# Generate CSR
openssl req -new -sha256 -key key.pem -out csr.csr

# Generate certificate
openssl req -x509 -sha256 -days 365 -key key.pem -in csr.csr -out cert.pem
```

Copy generated `cert.pem` and `key.pem` to [/serve](/serve) directory

### Launch

- Run locally

```shell script
npm run dev
```

- Build application

```shell script
npm run build
```

- Serve static files

```shell script
npm run serve
```

### Compile WASM binary

- Use Go **v1.24**

```shell script
gvm use go1.24
```

- Compile WASM binary using Node script

```shell script
npm run compile-wasm
```

Alternatively, compile WASM binary manually:

- Navigate to [/processing-wasm](/processing-wasm) directory

```shell script
cd ./processing-wasm
```

- Compile WASM binary

```shell script
GOOS=js GOARCH=wasm go build -o ../public/bin.wasm
```

### Linting

Using [ESLint](https://eslint.org)

### Cloud deployment

Application is automatically deployed to [Vercel](https://vercel.com) and is available at https://camera.dyum.in / https://filtering-camera-wasm.vercel.app

### License

[MIT](LICENSE.md)
