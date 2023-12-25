<script lang="ts">
	import type { PageData } from './$types';
	import { Order } from '$lib';
	export let data: PageData;
</script>

<Order
	orderInfoList={data.cartInfos.store_order_info_array}
	real_total_price={data.cartInfos.real_total_price}
	showDetail={false}
></Order>
