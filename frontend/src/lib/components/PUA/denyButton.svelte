<script lang="ts">
	type clickFunction = () => null;
	export let onclick: clickFunction;
</script>

<button
	class="rounded-3xl border-2 border-PUA-stone px-2 text-center text-xl font-semibold text-PUA-stone"
	on:click={onclick()}><slot /></button
>
