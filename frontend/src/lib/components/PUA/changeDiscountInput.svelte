<script lang="ts">
	import { ErrorMsg } from '$lib/index';
	export let title: string;
	export let value: string;
	export let text: string;
	export let error = true;
</script>

<div class="mt-2">
	<p class="text-3 font-bold text-red-900">{title}</p>
	<input
		type="text"
		class="w-full border-b-[2px] border-red-900 font-bold"
		bind:value
		placeholder={text}
	/>
	{#if error}
		<ErrorMsg width={'24'} height={'24'}></ErrorMsg>
	{/if}
</div>
