<script lang="ts">
	import BlueButton from './blueButton.svelte';

	let count = 0;
	export let allowNegative: boolean = false;
	function add() {
		count += 1;
		return null;
	}
	function sub() {
		if (!allowNegative && count == 0) {
			return null;
		}
		count -= 1;
		return null;
	}
</script>

<div class="m-3 p-3">
	<BlueButton text="-" onclick={sub} />
	<input
		class="m-3 rounded-full bg-slate-500 p-3 text-center text-2xl text-blue-100"
		type="text"
		bind:value={count}
	/>
	<BlueButton text="+" onclick={add} />
</div>
