import { exec } from 'child_process';

exec(
  'cd processing-wasm && GOOS=js GOARCH=wasm go build -o ../public/bin.wasm',
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
