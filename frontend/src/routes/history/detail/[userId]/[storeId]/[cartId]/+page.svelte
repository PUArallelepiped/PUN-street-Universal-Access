<script lang="ts">
	import Order from '$lib/components/PUA/order.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<Order
	showDetail={false}
	orderInfoList={[data.cartInfos]}
	real_total_price={data.cartInfos.total_price}
	type="history"
></Order>
