<script lang="ts">
	import { Input } from '$lib/components/ui/input';

	import StoreCard from '$lib/components/PUA/storeCard.svelte';
	import SortTag from '$lib/components/PUA/sortTag.svelte';
	import type { shopListResponse } from '$lib';
	import type { PageData } from './$types';
	import DualRangeSlider from '$lib/components/PUA/dualRangeSlider.svelte';
	import { deserialize } from '$app/forms';
	import type { ActionResult } from '@sveltejs/kit';
	let start = 0;
	let end = 1;
	let checkedTag: string[] = [];
	let searchString: string = '';

	export let data: PageData;
	let timer: number;
	let shopInfos = data.shopListResponses;
	async function handleSearchForm() {
		if (timer) clearTimeout(timer);
		timer = setTimeout(async () => {
			const form = document.getElementById('searchForm') as HTMLFormElement;
			let formData = new FormData(form);
			formData.append('start', start.toString());
			formData.append('end', end.toString());
			formData.append('checkedTag', checkedTag.toString());
			formData.append('searchString', searchString.toString());
			const resp = await fetch('?/search', {
				method: 'POST',
				body: formData
			});
			const actionResult: ActionResult = deserialize(await resp.text());
			if (actionResult.type === 'success') {
				shopInfos = actionResult.data as shopListResponse[];
			}
		}, 500);
	}
</script>

<div class="w-1/2"></div>
<div class="flex justify-start">
	<form action="?/search" id="searchForm" method="post" on:submit|preventDefault={handleSearchForm}>
		<div class="float-left m-5 flex flex-col items-center">
			<Input
				on:input={handleSearchForm}
				type="text"
				placeholder="QQㄋㄟㄋㄟ好喝到咩噗茶"
				class="m-5  max-w-xs bg-white"
				bind:value={searchString}
			/>
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
								<span class="text-base leading-tight text-PUA-dark-red"
									>{Math.floor(end * 1000)}</span
								>
							</div>
						</div>
						<div class="  w-full px-10">
							<DualRangeSlider bind:start bind:end on:changed={handleSearchForm}></DualRangeSlider>
						</div>
					</div>
				</div>
				<div class="flex flex-col items-center justify-center gap-5">
					<SortTag value="Tag"></SortTag>
					<div class="flex w-full max-w-xs flex-col justify-start gap-1 px-4">
						{#each data.categories as { category_id, category_name }}
							<div class="flex gap-3">
								<input
									bind:group={checkedTag}
									on:click={handleSearchForm}
									type="checkbox"
									id={category_id.toString()}
									name=""
									value={category_name + ':' + category_id.toString()}
									class="h-7 w-7 accent-PUA-dark-orange"
								/>
								<label for={category_id.toString()}>
									<div
										class="truncate text-center text-xl font-bold leading-relaxed text-PUA-dark-red"
									>
										{category_name}
									</div>
								</label>
							</div>
						{/each}
					</div>
				</div>
			</div>
		</div>
	</form>

	<div class=" m-4 flex h-min flex-wrap gap-3">
		{#each shopInfos as shop}
			<StoreCard
				name={shop.name}
				picture={shop.picture}
				rate={shop.rate}
				category_array={shop.category_array}
				id={shop.store_id}
			/>
		{/each}
	</div>
</div>
