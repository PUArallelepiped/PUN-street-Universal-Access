<script lang="ts">
	import { onMount } from 'svelte';
	import ErrorMsg from './errorMsg.svelte';
	export let value = '';
	export let width: string = '10';
	export let required: boolean = false;

	function handleInput(event: Event) {
		const target = event.target as HTMLTextAreaElement;
		value = target.value;
		if (target) {
			target.style.height = '0px';
			target.style.height = `${target.scrollHeight}px`;
		}
	}

	onMount(() => {
		const textarea = document.getElementById('svelteTextarea') as HTMLTextAreaElement;
		if (textarea) {
			textarea.style.height = '0px';
			textarea.style.height = `${textarea.scrollHeight}px`;
		}
	});
</script>

<textarea
	{required}
	bind:value
	on:input={handleInput}
	class="w-{width} peer h-7 overflow-hidden border-b-[1px] border-solid border-b-gray-400 bg-transparent text-base outline-none"
	placeholder="Enter text"
	id="svelteTextarea"
></textarea>
<div class="invisible w-64 peer-invalid:visible">
	<ErrorMsg width={'30'} height={'30'} text={'CANNOT BE EMPTY'}></ErrorMsg>
</div>
