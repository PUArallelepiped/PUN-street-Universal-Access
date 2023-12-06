<script>
	import OkButton from './OkButton.svelte';
	import Star from '$lib/assets/Star.svg';
	import noStar from '$lib/assets/noStar.svg';
	let rate = 4;
	function getRateList() {
		let rateList = [];
		for (let i = 0; i < 5; i++) {
			rateList.push(i < rate);
		}
		return rateList;
	}
</script>

<div class="flex gap-5 rounded-xl bg-white p-2">
	<div class="relative h-32 w-64">
		<img class=" rounded-xl shadow-inner" src="https://via.placeholder.com/256x128" alt="no pic" />
		<div
			class="absolute right-3 top-3 flex items-center gap-1 rounded-lg border-2
			 border-white border-opacity-50 bg-PUA-dark-red
			 p-1 text-base font-semibold leading-tight text-white"
		>
			<img src={Star} alt="" class=" h-4 w-4" />
			4.87
		</div>
	</div>
	<div class="flex w-96 flex-col">
		<div class=" text-2xl font-bold leading-relaxed text-PUA-dark-orange">Shop name</div>
		<div class=" text-base font-bold leading-relaxed text-PUA-dark-orange">2020-10-11</div>
		<div class=" text-base font-bold leading-relaxed text-PUA-dark-orange">NT$999</div>

		<div class="flex gap-2">
			<div class="text-2xl font-semibold leading-normal text-PUA-dark-red">Rate store</div>
			{#each getRateList() as r}
				{#if r}
					<img src={Star} alt="" class="h-7 w-7" />
				{:else}
					<img src={noStar} alt="" class="h-7 w-7" />
				{/if}
			{/each}
		</div>
	</div>
	<div class="flex items-center justify-center">
		<OkButton onclick={() => null} text="Detail"></OkButton>
	</div>
</div>
