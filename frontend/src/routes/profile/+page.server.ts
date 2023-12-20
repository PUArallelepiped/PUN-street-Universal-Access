import { backendPath } from '$lib/components/PUA/env.js';

export const actions = {
	monthStatistic: async ({ request }) => {
		const data = await request.formData();

		const resp = await fetch(
			backendPath + '/store/' + '1' + '/get-selling/' + data.get('year') + '/' + data.get('month')
		);
		return await resp.json();
	},
	yearStatistic: async ({ request }) => {
		const data = await request.formData();

		const resp = await fetch(backendPath + '/store/' + '1' + '/get-statistics/' + data.get('year'));
		return await resp.json();
	}
};
