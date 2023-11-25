<script lang="ts">
	let inputValue = '';

	export let width: string = 'null';

	function handleInput(event: Event) {
		const target = event.target as HTMLTextAreaElement;
		inputValue = target.value;
		adjustTextareaHeight();
	}

	function adjustTextareaHeight() {
		const textarea = document.getElementById('svelteTextarea') as HTMLTextAreaElement;
		if (textarea) {
			textarea.style.height = '0px';
			textarea.style.height = `${textarea.scrollHeight}px`;
		}
	}
</script>

<textarea
	bind:value={inputValue}
	on:input={handleInput}
	style={`width: ${width}px`}
	placeholder="Enter text"
	id="svelteTextarea"
></textarea>

<style>
	textarea {
		/* width: 200px;
		min-height: 30px; */
		height: 35px;
		font-size: 20px;
		padding: 0;
		background: transparent;
		border: none;
		border-bottom: 1px solid gray; /* Add a solid gray border at the bottom */
		outline: none; /* Remove default textarea outline */
		overflow: hidden;
	}
</style>
