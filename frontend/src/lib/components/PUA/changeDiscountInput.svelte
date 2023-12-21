<script lang="ts">
	import { ErrorMsg } from '$lib/index';
	export let title: string;
	export let value: string | number;
	export let type = 'text';
	export let text: string;
	export let name = '';
</script>

<div class="mt-2">
	<p class="text-3 text-PUA-stone font-bold">{title}</p>
	{#if type === 'text'}
		<input
			required={true}
			{name}
			type="text"
			class="... border-PUA-stone peer w-full border-b-[2px] font-bold"
			bind:value
			placeholder={text}
		/>
		<div class="invisible peer-invalid:visible">
			<ErrorMsg width={'24'} height={'24'} text={'CANNOT BE EMPTY'}></ErrorMsg>
		</div>
	{:else if type === 'number'}
		<input
			required={true}
			{name}
			min="1"
			max="99999"
			type="number"
			class="border-PUA-stone peer w-full border-b-[2px] font-bold"
			bind:value
			placeholder={text}
		/>
		<div class="invisible h-6 peer-invalid:visible">
			{#if !value}
				<ErrorMsg width={'24'} height={'24'} text={'CANNOT BE EMPTY'}></ErrorMsg>
			{:else if Number(value) > 99999 || Number(value) < 0}
				<ErrorMsg width={'24'} height={'24'} text={'EXCEEDS LIMIT'}></ErrorMsg>
			{/if}
		</div>
	{/if}
</div>
