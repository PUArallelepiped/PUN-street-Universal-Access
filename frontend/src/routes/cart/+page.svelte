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

	let showDetail = false;

	type cartInfo = {
		cart_id: number;
		customer_id: number;
		discount_id: number;
		product_id: number;
		product_quantity: number;
		store_id: number;
	};
	let cartInfos: cartInfo[] = [
		{
			cart_id: 1,
			customer_id: 1,
			discount_id: 9,
			product_id: 1,
			product_quantity: 100,
			store_id: 1
		}
	];
	let processedInfo: cartInfo[][];
	let processedIndex: number[] = [];
	onMount(async () => {
		const resp = await fetch(backendPath + '/customer/1/cart/1/carts');
		cartInfos = await resp.json();
		processedInfo = test(cartInfos);
		console.log(processedIndex);
		console.log(processedInfo);
	});
	function test(cartInfos: cartInfo[]) {
		let processedInfo: cartInfo[][] = [];
		cartInfos.forEach((element) => {
			let index = processedIndex.findIndex((v) => v == element.store_id);
			if (index == -1) {
				processedInfo.push([element]);
				processedIndex.push(element.store_id);
			} else {
				processedInfo[index].push(element);
			}
		});
		return processedInfo;
	}
	function checkout() {
		//TODO customer id need to change
		async () => {
			const resp = await fetch(backendPath + '/customer/1/cart/1/store/1/checkout', {
				method: 'POST',
				body: JSON.stringify({
					seasoning_discount_id: null,
					shipping_discount_id: null,
					taking_method: null
				})
			});
			let response: Promise<any> = await resp.json();
			console.log(response);
		};
		return null;
	}
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
			<div class="flex flex-wrap gap-3">
				<CartItemCard></CartItemCard>
				<CartItemCard></CartItemCard>
				<CartMoreItemCard></CartMoreItemCard>
			</div>
		</div>
		{#if processedInfo}
			{#each processedInfo as info}
				<div class="flex flex-col">
					<div class="flex justify-between">
						<div class="px-3 text-xl font-semibold leading-normal text-PUA-stone">
							{info[0].store_id}
						</div>
						<div class="flex items-baseline gap-2 px-7">
							<span class="text-xl font-semibold leading-normal text-PUA-stone">Shipping Fee </span>
							<span class="text-base font-semibold leading-tight text-PUA-stone">NT$</span>
							<span class="text-xl font-semibold leading-normal text-PUA-stone">null</span>
						</div>
					</div>
					<hr class="my-3 border-orange-950" />
					<div class="flex flex-wrap gap-3">
						{#each info as item}
							<CartItemCard {...item}></CartItemCard>
						{/each}
						<CartItemCard></CartItemCard>
						<CartItemCard></CartItemCard>
						<CartMoreItemCard></CartMoreItemCard>
					</div>
				</div>
			{/each}
		{/if}

		<div>
			<div class="px-3 text-xl font-semibold leading-normal text-PUA-stone">Discount</div>
			<hr class="my-3 border-orange-950" />
			<div class="flex flex-wrap gap-3">
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
		{#if showDetail}
			<CartLabelBox></CartLabelBox>
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
			<OkButton onclick={checkout}><div class="px-16 text-xl font-bold">Check out</div></OkButton>
		</div>
	</div>
</div>
