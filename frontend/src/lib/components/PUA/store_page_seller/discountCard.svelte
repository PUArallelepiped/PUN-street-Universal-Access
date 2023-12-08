<script lang="ts">
	import Mouseon from '$lib/assets/mouseon.svg';
	import Adddiscount from '$lib/assets/adddiscount.svg';
	import Transhcan from '$lib/assets/transhcan.svg';

	export let deleteDiscountCard: () => void;
	export let discountCardData: {
		haved: boolean;
		maxquantity: string;
		kind: string;
		how: string;
		way: string;
	} = {
		haved: false,
		maxquantity: 'null',
		kind: 'null',
		how: 'null',
		way: 'null'
	};
</script>

{#if !discountCardData.haved}
	<button
		on:click
		class="left-0 top-0 flex h-24 w-72 rounded-[10px] border-PUA-stone bg-gray-300 text-PUA-stone"
	>
		<div class="left-0 top-0 z-10 flex h-full w-[283px] items-center justify-center rounded-[8px]">
			<img
				src={Adddiscount}
				alt=""
				class="h-6 w-6 bg-transparent transition-opacity duration-300"
			/>
			<p class="p-2 font-bold text-PUA-stone">Add Discount</p>
		</div>
	</button>
{:else}
	<div class="flex items-center gap-4">
		<button on:click>
			<div
				class="left-0 top-0 flex h-24 w-72 rounded-[10px] border-[3px] border-PUA-stone text-PUA-stone"
			>
				<div class="relative bg-red-700">
					<div
						class=" group absolute left-0 top-0 z-10 flex h-full w-[283px] items-center justify-center rounded-[8px] hover:bg-neutral-400/75"
					>
						<img
							src={Mouseon}
							alt=""
							class="h-6 w-6 bg-transparent opacity-0 transition-opacity duration-300 group-hover:opacity-100"
						/>
						<p
							class="p-2 font-bold text-PUA-dark-red opacity-0 transition-opacity duration-300 group-hover:opacity-100"
						>
							Change Discount
						</p>
					</div>
				</div>
				<div class="flex p-2">
					<div class="w-2/5 p-3 text-center text-base font-bold">{discountCardData.kind}</div>
					<div class="ml-1 mr-1 border-r-[2px] border-PUA-stone p-3"></div>
					<div
						class="flex w-3/5 flex-wrap items-center justify-center p-3 text-center text-base font-bold"
					>
						<p>{discountCardData.how}</p>
						<p class="ml-1 mr-1 text-xl">{discountCardData.maxquantity}</p>
						<p>{discountCardData.way}</p>
					</div>
				</div>
			</div>
		</button>
		<button
			on:click={deleteDiscountCard}
			class=" rounded-lg bg-gray-200 px-6 py-3 text-black transition-transform duration-200 ease-in-out hover:scale-125 focus:animate-ping active:scale-90"
		>
			<img src={Transhcan} alt="" class="h-6 w-6" />
		</button>
	</div>
{/if}
