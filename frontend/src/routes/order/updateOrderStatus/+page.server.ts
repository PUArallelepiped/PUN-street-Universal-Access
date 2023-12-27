import { backendPath } from '$lib/components/PUA/env';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ cookies }) => {
	return {
		orderRespList: await getCart('1', '1')
	};
};

type orderRespTye = {
	store_id: number;
	cart_id: number;
	order_date: string;
	total_price: number;
	user_id: number;

	status: number;
	user_name: string;
};
async function getCart(cartId: string, storeId: string) {
	const Resp = await fetch(backendPath + '/seller/store/' + '1' + '/orders');
	if (Resp.status == 200) {
		let orderRespList = await Resp.json();

		console.log(orderRespList);
		return orderRespList as orderRespTye[];
	}
	return [];
}
