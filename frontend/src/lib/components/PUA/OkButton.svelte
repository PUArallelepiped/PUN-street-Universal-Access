<script lang="ts">
	type clickFunction = () => null;
	export let onclick: clickFunction;
	export let text: string;
</script>

<button
	class=" w-52 rounded-3xl bg-PUA-orange px-2 text-center text-lg font-semibold text-white"
	on:click={onclick()}>{text}</button
>
