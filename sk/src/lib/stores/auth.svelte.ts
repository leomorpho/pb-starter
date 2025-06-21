import { pb } from '$lib/pocketbase.js';
import type { AuthModel } from 'pocketbase';

interface AuthStore {
	isValid: boolean;
	token: string;
	model: AuthModel | null;
}

class AuthState {
	#auth = $state<AuthStore>({
		isValid: pb.authStore.isValid,
		token: pb.authStore.token,
		model: pb.authStore.model
	});

	constructor() {
		// Listen for auth changes
		pb.authStore.onChange(() => {
			this.#auth = {
				isValid: pb.authStore.isValid,
				token: pb.authStore.token,
				model: pb.authStore.model
			};
		});
	}

	get user() {
		return this.#auth.model;
	}

	get isLoggedIn() {
		return this.#auth.isValid && this.#auth.model !== null;
	}

	get token() {
		return this.#auth.token;
	}

	async login(email: string, password: string) {
		try {
			const authData = await pb.collection('users').authWithPassword(email, password);
			return { success: true, user: authData.record };
		} catch (error) {
			console.error('Login error:', error);
			return {
				success: false,
				error: error instanceof Error ? error.message : 'Login failed'
			};
		}
	}

	async signup(email: string, password: string, passwordConfirm: string, name?: string) {
		try {
			const data = {
				email,
				password,
				passwordConfirm,
				...(name && { name })
			};

			await pb.collection('users').create(data);

			// Auto-login after signup
			const authData = await pb.collection('users').authWithPassword(email, password);

			return { success: true, user: authData.record };
		} catch (error) {
			console.error('Signup error:', error);
			return {
				success: false,
				error: error instanceof Error ? error.message : 'Signup failed'
			};
		}
	}

	logout() {
		pb.authStore.clear();
	}

	async refreshAuth() {
		try {
			if (pb.authStore.isValid) {
				await pb.collection('users').authRefresh();
				return true;
			}
		} catch (error) {
			console.error('Auth refresh failed:', error);
			this.logout();
		}
		return false;
	}
}

export const authStore = new AuthState();
