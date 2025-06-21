<script lang="ts">
	import '../app.css';
	import Navigation from '$lib/components/Navigation.svelte';

	let { children } = $props();
</script>

<div class="min-h-screen bg-background text-foreground">
	<Navigation />

	<main class="container mx-auto px-4 py-8">
		{@render children()}
	</main>

	<footer class="border-t mt-auto">
		<div class="container mx-auto px-4 py-6 text-center text-muted-foreground">
			<p>&copy; 2024 App Name. All rights reserved.</p>
		</div>
	</footer>
</div>
