import cors from 'cors';
import express from 'express';

const app = express();

app.use(cors());
app.use(express.static('dist'));

app.get('/*', (_, res) => res.sendFile('index.html'));

app.listen(
  3000,
  () => console.log('Serving static files on port 3000'),
);
