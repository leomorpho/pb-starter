<script lang="ts">
	import ThemeToggle from '$lib/components/ThemeToggle.svelte';
	import { Home, Info, LogIn, LogOut, User } from 'lucide-svelte';
	import { authStore } from '$lib/stores/authClient.svelte.js';

	function handleLogout() {
		console.log('🚪 Logout clicked');
		authStore.logout();
		console.log('🚪 Logout completed');
	}

	// Debug: Log reactive updates
	$effect(() => {
		console.log('📱 Navigation reactive update:', {
			isLoggedIn: authStore.isLoggedIn,
			user: authStore.user?.email,
			initialized: authStore.initialized
		});
	});
</script>

<header class="border-b">
	<div class="container mx-auto px-4 py-4">
		<nav class="flex items-center justify-between">
			<div class="flex items-center space-x-4">
				<h1 class="text-xl font-semibold">
					<a href="/" class="hover:text-primary transition-colors"> App Name </a>
				</h1>
			</div>

			<div class="flex items-center space-x-2">
				<a
					href="/"
					class="hover:bg-accent hover:text-accent-foreground inline-flex h-9 w-9 items-center justify-center rounded-md text-sm font-medium whitespace-nowrap transition-colors"
					title="Home"
				>
					<Home class="h-4 w-4" />
				</a>
				<a
					href="/about"
					class="hover:bg-accent hover:text-accent-foreground inline-flex h-9 w-9 items-center justify-center rounded-md text-sm font-medium whitespace-nowrap transition-colors"
					title="About"
				>
					<Info class="h-4 w-4" />
				</a>

				{#if authStore.isLoggedIn}
					<div class="flex items-center space-x-2 text-sm">
						<User class="h-4 w-4" />
						<span>Welcome, {authStore.user?.name || authStore.user?.email}</span>
					</div>
					<button
						onclick={handleLogout}
						class="hover:bg-accent hover:text-accent-foreground inline-flex h-9 w-9 items-center justify-center rounded-md text-sm font-medium whitespace-nowrap transition-colors"
						title="Sign Out"
					>
						<LogOut class="h-4 w-4" />
					</button>
				{:else}
					<a
						href="/login"
						class="hover:bg-accent hover:text-accent-foreground inline-flex h-9 w-9 items-center justify-center rounded-md text-sm font-medium whitespace-nowrap transition-colors"
						title="Sign In"
					>
						<LogIn class="h-4 w-4" />
					</a>
				{/if}

				<ThemeToggle />
			</div>
		</nav>
	</div>
</header>
