<script lang="ts">
	import Dropdown from '$lib/assets/dropdown.svg';
	export let text = 'null';
	export let img_need = false;
	export let dropdown = true;
</script>

<div class="boder-b-1 mb-4 mt-4 flex justify-between border-b-[1px] border-b-PUA-stone">
	<div class="mx-5 text-xl font-bold text-PUA-stone">{text}</div>
	{#if img_need}
		<button on:click>
			<img
				src={Dropdown}
				alt=""
				class={` ${
					dropdown ? 'rotate-0  focus:rotate-0' : 'rotate-180 focus:rotate-180'
				}  transform transition-transform duration-300 ease-in  focus:outline-none`}
			/>
		</button>
	{/if}
</div>
