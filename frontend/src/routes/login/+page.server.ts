import { BACKEND_PATH as backendPath } from '$env/static/private';

export const actions = {
	login: async ({ request, cookies }) => {
		const data = await request.formData();
		const user_email = data.get('user_email');
		const password = data.get('password');
		const maxAge = 60 * 60 * 24; // 1 day
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
				.then((data) =>
					cookies.set('jwttoken', data, { maxAge: maxAge, sameSite: 'strict', httpOnly: false })
				);
			// console.log((await getId()).valueOf());
		}
		return res.status;
	}
};
