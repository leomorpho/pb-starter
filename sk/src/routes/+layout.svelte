<script lang="ts">
	import '../app.css';
	import ThemeToggle from '$lib/components/ThemeToggle.svelte';

	let { children } = $props();
</script>

<div class="min-h-screen bg-background text-foreground">
	<header class="border-b">
		<div class="container mx-auto px-4 py-4">
			<nav class="flex items-center justify-between">
				<div class="flex items-center space-x-4">
					<h1 class="text-xl font-semibold">
						<a href="/" class="hover:text-primary transition-colors">
							App Name
						</a>
					</h1>
				</div>
				
				<div class="flex items-center space-x-4">
					<a href="/" class="hover:text-primary transition-colors">Home</a>
					<a href="/about" class="hover:text-primary transition-colors">About</a>
					<ThemeToggle />
				</div>
			</nav>
		</div>
	</header>

	<main class="container mx-auto px-4 py-8">
		{@render children()}
	</main>

	<footer class="border-t mt-auto">
		<div class="container mx-auto px-4 py-6 text-center text-muted-foreground">
			<p>&copy; 2024 App Name. All rights reserved.</p>
		</div>
	</footer>
</div>
