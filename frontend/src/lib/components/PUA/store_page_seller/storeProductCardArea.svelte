<script lang="ts">
	import { StoreProductCard } from '$lib';
	import TagAdd from '../tag/tagAdd.svelte';
	export let productListResponse: {
		name: string;
		description: string;
		price: number;
		picture: string;
		product_id: number;
		status: number;
		stock: number;
		store_id: number;
	}[];

	let prodctListResponse_nextId = 0;
	if (prodctListResponse_nextId) {
		prodctListResponse_nextId = productListResponse.length + 1;
	}

	function removeStoreProductCard(id: number) {
		let result = productListResponse.find((item) => item.product_id === id);
		if (result !== undefined) {
			productListResponse = productListResponse.filter((cat) => cat.product_id !== id);
		}
	}
</script>

<div class="mx-5 flex-row space-y-2 p-2">
	{#each productListResponse as product}
		<StoreProductCard
			name={product.name}
			description={product.description}
			price={product.price}
			imgUrl={product.picture}
			href={'./' + 'store_page_seller/' + product.product_id}
			on:click={() => removeStoreProductCard(product.product_id)}
		/>
	{/each}
	<div class="flex h-20 items-center justify-center">
		<TagAdd href={'./' + 'store_page_seller/0'}></TagAdd>
	</div>
</div>
