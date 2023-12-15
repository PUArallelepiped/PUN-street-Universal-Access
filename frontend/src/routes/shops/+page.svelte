<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';

	//import * as Menubar from "$lib/components/ui/menubar";
	//import * as Select from "$lib/components/ui/select";
	import { Input } from '$lib/components/ui/input';
	import { Checkbox } from '$lib/components/ui/checkbox';
	import { Label } from '$lib/components/ui/label';
	//import { Slider } from '$lib/components/ui/slider';
	//import { Button } from "$lib/components/ui/button";
	//let checked = false;

	import DoubleRangeSlider from '$lib/components/ui/doubleSlider/doubleRangeSlider.svelte';
	import { backendPath } from '$lib/components/PUA/env';
	import StoreCard from '$lib/components/PUA/storeCard.svelte';
	import SortTag from '$lib/components/PUA/sortTag.svelte';
	import TagCheckbox from '$lib/components/PUA/tagCheckbox.svelte';
	import type { list } from 'postcss';
	import type { category } from '$lib';
	let start = 0;
	let end = 1;

	let sortCheckboxes = [
		{ id: 'Nice shop' },
		{ id: 'Fast shop' },
		{ id: 'Something nice for you' },
		{ id: 'Distance' }
	];

	let tagCheckboxes = [
		{ id: 'Fresh' },
		{ id: 'Organic' },
		{ id: 'Healthy' },
		{ id: 'Dairy-free' },
		{ id: 'Bakery' },
		{ id: 'Cafe' },
		{ id: 'Sugar-free' },
		{ id: 'Vegetarian' },
		{ id: 'Free-delivery' }
	];

	let shopListResponse: {
		address: string;
		description: string;
		name: string;
		picture: string;
		rate: number;
		rate_count: number;
		shipping_fee: number;
		status: number;
		store_id: number;
		category_array: category[];
	}[] = [
		{
			store_id: 1,
			shipping_fee: 100,
			address: 'pun street',
			rate_count: 10,
			rate: 5,
			category_array: [
				{
					category_name: 'drink',
					category_id: 1
				},
				{
					category_name: 'drink',
					category_id: 1
				}
			],
			name: "i'm pasta",
			description: 'good pasta',
			picture: 'https://i.imgur.com/1.jpg',
			status: 1
		}
	];

	async function getShopList() {
		const resp = await fetch(backendPath + '/stores');
		shopListResponse = await resp.json();
		return shopListResponse;
	}
</script>

<div class="flex justify-start">
	<div class="float-left m-5 flex flex-col items-center">
		<Input type="text" placeholder="QQㄋㄟㄋㄟ好喝到咩噗茶" class="m-5  max-w-xs bg-white" />
		<div class="flex w-96 flex-col justify-center gap-10 rounded-lg bg-white p-10">
			<div class="flex flex-col items-center gap-5">
				<SortTag value="Sort"></SortTag>
				<div class="flex w-full max-w-xs flex-col justify-start gap-1 px-4">
					{#each sortCheckboxes as { id }}
						<TagCheckbox {id}></TagCheckbox>
					{/each}
				</div>
			</div>

			<div class="flex flex-col items-center gap-5">
				<SortTag value="Price"></SortTag>
				<div class="flex flex-col">
					<div class=" flex w-80 justify-between">
						<div class="text-sm font-bold leading-tight text-PUA-dark-red">
							NT$
							<span class="text-base leading-tight text-PUA-dark-red">1</span>
						</div>
						<div class="text-sm font-bold leading-tight text-PUA-dark-red">
							NT$
							<span class="text-base leading-tight text-PUA-dark-red">1000</span>
						</div>
					</div>
					<div class="  w-full px-10">
						<DoubleRangeSlider bind:start bind:end />
					</div>
				</div>
			</div>
			<div class="flex flex-col items-center justify-center gap-5">
				<SortTag value="Tag"></SortTag>
				<div class="flex w-full max-w-xs flex-col justify-start gap-1 px-4">
					{#each tagCheckboxes as { id }}
						<TagCheckbox {id}></TagCheckbox>
					{/each}
				</div>
			</div>
		</div>
	</div>

	<div class=" m-4 flex h-min flex-wrap gap-3">
		{#await getShopList() then shops}
			{#each shops as shop}
				<StoreCard
					name={shop.name}
					picture={shop.picture}
					rate={shop.rate}
					category_array={shop.category_array}
				/>
			{/each}
		{/await}
	</div>
</div>
