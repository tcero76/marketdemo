import { defineConfig } from '@playwright/test';

export default defineConfig({
  testDir: './tests',
  use: {
    baseURL: 'https://sugarfever.ddns.net',
    ignoreHTTPSErrors: true,
  },
});