import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import { getIdByToken } from '$lib/components/PUA/getId';

export const load: PageServerLoad = async ({ params, cookies }) => {
	try {
		const jwttoken: string = cookies.get('jwttoken') || '';
		const user_id =  await getIdByToken(jwttoken);
		if (user_id != params.shop) {
			throw "not authorized"
		}
		return {
			shop: params.shop
		};
	}
	catch (e) {
		throw redirect(307, '/login');
	}
};
