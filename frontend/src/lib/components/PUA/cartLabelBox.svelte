<script lang="ts">
	import { PUALabel } from '$lib/index';
	let orderInfo = {
		name: 'Order Name',
		phone: 'Phone Number',
		method: 'Taking Method',
		totalShippingFee: 'Total Shipping Fee',
		address: 'Taking Address'
	};
</script>

<div class="flex flex-col text-xl font-semibold text-stone-600">
	<PUALabel labelName={orderInfo.name} value="0912345678"></PUALabel>
	<PUALabel labelName={orderInfo.phone} value="0912345678"></PUALabel>
	<PUALabel labelName={orderInfo.method} value="0912345678"></PUALabel>
	<PUALabel labelName={orderInfo.totalShippingFee} value="100"></PUALabel>
	<PUALabel labelName={orderInfo.address} value="a street b floor c city d country nnumber e"
	></PUALabel>
</div>
