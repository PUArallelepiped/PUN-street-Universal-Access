<script lang="ts">
	import { backendPath } from '$lib/components/PUA/env';
	import {
		OkButton,
		SeasoningCoupon,
		CartItemCard,
		DenyButton,
		CartMoreItemCard,
		CartLabelBox,
		ShippingCoupon,
		PUALabel
	} from '$lib/index';
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	import { CassetteTape } from 'lucide-svelte';
	import { invalidateAll } from '$app/navigation';
	export let data: PageData;
	console.log(data);

	let showDetail = false;

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
	<div class="flex w-full max-w-7xl flex-col gap-10">
		{#if data.cartInfos}
			{#if data.cartInfos.store_order_info_array != null}
				{#each data.cartInfos.store_order_info_array as storeInfo}
					<div class="flex flex-col">
						<div class="flex justify-between">
							<div class="px-3 text-xl font-semibold leading-normal text-PUA-stone">
								{storeInfo.store_name}
							</div>
							<div class="flex items-baseline gap-2 px-7">
								<span class="text-xl font-semibold leading-normal text-PUA-stone"
									>Shipping Fee
								</span>
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
									discountId={productInfo.event_discount_id}
									discountQuantity={productInfo.event_discount_max_quantity}
									on:clickDelete={async () => {
										console.log('click delete');
										const resp = await fetch(
											backendPath + '/customer/1/delete/product/' + productInfo.product_id,
											{ method: 'DELETE' }
										);
										console.log(resp.status);
										invalidateAll();
									}}
								></CartItemCard>
							{/each}
							<CartMoreItemCard></CartMoreItemCard>
						</div>
					</div>
				{/each}
			{:else}
				<div class="text-center text-5xl font-bold text-PUA-dark-red">No Product in your cart</div>
			{/if}
		{/if}

		<div>
			<div class="px-3 text-xl font-semibold leading-normal text-PUA-stone">Discount</div>
			<hr class="my-3 border-orange-950" />
			<div class="flex flex-wrap gap-3">
				{#if data.cartInfos.store_order_info_array != null}
					{#each data.cartInfos.store_order_info_array as storeInfo}
						<ShippingCoupon
							used={storeInfo.shipping_discount_bool}
							store_name={storeInfo.store_name}
							max_price={storeInfo.shipping_discount.discount_max_price}
							description={storeInfo.shipping_discount.discount_description}
						></ShippingCoupon>
					{/each}
					{#if data.cartInfos.store_order_info_array.length != 0}
						<SeasoningCoupon
							name={data.cartInfos.store_order_info_array[0].seasoning_discount.discount_name}
							percentage={data.cartInfos.store_order_info_array[0].seasoning_discount
								.discount_percentage}
							discount_start_date={data.cartInfos.store_order_info_array[0].seasoning_discount
								.discount_start_date}
							discount_end_date={data.cartInfos.store_order_info_array[0].seasoning_discount
								.discount_end_date}
							used={data.cartInfos.store_order_info_array[0].seasoning_discount_bool}
						/>
					{/if}
				{:else}
					<div class=" w-full text-center text-5xl font-bold text-PUA-dark-red">
						No Discount In cart
					</div>
				{/if}
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
			<!-- <DenyButton onclick={() => null}><div class="px-4">Delete All</div></DenyButton> -->
			<div class="flex place-items-baseline justify-between gap-7">
				<div class="text-4xl font-semibold text-orange-950">Total Price</div>
				<div class="flex items-baseline gap-3">
					<div class="text-2xl font-semibold text-orange-950">NT$</div>
					<div class="text-right text-3xl font-semibold text-orange-950">
						{data.cartInfos.real_total_price}
					</div>
				</div>
			</div>
			<OkButton onclick={checkout} text="Check Out"></OkButton>
		</div>
	</div>
</div>
