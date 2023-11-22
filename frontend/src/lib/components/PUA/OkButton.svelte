<script lang="ts">
	type clickFunction = () => null;
	export let onclick: clickFunction;
</script>

<button
	class=" bg-PUA-orange rounded-3xl px-2 text-center text-xl font-semibold text-white"
	on:click={onclick()}><slot /></button
>
