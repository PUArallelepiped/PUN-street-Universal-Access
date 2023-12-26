import { getIdByToken } from '$lib/components/PUA/getId';
import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params, cookies }) => {
	try {
		const jwttoken: string = cookies.get('jwttoken') || '';
		await getIdByToken(jwttoken);
		return {
			shop: params.shop,
			item: params.item
		};
	} catch (e) {
		throw redirect(307, '/login');
	}
};
