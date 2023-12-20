<script lang="ts">
	import DiscountButton from './discountButton.svelte';

	export let discount: {
		discount_max_quantity: number;
		product_id: number;
		discount_name: string;
		discount_description: string;
		discount_id: number;
		status: number;
	}[];
	export let toggleModel: () => null;
	export let addDiscountButton: (i: number) => null;
	export let addSign: boolean = true;

	export let type: boolean = false;

	$: last_index = discount.length;
</script>

<div class="relative h-full w-full">
	<div class="flex h-7 w-full items-center border-b-[1px] border-solid border-PUA-stone">
		<div class=" font-bold text-PUA-stone">Add Discount</div>
	</div>
	<div class="flex items-center justify-center">
		<div class="my-4 flex flex-wrap gap-1">
			{#each discount as { discount_id, discount_max_quantity }, index}
				<div class="flex justify-center">
					{#if type}
						<div class="">
							<DiscountButton
								text={'買' + { discount_max_quantity } + '送一'}
								id={discount_id.toString()}
							/>
						</div>
					{:else}
						<div class="">
							<button
								on:click={() => {
									toggleModel(), addDiscountButton(index);
								}}
								class="rounded-[10px] border-2 border-lime-800 px-2 py-0 font-bold text-lime-800 hover:bg-lime-800 hover:text-white active:bg-lime-800"
								>買{discount_max_quantity}送一</button
							>
						</div>
					{/if}
				</div>
			{/each}
			{#if addSign}
				<div class="flex items-center">
					<button
						on:click={addDiscountButton(last_index)}
						on:click={toggleModel}
						class="h-5 w-5 rounded-[10px] bg-lime-800 text-center text-[13px] font-bold text-white"
						>+</button
					>
				</div>
			{/if}
		</div>
	</div>
</div>
