<script lang="ts">
	import { authStore } from '$lib/stores/authClient.svelte.js';
	import { subscriptionStore } from '$lib/stores/subscription.svelte.js';
	import { config } from '$lib/config.js';
	import { pb } from '$lib/pocketbase.js';
	import AvatarUpload from '$lib/components/AvatarUpload.svelte';
	import { Crown, User, Mail, Calendar, Edit3, Upload } from 'lucide-svelte';
	import { onMount } from 'svelte';

	// State for editing profile
	let isEditingProfile = $state(false);
	let isEditingAvatar = $state(false);

	// Load subscription data on mount
	onMount(() => {
		subscriptionStore.loadData();
	});

	// Helper to get user avatar URL
	function getAvatarUrl(user: any): string | null {
		if (!user?.avatar) return null;
		return pb.files.getUrl(user, user.avatar, { thumb: '200x200' });
	}

	// Helper to format date
	function formatDate(dateString: string): string {
		return new Date(dateString).toLocaleDateString('en-US', {
			year: 'numeric',
			month: 'long',
			day: 'numeric'
		});
	}

	// Handle avatar upload completion
	function handleAvatarUpload() {
		// Refresh user data after avatar upload
		authStore.syncState();
		isEditingAvatar = false;
	}

	// Handle avatar upload error
	function handleAvatarError(error: Error) {
		console.error('Avatar upload error:', error);
		// You might want to show a toast notification here
	}
</script>

<svelte:head>
	<title>Dashboard - {config.app.name}</title>
	<meta name="description" content="User dashboard" />
</svelte:head>

