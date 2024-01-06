import { goto } from '$app/navigation';

export function Logout() {
	document.cookie = 'jwttoken=;' + 'expires=Thu, 04 Jun 1989 00:00:00 UTC; path=/;';
	goto('/login');
}
