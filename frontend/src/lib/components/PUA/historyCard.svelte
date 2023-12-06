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
	<div>
		<img
			class="h-32 w-64 rounded-xl shadow-inner"
			src="https://via.placeholder.com/256x128"
			alt="no pic"
		/>
	</div>
	<div class="flex w-96 flex-col">
		<div class=" text-2xl font-bold leading-relaxed text-PUA-dark-orange">Shop name</div>
		<div class=" text-base font-bold leading-relaxed text-PUA-dark-orange">2020-10-11</div>
		<div class=" text-base font-bold leading-relaxed text-PUA-dark-orange">NT$999</div>

		<div class="flex gap-2">
			<div class="text-2xl font-semibold leading-normal text-PUA-dark-red">Rate store</div>
			{#each getRateList() as r}
				{#if r}
					<img src={Star} alt="" />
				{:else}
					<img src={noStar} alt="" />
				{/if}
			{/each}
		</div>
	</div>
	<div class="flex items-center justify-center">
		<OkButton onclick={() => null} text="Detail"></OkButton>
	</div>
</div>
