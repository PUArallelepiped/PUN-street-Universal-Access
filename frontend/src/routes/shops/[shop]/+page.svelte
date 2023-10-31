<script lang="ts">
	import watermelon from '$lib/assets/watermelon.png';
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	import { ProductCard } from '$lib';
	export let data: PageData;
	let shopName = data.shop;
	let response: {
		id: string;
		name: string;
		address: string;
		email: string;
		phone: string;
	} = {
		id: '1',
		name: 'Im pasta',
		address: '100台灣台北市中正區八德路一段82巷9弄17號',
		email: 'impasta@pun.street.tw',
		phone: '09-1234-5678'
	};
	let prodctListResponce: {
		name: string;
		description: string;
		price: number;
		href: string;
	}[] = [
		{
			name: 'TEA EGG',
			description:
				'EGG of tea\n expensive\n also call putting, egg,egg,egg,egg, egg, egg,limit, limit, limit, limit, limit, limit, limit, limit, limit, ',
			price: 180,
			href: './' + shopName + '/product1'
		},
		{
			name: 'watermelon',
			description: 'a game',
			price: 0,
			href: './' + shopName + '/product2'
		},
		{
			name: 'swwika',
			description: 'praying',
			price: 102,
			href: './' + shopName + '/product3'
		}
	];
	onMount(async () => {
		const resp = await fetch(`http://localhost:5000/api/v1/store/1`);
		response = await resp.json();
		console.log(response);
		return response;
	});
</script>

<img src={watermelon} alt="" class="h-48 w-full" />
<div class="mx-5 lg:px-40">
	<div class="p-2">
		<div class="m-1 p-2 text-5xl font-bold">
			{response.name}
		</div>
		<div class="pl-5">
			{response.address}
		</div>
	</div>
	<div class="flex-row space-y-2 p-2">
		{#each prodctListResponce as product}
			<ProductCard
				name={product.name}
				href={product.href}
				description={product.description}
				price={product.price}
			/>
		{/each}
	</div>
</div>

<div class=" bg-orange-950 p-5 text-white">
	<div class="text-lg font-bold">DEBUG AREA</div>
	{#if response}
		<div class="text-center">
			{response.name}
		</div>
		<div class="text-center">
			{response.address}
		</div>
		<div class="text-center">
			{response.email}
		</div>
		<div class="text-center">
			id: {response.id}
		</div>
		<div class="text-center">
			{response.phone}
		</div>
	{:else}
		loading....
	{/if}
</div>
