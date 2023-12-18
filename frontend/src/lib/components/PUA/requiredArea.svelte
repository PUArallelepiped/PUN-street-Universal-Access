<script lang="ts">
	import { Redradiobox } from '$lib/index';
	export let nameYes: string;
	export let idYes: string;
	export let nameNo: string;
	export let idNo: string;
	export let group: boolean;
	export let product_required = true;
</script>

{#if product_required}
	<div class="ml-2 mr-2 flex items-center justify-center">
		<Redradiobox name={nameNo} id={idNo} checked={true} value={true} bind:group />
	</div>
	<label for="Choice" class="text-red-950">Yes</label>
	<div class="items-cener ml-2 mr-2 flex justify-center">
		<Redradiobox name={nameYes} id={idYes} checked={false} value={false} bind:group />
	</div>
	<label for="Choice" class="text-red-950">No</label>
{:else}
	<div class="ml-2 mr-2 flex items-center justify-center">
		<Redradiobox name={nameNo} id={idNo} checked={false} value={true} bind:group />
	</div>
	<label for="Choice" class="text-red-950">Yes</label>
	<div class="items-cener ml-2 mr-2 flex justify-center">
		<Redradiobox name={nameYes} id={idYes} checked={true} value={false} bind:group />
	</div>
	<label for="Choice" class="text-red-950">No</label>
{/if}
