import { getIdByToken } from '$lib/components/PUA/getId';
import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ cookies }) => {
	try {
		const jwttoken: string = cookies.get('jwttoken') || '';
		const user_id = await getIdByToken(jwttoken);
		if (user_id != '1') {
			throw new Error('Not admin');
		}
		return {};
	} catch (e) {
		throw redirect(307, '/');
	}
};
