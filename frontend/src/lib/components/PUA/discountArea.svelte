<script lang="ts">
	import DiscountButton from './discountButton.svelte';

	export let discount: {
		id: string;
		label: string;
	}[];
	export let toggleModel: () => null;
	export let addDiscountButton: () => null;
	export let addSign: boolean = true;

	export let type: boolean = false;
</script>

<div class="relative h-full w-full">
	<div class="flex h-7 w-full items-center border-b-[1px] border-solid border-PUA-stone">
		<div class=" font-bold text-PUA-stone">Add Discount</div>
	</div>
	<div class="flex items-center justify-center">
		<div class="my-4 flex flex-wrap gap-1">
			{#each discount as { id, label }}
				<div class="flex justify-center">
					{#if type}
						<div class="">
							<DiscountButton text={label} {id} />
						</div>
					{:else}
						<div class="">
							<button
								on:click={toggleModel}
								class="rounded-[10px] border-2 border-lime-800 px-2 py-0 font-bold text-lime-800 hover:bg-lime-800 hover:text-white active:bg-lime-800"
								>{label}</button
							>
						</div>
					{/if}
				</div>
			{/each}
			{#if addSign}
				<div class="flex items-center">
					<button
						on:click={addDiscountButton}
						on:click={toggleModel}
						class="h-5 w-5 rounded-[10px] bg-lime-800 text-center text-[13px] font-bold text-white"
						>+</button
					>
				</div>
			{/if}
		</div>
	</div>
</div>
