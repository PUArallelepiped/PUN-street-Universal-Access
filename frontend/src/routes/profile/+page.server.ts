import { backendPath } from '$lib/components/PUA/env.js';
import { getIdByToken } from '$lib/components/PUA/getId.js';
import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types.js';

export const actions = {
	monthStatistic: async ({ request, cookies }) => {
		try {
			const jwttoken: string = cookies.get('jwttoken') || '';
			const user_id = await getIdByToken(jwttoken);
			const data = await request.formData();

			const resp = await fetch(
				backendPath +
					'/store/' +
					user_id +
					'/get-selling/' +
					data.get('year') +
					'/' +
					data.get('month')
			);
			return await resp.json();
		} catch (e) {
			throw redirect(307, '/login');
		}
	},
	yearStatistic: async ({ request, cookies }) => {
		try {
			const jwttoken: string = cookies.get('jwttoken') || '';
			const user_id = await getIdByToken(jwttoken);
			const data = await request.formData();

			const resp = await fetch(
				backendPath + '/store/' + user_id + '/get-statistics/' + data.get('year')
			);
			return await resp.json();
		} catch (e) {
			throw redirect(307, '/login');
		}
	}
};

type userShort = {
	user_id: number;
	user_name: string;
	user_email: string;
	address: string;
	phone: string;
	birthday: string;
	password: string;
	authority: string;
};

export const load: PageServerLoad = async ({ cookies }) => {
	try {
		const jwttoken: string = cookies.get('jwttoken') || '';
		const user_id = await getIdByToken(jwttoken);
		console.log('user_id: ' + user_id);
		if (user_id == '1') {
			throw 'Admin';
		}
		return {
			userInfo: await getUserInfo(user_id)
		};
	} catch (e) {
		console.log(e);
		if (e == 'Admin') {
			throw redirect(307, '/admin');
		} else {
			throw redirect(307, '/login');
		}
	}
};

async function getUserInfo(id: string) {
	const resp = await fetch(backendPath + '/user/get-info/' + id);
	const json = await resp.json();
	return json as userShort;
}
