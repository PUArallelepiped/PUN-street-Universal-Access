<script lang="ts">
	export let text = 'null';
	export let id: string = 'null';

	let width: string = '50';
	let min_width: string = '0';
	let max_Width: string = '150';
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

<div
	class="flex h-fit w-fit items-center justify-center rounded-full bg-PUA-dark-red px-4 text-white"
>
	<input
		bind:value={text}
		on:input={handleInput}
		on:keydown
		maxlength="20"
		class=" w-${width} h-7 overflow-hidden bg-transparent text-white underline outline-none placeholder:text-white"
		placeholder="Enter Hashtag"
		{id}
	/>
</div>
