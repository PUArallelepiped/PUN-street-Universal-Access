import { twpUrl } from '$lib/components/PUA/env';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ url }) => {
	const state = url.searchParams.get('state');
	const code = url.searchParams.get('code');

	return { state: state, code: code };
};
export const actions = {
	getToken: async ({ request }) => {
		const data = await request.formData();
		const code = data.get('code');
		const verifier = data.get('verifier');
		const tokenUrl = twpUrl + '/api/oauth/token';
		const resp = await fetch(tokenUrl, {
			headers: {
				Accept: 'application/json',
				'Content-Type': 'application/json'
			},
			method: 'POST',
			body: JSON.stringify({
				code: code,
				code_verifier: verifier
			})
		});
		const json = await resp.json();
		console.log(json);
		return json;
	}
};
