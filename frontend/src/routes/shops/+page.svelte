<script lang="ts">
	// import { page } from '$app/stores';

	import { Input } from '$lib/components/ui/input';

	import StoreCard from '$lib/components/PUA/storeCard.svelte';
	import SortTag from '$lib/components/PUA/sortTag.svelte';
	import { CheckBox } from '$lib';
	import type { PageData } from './$types';
	import DualRangeSlider from '$lib/components/PUA/dualRangeSlider.svelte';
	let start = 0;
	let end = 1;

	export let data: PageData;
</script>

<div class="w-1/2"></div>
<div class="flex justify-start">
	<div class="float-left m-5 flex flex-col items-center">
		<Input type="text" placeholder="QQㄋㄟㄋㄟ好喝到咩噗茶" class="m-5  max-w-xs bg-white" />
		<div class="flex w-96 flex-col justify-center gap-10 rounded-lg bg-white p-10">
			<div class="flex flex-col items-center gap-5">
				<SortTag value="Price"></SortTag>
				<div class="flex flex-col">
					<div class=" flex w-80 justify-between">
						<div class="text-sm font-bold leading-tight text-PUA-dark-red">
							NT$
							<span class="text-base leading-tight text-PUA-dark-red"
								>{Math.floor(start * 1000)}</span
							>
						</div>
						<div class="text-sm font-bold leading-tight text-PUA-dark-red">
							NT$
							<span class="text-base leading-tight text-PUA-dark-red">{Math.floor(end * 1000)}</span
							>
						</div>
					</div>
					<div class="  w-full px-10">
						<DualRangeSlider bind:start bind:end></DualRangeSlider>
					</div>
				</div>
			</div>
			<div class="flex flex-col items-center justify-center gap-5">
				<SortTag value="Tag"></SortTag>
				<div class="flex w-full max-w-xs flex-col justify-start gap-1 px-4">
					{#each data.categories as { category_id, category_name }}
						<CheckBox id={category_id.toString()} text={category_name} value={category_name}
						></CheckBox>
					{/each}
				</div>
			</div>
		</div>
	</div>

	<div class=" m-4 flex h-min flex-wrap gap-3">
		{#each data.shopListResponses as shop}
			<StoreCard
				name={shop.name}
				picture={shop.picture}
				rate={shop.rate}
				category_array={shop.category_array}
			/>
		{/each}
	</div>
</div>
