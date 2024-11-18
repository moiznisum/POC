// tools/executors/go/go.executor.ts

import { ExecutorContext } from '@nrwl/devkit';
import { execSync } from 'child_process';

export default async function runExecutor(options: any, context: ExecutorContext) {
  console.log('Building Go backend...');

  try {
    execSync('cd server && go build -o instance -v', { stdio: 'inherit' });

    execSync('./server/instance', { stdio: 'inherit' });
    return { success: true };
  } catch (err) {
    console.error('Error running Go application:', err);
    return { success: false };
  }
}
