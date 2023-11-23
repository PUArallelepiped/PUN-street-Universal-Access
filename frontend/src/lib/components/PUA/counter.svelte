<script lang="ts">
	import triangle_sub from '$lib/assets/triangle_sub.png';
	import triangle_add from '$lib/assets/triangle_add.png';
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

<div class="flex items-center justify-center">
	<button class="flex items-center" on:click={sub}>
		<img src={triangle_sub} alt="" class="h-[25px] object-cover" />
	</button>

	<input
		class="ml-[12px] mr-[12px] rounded-[20px] border-[3px] border-red-900 bg-transparent text-center"
		type="text"
		bind:value={count}
	/>

	<button class="triangle-button-add" on:click={add}>
		<img src={triangle_add} alt="" class="h-[25px] object-cover" />
	</button>
</div>

<style>
</style>
