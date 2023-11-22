<script lang="ts">
	import {
		OkButton,
		SeasoningCoupon,
		CartItemCard,
		DenyButton,
		CartLabelBox,
		ShippingCoupon
	} from '$lib/index';
	// let orderInfo = {
	// 	name: 'Order Name',
	// 	phone: 'Phone Number',
	// 	method: 'Taking Method',
	// 	totalShippingFee: 'Total Shipping Fee',
	// 	address: 'Taking Address'
	// };
</script>

<div class="flex flex-col items-center justify-center gap-5 p-8">
	<div class="flex max-w-7xl flex-col gap-10">
		<div class="flex flex-col">
			<div class="flex justify-between">
				<div class="px-3 text-xl font-semibold leading-normal text-stone-600">IM pasta</div>
				<div class="px-7">
					<span class="text-xl font-semibold leading-normal text-stone-600"
						>Shipping Fee
					</span><span class="text-base font-semibold leading-tight text-stone-600">NT$</span><span
						class="text-xl font-semibold leading-normal text-stone-600"
					>
						60</span
					>
				</div>
			</div>
			<hr class="my-3 border-orange-950" />
			<div class="grid grid-cols-3 gap-3">
				<CartItemCard></CartItemCard>
				<CartItemCard></CartItemCard>
			</div>
		</div>
		<div>
			<div class="px-3 text-xl font-semibold leading-normal text-stone-600">Discount</div>
			<hr class="my-3 border-orange-950" />
			<div class="grid grid-cols-3 gap-3">
				<ShippingCoupon used={false} />
				<ShippingCoupon used={false} />
				<ShippingCoupon used={true} />
				<ShippingCoupon used={true} />
				<ShippingCoupon />
				<ShippingCoupon />
				<ShippingCoupon />
				<SeasoningCoupon></SeasoningCoupon>
			</div>
		</div>
		<CartLabelBox></CartLabelBox>
		<div class="flex justify-between">
			<DenyButton onclick={() => null}><div class="px-4">Delete All</div></DenyButton>
			<div class="flex place-items-baseline justify-between gap-7">
				<div class="text-4xl font-semibold text-orange-950">Total Price</div>
				<div class="flex items-baseline gap-3">
					<div class="text-2xl font-semibold text-orange-950">NT$</div>
					<div class="text-right text-3xl font-semibold text-orange-950">9999</div>
				</div>
			</div>
			<OkButton onclick={() => null}><div class="px-16 text-xl font-bold">Check out</div></OkButton>
		</div>
	</div>
</div>
