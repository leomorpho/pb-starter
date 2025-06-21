<script lang="ts">
	import { authStore } from '$lib/stores/auth.svelte.js';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';

	// Redirect to login if not authenticated
	onMount(() => {
		if (!authStore.isLoggedIn) {
			goto('/login');
		}
	});
</script>

<svelte:head>
	<title>Dashboard - App Name</title>
	<meta name="description" content="User dashboard" />
</svelte:head>

{#if authStore.isLoggedIn}
	<div class="container mx-auto px-4 py-8">
		<div class="mx-auto max-w-4xl">
			<div class="rounded-lg border bg-white p-6 shadow-sm">
				<h1 class="mb-6 text-3xl font-bold">Dashboard</h1>

				<div class="grid gap-6 md:grid-cols-2">
					<div class="rounded-lg bg-blue-50 p-4">
						<h2 class="mb-2 text-lg font-semibold text-blue-900">Welcome back!</h2>
						<p class="text-blue-700">
							Hello, {authStore.user?.name || authStore.user?.email}! You're successfully logged in
							to your account.
						</p>
					</div>

					<div class="rounded-lg bg-green-50 p-4">
						<h2 class="mb-2 text-lg font-semibold text-green-900">User Information</h2>
						<div class="space-y-2 text-green-700">
							<p><strong>Email:</strong> {authStore.user?.email}</p>
							{#if authStore.user?.name}
								<p><strong>Name:</strong> {authStore.user.name}</p>
							{/if}
							<p><strong>ID:</strong> {authStore.user?.id}</p>
						</div>
					</div>
				</div>

				<div class="mt-6 rounded-lg bg-gray-50 p-4">
					<h2 class="mb-2 text-lg font-semibold">Protected Content</h2>
					<p class="text-gray-700">
						This is a protected area that's only accessible to authenticated users. You can add your
						app's main functionality here.
					</p>
				</div>
			</div>
		</div>
	</div>
{:else}
	<div class="container mx-auto px-4 py-8">
		<div class="text-center">
			<p>Redirecting to login...</p>
		</div>
	</div>
{/if}
