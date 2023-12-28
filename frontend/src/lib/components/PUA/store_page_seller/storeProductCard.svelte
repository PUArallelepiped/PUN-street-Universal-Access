<script lang="ts">
	import transhcan from '$lib/assets/transhcan.svg';
	import watermelon from '$lib/assets/watermelon.png';
	import { BuyNforOneFree } from '$lib';
	export let description: string = 'the shopkeeper is lazy';
	export let price: number = 123;
	export let name: string = 'TEA EGG';
	export let imgUrl: string = watermelon;
	export let href: string = './';
	export let event_discount_array: { discount_max_quantity: number }[] = [
		{ discount_max_quantity: 0 }
	];
</script>

<a
	{href}
	class="flex items-center justify-between overflow-hidden rounded-3xl bg-white duration-150 hover:scale-[1.02]"
>
	<div class="m-4 flex h-24 items-center">
		<div class="h-24 w-24 overflow-hidden rounded-lg">
			<img src={imgUrl} alt="" class="h-full object-cover object-center" />
		</div>
		<div class="mx-5 flex h-full flex-col content-start gap-1">
			<div
				class="flex flex-col content-start text-xl md:flex-row md:items-center md:text-2xl lg:flex-row lg:items-center lg:text-2xl"
			>
				<div class="text-PUA-dark-orange mr-1 font-semibold">{name}</div>
				<div class="md:flew-row flex flex-col sm:flex-row lg:flex-row">
					{#each event_discount_array as { discount_max_quantity }}
						<BuyNforOneFree quantity={discount_max_quantity}></BuyNforOneFree>
					{/each}
				</div>
			</div>
			<div class="text-PUA-dark-gray whitespace-pre-line text-base font-normal">
				{description}
			</div>
		</div>
	</div>
	<div
		class="text-PUA-dark-orange mr-5 flex w-40 items-center justify-end gap-8 text-sm font-semibold"
	>
		<div class="flex items-center">
			<p>NT$</p>

			<p class="w-10 text-end text-xl">{price}</p>
		</div>
		<button
			on:click|preventDefault
			class="w-28 transition-transform duration-200 hover:scale-125 active:scale-90"
		>
			<img src={transhcan} width="21" height="24" alt="" />
		</button>
	</div>
</a>
