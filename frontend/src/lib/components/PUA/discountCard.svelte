<script lang="ts">
	import Mouseon from '$lib/assets/mouseon.svg';
	// export let text = 'null';
	import Adddiscount from '$lib/assets/adddiscount.svg';
	export let type = 'null';
</script>

{#if type === 'null'}
	<button
		class="border-PUA-stone text-PUA-stone left-0 top-0 flex h-24 w-72 rounded-[10px] bg-gray-300"
	>
		<div class="left-0 top-0 z-10 flex h-full w-[283px] items-center justify-center rounded-[8px]">
			<img
				src={Adddiscount}
				alt=""
				class="h-6 w-6 bg-transparent transition-opacity duration-300"
			/>
			<p class="text-PUA-stone p-2 font-bold">Add Discount</p>
		</div>
	</button>
{:else}
	<button>
		<div
			class="border-PUA-stone text-PUA-stone left-0 top-0 flex h-24 w-72 rounded-[10px] border-[3px]"
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
						class="text-PUA-dark-red p-2 font-bold opacity-0 transition-opacity duration-300 group-hover:opacity-100"
					>
						Change Discount
					</p>
				</div>
			</div>
			<div class="flex p-2">
				<div class="w-2/5 p-3 text-center font-bold">Shipping Discount</div>
				<div class="border-PUA-stone ml-1 mr-1 border-r-[2px] p-3"></div>
				<div class="flex w-3/5 flex-wrap items-center justify-center p-3 text-center font-bold">
					<p>NT$</p>
					<p class="ml-1 mr-1 text-xl">1000</p>
					<p>free shipping</p>
				</div>
			</div>
		</div>
	</button>
{/if}
