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

<div class="flex justify-start bg-gray-200">
	<div class="w-30 float-left mt-[15px]">
		<Input type="text" placeholder="QQㄋㄟㄋㄟ好喝到咩噗茶" class="my-3 ml-3 max-w-xs bg-white" />
		<div
			class="my-3 inline-flex h-[25px] items-center justify-center gap-2.5 rounded-full bg-PUA-red px-[23px] py-1"
		>
			<div
				class="mt-2 h-[43px] w-[300px] text-center font-['Inter'] text-xl font-bold leading-relaxed text-white"
			>
				Sort
			</div>
		</div>
		<div>
			{#each sortCheckboxes as { id }}
				<div class="ml-2 h-[30px]">
					<Checkbox {id} />
					<Label for={id}>{id}</Label>
				</div>
			{/each}
		</div>

		<div
			class="my-3 inline-flex h-[25px] items-center justify-center gap-2.5 rounded-full bg-PUA-red px-[23px] py-1"
		>
			<div
				class="mt-2 h-[43px] w-[300px] text-center font-['Inter'] text-xl font-bold leading-relaxed text-white"
			>
				Price
			</div>
		</div>
		<div class="ml-3 flex w-[350px] justify-between">
			<div>NT$1</div>
			<div>NT$1000</div>
		</div>
		<div class="my-3 ml-3 h-[30px] w-[325px]">
			<DoubleRangeSlider bind:start bind:end />
		</div>

		<div
			class="my-3 inline-flex h-[25px] items-center justify-center gap-2.5 rounded-full bg-PUA-red px-[23px] py-1"
		>
			<div
				class="mt-2 h-[43px] w-[300px] text-center font-['Inter'] text-xl font-bold leading-relaxed text-white"
			>
				Tag
			</div>
		</div>
		<div>
			<div class="max-w-60 pr-40">
				{#each tagCheckboxes as { id }}
					<div class="ml-2 flex-wrap">
						<Checkbox {id} />
						<Label for={id}>{id}</Label>
					</div>
				{/each}
			</div>
		</div>
	</div>

	<div class="gap-x-25 ml-20 grid w-[900px] grid-cols-2 items-start justify-start">
		{#each shopListResponse as shop}
			<div
				class="inline-flex flex-col items-start justify-start gap-3 rounded px-6 py-3 hover:bg-gray-300"
			>
				<StoreCard
					name={shop.name}
					description={shop.description}
					picture={shop.picture}
					address={shop.address}
					rate={shop.rate}
				/>
			</div>
		{/each}
	</div>
</div>

<h1>here is shop list</h1>
<p>Visit <a href="{$page.route.id}/impasta" class="text-blue-400"> Shop </a></p>