<div class="container mx-auto px-4 py-8">
	<div class="mx-auto max-w-6xl">
		<!-- Header -->
		<div class="mb-8">
			<h1 class="text-4xl font-bold text-foreground mb-2">Dashboard</h1>
			<p class="text-muted-foreground">Welcome back to your personal space</p>
		</div>

		<div class="grid gap-8 lg:grid-cols-3">
			<!-- Profile Section -->
			<div class="lg:col-span-1">
				<div class="bg-card rounded-xl border border-border p-6 shadow-sm">
					<div class="text-center">
						<!-- Avatar Section -->
						<div class="relative mb-6">
							<div class="relative inline-block">
								{#if getAvatarUrl(authStore.user)}
									<img
										src={getAvatarUrl(authStore.user)}
										alt="Profile"
										class="w-24 h-24 rounded-full object-cover border-4 border-background shadow-lg"
									/>
								{:else}
									<div class="w-24 h-24 rounded-full bg-muted border-4 border-background shadow-lg flex items-center justify-center">
										<User class="w-8 h-8 text-muted-foreground" />
									</div>
								{/if}
								
								<!-- Edit Avatar Button -->
								<button
									onclick={() => isEditingAvatar = !isEditingAvatar}
									class="absolute -bottom-1 -right-1 w-8 h-8 bg-primary text-primary-foreground rounded-full shadow-lg hover:bg-primary/90 transition-colors flex items-center justify-center"
									title="Edit avatar"
								>
									<Edit3 class="w-4 h-4" />
								</button>
							</div>
						</div>

						<!-- Avatar Upload Component (when editing) -->
						{#if isEditingAvatar}
							<div class="mb-6 p-4 bg-muted/50 rounded-lg border border-dashed border-border">
								<AvatarUpload
									size="sm"
									onUploadComplete={handleAvatarUpload}
									onUploadError={handleAvatarError}
								/>
								<button
									onclick={() => isEditingAvatar = false}
									class="mt-3 text-sm text-muted-foreground hover:text-foreground"
								>
									Cancel
								</button>
							</div>
						{/if}

						<!-- User Name & Email -->
						<div class="space-y-2">
							<h2 class="text-xl font-semibold text-foreground">
								{authStore.user?.name || 'User'}
							</h2>
							<p class="text-muted-foreground flex items-center justify-center gap-2">
								<Mail class="w-4 h-4" />
								{authStore.user?.email}
							</p>
							{#if authStore.user?.created}
								<p class="text-sm text-muted-foreground flex items-center justify-center gap-2">
									<Calendar class="w-4 h-4" />
									Member since {formatDate(authStore.user.created)}
								</p>
							{/if}
						</div>

						<!-- Subscription Status -->
						{#if subscriptionStore.isSubscribed}
							<div class="mt-4 inline-flex items-center gap-2 px-3 py-1 bg-yellow-100 dark:bg-yellow-900/30 text-yellow-800 dark:text-yellow-200 rounded-full text-sm font-medium">
								<Crown class="w-4 h-4" />
								Premium Member
							</div>
						{:else}
							<div class="mt-4">
								<a
									href="/pricing"
									class="inline-flex items-center gap-2 px-4 py-2 bg-primary text-primary-foreground rounded-md hover:bg-primary/90 transition-colors text-sm font-medium"
								>
									<Crown class="w-4 h-4" />
									Upgrade to Premium
								</a>
							</div>
						{/if}
					</div>
				</div>
			</div>

			<!-- Main Content -->
			<div class="lg:col-span-2 space-y-6">
				<!-- Quick Actions -->
				<div class="bg-card rounded-xl border border-border p-6 shadow-sm">
					<h3 class="text-lg font-semibold text-foreground mb-4">Quick Actions</h3>
					<div class="grid gap-4 sm:grid-cols-2">
						<a
							href="/files"
							class="flex items-center gap-3 p-4 bg-blue-50 dark:bg-blue-950/50 rounded-lg border border-blue-200 dark:border-blue-800/50 hover:bg-blue-100 dark:hover:bg-blue-950/70 transition-colors group"
						>
							<div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
								<Upload class="w-5 h-5 text-white" />
							</div>
							<div>
								<h4 class="font-medium text-blue-900 dark:text-blue-100">Upload Files</h4>
								<p class="text-sm text-blue-700 dark:text-blue-300">Manage your uploads</p>
							</div>
						</a>

						{#if subscriptionStore.isSubscribed}
							<a
								href="/billing"
								class="flex items-center gap-3 p-4 bg-green-50 dark:bg-green-950/50 rounded-lg border border-green-200 dark:border-green-800/50 hover:bg-green-100 dark:hover:bg-green-950/70 transition-colors group"
							>
								<div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
									<Crown class="w-5 h-5 text-white" />
								</div>
								<div>
									<h4 class="font-medium text-green-900 dark:text-green-100">Manage Billing</h4>
									<p class="text-sm text-green-700 dark:text-green-300">View subscription</p>
								</div>
							</a>
						{:else}
							<a
								href="/pricing"
								class="flex items-center gap-3 p-4 bg-purple-50 dark:bg-purple-950/50 rounded-lg border border-purple-200 dark:border-purple-800/50 hover:bg-purple-100 dark:hover:bg-purple-950/70 transition-colors group"
							>
								<div class="w-10 h-10 bg-purple-500 rounded-lg flex items-center justify-center">
									<Crown class="w-5 h-5 text-white" />
								</div>
								<div>
									<h4 class="font-medium text-purple-900 dark:text-purple-100">Go Premium</h4>
									<p class="text-sm text-purple-700 dark:text-purple-300">Unlock all features</p>
								</div>
							</a>
						{/if}
					</div>
				</div>

				<!-- Account Overview -->
				<div class="bg-card rounded-xl border border-border p-6 shadow-sm">
					<div class="flex items-center justify-between mb-4">
						<h3 class="text-lg font-semibold text-foreground">Account Overview</h3>
						<button
							onclick={() => isEditingProfile = !isEditingProfile}
							class="text-sm text-muted-foreground hover:text-foreground flex items-center gap-1"
						>
							<Edit3 class="w-4 h-4" />
							Edit
						</button>
					</div>

					<div class="grid gap-4 sm:grid-cols-2">
						<div class="p-4 bg-muted/50 rounded-lg">
							<h4 class="text-sm font-medium text-muted-foreground uppercase tracking-wide mb-1">Email</h4>
							<p class="text-foreground">{authStore.user?.email}</p>
						</div>
						
						<div class="p-4 bg-muted/50 rounded-lg">
							<h4 class="text-sm font-medium text-muted-foreground uppercase tracking-wide mb-1">Name</h4>
							<p class="text-foreground">{authStore.user?.name || 'Not set'}</p>
						</div>
						
						<div class="p-4 bg-muted/50 rounded-lg">
							<h4 class="text-sm font-medium text-muted-foreground uppercase tracking-wide mb-1">User ID</h4>
							<p class="text-foreground font-mono text-sm">{authStore.user?.id}</p>
						</div>
						
						<div class="p-4 bg-muted/50 rounded-lg">
							<h4 class="text-sm font-medium text-muted-foreground uppercase tracking-wide mb-1">Status</h4>
							<p class="text-foreground">
								{subscriptionStore.isSubscribed ? 'Premium Member' : 'Free Account'}
							</p>
						</div>
					</div>
				</div>

				<!-- Recent Activity / Welcome Message -->
				<div class="bg-gradient-to-r from-blue-50 to-indigo-50 dark:from-blue-950/30 dark:to-indigo-950/30 rounded-xl border border-blue-200 dark:border-blue-800/50 p-6">
					<h3 class="text-lg font-semibold text-blue-900 dark:text-blue-100 mb-2">Welcome to your Dashboard!</h3>
					<p class="text-blue-700 dark:text-blue-300 mb-4">
						This is your personal space where you can manage your account, upload files, and access all premium features.
					</p>
					<div class="flex flex-wrap gap-2">
						<span class="px-3 py-1 bg-blue-100 dark:bg-blue-900/50 text-blue-800 dark:text-blue-200 rounded-full text-sm">
							✨ File Uploads
						</span>
						<span class="px-3 py-1 bg-blue-100 dark:bg-blue-900/50 text-blue-800 dark:text-blue-200 rounded-full text-sm">
							🔒 Secure Storage
						</span>
						<span class="px-3 py-1 bg-blue-100 dark:bg-blue-900/50 text-blue-800 dark:text-blue-200 rounded-full text-sm">
							⚡ Fast Performance
						</span>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>
