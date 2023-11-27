<script lang="ts">
	import DiscountButton from './discountButton.svelte';

	export let discount: {
		id: number;
		label: string;
	}[];
	export let toggleModal: () => null;
	export let addDiscountButton: () => null;
	export let addSign: boolean = true;

	export let type: boolean = false;
</script>

<div class="relative h-full w-full">
	<div class="flex h-[30px] w-full items-center border-b-[1px] border-solid border-PUA-stone">
		<div class=" font-bold text-PUA-stone">Add Discount</div>
	</div>
	<div class="flex items-center justify-center">
		<div class="flex flex-wrap">
			{#each discount as { label }}
				<div class="flex justify-center">
					{#if type}
						<div class="mb-2 ml-1 mt-2">
							<DiscountButton text={label} />
						</div>
					{:else}
						<button
							on:click={toggleModal}
							class="mb-2 ml-1 mt-2 rounded-[10px] border-2 border-lime-800 p-5 px-2 py-0 font-bold text-lime-800 hover:bg-lime-800 hover:text-white active:bg-lime-800"
							>{label}</button
						>
					{/if}
				</div>
			{/each}
			{#if addSign}
				<div class="flex w-[90px] items-center">
					<button
						on:click={addDiscountButton}
						on:click={toggleModal}
						class="ml-1 h-[20px] w-[20px] rounded-[10px] bg-lime-800 text-center text-[13px] font-bold text-white"
						>+</button
					>
				</div>
			{/if}
		</div>
	</div>
</div>
