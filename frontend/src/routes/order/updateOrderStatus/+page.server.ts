import { backendPath } from '$lib/components/PUA/env';
import { getIdByToken } from '$lib/components/PUA/getId';
import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ cookies }) => {
	try{
		const jwttoken: string = cookies.get('jwttoken') || '';
		const user_id = await getIdByToken(jwttoken);
		return {
			orderRespList: await getCart(user_id)
		};
	}
	catch(e){
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
		let orderRespList = await Resp.json();

		console.log(orderRespList);
		return orderRespList as orderRespTye[];
	}
	return [];
}
