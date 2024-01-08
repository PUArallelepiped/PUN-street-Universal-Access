import type { cartInfo } from '$lib';
import { getIdByToken } from '$lib/components/PUA/getId.js';
import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types.js';
import { BACKEND_PATH as backendPath } from '$env/static/private';

export const load: PageServerLoad = async ({ cookies }) => {
	try {
		const jwttoken: string = cookies.get('jwttoken') || '';
		const user_id = await getIdByToken(jwttoken);
		return {
			cartInfos: await getCart(user_id)
		};
	} catch (e) {
		throw redirect(307, '/login');
	}
};
async function getCart(id: string) {
	try {
		const resp = await fetch(backendPath + '/customer/' + id + '/carts');
		if (resp.ok) {
			return (await resp.json()) as cartInfo;
		}
		return (await (await fetch('/cart.json')).json()) as cartInfo;
	} catch {
		return (await (await fetch('/cart.json')).json()) as cartInfo;
	}
}
