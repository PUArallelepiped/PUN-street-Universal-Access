import { backendPath } from '$lib/components/PUA/env.js';
import type { PageServerLoad } from './$types';

type historyInfoType = {
	customer_id: number;
	cart_id: number;
	store_id: number;
	order_date: string;
	total_price: number;
	store_name: string;
	store_rate: number;
	store_picture: string;
	status: number;
};
export const load: PageServerLoad = async () => {
	return {
		history: await getHistory()
	};
};
async function getHistory() {
	try {
		const resp = await fetch(backendPath + '/customer/1/get-history');
		return (await resp.json()) as historyInfoType[];
	} catch {
		return [] as historyInfoType[];
	}
}
