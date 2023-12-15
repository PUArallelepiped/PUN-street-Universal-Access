<script lang="ts">
	export let inputValue = '';

	export let width: string = '10';

	function handleInput(event: Event) {
		const target = event.target as HTMLTextAreaElement;
		inputValue = target.value;
		if (target) {
			target.style.height = '0px';
			target.style.height = `${target.scrollHeight}px`;
		}
	}
</script>

<textarea
	bind:value={inputValue}
	on:input={handleInput}
	class="w-{width}  h-7 overflow-hidden border-b-[1px] border-solid border-b-gray-400 bg-transparent text-xl outline-none"
	placeholder="Enter text"
	id="svelteTextarea"
></textarea>
