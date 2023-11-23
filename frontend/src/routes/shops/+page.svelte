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
	import StoreCard from '$lib/components/ui/storeCard/storeCard.svelte';
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

<div class="flex justify-start bg-gradient-to-b from-zinc-500 to-neutral-100 shadow">
	<div class="float-left w-[30%] pt-[15px]">
		<Input type="text" placeholder="QQㄋㄟㄋㄟ好喝到咩噗茶" class="max-w-xs" />

		<p class="h-[30px] w-[325px] text-center text-lg">Sort</p>
		<div>
			{#each sortCheckboxes as { id }}
				<div class="h-[30px]">
					<Checkbox {id} />
					<Label for={id}>{id}</Label>
				</div>
			{/each}
		</div>

		<p class="h-[30px] w-[325px] text-center text-lg">Price</p>
		<div class="flex w-[325px] justify-between">
			<div>NT$1</div>
			<div>NT$700</div>
			<div>NT$1000</div>
		</div>
		<div class="h-[30px] w-[325px]">
			<DoubleRangeSlider bind:start bind:end />
		</div>

		<p class="h-[30px] w-[325px] text-center text-lg">Tag</p>
		<div>
			<div class="max-w-60 pr-40">
				{#each tagCheckboxes as { id }}
					<div class="flex-wrap">
						<Checkbox {id} />
						<Label for={id}>{id}</Label>
					</div>
				{/each}
			</div>
		</div>
	</div>

	<div class="grid gap-x-25 grid-cols-2 w-[900px] justify-start items-start">
		{#each shopListResponse as shop}
			<div class="px-6 py-3 rounded-[10px] flex-col justify-start items-start gap-3 inline-flex">
				<StoreCard
					name={shop.name}
					description={shop.description}
					picture={shop.picture}
					address={shop.address}
				/>
			</div>
		{/each}
	</div>
</div>

<h1>here is shop list</h1>
<p>Visit <a href="{$page.route.id}/impasta" class="text-blue-400"> Shop </a></p>
