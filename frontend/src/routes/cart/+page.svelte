<script lang="ts">
	import { backendPath } from '$lib/components/PUA/env';
	import {
		OkButton,
		SeasoningCoupon,
		CartItemCard,
		DenyButton,
		CartMoreItemCard,
		CartLabelBox,
		ShippingCoupon
	} from '$lib/index';
	import { onMount } from 'svelte';
	// let orderInfo = {
	// 	name: 'Order Name',
	// 	phone: 'Phone Number',
	// 	method: 'Taking Method',
	// 	totalShippingFee: 'Total Shipping Fee',
	// 	address: 'Taking Address'
	// };
	type seasonindDiscount = {
		discount_start_date: string;
		discount_name: string;
		discount_end_date: string;
		discount_description: string;
		discount_percentage: number;
		discount_id: number;
		status: number;
	};
	let seasoningDiscountList: seasonindDiscount[] = [];
	onMount(async () => {
		const init = await fetch(`/seasoningDiscount.json`);
		console.log(init);
		if (init.status == 200) {
			seasoningDiscountList = await init.json();
		}
		const resp = await fetch(backendPath + `/seasoning-discounts`);
		if (resp.status == 200) {
			seasoningDiscountList = await resp.json();
		}
	});
</script>

<div class="flex flex-col items-center justify-center gap-5 p-8">
	<div class="flex max-w-7xl flex-col gap-10">
		<div class="flex flex-col">
			<div class="flex justify-between">
				<div class="px-3 text-xl font-semibold leading-normal text-PUA-stone">IM pasta</div>
				<div class="flex items-baseline gap-2 px-7">
					<span class="text-xl font-semibold leading-normal text-PUA-stone">Shipping Fee </span>
					<span class="text-base font-semibold leading-tight text-PUA-stone">NT$</span>
					<span class="text-xl font-semibold leading-normal text-PUA-stone">60</span>
				</div>
			</div>
			<hr class="my-3 border-orange-950" />
			<div class="grid grid-cols-3 gap-3">
				<CartItemCard></CartItemCard>
				<CartItemCard></CartItemCard>
				<CartMoreItemCard></CartMoreItemCard>
			</div>
		</div>
		<div>
			<div class="px-3 text-xl font-semibold leading-normal text-PUA-stone">Discount</div>
			<hr class="my-3 border-orange-950" />
			<div class="grid grid-cols-3 gap-3">
				<ShippingCoupon used={false} />
				<ShippingCoupon used={false} />
				<ShippingCoupon used={true} />
				<ShippingCoupon used={true} />
				<ShippingCoupon />
				<ShippingCoupon />
				<ShippingCoupon />
				{#each seasoningDiscountList as SeasoningDiscount}
					<SeasoningCoupon
						name={SeasoningDiscount.discount_name}
						percentage={SeasoningDiscount.discount_percentage}
						discount_start_date={SeasoningDiscount.discount_start_date}
						discount_end_date={SeasoningDiscount.discount_end_date}
						used={SeasoningDiscount.status === 1}
					/>
				{/each}
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
			<OkButton onclick={() => null} text="Check Out"></OkButton>
		</div>
	</div>
</div>
