<script lang="ts">
	import OkButton from './OkButton.svelte';
	import Star from '$lib/assets/Star.svg';
	import noStar from '$lib/assets/noStar.svg';
	import { onMount } from 'svelte';
	import { backendPath } from './env';
	import { goto } from '$app/navigation';
	import TagStar from './tag/tagStar.svelte';
	export let avgRate: number = 4.87;
	export let shopName: string;
	export let date: string;
	export let cost: number;
	export let img: string;
	export let storeId: number;
	export let cartId: number;
	export let userID: number;
	let rate = 2;
	let rateList: boolean[] = [];
	onMount(() => {
		getRateList();
	});
	function getRateList() {
		rateList = [];
		for (let i = 0; i < 5; i++) {
			rateList.push(i < rate);
		}
		return rateList;
	}
	async function rateStore(index: number) {
		try {
			fetch(backendPath + '/store/' + storeId + '/rate', {
				method: 'POST',
				body: JSON.stringify({
					rate: index
				})
			});
			rate = index;
		} catch {
			rate = index;
		}

		getRateList();
		return null;
	}
</script>

<div class="flex gap-5 rounded-xl bg-white p-2">
	<div class="relative h-32 w-64">
		<img class="flex h-full w-full rounded-xl object-cover shadow-inner" src={img} alt="no pic" />

		<div class="absolute right-3 top-3 flex">
			<TagStar text={avgRate}></TagStar>
		</div>
	</div>
	<div class="flex w-96 flex-col">
		<div class=" text-PUA-dark-orange text-2xl font-bold leading-relaxed">{shopName}</div>
		<div class=" text-PUA-dark-orange text-base font-bold leading-relaxed">{date}</div>
		<div class=" text-PUA-dark-orange text-base font-bold leading-relaxed">NT${cost}</div>

		<div class="flex gap-2">
			<div class="text-PUA-dark-red text-2xl font-semibold leading-normal">Rate store</div>
			{#if rate >= 0}
				{#each rateList as r, index}
					<button
						on:click={() => {
							rateStore(index + 1);
						}}
					>
						{#if r}
							<img src={Star} alt="" class="h-7 w-7" />
						{:else}
							<img src={noStar} alt="" class="h-7 w-7" />
						{/if}
					</button>
				{/each}
			{/if}
		</div>
	</div>
	<div class="flex items-center justify-center">
		<OkButton
			onclick={() => goto('/history/detail/' + userID + '/' + storeId + '/' + cartId)}
			text="Detail"
		></OkButton>
	</div>
</div>
