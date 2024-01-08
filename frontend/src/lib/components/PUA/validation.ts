import { goto } from '$app/navigation';
import { BACKEND_PATH as backendPath } from '$env/static/private';

export async function validateToken() {
	const response = await fetch(backendPath + '/validate', {
		method: 'GET',
		credentials: 'include'
	});
	if (response.status !== 200) {
		goto('/login');
		return;
	}
}
