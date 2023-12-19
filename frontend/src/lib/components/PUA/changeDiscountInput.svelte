<script lang="ts">
	import { ErrorMsg } from '$lib/index';
	export let title: string;
	export let value: string | number;
	export let type = 'text';
	export let text: string;
	export let name = '';
</script>

<div class="mt-2">
	<p class="text-3 font-bold text-PUA-stone">{title}</p>
	{#if type === 'text'}
		<input
			required
			{name}
			type="text"
			class="... peer w-full border-b-[2px] border-PUA-stone font-bold"
			bind:value
			placeholder={text}
		/>
	{:else if type === 'number'}
		<input
			required
			{name}
			min="1"
			max="99999"
			type="number"
			class="peer w-full border-b-[2px] border-PUA-stone font-bold"
			bind:value
			placeholder={text}
		/>
	{/if}
	<div class="invisible peer-invalid:visible">
		<ErrorMsg width={'24'} height={'24'}></ErrorMsg>
	</div>
</div>
