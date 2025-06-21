<script lang="ts">
	import ThemeToggle from '$lib/components/ThemeToggle.svelte';
	import { Home, Info, LogIn } from 'lucide-svelte';
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
			
			<div class="flex items-center space-x-2">
				<a href="/" class="inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium transition-colors hover:bg-accent hover:text-accent-foreground h-9 w-9" title="Home">
					<Home class="h-4 w-4" />
				</a>
				<a href="/about" class="inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium transition-colors hover:bg-accent hover:text-accent-foreground h-9 w-9" title="About">
					<Info class="h-4 w-4" />
				</a>
				<a href="/login" class="inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium transition-colors hover:bg-accent hover:text-accent-foreground h-9 w-9" title="Sign In">
					<LogIn class="h-4 w-4" />
				</a>
				<ThemeToggle />
			</div>
		</nav>
	</div>
</header>