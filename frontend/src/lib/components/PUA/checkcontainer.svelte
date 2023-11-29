<script lang="ts">
	import Redradiobox from './redRadioBox.svelte';
	export let category: string = '';
	export let subcategory: string = '';
</script>

<div class="flex h-[30px] w-full items-center justify-center hover:bg-gray-400">
	<Redradiobox name={category} id={category + subcategory} />
	<label
		for={category + subcategory}
		class="flex h-full w-full items-center justify-end font-bold text-primary text-red-950"
	>
		{subcategory}
	</label>
</div>
