<script lang="ts">
	import '../app.css';
	import Navigation from '$lib/components/Navigation.svelte';
	import { onMount } from 'svelte';
	import { authStore } from '$lib/stores/auth.svelte.js';

	let { children } = $props();

	// Initialize auth state on app startup
	onMount(() => {
		authStore.refreshAuth();
	});
</script>

<div class="bg-background text-foreground min-h-screen">
	<Navigation />

	<main class="container mx-auto px-4 py-8">
		{@render children()}
	</main>

	<footer class="mt-auto border-t">
		<div class="text-muted-foreground container mx-auto px-4 py-6 text-center">
			<p>&copy; 2024 App Name. All rights reserved.</p>
		</div>
	</footer>
</div>
