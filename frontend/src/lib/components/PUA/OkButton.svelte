<script lang="ts">
	type clickFunction = () => null;
	export let onclick: clickFunction;
</script>

<button
	class=" rounded-3xl bg-PUA-orange px-2 text-center text-xl font-semibold text-white"
	on:click={onclick()}><slot /></button
>
