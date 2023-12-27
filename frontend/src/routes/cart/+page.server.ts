import type { cartInfo } from '$lib';
import { backendPath } from '$lib/components/PUA/env';
import type { PageServerLoad } from './$types.js';

export const load: PageServerLoad = async () => {
	return {
		cartInfos: await getCart()
	};
};
async function getCart() {
	try {
		const resp = await fetch(backendPath + '/customer/1/carts');
		if (resp.ok) {
			return (await resp.json()) as cartInfo;
		}
		return (await (await fetch('/cart.json')).json()) as cartInfo;
	} catch {
		return (await (await fetch('/cart.json')).json()) as cartInfo;
	}
}
