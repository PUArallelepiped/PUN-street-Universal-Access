<script lang="ts">
	import Mouseon from '$lib/assets/mouseon.svg';
	import Adddiscount from '$lib/assets/adddiscount.svg';
	import Transhcan from '$lib/assets/transhcan.svg';

	export let deleteDiscountCard: () => void;
	export let discountCardData: {
		discount_name: string;
		discount_description: string;
		discount_max_quantity: number;
		discount_id: number;
		status: number;
	} = {
		discount_name: 'free shipping',
		discount_description: 'free shipping when total price over 1000',
		discount_max_quantity: 1000,
		discount_id: 1,
		status: 1
	};
	export let dis_haved: boolean = false;
</script>

{#if !dis_haved}
	<button on:click class=" text-PUA-stone left-0 top-0 flex h-24 w-72 rounded-xl bg-gray-300">
		<div class="left-0 top-0 z-10 flex h-full w-72 items-center justify-center rounded-[8px]">
			<img
				src={Adddiscount}
				alt=""
				class="h-6 w-6 bg-transparent transition-opacity duration-300"
			/>
			<p class="text-PUA-stone p-2 font-bold">Add Discount</p>
		</div>
	</button>
{:else}
	<div class="flex items-center gap-4">
		<button on:click>
			<div class="text-PUA-stone left-0 top-0 flex h-24 w-80 rounded-xl">
				<div class="relative">
					<div
						class=" border-PUA-stone group absolute left-0 top-0 z-10 flex h-full w-80 items-center justify-center rounded-xl border-4 hover:bg-neutral-400/75"
					>
						<img
							src={Mouseon}
							alt=""
							class="h-6 w-6 bg-transparent opacity-0 transition-opacity duration-300 group-hover:opacity-100"
						/>
						<p
							class="text-PUA-dark-red p-2 font-bold opacity-0 transition-opacity duration-300 group-hover:opacity-100"
						>
							Change Discount
						</p>
					</div>
				</div>

				<div
					class="bg-transport border-PUA-stone text-PUA-stone flex w-full items-center rounded-xl border-4 text-center font-semibold"
				>
					<div class=" px-4 py-4 text-base font-semibold">Shipping Discount</div>
					<div class="bg-PUA-stone h-12 w-1"></div>
					<div class="relative w-full px-4">
						<div class="flex items-baseline justify-center">
							<span class="ml-1 mr-1 text-base">NT$</span>
							<span class="text-2xl">{discountCardData.discount_max_quantity}</span>
						</div>
						<div class="whitespace-nowrap text-xl">free shipping</div>
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
