import { pb } from '$lib/pocketbase.js';
import { authStore } from '$lib/stores/auth.svelte.js';

// Initialize auth on client startup
pb.authStore.loadFromCookie(document.cookie);

// Refresh auth token periodically
if (pb.authStore.isValid) {
	authStore.refreshAuth();
}
