<script lang="ts">
	import watermelon from '$lib/assets/watermelon.png';
	import { BuyNforOneFree } from '$lib';
	export let href: string = './';
	export let description: string = 'the shopkeeper is lazy';
	export let price: number = 123;
	export let name: string = 'TEA EGG';
	export let imgUrl: string = watermelon;
	export let event_discount_array: { discount_max_quantity: number }[] = [
		{ discount_max_quantity: 0 }
	];
	export let status = 1;
</script>

<a
	{href}
	class:bg-gray-300={status == 2}
	class:pointer-events-none={status == 2}
	class:bg-white={status != 2}
	class:hover:scale-105={status != 2}
	class="flex items-center justify-between rounded-3xl duration-150"
>
	<div class="m-4 flex h-24 items-center">
		<div class="h-24 w-24 overflow-hidden rounded-lg">
			<img src={imgUrl} alt="" class="h-full object-cover object-center" />
		</div>
		<div class="mx-5 flex h-full flex-col content-start gap-1">
			<div
				class="flex flex-col content-start text-xl md:flex-row md:items-center md:text-2xl lg:flex-row lg:items-center lg:text-2xl"
			>
				<div class="mr-1 font-semibold text-PUA-dark-orange">{name}</div>
				<div class="md:flew-row flex flex-col sm:flex-row lg:flex-row">
					{#each event_discount_array as { discount_max_quantity }}
						<BuyNforOneFree quantity={discount_max_quantity}></BuyNforOneFree>
					{/each}
					{#if status === 2}
						<div
							class="mx-2 rounded-lg border-2 border-PUA-red px-1 text-base font-semibold text-PUA-red"
						>
							no storage
						</div>
					{/if}
				</div>
			</div>
			<div class="whitespace-pre-line text-base font-normal text-PUA-dark-gray">
				{description}
			</div>
		</div>
	</div>
	<div class="mr-5 flex w-24 items-baseline justify-end text-sm font-semibold text-PUA-dark-orange">
		NT$ <p class="w-10 text-end text-xl">{price}</p>
	</div>
</a>
