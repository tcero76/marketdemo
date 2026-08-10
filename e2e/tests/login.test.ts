import { test, expect } from '@playwright/test';

test('login test', async ({ page }) => {
    await page.goto('https://sugarfever.ddns.net/demo');
    await expect(
        page.getByText('Marketplace', { exact: true })
    ).toBeVisible();
    await page.getByRole('button', { name: 'login' }).click();
    await page.waitForLoadState('load');
    const email = page.locator('#email');
    const password = page.locator('#password');
    await email.pressSequentially('leonardo', { delay: 100 });
    await password.pressSequentially('password', { delay: 100 });
    await page.getByRole('button', { name: 'Login', exact: true }).click();
    await page.waitForTimeout(3000);
});