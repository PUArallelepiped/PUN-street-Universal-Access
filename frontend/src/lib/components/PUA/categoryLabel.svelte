<script lang="ts">
	import Dropdown from '$lib/assets/dropdown.svg';
	export let text = 'null';
	export let img_nedd = false;
</script>

<div class="boder-b-1 mb-4 mt-4 flex justify-between border-b-[1px] border-b-PUA-stone">
	<div class="mx-5 text-xl font-bold text-PUA-stone">{text}</div>
	{#if img_nedd}
		<img src={Dropdown} alt="" />
	{/if}
</div>
