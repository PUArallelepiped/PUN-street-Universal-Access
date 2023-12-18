import { backendPath } from '$lib/components/PUA/env.js';

export const actions = {
	statistic: async ({ request }) => {
		const data = await request.formData();

		const resp = await fetch(
			backendPath + '/store/' + '1' + '/get-selling/' + '2020/' + data.get('month')
		);
		const j = await resp.json();
		return j;
	}
};
