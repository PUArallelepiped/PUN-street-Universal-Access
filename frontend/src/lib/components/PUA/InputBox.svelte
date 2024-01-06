<script lang="ts">
	export let label: string;
	export let value: string | number | null = '';
	export let type: '' | 'input' | 'password' | 'date' = 'input';
	export let onInput: () => void = () => {};
</script>

<div class="">
	<div class="text-left text-xl font-bold leading-relaxed text-orange-950">{label}</div>
	{#if type == 'input' || type == ''}
		<input
			class="h-10 w-96 rounded-lg border-2 bg-gray-200 px-2 text-xl font-medium leading-relaxed text-orange-950 shadow-inner"
			type="input"
			bind:value
			on:input={onInput}
		/>
	{:else if type == 'password'}
		<input
			class="h-10 w-96 rounded-lg border-2 bg-gray-200 px-2 text-xl font-medium leading-relaxed text-orange-950 shadow-inner"
			type="password"
			bind:value
			on:input={onInput}
		/>
	{:else if type == 'date'}
		<input
			class="h-10 w-96 rounded-lg border-2 bg-gray-200 px-2 text-xl font-medium leading-relaxed text-orange-950 shadow-inner"
			type="date"
			bind:value
			on:input={onInput}
		/>
	{/if}
</div>
