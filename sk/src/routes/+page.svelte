<script lang="ts">
	import { onMount } from 'svelte';
	import { gsap } from 'gsap';
	import { ScrollTrigger } from 'gsap/dist/ScrollTrigger';
	import { CustomEase } from 'gsap/dist/CustomEase';
	import { Card, CardContent } from '$lib/components/ui/card';
	import { Button } from '$lib/components/ui/button';
	import { Badge } from '$lib/components/ui/badge';
	import { config } from '$lib/config.js';
	import { authStore } from '$lib/stores/authClient.svelte.js';
	import { subscriptionStore } from '$lib/stores/subscription.svelte.js';
	import { Zap, Shield, Users, Rocket, Star, CheckCircle } from 'lucide-svelte';

	gsap.registerPlugin(ScrollTrigger, CustomEase);

	// Initialize subscription store in component context
	subscriptionStore.initialize();

	const features = [
		{
			icon: Rocket,
			title: 'Lightning Fast',
			subtitle: 'Built for Speed',
			description: 'Optimized performance with modern architecture and static site generation for blazing fast load times.',
			benefits: ['Sub-second page loads', 'Global CDN distribution', 'Optimized images & assets'],
			gradient: 'from-blue-500 to-purple-600'
		},
		{
			icon: Shield,
			title: 'Enterprise Security',
			subtitle: 'Bank-Grade Protection',
			description: 'Advanced security features with encrypted data, secure authentication, and compliance standards.',
			benefits: ['End-to-end encryption', 'Multi-factor authentication', 'GDPR compliant'],
			gradient: 'from-green-500 to-teal-600'
		},
		{
			icon: Users,
			title: 'Team Collaboration',
			subtitle: 'Work Together',
			description: 'Powerful collaboration tools that help teams stay synchronized and productive.',
			benefits: ['Real-time updates', 'Team workspaces', 'Activity tracking'],
			gradient: 'from-pink-500 to-rose-600'
		},
		{
			icon: Zap,
			title: 'Smart Automation',
			subtitle: 'AI-Powered',
			description: 'Intelligent automation that learns from your patterns and optimizes workflows automatically.',
			benefits: ['Smart suggestions', 'Automated workflows', 'Predictive analytics'],
			gradient: 'from-yellow-500 to-orange-600'
		}
	];

	const stats = [
		{ number: '99.9%', label: 'Uptime Guarantee' },
		{ number: '10x', label: 'Faster Development' },
		{ number: '500K+', label: 'Active Users' },
		{ number: '24/7', label: 'Support Available' }
	];

	onMount(() => {
		// Hero animations
		gsap
			.timeline()
			.from('.hero-badge', { scale: 0, opacity: 0, duration: 1, ease: 'back.out(2)' })
			.from('.hero-title', { y: 100, opacity: 0, duration: 1.2, ease: 'power3.out' }, '-=0.5')
			.from('.hero-subtitle', { y: 50, opacity: 0, duration: 0.8 }, '-=0.6')
			.from('.hero-desc', { opacity: 0, duration: 0.6 }, '-=0.4')
			.from('.hero-buttons', { y: 30, opacity: 0, duration: 0.6 }, '-=0.2')
			.from('.hero-stats', { y: 40, opacity: 0, duration: 0.8, stagger: 0.1 }, '-=0.4');

		// Animate floating blobs
		gsap.to('.blob-1', {
			x: 'random(-100, 100)',
			y: 'random(-100, 100)',
			rotation: 'random(-45, 45)',
			duration: 'random(8, 12)',
			repeat: -1,
			yoyo: true,
			ease: 'sine.inOut'
		});

		gsap.to('.blob-2', {
			x: 'random(-150, 150)',
			y: 'random(-80, 80)',
			rotation: 'random(-45, 45)',
			duration: 'random(10, 14)',
			repeat: -1,
			yoyo: true,
			ease: 'sine.inOut',
			delay: 2
		});

		gsap.to('.blob-3', {
			x: 'random(-80, 80)',
			y: 'random(-120, 120)',
			rotation: 'random(-45, 45)',
			duration: 'random(6, 10)',
			repeat: -1,
			yoyo: true,
			ease: 'sine.inOut',
			delay: 4
		});

		// Scroll animations
		gsap.utils.toArray('.fade-up').forEach((element: any) => {
			gsap.from(element, {
				y: 60,
				opacity: 0,
				duration: 0.8,
				scrollTrigger: {
					trigger: element,
					start: 'top 85%'
				}
			});
		});

		// Feature cards stagger
		gsap.utils.toArray('.feature-card').forEach((card: any, i: number) => {
			gsap.from(card, {
				y: 100,
				opacity: 0,
				duration: 1,
				delay: i * 0.2,
				scrollTrigger: {
					trigger: card,
					start: 'top 80%'
				}
			});
		});

		// Floating animation for feature icons
		gsap.utils.toArray('.feature-icon').forEach((icon: any) => {
			gsap.to(icon, {
				y: -10,
				duration: 2,
				repeat: -1,
				yoyo: true,
				ease: 'sine.inOut',
				delay: gsap.utils.random(0, 2)
			});
		});
	});
