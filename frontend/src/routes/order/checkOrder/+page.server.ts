import { backendPath } from '$lib/components/PUA/env.js';
import type { PageServerLoad } from './$types';
type orderInfoType = {
	user_id: number;
	cart_id: number;
	store_id: number;
	status: number;
	store_name: string;
	store_picture: string;
};
export const load: PageServerLoad = async () => {
	return {
		orderInfoList: await getOrder()
	};
};

// todo custumer id
async function getOrder() {
	try {
		const customerId = 1;
		const resp = await fetch(backendPath + '/customer/' + customerId + '/order-status');
		return (await resp.json()) as orderInfoType[];
	} catch {
		return [];
	}
}
