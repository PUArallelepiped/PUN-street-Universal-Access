<script lang="ts">
	import type { PageData } from './$types';
	import { ProductCard } from '$lib';
	import TagLabelAreaForShow from '$lib/components/PUA/tagLabelAreaForShow.svelte';
	export let data: PageData;
	let shopName = data.shop;
	console.log(data);
</script>

<div class="h-48 w-full overflow-hidden shadow">
	<img src={data.shopInfo.picture} alt="" class="w-full object-cover" />
</div>
<div class="mx-5 mt-10 lg:px-40">
	<div class="mx-5 space-y-2">
		<div class="text-PUA-stone text-5xl font-bold">{data.shopInfo.name}</div>
		<div class="font-bold text-red-950">{data.shopInfo.address}</div>
		<div class="flex w-full justify-start gap-6">
			<TagLabelAreaForShow
				bind:tagText={data.shopInfo.category_array}
				bind:star_score={data.shopInfo.rate}
			></TagLabelAreaForShow>
		</div>
	</div>
	<div class="flex-row space-y-2 p-2">
		{#each data.productList as product}
			<ProductCard
				name={product.name}
				href={'./' + shopName + '/' + product.product_id}
				description={product.description}
				price={product.price}
				imgUrl={product.picture}
				event_discount_array={product.event_discount_array}
				status={product.status}
			/>
		{/each}
	</div>
</div>
