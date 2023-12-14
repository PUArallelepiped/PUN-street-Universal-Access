<script lang="ts">
	type clickFunction = () => null;
	type clickFunctionAsync = () => Promise<void>;
	export let onclick: clickFunction | clickFunctionAsync;
	export let text: string;
	export let disabled: boolean;
</script>

<button
	class=" bg-PUA-orange w-52 rounded-3xl px-2 text-center text-lg font-semibold text-white disabled:opacity-50 disabled:cursor-not-allowed"
	on:click={onclick}
	disabled={disabled}
>
	{text}
</button>
