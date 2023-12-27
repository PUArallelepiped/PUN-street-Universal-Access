<script lang="ts">
	import { goto, invalidateAll } from '$app/navigation';

	import {
		CartItemCard,
		CartLabelBox,
		CartMoreItemCard,
		OkButton,
		SeasoningCoupon,
		ShippingCoupon
	} from '$lib';
	import type { storeOrderInfo } from '$lib';
	import { backendPath } from './env';
	import { getId } from './getId';

	export let showDetail: boolean;
	export let orderInfoList: storeOrderInfo[];
	export let real_total_price: number;
	export let type: string = 'cart';

	async function checkout() {
		try {
			const user_id = (await getId()).valueOf();
			await fetch(backendPath + '/customer/' + user_id + '/checkout', {
				method: 'POST',
				body: JSON.stringify({
					seasoning_discount_id: 1,
					shipping_discount_id: 1,
					taking_method: 1
				})
			});
			invalidateAll();
			return null;
		} catch (e) {
			goto('/login');
		}
	}
</script>

<div class="flex flex-col items-center justify-center gap-5 p-8">
	<div class="flex w-full max-w-7xl flex-col gap-10">
		{#if orderInfoList}
			{#if orderInfoList != null}
				{#each orderInfoList as storeInfo}
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
									product_id={productInfo.product_id}
									store_id={storeInfo.store_id}
									product_picture={productInfo.product_picture}
									description={'null'}
									product_name={productInfo.product_name}
									product_quantity={productInfo.product_quantity}
									product_price={productInfo.product_price}
									discountId={productInfo.event_discount_id}
									discountQuantity={productInfo.event_discount_max_quantity}
									canDelete={type == 'cart'}
									on:clickDelete={async () => {
										try {
											const user_id = (await getId()).valueOf();
											await fetch(
												backendPath +
													'/customer/' +
													user_id +
													'/delete/product/' +
													productInfo.product_id,
												{ method: 'DELETE' }
											);
											invalidateAll();
										} catch (e) {
											goto('/login');
										}
									}}
								></CartItemCard>
							{/each}
							{#if type == 'cart'}
								<CartMoreItemCard
									on:click={() => {
										goto('/shops/' + storeInfo.store_id);
									}}
								></CartMoreItemCard>
							{/if}
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
				{#if orderInfoList != null}
					{#each orderInfoList as storeInfo}
						<ShippingCoupon
							used={storeInfo.shipping_discount_bool}
							store_name={storeInfo.store_name}
							max_price={storeInfo.shipping_discount.discount_max_price}
							description={storeInfo.shipping_discount.discount_description}
						></ShippingCoupon>
					{/each}
					{#if orderInfoList.length != 0}
						{#if orderInfoList[0].seasoning_discount.discount_id != 1}
							<SeasoningCoupon
								name={orderInfoList[0].seasoning_discount.discount_name}
								percentage={orderInfoList[0].seasoning_discount.discount_percentage}
								discount_start_date={orderInfoList[0].seasoning_discount.discount_start_date}
								discount_end_date={orderInfoList[0].seasoning_discount.discount_end_date}
								used={orderInfoList[0].seasoning_discount_bool}
							/>
						{/if}
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
						{real_total_price}
					</div>
				</div>
			</div>
			{#if type == 'cart'}
				<OkButton onclick={checkout} text="Check Out"></OkButton>
			{/if}
		</div>
	</div>
</div>
