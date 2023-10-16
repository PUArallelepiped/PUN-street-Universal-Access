<script lang="ts">
	export let text: string = 'null';
	export let onclick: Function;
</script>

<button class="rounded-full bg-cyan-200 p-3 text-2xl" on:click={onclick()}>{text}</button>
