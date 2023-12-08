<script lang="ts">
	import { HashtagLabel, StoreProductCard } from '$lib';

	export let prodctListResponse: {
		name: string;
		description: string;
		price: number;
		picture: string;
		product_id: number;
		status: number;
		stock: number;
		store_id: number;
	}[];

	let prodctListResponse_nextId = prodctListResponse.length + 1;

	function addStoreProductCard() {
		prodctListResponse = [
			...prodctListResponse,
			{
				status: 1,
				stock: 100,
				store_id: 2,
				name: 'watermelon',
				description: 'a game',
				price: 0,
				picture: 'https://i.imgur.com/3i3tyXJ.gif',
				product_id: prodctListResponse_nextId
			}
		];
		prodctListResponse_nextId = prodctListResponse_nextId + 1;
	}
	function removeStoreProductCard(id: number) {
		let result = prodctListResponse.find((item) => item.product_id === id);
		if (result !== undefined) {
			prodctListResponse = prodctListResponse.filter((cat) => cat.product_id !== id);
		}
	}
</script>

<div class="mx-5 flex-row space-y-2 p-2">
	{#each prodctListResponse as product}
		<StoreProductCard
			name={product.name}
			description={product.description}
			price={product.price}
			imgUrl={product.picture}
			on:click={() => removeStoreProductCard(product.product_id)}
		/>
	{/each}
	<div class="flex h-20 items-center justify-center">
		<HashtagLabel on:click={addStoreProductCard} type={'add'}></HashtagLabel>
	</div>
</div>
