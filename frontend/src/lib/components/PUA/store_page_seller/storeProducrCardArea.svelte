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

	function addStoreProductCard() {
		let prodctListResponse_nextId = prodctListResponse.length;
		prodctListResponse = [
			...prodctListResponse,
			{
				status: 1,
				stock: 100,
				store_id: prodctListResponse_nextId,
				name: 'watermelon',
				description: 'a game',
				price: 0,
				picture: 'https://i.imgur.com/3i3tyXJ.gif',
				product_id: 2
			}
		];
	}
	function removeStoreProductCard(id: number) {
		let result = prodctListResponse.find((item) => item.store_id === id);
		if (result !== undefined) {
			prodctListResponse = prodctListResponse.filter((cat) => cat.store_id !== id);
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
			on:click={() => removeStoreProductCard(product.store_id)}
		/>
	{/each}
	<div class="flex h-20 items-center justify-center">
		<HashtagLabel on:click={addStoreProductCard} type={'add'}></HashtagLabel>
	</div>
</div>
