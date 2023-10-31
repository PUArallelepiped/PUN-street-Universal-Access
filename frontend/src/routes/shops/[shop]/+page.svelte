<script lang="ts">
	import watermelon from '$lib/assets/watermelon.png';
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	import { ProductCard } from '$lib';
	import { backendPath } from '$lib/components/PUA/env';
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
	let prodctListResponse: {
		name: string;
		description: string;
		price: number;
		href: string;
		picture: string;
		product_id: number;
		status: number;
		stock: number;
		store_id: number;
	}[] = [
		{
			status: 1,
			name: 'TEA EGG',
			description:
				'EGG of tea\n expensive\n also call putting, egg,egg,egg,egg, egg, egg,limit, limit, limit, limit, limit, limit, limit, limit, limit, ',
			price: 180,
			href: './' + shopName + '/product1',
			picture: 'https://i.imgur.com/3i3tyXJ.gif',
			product_id: 1,
			store_id: 1,

			stock: 100
		},
		{
			status: 1,
			stock: 100,
			store_id: 1,
			name: 'watermelon',
			description: 'a game',
			price: 0,
			href: './' + shopName + '/product2',
			picture: 'https://i.imgur.com/3i3tyXJ.gif',
			product_id: 2
		},
		{
			status: 1,
			stock: 100,
			store_id: 1,
			product_id: 3,
			picture: 'https://i.imgur.com/3i3tyXJ.gif',
			name: 'swwika',
			description: 'praying',
			price: 102,
			href: './' + shopName + '/product3'
		}
	];
	onMount(async () => {
		const resp = await fetch(backendPath + `/store/1/products`);
		prodctListResponse = await resp.json();
		console.log(prodctListResponse);
	});
	onMount(async () => {
		const resp = await fetch(backendPath + `/store/1`);
		shopInfoResponse = await resp.json();
	});
</script>

<img src={watermelon} alt="" class="h-48 w-full" />
<div class="mx-5 lg:px-40">
	<div class="p-2">
		<div class="m-1 p-2 text-5xl font-bold">
			{shopInfoResponse.name}
		</div>
		<div class="pl-5">
			{shopInfoResponse.address}
		</div>
	</div>
	<div class="flex-row space-y-2 p-2">
		{#each prodctListResponse as product}
			<ProductCard
				name={product.name}
				href={product.href}
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
