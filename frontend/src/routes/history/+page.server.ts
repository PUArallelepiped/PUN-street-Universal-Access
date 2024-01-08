import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import { getIdByToken } from '$lib/components/PUA/getId';
import { BACKEND_PATH as backendPath } from '$env/static/private';

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
export const load: PageServerLoad = async ({ cookies }) => {
	try {
		const jwttoken: string = cookies.get('jwttoken') || '';
		const user_id = await getIdByToken(jwttoken);
		return {
			history: await getHistory(user_id)
		};
	} catch (e) {
		throw redirect(307, '/login');
	}
};
async function getHistory(id: string) {
	try {
		const resp = await fetch(backendPath + '/customer/' + id + '/get-history');
		return (await resp.json()) as historyInfoType[];
	} catch {
		return [] as historyInfoType[];
	}
}
