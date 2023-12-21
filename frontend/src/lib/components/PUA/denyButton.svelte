<script lang="ts">
	type clickFunction = () => null;
	export let onclick: clickFunction;
</script>

<button
	class="border-PUA-stone text-PUA-stone w-52 rounded-3xl border-2 px-2 text-center text-xl font-semibold"
	on:click={onclick()}><slot /></button
>
