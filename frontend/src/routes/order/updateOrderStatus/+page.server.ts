import { getIdByToken } from '$lib/components/PUA/getId';
import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import { BACKEND_PATH as backendPath } from '$env/static/private';

export const load: PageServerLoad = async ({ cookies }) => {
	try {
		const jwttoken: string = cookies.get('jwttoken') || '';
		const user_id = await getIdByToken(jwttoken);
		return {
			orderRespList: await getCart(user_id)
		};
	} catch (e) {
		throw redirect(307, '/login');
	}
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

async function getCart(storeId: string) {
	const Resp = await fetch(backendPath + '/seller/store/' + storeId + '/orders');
	if (Resp.status == 200) {
		const orderRespList = await Resp.json();

		console.log(orderRespList);
		return orderRespList as orderRespTye[];
	}
	return [];
}

export const actions = {
	updateOrderStatus: async ({ request }) => {
		try {
			const data = await request.formData();
			const store_id = data.get('store_id');
			const cart_id = data.get('cart_id');
			const user_id = data.get('user_id');

			await fetch(
				backendPath +
					`/seller/update-order-status/customer/` +
					user_id +
					`/cart/` +
					cart_id +
					`/store/` +
					store_id,
				{
					method: 'PUT'
				}
			);
		} catch {
			throw redirect(307, '/login');
		}
	}
};
