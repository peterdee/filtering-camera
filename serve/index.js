import compression from 'compression';
import { createServer } from 'node:https';
import express from 'express';
import { readFileSync } from 'node:fs';

const app = express();
const PORT = Number(process.env.PORT) || 3000;

app.use(compression());
app.use(express.static(`${process.cwd()}/dist`));

app.get('/*', (_, res) => res.sendFile('index.html'));

createServer(
  {
    cert: readFileSync(`${process.cwd()}/serve/cert.pem`),
    key: readFileSync(`${process.cwd()}/serve/key.pem`),
  },
  app,
).listen(
  PORT,
  () => console.log(`Serving static files on port ${PORT}`),
);
