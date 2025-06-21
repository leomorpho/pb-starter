<script lang="ts">
	import ThemeToggle from '$lib/components/ThemeToggle.svelte';
</script>

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
				<a href="/login" class="hover:text-primary transition-colors">Login</a>
				<a href="/signup" class="hover:text-primary transition-colors">Sign Up</a>
				<ThemeToggle />
			</div>
		</nav>
	</div>
</header>