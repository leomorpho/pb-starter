import { expect, test } from '@playwright/test';

test('home page has expected h1', async ({ page }) => {
	await page.goto('/');
	// Check for the main hero heading specifically
	await expect(page.getByRole('heading', { name: 'MODERN SAAS PLATFORM' })).toBeVisible();
	// Also check that the navigation heading is present
	await expect(page.getByRole('heading', { name: 'My SvelteKit App' })).toBeVisible();
});
