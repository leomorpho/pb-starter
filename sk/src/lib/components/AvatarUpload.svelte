<script lang="ts">
	import { onMount } from 'svelte';
	import FileUpload from './FileUpload.svelte';
	import { getFileUrl, type FileUploadRecord } from '$lib/files.js';
	import { authStore } from '$lib/stores/authClient.svelte.js';
	import { pb } from '$lib/pocketbase.js';
	import { User, Camera } from 'lucide-svelte';

	// Props
	let {
		size = 'lg',
		showUploadButton = true,
		autoUpdateProfile = true,
		onUploadComplete = null,
		onUploadError = null,
		className = ''
	}: {
		size?: 'sm' | 'md' | 'lg' | 'xl';
		showUploadButton?: boolean;
		autoUpdateProfile?: boolean;
		onUploadComplete?: ((record: FileUploadRecord) => void) | null;
		onUploadError?: ((error: Error) => void) | null;
		className?: string;
	} = $props();

	// State
	let showUploadDialog = $state(false);
	let currentAvatarUrl = $state<string | null>(null);
	let isUpdatingProfile = $state(false);

	// Size configurations
	const sizeClasses = {
		sm: 'w-8 h-8',
		md: 'w-12 h-12', 
		lg: 'w-16 h-16',
		xl: 'w-24 h-24'
	};

	const iconSizes = {
		sm: 'h-3 w-3',
		md: 'h-4 w-4',
		lg: 'h-5 w-5', 
		xl: 'h-6 w-6'
	};

	// Update avatar URL when user changes
	$effect(() => {
		if (authStore.user?.avatar) {
			currentAvatarUrl = pb.files.getUrl(authStore.user, authStore.user.avatar);
		} else {
			currentAvatarUrl = null;
		}
	});

	// Handle avatar upload completion
	async function handleAvatarUpload(record: FileUploadRecord) {
		try {
			if (autoUpdateProfile && authStore.user) {
				isUpdatingProfile = true;
				
				// Update user record with new avatar
				const updatedUser = await pb.collection('users').update(authStore.user.id, {
					avatar: record.file
				});

				// Update auth store
				authStore.user = updatedUser;
				currentAvatarUrl = getFileUrl(record);
			}

			if (onUploadComplete) {
				onUploadComplete(record);
			}

			showUploadDialog = false;
		} catch (error) {
			console.error('Failed to update profile avatar:', error);
			if (onUploadError) {
				onUploadError(error instanceof Error ? error : new Error('Failed to update profile'));
			}
		} finally {
			isUpdatingProfile = false;
		}
	}

	// Handle upload error
	function handleUploadError(error: Error) {
		console.error('Avatar upload error:', error);
		if (onUploadError) {
			onUploadError(error);
		}
	}
</script>

<!-- Avatar Display -->
<div class="avatar-container {className}">
	<div class="relative inline-block">
		<!-- Avatar Image -->
		<div class="relative {sizeClasses[size]} rounded-full overflow-hidden bg-muted border-2 border-background shadow-sm">
			{#if currentAvatarUrl}
				<img 
					src={currentAvatarUrl} 
					alt="Avatar"
					class="w-full h-full object-cover"
				/>
			{:else}
				<div class="w-full h-full flex items-center justify-center bg-gradient-to-br from-primary/20 to-primary/10">
					<User class="{iconSizes[size]} text-primary" />
				</div>
			{/if}
			
			<!-- Loading overlay -->
			{#if isUpdatingProfile}
				<div class="absolute inset-0 bg-black/50 flex items-center justify-center">
					<div class="animate-spin rounded-full h-1/2 w-1/2 border-2 border-white border-t-transparent"></div>
				</div>
			{/if}
		</div>

		<!-- Upload Button -->
		{#if showUploadButton && authStore.isLoggedIn}
			<button
				onclick={() => showUploadDialog = true}
				class="absolute -bottom-1 -right-1 bg-primary text-primary-foreground rounded-full p-1.5 shadow-lg hover:bg-primary/90 transition-colors"
				title="Change avatar"
			>
				<Camera class="h-3 w-3" />
			</button>
		{/if}
	</div>
</div>

<!-- Upload Dialog -->
{#if showUploadDialog}
	<div class="fixed inset-0 bg-black/50 flex items-center justify-center z-50" onclick={(e) => e.target === e.currentTarget && (showUploadDialog = false)}>
		<div class="bg-background rounded-lg p-6 max-w-md w-full mx-4">
			<h3 class="text-lg font-semibold mb-4">Upload Avatar</h3>
			
			<FileUpload
				fileType="avatar"
				category="profile_image"
				maxFiles={1}
				allowedTypes={['image/jpeg', 'image/png', 'image/webp', 'image/gif']}
				maxFileSize={5 * 1024 * 1024}
				processAfterUpload={['resize:200x200', 'thumbnail:50x50']}
				visibility="public"
				placeholder="Upload a profile picture"
				showPreview={true}
				onUploadComplete={handleAvatarUpload}
				onUploadError={handleUploadError}
				className="mb-4"
			/>

			<div class="flex justify-end space-x-2">
				<button
					onclick={() => showUploadDialog = false}
					class="px-4 py-2 text-sm border border-muted-foreground/20 rounded-md hover:bg-muted transition-colors"
				>
					Cancel
				</button>
			</div>
		</div>
	</div>
{/if}

<style>
	.avatar-container {
		@apply inline-block;
	}
</style>