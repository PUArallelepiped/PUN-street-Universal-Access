<script lang="ts">
	import BuyNforMFree from '$lib/components/PUA/buyNforMFree.svelte';
	import { createEventDispatcher } from 'svelte';
	export let product_picture: string;
	export let product_price: number = 160;
	export let product_name: string = '茶碗蒸';
	export let product_quantity: number = 1;
	export let discountQuantity: number;
	export let discountId: number;
	export let store_id: number;
	export let product_id: number;
	export let canDelete: boolean;
	const dispatch = createEventDispatcher();
</script>

<div class="flex h-32 w-96 flex-col rounded-xl bg-white p-2.5">
	<a href={'/shops/' + store_id + '/' + product_id}>
		<div class="flex gap-4">
			<div class="h-20 w-20">
				<img
					class="flex h-20 w-20 truncate rounded-xl object-cover shadow-inner"
					src={product_picture}
					alt="img"
				/>
			</div>
			<div class="flex w-3/4 flex-col gap-3">
				<div class="line-clamp-2 text-ellipsis break-words text-2xl font-bold text-orange-950">
					{product_name}
				</div>
			</div>
		</div>
	</a>
	<div class="flex items-center justify-between">
		<div class="flex items-center gap-4 px-3">
			{#if canDelete}
				<button
					class="p-2"
					on:click={() => {
						dispatch('clickDelete', {});
					}}
				>
					<svg width="16" height="18" fill="none" xmlns="http://www.w3.org/2000/svg">
						<path
							id="Vector"
							d="M6.85714 0C5.6 0 4.57143 1.0125 4.57143 2.25H2.28571C1.02857 2.25 0 3.2625 0 4.5H16C16 3.2625 14.9714 2.25 13.7143 2.25H11.4286C11.4286 1.0125 10.4 0 9.14286 0H6.85714ZM2.28571 6.75V17.5725C2.28571 17.82 2.46857 18 2.72 18H13.3029C13.5543 18 13.7371 17.82 13.7371 17.5725V6.75H11.4514V14.625C11.4514 15.255 10.9486 15.75 10.3086 15.75C9.66857 15.75 9.16572 15.255 9.16572 14.625V6.75H6.88V14.625C6.88 15.255 6.37714 15.75 5.73714 15.75C5.09714 15.75 4.59429 15.255 4.59429 14.625V6.75H2.30857H2.28571Z"
							fill="#755555"
						/>
					</svg>
				</button>
			{/if}
			<div class="flex items-end gap-4">
				<span class=" text-base font-semibold text-red-900">NT$ </span>
				<span class=" text-xl font-semibold text-red-900">{product_price}</span>
				<span class=" text-xl font-semibold text-red-900">x{product_quantity}</span>
			</div>
		</div>
		{#if discountId !== 0}
			<BuyNforMFree quantity={discountQuantity}></BuyNforMFree>
		{/if}
	</div>
</div>
