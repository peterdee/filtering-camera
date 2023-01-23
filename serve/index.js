import cors from 'cors';
import { createServer } from 'node:https';
import express from 'express';
import { readFileSync } from 'node:fs';

const app = express();

app.use(cors());
app.use(express.static(`${process.cwd()}/dist`));

app.get('/*', (_, res) => res.sendFile('index.html'));

createServer(
  {
    cert: readFileSync(`${process.cwd()}/serve/cert.pem`),
    key: readFileSync(`${process.cwd()}/serve/key.pem`),
  },
  app,
).listen(
  3000,
  () => console.log('Serving static files on port 3000'),
);
