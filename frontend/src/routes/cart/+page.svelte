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
	import type { PageData } from './$types';
	import { CassetteTape } from 'lucide-svelte';
	export let data: PageData;
	console.log(data);
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
		if (init.status == 200) {
			seasoningDiscountList = await init.json();
		}
		const resp = await fetch(backendPath + `/seasoning-discounts`);
		if (resp.status == 200) {
			seasoningDiscountList = await resp.json();
		}
	});

	let showDetail = true;

	function checkout() {
		//TODO customer id need to change
		fetch(backendPath + '/customer/1/cart/1/store/1/checkout', {
			method: 'POST',
			body: JSON.stringify({
				seasoning_discount_id: 1,
				shipping_discount_id: 1,
				taking_method: 1
			})
		});
		return null;
	}
</script>

<div class="flex flex-col items-center justify-center gap-5 p-8">
	<div class="flex max-w-7xl flex-col gap-10">
		{#if data.cartInfos}
			{#each data.cartInfos.store_order_info_array as storeInfo}
				<div class="flex flex-col">
					<div class="flex justify-between">
						<div class="px-3 text-xl font-semibold leading-normal text-PUA-stone">
							{storeInfo.store_name}
						</div>
						<div class="flex items-baseline gap-2 px-7">
							<span class="text-xl font-semibold leading-normal text-PUA-stone">Shipping Fee </span>
							<span class="text-base font-semibold leading-tight text-PUA-stone">NT$</span>
							<span class="text-xl font-semibold leading-normal text-PUA-stone"
								>{storeInfo.store_shipping_fee}</span
							>
						</div>
					</div>
					<hr class="my-3 border-orange-950" />
					<div class="flex flex-wrap gap-3">
						{#each storeInfo.product_order as productInfo}
							<CartItemCard
								product_picture={productInfo.product_picture}
								description={'null'}
								product_name={productInfo.product_name}
								product_quantity={productInfo.product_quantity}
								product_price={productInfo.product_price}
							></CartItemCard>
						{/each}
						<CartMoreItemCard></CartMoreItemCard>
					</div>
				</div>
			{/each}
		{/if}

		<div>
			<div class="px-3 text-xl font-semibold leading-normal text-PUA-stone">Discount</div>
			<hr class="my-3 border-orange-950" />
			<div class="flex flex-wrap gap-3">
				<!-- {#each seasoningDiscountList as SeasoningDiscount}
					<SeasoningCoupon
						name={SeasoningDiscount.discount_name}
						percentage={SeasoningDiscount.discount_percentage}
						discount_start_date={SeasoningDiscount.discount_start_date}
						discount_end_date={SeasoningDiscount.discount_end_date}
						used={SeasoningDiscount.status === 1}
					/>
				{/each} -->
			</div>
		</div>
		{#if showDetail}
			<div class="flex justify-end">
				<div class="w-1/2">
					<CartLabelBox></CartLabelBox>
				</div>
			</div>
		{/if}
		<div class="flex justify-between">
			<DenyButton onclick={() => null}><div class="px-4">Delete All</div></DenyButton>
			<div class="flex place-items-baseline justify-between gap-7">
				<div class="text-4xl font-semibold text-orange-950">Total Price</div>
				<div class="flex items-baseline gap-3">
					<div class="text-2xl font-semibold text-orange-950">NT$</div>
					<div class="text-right text-3xl font-semibold text-orange-950">9999</div>
				</div>
			</div>
			<OkButton onclick={checkout} text="Check Out"></OkButton>
		</div>
	</div>
</div>
