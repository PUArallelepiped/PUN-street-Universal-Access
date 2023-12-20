<script lang="ts">
	import watermelon from '$lib/assets/watermelon.png';
	import type { PageData } from './$types';
	import { ProductCard } from '$lib';
	import TagLabelAreaForShow from '$lib/components/PUA/tagLabelAreaForShow.svelte';
	export let data: PageData;
	let shopName = data.shop;
</script>

<div class="h-48 w-full overflow-hidden">
	<img src={watermelon} alt="" class="w-full object-cover" />
</div>
<div class="mx-5 mt-10 lg:px-40">
	<div class="mx-5 space-y-2">
		<div class="text-5xl font-bold text-PUA-stone">{data.shopInfo.name}</div>
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
			/>
		{/each}
	</div>
</div>

<div class=" bg-orange-950 p-5 text-white">
	<div class="text-lg font-bold">DEBUG AREA</div>
	{#if data.shopInfo}
		<div class="flex-col text-center">
			<div>
				{data.shopInfo.store_id}
			</div>
			<div>
				fee:{data.shopInfo.shipping_fee}
			</div>
			<div>
				{data.shopInfo.address}
			</div>
			<div>
				rate_count: {data.shopInfo.rate_count}
			</div>
			<div>
				rate:{data.shopInfo.rate}
			</div>
			<div>
				{data.shopInfo.name}
			</div>
			<div>
				{data.shopInfo.description}
			</div>
			<div>
				{data.shopInfo.picture}
			</div>
			<div>
				status:{data.shopInfo.status}
			</div>
		</div>
	{:else}
		loading....
	{/if}
</div>
