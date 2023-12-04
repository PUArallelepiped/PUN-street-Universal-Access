<script lang="ts">
	import start from '$lib/assets/start.svg';
	export let text = 'null';
	export let type = 'text';
</script>

{#if type === 'text'}
	<div
		class="bg-PUA-dark-red justify-center rounded-[20px] px-5 py-1 text-center text-xl font-bold text-white"
	>
		{text}
	</div>
{:else if type === 'add'}
	<button
		class="bg-PUA-dark-red h-9 w-9 items-center justify-center rounded-[18px] text-center text-3xl font-bold text-white"
	>
		+
	</button>
{:else if type == 'start'}
	<div
		class="bg-PUA-dark-red flex items-baseline justify-center gap-1 rounded-[10px] border-4 border-amber-700/75 px-3 text-center text-xl font-bold text-white"
	>
		<img src={start} alt="" />
		{text}
	</div>
{/if}
