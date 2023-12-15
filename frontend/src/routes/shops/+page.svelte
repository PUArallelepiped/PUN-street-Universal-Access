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
	}[] = [
		{
			store_id: 1,
			description: 'you are not pasta',
			name: 'im pasta',
			address: 'pun street',
			rate: 4,
			rate_count: 100,
			picture: 'https://i.imgur.com/3i3tyXJ.gif',
			status: 1,
			shipping_fee: 35
		},
		{
			store_id: 2,
			description: 'special meal',
			name: 'number five',
			address: 'pun street',
			rate: 3.5,
			rate_count: 5,
			picture: 'https://i.imgur.com/3i3tyXJ.gif',
			status: 1,
			shipping_fee: 15
		},
		{
			store_id: 3,
			description: 'moooooooooooooos',
			name: 'mos burger',
			address: 'NTUT',
			rate: 4,
			rate_count: 123,
			picture: 'https://i.imgur.com/3i3tyXJ.gif',
			status: 1,
			shipping_fee: 55
		}
	];

	onMount(async () => {
		const resp = await fetch(backendPath + '/stores');
		shopListResponse = await resp.json();
		console.log(shopListResponse);
	});
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
			<div class="flex flex-col items-center gap-5">
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
		{#each shopListResponse as shop}
			<StoreCard
				name={shop.name}
				description={shop.description}
				picture={shop.picture}
				address={shop.address}
				rate={shop.rate}
			/>
		{/each}
	</div>
</div>
