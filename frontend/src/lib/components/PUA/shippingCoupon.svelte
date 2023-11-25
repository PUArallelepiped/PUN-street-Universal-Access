<script lang="ts">
	async function changeColor(reverse: boolean) {
		let bgColor = 'PUA-gray';
		let textColor = 'PUA-stone';
		if (reverse) {
			bgColor = [textColor, (textColor = bgColor)][0];
		}
		return { textColor, bgColor };
	}
	export let used: boolean = true;
</script>

{#await changeColor(used)}
	loading
	<p class="w-0 border-PUA-stone bg-PUA-gray text-PUA-stone"></p>
	<p class="w-0 border-PUA-gray bg-PUA-stone text-PUA-gray"></p>
{:then color}
	<div class="flex w-96 flex-col">
		<div class="py-1 text-xl font-semibold text-PUA-stone">IM pasta</div>
		<div
			class="text-{color?.textColor} bg-{color?.bgColor} border-{color?.textColor} flex items-center rounded-xl border-2 text-center font-semibold"
		>
			<div class="px-6 py-5 text-base font-semibold">Shipping Discount</div>
			<div class="bg-{color?.textColor} h-12 w-1"></div>
			<div class="w-full px-11">
				<div><span class="text-base">NT$</span><span class="text-2xl"> 1000</span></div>
				<div class="text-xl">free shipping</div>
			</div>
		</div>
	</div>
{/await}
