<script lang="ts">
	import { CartItemCard, CartLabelBox } from '$lib/index';
	let orderInfo = {
		name: 'Order Name',
		phone: 'Phone Number',
		method: 'Taking Method',
		totalShippingFee: 'Total Shipping Fee',
		address: 'Taking Address'
	};
</script>

<div class="flex flex-col items-center justify-center gap-5">
	<div class="flex max-w-7xl flex-col gap-3">
		<div class="flex flex-col">
			<div class="flex justify-between">
				<div class="text-xl font-semibold leading-normal text-stone-600">IM pasta</div>
				<div>
					<span style="text-stone-600 text-xl font-semibold  leading-normal"
						>Shipping Fee
					</span><span style="text-stone-600 text-base font-semibold  leading-tight">NT$</span><span
						style="text-stone-600 text-xl font-semibold  leading-normal"
					>
						60</span
					>
				</div>
			</div>
			<hr class="border-orange-950" />
		</div>
		<div class="grid grid-cols-3 gap-3">
			<CartItemCard></CartItemCard>
			<CartItemCard></CartItemCard>
		</div>
		<div>
			<div class="text-xl font-semibold leading-normal text-stone-600">Discount</div>
			<hr class="border-orange-950" />
		</div>
		<CartLabelBox></CartLabelBox>
	</div>
</div>
