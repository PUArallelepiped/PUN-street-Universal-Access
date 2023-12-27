import { backendPath } from '$lib/components/PUA/env';
import { getIdByToken } from '$lib/components/PUA/getId';
import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

type shopInfoType = {
	store_id: number;
	shipping_fee: number;
	address: string;
	rate_count: number;
	rate: number;
	name: string;
	description: string;
	picture: string;
	status: number;
	category_array: { category_name: string; category_id: number }[];
};
type ProductInfoType = {
	name: string;
	description: string;
	price: number;
	picture: string;
	product_id: number;
	status: number;
	stock: number;
	store_id: number;
	event_discount_array: {
		discount_max_quantity: number;
	}[];
};
export const load: PageServerLoad = async ({ params, cookies }) => {
	try {
		const jwttoken: string = cookies.get('jwttoken') || '';
		await getIdByToken(jwttoken);
		return {
			shop: params.shop,
			shopInfo: await getShopInfo(params.shop),
			productList: await getProductList(params.shop)
		};
	} catch (e) {
		throw redirect(307, '/login');
	}
};
async function getShopInfo(shop: string) {
	try {
		const resp = await fetch(backendPath + '/store/' + shop);
		if (resp.ok) {
			return (await resp.json()) as shopInfoType;
		}
		return {
			store_id: 1,
			shipping_fee: 100,
			address: '100台灣台北市中正區八德路一段82巷9弄17號',
			rate_count: 10,
			rate: 5,
			name: 'Im pasta',
			description: 'good pasta',
			picture: 'https://i.imgur.com/1.jpg',
			status: 1
		} as shopInfoType;
	} catch {
		return {
			store_id: 1,
			shipping_fee: 100,
			address: '100台灣台北市中正區八德路一段82巷9弄17號',
			rate_count: 10,
			rate: 5,
			name: 'Im pasta',
			description: 'good pasta',
			picture: 'https://i.imgur.com/1.jpg',
			status: 1
		} as shopInfoType;
	}
}
async function getProductList(shop: string) {
	try {
		const resp = await fetch(backendPath + '/store/' + shop + '/products');
		if (resp.ok) {
			return (await resp.json()) as ProductInfoType[];
		}
		return (await (await fetch('/shop.json')).json()) as ProductInfoType[];
	} catch {
		return (await (await fetch('/shop.json')).json()) as ProductInfoType[];
	}
}
