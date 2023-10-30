<script lang="ts">
	import watermelon from '$lib/assets/watermelon.png';
	export let href: string = './';
	export let description: string = 'the shopkeeper is lazy';
	export let price: number = 123;
</script>

<a {href} class="m-4 flex items-center justify-between rounded-3xl bg-white">
	<div class="flex items-center">
		<img src={watermelon} alt="" class="m-4 h-24 w-24 rounded-xl" />
		<div class="p-3">
			<div class="m-2 flex items-center gap-2 text-2xl">
				<div class="">TEA EGG</div>
				<div class="rounded-lg border-2 border-green-900 px-2 text-green-900">buy n for n free</div>
			</div>
			<div class="text-gray-500">
				{description}
			</div>
		</div>
	</div>
	<div class="w-28">NT$ {price}</div>
</a>
