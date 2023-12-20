<script lang="ts">
	import OkButton from './OkButton.svelte';
	import Star from '$lib/assets/Star.svg';
	import noStar from '$lib/assets/noStar.svg';
	export let avgRate: number = 4.87;
	export let shopName: string;
	export let date: string;
	export let cost: number;
	let rate = 2;
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
			<div>
				<svg
					width="26"
					height="24"
					viewBox="0 0 26 24"
					fill="none"
					xmlns="http://www.w3.org/2000/svg"
					class="fill-yellow-400"
				>
					<g id="Star 6">
						<path
							d="M12.0489 0.927052C12.3483 0.00574112 13.6517 0.00573993 13.9511 0.927051L16.1432 7.67376C16.2771 8.08578 16.661 8.36475 17.0943 8.36475H24.1882C25.1569 8.36475 25.5597 9.60436 24.7759 10.1738L19.0369 14.3435C18.6864 14.5981 18.5397 15.0495 18.6736 15.4615L20.8657 22.2082C21.1651 23.1295 20.1106 23.8956 19.3269 23.3262L13.5878 19.1565C13.2373 18.9019 12.7627 18.9019 12.4122 19.1565L6.67312 23.3262C5.88941 23.8956 4.83493 23.1295 5.13428 22.2082L7.32642 15.4615C7.46029 15.0495 7.31363 14.5981 6.96315 14.3435L1.22405 10.1738C0.440337 9.60436 0.843112 8.36475 1.81184 8.36475H8.90575C9.33897 8.36475 9.72293 8.08578 9.8568 7.67376L12.0489 0.927052Z"
							fill="#C65633"
						/>
						<path
							d="M13 1.23607L15.1921 7.98278C15.4599 8.80682 16.2278 9.36475 17.0943 9.36475H24.1882L18.4491 13.5344C17.7481 14.0437 17.4548 14.9465 17.7225 15.7705L19.9147 22.5172L14.1756 18.3475C13.4746 17.8382 12.5254 17.8382 11.8244 18.3475L6.08533 22.5172L8.27747 15.7705C8.54522 14.9465 8.25191 14.0437 7.55093 13.5344L1.81184 9.36475H8.90575C9.7722 9.36475 10.5401 8.80683 10.8079 7.98278L13 1.23607Z"
							stroke="#C65633"
							stroke-opacity="0.5"
							stroke-width="2"
						/>
					</g>
				</svg>
			</div>
			{avgRate}
		</div>
	</div>
	<div class="flex w-96 flex-col">
		<div class=" text-2xl font-bold leading-relaxed text-PUA-dark-orange">{shopName}</div>
		<div class=" text-base font-bold leading-relaxed text-PUA-dark-orange">{date}</div>
		<div class=" text-base font-bold leading-relaxed text-PUA-dark-orange">NT${cost}</div>

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
