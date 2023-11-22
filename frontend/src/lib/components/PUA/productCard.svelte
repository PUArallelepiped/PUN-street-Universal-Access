<script lang="ts">
	import watermelon from '$lib/assets/watermelon.png';
	export let href: string = './';
	export let description: string = 'the shopkeeper is lazy';
	export let price: number = 123;
	export let name: string = 'TEA EGG';
	export let imgUrl: string = watermelon;
</script>

<a {href} class="flex items-center justify-between rounded-3xl bg-white">
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
					<div
						class="border-PUA-green text-PUA-green mx-2 rounded-lg border-2 px-1 text-base font-semibold"
					>
						buy n for n free
					</div>
					<div
						class="border-PUA-red text-PUA-red mx-2 rounded-lg border-2 px-1 text-base font-semibold"
					>
						no storange
					</div>
				</div>
			</div>
			<div class="text-PUA-dark-gray whitespace-pre-line text-base font-normal">
				{description}
			</div>
		</div>
	</div>
	<div class="text-PUA-dark-orange mr-5 flex w-24 items-baseline justify-end text-sm font-semibold">
		NT$ <p class="w-10 text-end text-xl">{price}</p>
	</div>
</a>