</script>

<svelte:head>
	<title>{config.app.name} • Modern SaaS Platform</title>
	<meta
		name="description"
		content="Transform your workflow with our modern, secure, and lightning-fast platform. Built for teams that demand excellence."
	/>
</svelte:head>

<!-- Hero Section -->
<section class="relative min-h-screen flex items-center overflow-hidden bg-background">
	<!-- Background Effects -->
	<div class="absolute inset-0 overflow-hidden">
		<div class="blob-1 absolute top-20 left-20 w-96 h-96 bg-gradient-to-r from-blue-400/20 to-purple-600/20 rounded-full blur-3xl"></div>
		<div class="blob-2 absolute bottom-20 right-20 w-80 h-80 bg-gradient-to-r from-pink-400/20 to-rose-600/20 rounded-full blur-3xl"></div>
		<div class="blob-3 absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-[600px] h-[600px] bg-gradient-to-r from-green-400/10 to-teal-600/10 rounded-full blur-3xl"></div>
		
		<!-- Grid Pattern -->
		<div class="absolute inset-0 bg-[linear-gradient(to_right,#80808012_1px,transparent_1px),linear-gradient(to_bottom,#80808012_1px,transparent_1px)] bg-[size:24px_24px]"></div>
	</div>

	<div class="relative max-w-7xl mx-auto px-6 grid lg:grid-cols-12 gap-12 items-center">
		<!-- Left Side -->
		<div class="lg:col-span-7 space-y-8">
			<div class="hero-badge">
				<Badge class="bg-gradient-to-r from-blue-500 to-purple-600 text-white border-0">
					{#if subscriptionStore.isSubscribed}
						✨ Premium User
					{:else}
						🚀 Now Available
					{/if}
				</Badge>
			</div>

			<div class="space-y-6">
				<h1 class="hero-title text-6xl lg:text-8xl xl:text-9xl font-black leading-none">
					<span class="block text-foreground">MODERN</span>
					<span class="block bg-gradient-to-r from-blue-600 via-purple-600 to-pink-600 bg-clip-text text-transparent">
						SAAS
					</span>
					<span class="block text-foreground">PLATFORM</span>
				</h1>

				<p class="hero-subtitle text-xl lg:text-2xl text-muted-foreground font-light max-w-2xl">
					Transform your workflow with our secure, lightning-fast platform. Built for teams that demand excellence and innovation.
				</p>
			</div>

			<div class="hero-desc">
				<p class="text-muted-foreground/80 max-w-xl text-sm leading-relaxed">
					Join thousands of teams using our platform to streamline operations, boost productivity, and scale their business with confidence.
				</p>
			</div>

			<div class="hero-buttons flex flex-wrap gap-4">
				{#if authStore.isLoggedIn}
					<Button href="/dashboard" size="lg" class="bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700">
						Go to Dashboard
					</Button>
				{:else}
					<Button href="/signup" size="lg" class="bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700">
						Get Started Free
					</Button>
				{/if}
				<Button href="/pricing" variant="outline" size="lg">
					View Pricing
				</Button>
			</div>

			<!-- Stats -->
			<div class="hero-stats grid grid-cols-2 lg:grid-cols-4 gap-6 pt-8 border-t border-border/50">
				{#each stats as stat}
					<div class="text-center lg:text-left">
						<div class="text-2xl lg:text-3xl font-bold bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent">
							{stat.number}
						</div>
						<div class="text-sm text-muted-foreground">{stat.label}</div>
					</div>
				{/each}
			</div>
		</div>

		<!-- Right Side - Visual Element -->
		<div class="hidden lg:block lg:col-span-5">
			<div class="relative">
				<div class="absolute inset-0 bg-gradient-to-r from-blue-500/20 to-purple-600/20 rounded-3xl transform rotate-6"></div>
				<div class="relative bg-background/90 backdrop-blur-xl border border-border/50 rounded-3xl p-8 transform -rotate-3 hover:rotate-0 transition-transform duration-700">
					<div class="space-y-4">
						<div class="flex items-center space-x-3">
							<div class="w-3 h-3 bg-red-500 rounded-full"></div>
							<div class="w-3 h-3 bg-yellow-500 rounded-full"></div>
							<div class="w-3 h-3 bg-green-500 rounded-full"></div>
						</div>
						<div class="space-y-3">
							<div class="h-4 bg-gradient-to-r from-blue-500 to-purple-600 rounded w-3/4"></div>
							<div class="h-3 bg-muted rounded w-1/2"></div>
							<div class="h-3 bg-muted rounded w-2/3"></div>
							<div class="grid grid-cols-2 gap-3 mt-6">
								<div class="h-20 bg-gradient-to-br from-green-400 to-teal-600 rounded-lg"></div>
								<div class="h-20 bg-gradient-to-br from-pink-400 to-rose-600 rounded-lg"></div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</section>

<!-- Features Section -->
<section class="py-32 bg-background">
	<div class="max-w-7xl mx-auto px-6">
		<div class="text-center mb-20">
			<div class="fade-up space-y-4">
				<Badge variant="secondary">Core Features</Badge>
				<h2 class="text-4xl lg:text-6xl font-black text-foreground">
					EVERYTHING YOU<br />
					<span class="bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent">NEED TO SUCCEED</span>
				</h2>
				<p class="text-muted-foreground text-lg max-w-2xl mx-auto">
					Powerful features designed to streamline your workflow and accelerate growth
				</p>
			</div>
		</div>

		<div class="grid md:grid-cols-2 gap-8">
			{#each features as feature, i}
				<Card class="feature-card group hover:shadow-2xl transition-all duration-500 border-0 bg-gradient-to-br from-background to-muted/30">
					<CardContent class="p-8">
						<div class="space-y-6">
							<!-- Icon -->
							<div class="feature-icon relative">
								<div class="w-16 h-16 bg-gradient-to-r {feature.gradient} rounded-2xl flex items-center justify-center transform group-hover:scale-110 transition-transform duration-300">
									<svelte:component this={feature.icon} class="h-8 w-8 text-white" />
								</div>
								<div class="absolute -top-1 -right-1 w-6 h-6 bg-gradient-to-r {feature.gradient} rounded-full opacity-60 animate-pulse"></div>
							</div>

							<!-- Content -->
							<div class="space-y-4">
								<div>
									<h3 class="text-2xl font-bold text-foreground group-hover:bg-gradient-to-r group-hover:{feature.gradient} group-hover:bg-clip-text group-hover:text-transparent transition-all duration-300">
										{feature.title}
									</h3>
									<p class="text-sm font-medium bg-gradient-to-r {feature.gradient} bg-clip-text text-transparent">
										{feature.subtitle}
									</p>
								</div>
								<p class="text-muted-foreground leading-relaxed">
									{feature.description}
								</p>
							</div>

							<!-- Benefits -->
							<div class="space-y-2">
								{#each feature.benefits as benefit}
									<div class="flex items-center space-x-3">
										<CheckCircle class="h-4 w-4 text-green-500" />
										<span class="text-sm text-muted-foreground">{benefit}</span>
									</div>
								{/each}
							</div>
						</div>
					</CardContent>
				</Card>
			{/each}
		</div>
	</div>
</section>

<!-- Social Proof Section -->
<section class="py-32 bg-muted/20">
	<div class="max-w-7xl mx-auto px-6">
		<div class="fade-up text-center mb-16">
			<Badge variant="secondary">Trusted by Teams</Badge>
			<h2 class="text-4xl lg:text-5xl font-black text-foreground mt-8 mb-6">
				JOIN THOUSANDS OF<br />
				<span class="bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent">SATISFIED USERS</span>
			</h2>
		</div>

		<div class="grid md:grid-cols-3 gap-8">
			{#each Array(3) as _, i}
				<Card class="fade-up group hover:shadow-xl transition-all duration-300" style="animation-delay: {i * 0.2}s">
					<CardContent class="p-6">
						<div class="space-y-4">
							<div class="flex space-x-1">
								{#each Array(5) as _}
									<Star class="h-4 w-4 fill-yellow-400 text-yellow-400" />
								{/each}
							</div>
							<p class="text-muted-foreground italic">
								"This platform has completely transformed how our team collaborates. The speed and reliability are unmatched."
							</p>
							<div class="flex items-center space-x-3">
								<div class="w-10 h-10 bg-gradient-to-r from-blue-500 to-purple-600 rounded-full flex items-center justify-center">
									<span class="text-white font-bold text-sm">{String.fromCharCode(65 + i)}</span>
								</div>
								<div>
									<div class="font-semibold">Team Lead</div>
									<div class="text-sm text-muted-foreground">Growing Startup</div>
								</div>
							</div>
						</div>
					</CardContent>
				</Card>
			{/each}
		</div>
	</div>
</section>

<!-- CTA Section -->
<section class="py-32 bg-background">
	<div class="max-w-5xl mx-auto px-6 text-center">
		<div class="fade-up space-y-12">
			<Badge variant="secondary">Ready to Start?</Badge>

			<h2 class="text-5xl lg:text-7xl font-black text-foreground">
				START YOUR<br />
				<span class="bg-gradient-to-r from-blue-600 via-purple-600 to-pink-600 bg-clip-text text-transparent">
					JOURNEY TODAY
				</span>
			</h2>

			<p class="text-muted-foreground text-lg max-w-2xl mx-auto">
				Join thousands of teams already using our platform to streamline their workflow and accelerate growth.
			</p>

			<div class="flex flex-col sm:flex-row gap-6 justify-center">
				{#if authStore.isLoggedIn}
					{#if subscriptionStore.isSubscribed}
						<Button href="/dashboard" size="lg" class="bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700">
							Go to Dashboard
						</Button>
						<Button href="/premium" variant="outline" size="lg">
							Explore Premium
						</Button>
					{:else}
						<Button href="/pricing" size="lg" class="bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700">
							Upgrade Now
						</Button>
						<Button href="/dashboard" variant="outline" size="lg">
							Go to Dashboard
						</Button>
					{/if}
				{:else}
					<Button href="/signup" size="lg" class="bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700">
						Get Started Free
					</Button>
					<Button href="/pricing" variant="outline" size="lg">
						View Pricing
					</Button>
				{/if}
			</div>

			<!-- Final Stats -->
			<div class="grid grid-cols-2 md:grid-cols-4 gap-8 pt-16 border-t border-border/50">
				{#each stats as stat}
					<div>
						<div class="text-3xl font-bold bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent">
							{stat.number}
						</div>
						<div class="text-sm text-muted-foreground mt-1">{stat.label}</div>
					</div>
				{/each}
			</div>
		</div>
	</div>
</section>

<style>
	:global(.fade-up) {
		opacity: 0;
	}
	
	:global(.feature-card) {
		opacity: 0;
	}
	
	:global(.hero-badge),
	:global(.hero-title),
	:global(.hero-subtitle),
	:global(.hero-desc),
	:global(.hero-buttons),
	:global(.hero-stats) {
		opacity: 0;
	}
</style>