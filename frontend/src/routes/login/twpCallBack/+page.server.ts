import type { PageServerLoad } from './$types';
import type { Cookies } from '@sveltejs/kit';
import { PUBLIC_TWPURL as twpUrl } from '$env/static/public';
import { BACKEND_PATH as backendPath } from '$env/static/private';

export const load: PageServerLoad = async ({ url }) => {
	const state = url.searchParams.get('state');
	const code = url.searchParams.get('code');

	return { state: state, code: code };
};
interface StoreInfo {
	name: string;
	description: string;
	address: string;
	shipping_fee: number | null;
	picture: string;
}
interface UserInfo {
	user_name: string;
	user_email: string;
	password: string;
	phone: string;
	address: string;
	birthday: string;
	StoreRegisterInfo: StoreInfo | null;
}
export const actions = {
	getToken: async ({ request, cookies }) => {
		const data = await request.formData();
		const code = data.get('code');
		const verifier = data.get('verifier');
		const tokenUrl = twpUrl + '/api/oauth/token';
		const resp = await fetch(tokenUrl, {
			headers: {
				Accept: 'application/json',
				'Content-Type': 'application/json'
			},
			method: 'POST',
			body: JSON.stringify({
				code: code,
				code_verifier: verifier
			})
		});
		const cookieStr = resp.headers.get('set-cookie');
		if (cookieStr === null) {
			console.log('redirect to /login');
			return '/login';
		}

		type twpLoginResp = { access_token: string };
		const json = (await resp.json()) as twpLoginResp;
		const access_token = json.access_token;
		const twpUserInfo = await getTwpUserInfo(access_token);
		const isRegister: boolean = await Login(twpUserInfo.email, cookies);
		if (!isRegister) {
			const storeInfo: StoreInfo = {
				name: twpUserInfo.name,
				description: 'twp user',
				address: twpUserInfo.address,
				shipping_fee: 100,
				picture: twpUrl + twpUserInfo.image_url
			};
			const userInfo: UserInfo = {
				user_name: twpUserInfo.name,
				user_email: twpUserInfo.email,
				password: 'twp',
				phone: '0912345678',
				address: twpUserInfo.address,
				birthday: '1970-01-01',
				StoreRegisterInfo: storeInfo
			};
			const res = await fetch(backendPath + '/register', {
				method: 'POST',
				body: JSON.stringify(userInfo)
			});
			console.log(res.status);
			if (res.status !== 200) {
				return '/login';
			}
			const registerSucc = await Login(twpUserInfo.email, cookies);
			if (registerSucc) {
				return '/shops';
			} else {
				return '/login';
			}
		} else {
			return '/shops';
		}
		return '/login';
	}
};

type twpUserInfo = {
	address: 'string';
	email: 'string';
	image_url: 'string';
	name: 'string';
};
const getTwpUserInfo = async (access_token: string) => {
	const resp = await fetch(twpUrl + '/api/user/info', {
		headers: { Authorization: 'Bearer ' + access_token }
	});
	return (await resp.json()) as twpUserInfo;
};

const Login = async (email: string, cookies: Cookies) => {
	const res = await fetch(backendPath + '/login', {
		method: 'POST',
		credentials: 'include',
		body: JSON.stringify({
			user_email: email,
			password: 'twp'
		})
	});
	if (res.status == 200) {
		await res.json().then((data) => {
			cookies.set('jwttoken', data, { path: '/', sameSite: 'strict' });
		});
		return true;
	} else {
		return false;
	}
};
