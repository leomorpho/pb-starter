<script lang="ts">
	import { Label } from '$lib/components/ui/label/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Edit3 } from 'lucide-svelte';
	import { pb } from '$lib/pocketbase.js';
	import { authStore } from '$lib/stores/authClient.svelte.js';

	// Personal account editing state
	let isEditingProfile = $state(false);
	let editPersonalName = $state(authStore.user?.name || '');
	let editPersonalEmail = $state(authStore.user?.email || '');
	let isSavingPersonal = $state(false);
	let personalEditError = $state<string | null>(null);

	// Update edit values when user data changes
	$effect(() => {
		if (authStore.user) {
			editPersonalName = authStore.user.name || '';
			editPersonalEmail = authStore.user.email || '';
		}
	});

	// Handle personal account save
	async function handleSavePersonal() {
		if (!editPersonalName.trim()) {
			personalEditError = 'Name is required';
			return;
		}

		if (!editPersonalEmail.trim()) {
			personalEditError = 'Email is required';
			return;
		}

		isSavingPersonal = true;
		personalEditError = null;

		try {
			// Update user data
			const updatedUser = await pb.collection('users').update(authStore.user!.id, {
				name: editPersonalName.trim(),
				email: editPersonalEmail.trim()
			});

			// Update auth store
			authStore.syncState();
			isEditingProfile = false;
		} catch (error: any) {
			console.error('Failed to update profile:', error);
			personalEditError = error.message || 'Failed to update profile';
		} finally {
			isSavingPersonal = false;
		}
	}

	// Handle cancel personal edit
	function handleCancelPersonalEdit() {
		isEditingProfile = false;
		editPersonalName = authStore.user?.name || '';
		editPersonalEmail = authStore.user?.email || '';
		personalEditError = null;
	}
</script>

<div class="bg-card rounded-xl border border-border p-6 shadow-sm">
	<div class="flex items-center justify-between mb-4">
		<h3 class="text-lg font-semibold text-foreground">Personal Account</h3>
		<button
			onclick={() => isEditingProfile = !isEditingProfile}
			class="text-sm text-muted-foreground hover:text-foreground flex items-center gap-1"
			disabled={isSavingPersonal}
		>
			<Edit3 class="w-4 h-4" />
			Edit
		</button>
	</div>

	{#if isEditingProfile}
		<!-- Edit Form -->
		<form onsubmit={(e) => { e.preventDefault(); handleSavePersonal(); }} class="space-y-4">
			{#if personalEditError}
				<div class="p-3 bg-red-50 dark:bg-red-950/50 border border-red-200 dark:border-red-800 rounded-lg">
					<p class="text-sm text-red-600 dark:text-red-400">{personalEditError}</p>
				</div>
			{/if}

			<div class="space-y-2">
				<Label for="edit-personal-name">Name</Label>
				<Input
					id="edit-personal-name"
					type="text"
					bind:value={editPersonalName}
					disabled={isSavingPersonal}
					required
				/>
			</div>

			<div class="space-y-2">
				<Label for="edit-personal-email">Email</Label>
				<Input
					id="edit-personal-email"
					type="email"
					bind:value={editPersonalEmail}
					disabled={isSavingPersonal}
					required
				/>
			</div>

			<div class="flex gap-2">
				<Button
					type="submit"
					size="sm"
					disabled={isSavingPersonal}
				>
					{isSavingPersonal ? 'Saving...' : 'Save Changes'}
				</Button>
				<Button
					type="button"
					variant="outline"
					size="sm"
					onclick={handleCancelPersonalEdit}
					disabled={isSavingPersonal}
				>
					Cancel
				</Button>
			</div>
		</form>
	{:else}
		<!-- Display Mode -->
		<div class="grid gap-4 sm:grid-cols-2">
			<div class="p-4 bg-muted/50 rounded-lg">
				<h4 class="text-sm font-medium text-muted-foreground uppercase tracking-wide mb-1">Email</h4>
				<p class="text-foreground">{authStore.user?.email}</p>
			</div>
			
			<div class="p-4 bg-muted/50 rounded-lg">
				<h4 class="text-sm font-medium text-muted-foreground uppercase tracking-wide mb-1">Name</h4>
				<p class="text-foreground">{authStore.user?.name || 'Not set'}</p>
			</div>
		</div>
	{/if}
</div>