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
	export let group = 1;

	$: last_index = discount.length;
</script>

<div class="relative h-full w-full">
	<div class="border-PUA-stone flex h-7 w-full items-center border-b-[1px] border-solid">
		<div class=" text-PUA-stone font-bold">Add Discount</div>
	</div>
	{#if discount.length === 0 && type}
		<div class=" my-4 flex w-full animate-pulse items-center justify-center">
			<p class=" border-PUA-stone text-PUA-stone rounded-lg border-2 px-4 text-base font-bold">
				No Event Discount
			</p>
		</div>
	{/if}

	<div class="flex items-center justify-center">
		<div class="my-4 flex flex-wrap gap-1">
			{#each discount as { discount_id, discount_max_quantity }, index}
				<div class="flex justify-center">
					{#if type}
						<div class="">
							<DiscountButton
								id={discount_id}
								bind:value={discount_id}
								bind:group
								text={'買' + discount_max_quantity.toString() + '送一'}
							/>
						</div>
					{:else}
						<div class="">
							<button
								type="button"
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
						type="button"
						on:click={addDiscountButton(last_index)}
						on:click={toggleModel}
						class="flex h-5 w-5 items-center justify-center rounded-full bg-lime-800 text-center text-[13px] text-xs font-bold text-white duration-150 hover:scale-125 hover:border-2 hover:border-lime-800 hover:bg-transparent hover:text-lime-800"
						>+</button
					>
				</div>
			{/if}
		</div>
	</div>
</div>
