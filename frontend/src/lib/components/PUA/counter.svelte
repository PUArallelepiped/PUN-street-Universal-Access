<script lang="ts">
	import triangle_sub from '$lib/assets/triangle_sub.svg';
	import triangle_add from '$lib/assets/triangle_add.svg';
	export let count = 1;
	export let allowNegative: boolean = false;
	function add() {
		count = Number(count) + 1;
		return null;
	}
	function sub() {
		if (!allowNegative && Number(count) <= 1) {
			count = 1;
			return null;
		}
		count = Number(count) - 1;
		return null;
	}
	function handleInput() {
		count = Number(count.toString().replace(/\D/g, ''));
		if (count < 1) {
			count = 1;
		}
	}
</script>

<div class="flex items-center justify-center">
	<button class="flex items-center" on:click={sub}>
		<img src={triangle_sub} alt="" class="h-[25px] object-cover" />
	</button>

	<input
		class="ml-[12px] mr-[12px] appearance-none rounded-[20px] border-[3px] border-red-900 bg-transparent text-center"
		type="text"
		bind:value={count}
		on:input={handleInput}
	/>

	<button class="triangle-button-add" on:click={add}>
		<img src={triangle_add} alt="" class="h-[25px] object-cover" />
	</button>
</div>
