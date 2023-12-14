<script lang="ts">
	type clickFunction = () => null;
	type clickFunctionAsync = () => Promise<void>;
	export let onclick: clickFunction | clickFunctionAsync;
	export let text: string;
	export let disabled: boolean;
</script>

<button
	class=" w-52 rounded-3xl bg-PUA-orange px-2 text-center text-lg font-semibold text-white disabled:cursor-not-allowed disabled:opacity-50"
	on:click={onclick}
	{disabled}
>
	{text}
</button>
