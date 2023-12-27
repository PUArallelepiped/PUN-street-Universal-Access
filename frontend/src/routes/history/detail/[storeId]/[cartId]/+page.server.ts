import type { storeOrderInfo } from '$lib';
import { backendPath } from '$lib/components/PUA/env';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params }) => {
	return {
		cartInfos: await getCart(params.cartId, params.storeId)
	};
};
async function getCart(cartId: string, storeId: string) {
	try {
		const resp = await fetch(
			backendPath + '/customer/1/cart/' + cartId + '/store/' + storeId + '/carts'
		);
		if (resp.ok) {
			return (await resp.json()) as storeOrderInfo;
		}
		return (await (await fetch('/updateOrderDetail.json')).json()) as storeOrderInfo;
	} catch {
		return (await (await fetch('/updateOrderDetail.json')).json()) as storeOrderInfo;
	}
}
