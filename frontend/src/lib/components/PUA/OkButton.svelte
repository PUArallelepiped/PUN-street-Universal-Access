<script lang="ts">
	type clickFunction = () => void;
	type clickFunctionAsync = () => Promise<void>;
	export let onclick: clickFunction | clickFunctionAsync = () => null;
	export let text: string;
	export let disabled: boolean = false;
</script>

<button
	class=" w-52 rounded-3xl bg-PUA-orange px-2 text-center text-lg font-semibold text-white disabled:cursor-not-allowed disabled:opacity-50"
	on:click={onclick}
	{disabled}
>
	{text}
</button>
