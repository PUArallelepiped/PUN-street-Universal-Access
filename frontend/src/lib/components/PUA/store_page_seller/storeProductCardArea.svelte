<script lang="ts">
	import { StoreProductCard } from '$lib';
	import TagAdd from '../tag/tagAdd.svelte';
	import { backendPath } from '../env';
	export let productListResponse: {
		name: string;
		description: string;
		price: number;
		picture: string;
		product_id: number;
		status: number;
		stock: number;
		store_id: number;
		event_discount_array: {
			discount_max_quantity: number;
			product_id: number;
			discount_name: string;
			discount_description: string;
			discount_id: number;
			status: number;
		}[];
	}[];
	export let shop_id = '1';

	let prodctListResponse_nextId = 0;
	if (prodctListResponse_nextId) {
		prodctListResponse_nextId = productListResponse.length + 1;
	}

	async function removeStoreProductCard(id: number) {
		const resp = await fetch(backendPath + '/product/' + id + '/delete-product', {
			method: 'PUT'
		}).then(() => {
			return fetch(backendPath + `/store/` + shop_id + `/products`);
		});
		console.log();
		productListResponse = await resp.json();
	}
</script>

<div class="mx-5 flex-row space-y-2 p-2">
	{#if productListResponse}
		{#each productListResponse as product}
			<StoreProductCard
				event_discount_array={product.event_discount_array}
				name={product.name}
				description={product.description}
				price={product.price}
				imgUrl={product.picture}
				href={'./' + 'store_page_seller/' + product.product_id}
				on:click={() => removeStoreProductCard(product.product_id)}
			/>
		{/each}
	{/if}
	<div class="flex h-20 items-center justify-center">
		<TagAdd href={'./' + 'store_page_seller/0'}></TagAdd>
	</div>
</div>
