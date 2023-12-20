<script lang="ts">
	import watermelon from '$lib/assets/watermelon.png';
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	import { ProductCard } from '$lib';
	import { backendPath } from '$lib/components/PUA/env';
	import TagLabelAreaForShow from '$lib/components/PUA/tagLabelAreaForShow.svelte';
	export let data: PageData;
	let shopName = data.shop;

	let shopInfoResponse: {
		store_id: number;
		shipping_fee: number;
		address: string;
		rate_count: number;
		rate: number;
		name: string;
		description: string;
		picture: string;
		status: number;
	} = {
		store_id: 1,
		shipping_fee: 100,
		address: '100台灣台北市中正區八德路一段82巷9弄17號',
		rate_count: 10,
		rate: 5,
		name: 'Im pasta',
		description: 'good pasta',
		picture: 'https://i.imgur.com/1.jpg',
		status: 1
	};
	type ProductResponse = {
		name: string;
		description: string;
		price: number;
		picture: string;
		product_id: number;
		status: number;
		stock: number;
		store_id: number;
	};
	let prodctListResponse: ProductResponse[] = [];
	onMount(async () => {
		const resp = await fetch('/shop.json');
		prodctListResponse = await resp.json();

		const backResp = await fetch(backendPath + `/store/1/products`);
		if (backResp.status == 200) {
			prodctListResponse = await backResp.json();
		}
	});
	let tagList = [
		{ category_id: 0, category_name: 'free' },
		{ category_id: 1, category_name: 'drink' },
		{ category_id: 2, category_name: 'drink' },
		{ category_id: 3, category_name: 'local' },
		{ category_id: 4, category_name: 'free delivery hahahaha' },
		{ category_id: 5, category_name: 'color' },
		{ category_id: 6, category_name: 'school free' },
		{ category_id: 7, category_name: 'for student' },
		{ category_id: 8, category_name: 'delicious' }
	];
	onMount(async () => {
		const resp = await fetch(backendPath + `/store/1`);
		shopInfoResponse = await resp.json();
		// const tag_category = await fetch(`/categories.json`);
		// tagList = await tag_category.json();
	});
</script>

<div class="h-48 w-full overflow-hidden">
	<img src={watermelon} alt="" class="w-full object-cover" />
</div>
<div class="mx-5 mt-10 lg:px-40">
	<!-- <div class="p-2">
		<div class="m-1 p-2 text-5xl font-bold">
			{shopInfoResponse.name}
		</div>
		<div class="pl-5">
			{shopInfoResponse.address}
		</div>
	</div> -->
	<div class="mx-5 space-y-2">
		<div class="text-PUA-stone text-5xl font-bold">{shopInfoResponse.name}</div>
		<div class="font-bold text-red-950">{shopInfoResponse.address}</div>
		<div class="flex w-full justify-start gap-6">
			<TagLabelAreaForShow tagText={tagList} star_score={shopInfoResponse.rate.toString()}
			></TagLabelAreaForShow>
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
	{#if shopInfoResponse}
		<div class="flex-col text-center">
			<div>
				{shopInfoResponse.store_id}
			</div>
			<div>
				fee:{shopInfoResponse.shipping_fee}
			</div>
			<div>
				{shopInfoResponse.address}
			</div>
			<div>
				rate_count: {shopInfoResponse.rate_count}
			</div>
			<div>
				rate:{shopInfoResponse.rate}
			</div>
			<div>
				{shopInfoResponse.name}
			</div>
			<div>
				{shopInfoResponse.description}
			</div>
			<div>
				{shopInfoResponse.picture}
			</div>
			<div>
				status:{shopInfoResponse.status}
			</div>
		</div>
	{:else}
		loading....
	{/if}
</div>
