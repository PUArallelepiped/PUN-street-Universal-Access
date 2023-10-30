<script lang="ts">
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	import ProductCard from '$lib/components/productCard.svelte';
	export let data: PageData;
	let shopName = data.shop;
	let response: {
		id: string;
		name: string;
		address: string;
		email: string;
		phone: string;
	};
	let prodctListResponce: {
		name: string;
		description: string;
		price: number;
		href: string;
	}[] = [
		{
			name: 'TEA EGG',
			description: 'EGG of tea',
			price: 180,
			href: './{shopName}/product3'
		},
		{
			name: 'watermelon',
			description: 'a game',
			price: 0,
			href: './{shopName}/product3'
		},
		{
			name: 'swwika',
			description: 'praying',
			price: 102,
			href: './{shopName}/product3'
		}
	];
	onMount(async () => {
		const resp = await fetch(`http://localhost:5000/api/v1/store/1`);
		response = await resp.json();
		console.log(response);
		return response;
	});
</script>

<div>
	{#if response}
		<div class="text-center text-2xl">
			{response.name}
		</div>
		<div class="text-center text-2xl">
			{response.address}
		</div>
		<div class="text-center text-2xl">
			{response.email}
		</div>
		<div class="text-center text-2xl">
			id: {response.id}
		</div>
		<div class="text-center text-2xl">
			{response.phone}
		</div>
	{:else}
		loading....
	{/if}
</div>
<h1>
	here is {shopName}
</h1>
{#each prodctListResponce as product}
	<ProductCard href={product.href} description={product.description} price={product.price} />
{/each}
