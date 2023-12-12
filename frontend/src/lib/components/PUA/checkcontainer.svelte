<script lang="ts">
	import Redradiobox from './redRadioBox.svelte';
	export let category: string = '';
	export let subcategory: string = '';
</script>

<div class="flex w-full items-center space-x-2">
	<div class="flex h-8 w-full items-center justify-center rounded-full px-5 hover:bg-gray-300">
		<Redradiobox name={category} id={category + subcategory} />
		<label
			for={category + subcategory}
			class="flex h-full w-full cursor-pointer items-center justify-end font-bold text-primary text-red-950"
		>
			{subcategory}
		</label>
	</div>
</div>
