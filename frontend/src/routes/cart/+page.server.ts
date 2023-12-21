import { backendPath } from '$lib/components/PUA/env';
import type { PageServerLoad } from './$types.js';

export const load: PageServerLoad = async () => {
	return {
		cartInfos: await getCart()
	};
};
type product_order = {
	event_discount_max_quantity: number;
	event_discount_id: number;
	product_id: number;
	product_price: number;
	product_name: string;
	product_quantity: number;
	product_picture: string;
};
type shipping_discount = {
	discount_name: string;
	discount_description: string;
	discount_max_price: number;
	discount_id: number;
	status: number;
};
type seasoning_discount = {
	discount_start_date: string;
	discount_name: string;
	discount_end_date: string;
	discount_description: string;
	discount_percentage: number;
	discount_id: number;
	status: number;
};
type storeOrderInfo = {
	store_id: 1;
	store_name: string;
	store_shipping_fee: number;
	product_order: product_order[];
	shipping_discount: shipping_discount;
	shipping_discount_bool: true;
	seasoning_discount: seasoning_discount;
	seasoning_discount_bool: true;
	total_price: number;
};
type cartInfo = {
	real_total_price: number;
	store_order_info_array: storeOrderInfo[];
};
async function getCart() {
	try {
		const resp = await fetch(backendPath + '/customer/1/carts');
		if (resp.ok) {
			return (await resp.json()) as cartInfo;
		}
	} catch {
		return undefined;
	}
}
