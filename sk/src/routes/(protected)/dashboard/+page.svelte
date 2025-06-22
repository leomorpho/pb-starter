<script lang="ts">
	import { authStore } from '$lib/stores/authClient.svelte.js';
	import { config } from '$lib/config.js';
</script>

<svelte:head>
	<title>Dashboard - {config.app.name}</title>
	<meta name="description" content="User dashboard" />
</svelte:head>

<div class="container mx-auto px-4 py-8">
	<div class="mx-auto max-w-4xl">
		<div class="rounded-lg border border-border bg-background p-6 shadow-sm">
			<h1 class="mb-6 text-3xl font-bold text-foreground">Dashboard</h1>

			<div class="grid gap-6 md:grid-cols-2">
				<div class="rounded-lg bg-blue-50 dark:bg-blue-950 p-4 border border-blue-200 dark:border-blue-800">
					<h2 class="mb-2 text-lg font-semibold text-blue-900 dark:text-blue-100">Welcome back!</h2>
					<p class="text-blue-700 dark:text-blue-300">
						Hello, {authStore.user?.name || authStore.user?.email}! You're successfully logged in
						to your account.
					</p>
				</div>

				<div class="rounded-lg bg-green-50 dark:bg-green-950 p-4 border border-green-200 dark:border-green-800">
					<h2 class="mb-2 text-lg font-semibold text-green-900 dark:text-green-100">User Information</h2>
					<div class="space-y-2 text-green-700 dark:text-green-300">
						<p><strong>Email:</strong> {authStore.user?.email}</p>
						{#if authStore.user?.name}
							<p><strong>Name:</strong> {authStore.user.name}</p>
						{/if}
						<p><strong>ID:</strong> {authStore.user?.id}</p>
					</div>
				</div>
			</div>

			<div class="mt-6 rounded-lg bg-muted p-4 border border-border">
				<h2 class="mb-2 text-lg font-semibold text-foreground">Protected Content</h2>
				<p class="text-muted-foreground">
					This is a protected area that's only accessible to authenticated users. You can add your
					app's main functionality here.
				</p>
			</div>
		</div>
	</div>
</div>
