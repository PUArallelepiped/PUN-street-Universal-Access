import { backendPath } from '$lib/components/PUA/env';
import type { shopListResponse } from '$lib';
import type { PageServerLoad } from './$types.js';
import { redirect } from '@sveltejs/kit';
import { getIdByToken } from '$lib/components/PUA/getId.js';

export const load: PageServerLoad = async ({ cookies }) => {
	try {
		const jwttoken: string = cookies.get('jwttoken') || '';
		await getIdByToken(jwttoken);
		return {
			shopListResponses: await shopListResponses(),
			categories: await getCategory()
		};
	} catch (e) {
		throw redirect(307, '/login');
	}
};
type category = {
	category_id: number;
	category_name: string;
};
async function getCategory() {
	try {
		const resp = await fetch(backendPath + '/categories');
		if (resp.ok) {
			return (await resp.json()) as category[];
		}
		return [
			{
				category_id: 1,
				category_name: 'test'
			}
		] as category[];
	} catch {
		return [
			{
				category_id: 1,
				category_name: 'test'
			}
		] as category[];
	}
}

const shopListResponses = async () => {
	let result: shopListResponse[] = [
		{
			store_id: 1,
			shipping_fee: 100,
			address: 'null',
			rate_count: 10,
			rate: 5,
			category_array: [
				{
					category_name: 'drink',
					category_id: 1
				},
				{
					category_name: 'drink',
					category_id: 1
				}
			],
			name: 'null',
			description: 'good pasta',
			picture: 'https://i.imgur.com/1.jpg',
			status: 1
		}
	];

	try {
		const resp = await fetch(backendPath + '/stores', {
			method: 'POST',
			body: JSON.stringify({
				category_array: [],
				price_low: 0,
				price_high: 1000,
				search_string: ''
			})
		});
		if (resp.status === 200) {
			result = (await resp.json()) as shopListResponse[];
		}
		return result;
	} catch {
		return result;
	}
};

export const actions = {
	search: async ({ request, cookies }) => {
		try {
			const jwttoken: string = cookies.get('jwttoken') || '';
			await getIdByToken(jwttoken);
			const data = await request.formData();
			const start = data.get('start') as string;
			const end = data.get('end') as string;
			const startInt = Math.floor(parseFloat(start) * 1000);
			const endInt = Math.floor(parseFloat(end) * 1000);
			const checkedTagString = data.get('checkedTag') as string;
			const checkedTagStringList = checkedTagString.split(',');
			const checkedTagList: { category_name: string; category_id: number }[] = [];

			checkedTagStringList.forEach((element) => {
				const context = element.split(':');
				if (context[0] === '') {
					return;
				}
				checkedTagList.push({
					category_name: context[0],
					category_id: parseInt(context[1])
				});
			});

			const resp = await fetch(backendPath + '/stores', {
				method: 'POST',
				body: JSON.stringify({
					category_array: checkedTagList,
					price_low: startInt,
					price_high: endInt,
					search_string: data.get('searchString')
				})
			});
			return await resp.json();
		} catch (e) {
			throw redirect(307, '/login');
		}
	}
};
