<script lang="ts">
	import { onMount } from 'svelte';

	let inputValue = '';
	export let width: string = 'null';
	export let min_width: string = 'null';
	export let max_Width: string = 'null';
	export let id: string = 'null';
	export let text_size: string = '14';

	let max = Number(max_Width);

	onMount(() => {
		adjustInputWidth();
	});

	function handleInput(event: Event) {
		inputValue = (event.target as HTMLInputElement).value;
		adjustInputWidth();
	}

	function adjustInputWidth() {
		const input = document.getElementById(id);
		if (input) {
			const minWidth = Number(min_width);
			const textWidth = getTextWidth(inputValue, text_size);
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

<input
	bind:value={inputValue}
	on:input={handleInput}
	style={`min-width: ${min_width}px; width: ${width}; overflow: hidden;`}
	placeholder="Enter label"
	{id}
/>

<style>
	input {
		padding: 0;
		background: transparent;
		border: none;
		border-bottom: 1px solid gray;
		outline: none;
		overflow: hidden;
	}
</style>
