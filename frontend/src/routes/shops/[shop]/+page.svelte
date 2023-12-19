<script lang="ts">
	import watermelon from '$lib/assets/watermelon.png';
	import type { PageData } from './$types';
	import { ProductCard } from '$lib';
	export let data: PageData;
	let shopName = data.shop;
</script>

<div class="h-48 w-full overflow-hidden">
	<img src={watermelon} alt="" class="w-full object-cover" />
</div>
<div class="mx-5 lg:px-40">
	<div class="p-2">
		<div class="m-1 p-2 text-5xl font-bold">
			{data.shopInfo.name}
		</div>
		<div class="pl-5">
			{data.shopInfo.address}
		</div>
	</div>
	<div class="flex-row space-y-2 p-2">
		{#each prodctListResponse as product}
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
