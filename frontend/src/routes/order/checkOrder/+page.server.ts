import type { PageServerLoad } from './$types';
import { redirect } from '@sveltejs/kit';
import { getIdByToken } from '$lib/components/PUA/getId';
import { BACKEND_PATH as backendPath } from '$env/static/private';
type orderInfoType = {
	user_id: number;
	cart_id: number;
	store_id: number;
	status: number;
	store_name: string;
	store_picture: string;
};
export const load: PageServerLoad = async ({ cookies }) => {
	try {
		const jwttoken: string = cookies.get('jwttoken') || '';
		const user_id = await getIdByToken(jwttoken);
		return {
			orderInfoList: await getOrder(user_id)
		};
	} catch (e) {
		throw redirect(307, '/login');
	}
};

// todo custumer id
async function getOrder(id: string) {
	try {
		const resp = await fetch(backendPath + '/customer/' + id + '/order-status');
		return (await resp.json()) as orderInfoType[];
	} catch {
		return [];
	}
}
