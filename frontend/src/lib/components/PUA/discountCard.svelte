<script lang="ts">
	import Mouseon from '$lib/assets/mouseon.svg';
	import Adddiscount from '$lib/assets/adddiscount.svg';
	export let text = 'null';
	export let type = 'null';
	export let kind = 'null';
	export let how = 'null';
	export let way = 'null';
</script>

{#if type === 'null'}
	<button
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
	<button>
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
				<div class="w-2/5 p-3 text-center text-base font-bold">{kind}</div>
				<div class="ml-1 mr-1 border-r-[2px] border-PUA-stone p-3"></div>
				<div
					class="flex w-3/5 flex-wrap items-center justify-center p-3 text-center text-base font-bold"
				>
					<p>{how}</p>
					<p class="ml-1 mr-1 text-xl">{text}</p>
					<p>{way}</p>
				</div>
			</div>
		</div>
	</button>
{/if}
