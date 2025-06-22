<script lang="ts">
	import { subscriptionStore } from '$lib/stores/subscription.svelte.js';
	import { authStore } from '$lib/stores/authClient.svelte.js';
	import { createCheckoutSession, formatPrice, formatInterval } from '$lib/stripe.js';
	import { config } from '$lib/config.js';
	import { Check, Loader2 } from 'lucide-svelte';
	import { onMount } from 'svelte';
	
	let isLoading = $state(false);
	let checkoutLoading = $state<string | null>(null);

	onMount(() => {
		subscriptionStore.refresh();
	});

	async function handleSubscribe(priceId: string) {
		if (!authStore.isLoggedIn) {
			// Redirect to login
			window.location.href = '/login?redirect=/pricing';
			return;
		}

		checkoutLoading = priceId;
		try {
			await createCheckoutSession(priceId, 'subscription');
		} catch (error) {
			console.error('Error creating checkout session:', error);
			alert('Failed to start checkout. Please try again.');
		} finally {
			checkoutLoading = null;
		}
	}

	// Group prices by product
	function getProductsWithPrices() {
		return subscriptionStore.products.map(product => ({
			...product,
			prices: subscriptionStore.getPricesForProduct(product.product_id)
				.filter(price => price.type === 'recurring')
				.sort((a, b) => a.unit_amount - b.unit_amount)
		})).filter(product => product.prices.length > 0);
	}

	function isCurrentPlan(priceId: string): boolean {
		return subscriptionStore.userSubscription?.price_id === priceId;
	}

	function getButtonText(priceId: string): string {
		if (checkoutLoading === priceId) return 'Processing...';
		if (isCurrentPlan(priceId)) return 'Current Plan';
		if (!authStore.isLoggedIn) return 'Sign Up to Subscribe';
		if (subscriptionStore.isSubscribed) return 'Switch Plan';
		return 'Subscribe';
	}

	function isButtonDisabled(priceId: string): boolean {
		return checkoutLoading !== null || isCurrentPlan(priceId);
	}
</script>

<svelte:head>
	<title>Pricing - {config.app.name}</title>
	<meta name="description" content="Choose the perfect plan for your needs" />
</svelte:head>

<div class="container mx-auto px-4 py-8">
	<div class="mx-auto max-w-6xl">
		<div class="text-center mb-12">
			<h1 class="text-4xl font-bold mb-4">Choose Your Plan</h1>
			<p class="text-xl text-muted-foreground">
				Select the perfect plan for your needs. Cancel or change anytime.
			</p>
		</div>

		{#if subscriptionStore.isLoading}
			<div class="text-center py-12">
				<Loader2 class="h-8 w-8 animate-spin mx-auto mb-4" />
				<p>Loading pricing plans...</p>
			</div>
		{:else}
			<div class="grid gap-8 md:grid-cols-2 lg:grid-cols-3">
				{#each getProductsWithPrices() as product (product.id)}
					<div class="rounded-lg border bg-card p-6 shadow-sm">
						<div class="mb-6">
							<h3 class="text-2xl font-bold mb-2">{product.name}</h3>
							{#if product.description}
								<p class="text-muted-foreground">{product.description}</p>
							{/if}
						</div>

						<div class="space-y-4 mb-6">
							{#each product.prices as price (price.id)}
								<div class="border rounded-lg p-4">
									<div class="flex items-baseline mb-2">
										<span class="text-3xl font-bold">
											{formatPrice(price.unit_amount, price.currency)}
										</span>
										<span class="text-muted-foreground ml-2">
											{formatInterval(price.interval || 'month', price.interval_count)}
										</span>
									</div>
									
									{#if price.trial_period_days}
										<p class="text-sm text-green-600 mb-3">
											{price.trial_period_days} day free trial
										</p>
									{/if}

									<button
										onclick={() => handleSubscribe(price.price_id)}
										disabled={isButtonDisabled(price.price_id)}
										class="w-full rounded-md px-4 py-2 text-sm font-medium transition-colors
											{isCurrentPlan(price.price_id) 
												? 'bg-green-100 text-green-800 cursor-not-allowed' 
												: 'bg-primary text-primary-foreground hover:bg-primary/90 disabled:opacity-50 disabled:cursor-not-allowed'
											}"
									>
										{#if checkoutLoading === price.price_id}
											<Loader2 class="h-4 w-4 animate-spin inline mr-2" />
										{/if}
										{getButtonText(price.price_id)}
									</button>

									{#if isCurrentPlan(price.price_id)}
										<div class="flex items-center justify-center mt-2 text-green-600">
											<Check class="h-4 w-4 mr-1" />
											<span class="text-sm">Active subscription</span>
										</div>
									{/if}
								</div>
							{/each}
						</div>
					</div>
				{/each}
			</div>

			{#if getProductsWithPrices().length === 0}
				<div class="text-center py-12">
					<p class="text-muted-foreground">No pricing plans available at the moment.</p>
					<p class="text-sm text-muted-foreground mt-2">
						Please check back later or contact support.
					</p>
				</div>
			{/if}
		{/if}

		{#if subscriptionStore.isSubscribed}
			<div class="mt-12 text-center">
				<div class="rounded-lg bg-green-50 border border-green-200 p-6 inline-block">
					<h3 class="text-lg font-semibold text-green-800 mb-2">You're subscribed!</h3>
					<p class="text-green-700 mb-4">
						Manage your subscription, update payment methods, and view billing history.
					</p>
					<a 
						href="/billing" 
						class="inline-flex items-center rounded-md bg-green-600 px-4 py-2 text-sm font-medium text-white hover:bg-green-700 transition-colors"
					>
						Manage Subscription
					</a>
				</div>
			</div>
		{/if}

		<div class="mt-16 text-center">
			<h3 class="text-lg font-semibold mb-4">Questions?</h3>
			<p class="text-muted-foreground">
				Need help choosing the right plan? 
				<a href="/about" class="text-primary hover:underline">Contact us</a>
				and we'll help you find the perfect fit.
			</p>
		</div>
	</div>
</div>