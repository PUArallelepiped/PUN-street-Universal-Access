<script lang="ts">
	import start from '$lib/assets/start.svg';
	export let text = 'null';
	export let type = 'text';

	let width: string = '50';
	let min_width: string = '0';
	let max_Width: string = '150';
	export let id: string = 'null';
	let text_size: string = '14';

	let max = Number(max_Width);

	function handleInput(event: Event) {
		text = (event.target as HTMLInputElement).value;
		adjustInputWidth();
	}

	function adjustInputWidth() {
		const input = document.getElementById(id);
		if (input) {
			const minWidth = Number(min_width);
			const textWidth = getTextWidth(text, text_size);
			let newWidth = Math.max(minWidth, textWidth + 10);
			newWidth = Math.min(newWidth, max);
			input.style.width = `${newWidth}px`;
		}
	}

	function getTextWidth(text: string, font: string): number {
		const span = document.createElement('span');
		span.style.font = font;
		span.style.visibility = 'hidden';
		span.style.position = 'absolute';
		span.style.whiteSpace = 'nowrap';
		span.innerHTML = text;

		document.body.appendChild(span);
		const width = span.getBoundingClientRect().width;
		document.body.removeChild(span);

		return width;
	}
</script>

{#if type === 'text'}
	<div
		class="flex h-7 items-center justify-center rounded-[20px] bg-PUA-dark-red px-4 py-0 text-center text-base font-bold text-white"
	>
		{text}
	</div>
{:else if type === 'add'}
	<button
		on:click
		class="flex h-7 w-7 items-center justify-center rounded-full bg-PUA-dark-red text-center text-xl font-bold text-white"
	>
		+
	</button>
{:else if type == 'star'}
	<div
		class=" flex h-7 w-16 items-center justify-center gap-1 rounded-[10px] border-4 border-red-200/25 bg-PUA-dark-red px-3 text-center text-base font-bold text-white"
	>
		<img src={start} alt="" />
		{text}
	</div>
{:else if type == 'input'}
	<div
		class="flex h-fit w-fit items-center justify-center rounded-full bg-PUA-dark-red px-4 text-white"
	>
		<input
			bind:value={text}
			on:input={handleInput}
			on:keydown
			class=" w-${width}  h-7 overflow-hidden bg-transparent text-white underline outline-none placeholder:text-white"
			placeholder="Enter Hashtag"
			{id}
		/>
	</div>
{/if}
