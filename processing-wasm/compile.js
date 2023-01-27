import { exec } from 'child_process';

exec(
  'cd processing-wasm && tinygo build -o ../public/bin.wasm -target wasm main.go',
  (error, _, stderr) => {
    if (error) {
      throw error;
    }
    if (stderr) {
      throw stderr;
    }
    return process.exit(0);
  },
);
