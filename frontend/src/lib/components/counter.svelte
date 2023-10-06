<script lang="ts">
	let count = 0;
	export let allowNegative: boolean = false;
	function add() {
		count += 1;
	}
	function sub() {
		if (!allowNegative && count == 0) {
			return;
		}
		count -= 1;
	}
</script>

<button on:click={add}>+</button>
<input type="text" bind:value={count} />
<button on:click={sub}>-</button>
