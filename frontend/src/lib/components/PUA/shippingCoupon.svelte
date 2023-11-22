<script lang="ts">
	async function changeColor(reverse: boolean) {
		let bgColor = 'PUA-gray';
		let textColor = 'PUA-stone';
		if (reverse) {
			bgColor = [textColor, (textColor = bgColor)][0];
		}
		console.log(textColor);
		return { textColor, bgColor };
	}
	export let used: boolean = true;
</script>

{#await changeColor(used)}
	loading
	<p class="text-PUA-stone border-PUA-stone bg-PUA-gray w-0"></p>
	<p class="text-PUA-gray border-PUA-gray bg-PUA-stone w-0"></p>
{:then color}
	<div class="flex flex-col">
		<div class="text-PUA-stone text-xl font-semibold">IM pasta</div>
		<div
			class="text-{color?.textColor} bg-{color?.bgColor} border-{color?.textColor} flex w-96 items-center rounded-xl border-2 text-center font-semibold"
		>
			<div class=" px-6 py-5 text-base">Shipping Discount</div>
			<div class="bg-{color?.textColor} h-12 w-1"></div>
			<div class="w-full px-11">
				<div><span class="text-base">NT$</span><span class="text-2xl"> 1000</span></div>
				<div class="text-xl">free shipping</div>
			</div>
		</div>
	</div>
{/await}
