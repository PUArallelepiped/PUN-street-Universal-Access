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
	<div class="flex h-32 w-96 flex-col">
		<div
			class="text-{color?.textColor} bg-{color?.bgColor} border-{color?.textColor} flex h-full items-center rounded-xl border-2 text-center font-semibold"
		>
			<div class=" px-6 py-5 text-base">Seasoning Discount</div>
			<div class="bg-{color?.textColor} h-12 w-1"></div>
			<div class="w-full">
				<div class="text-3xl">80%</div>
				<div>2023/01/01~2023/12/12</div>
			</div>
		</div>
	</div>
{/await}
