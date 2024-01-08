import { PUBLIC_TWPURL as twpUrl, PUBLIC_BACKEND_PATH as backendPath } from '$env/static/public';

export const actions = {
    login: async ({ request, cookies }) => {
        const data = await request.formData();
		const user_email = data.get('user_email');
		const password = data.get('password');
		let maxAge = 60 * 60 * 24; // 1 day
		const res = await fetch(backendPath + '/login', {
			method: 'POST',
			credentials: 'include',
			body: JSON.stringify({
				user_email,
				password
			})
		});
		if (res.status == 200) {
			await res
				.json()
				.then(
					(data) =>
					cookies.set('jwttoken', data, { maxAge: maxAge, path: '/' , sameSite: 'strict'})
				);
			// console.log((await getId()).valueOf());
		}
		return res.status;
}}