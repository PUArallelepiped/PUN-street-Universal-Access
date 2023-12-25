<script lang="ts">
	import { backendPath } from '$lib/components/PUA/env';
	import type { PageData } from './$types';
	import { goto, invalidateAll } from '$app/navigation';
	import { Order } from '$lib';
	export let data: PageData;

	let showDetail = false;

	async function checkout() {
		//TODO customer id need to change
		await fetch(backendPath + '/customer/1/checkout', {
			method: 'POST',
			body: JSON.stringify({
				seasoning_discount_id: 1,
				shipping_discount_id: 1,
				taking_method: 1
			})
		});
		invalidateAll();
		return null;
	}
</script>

<Order
	orderInfoList={data.cartInfos.store_order_info_array}
	real_total_price={data.cartInfos.real_total_price}
	showDetail={false}
></Order>
