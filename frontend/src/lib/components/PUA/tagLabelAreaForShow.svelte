<script lang="ts">
	import TagStar from './tag/tagStar.svelte';
	import TagLabel from './tag/tagLabel.svelte';
	export let tagText: { category_id: number; category_name: string }[];
	export let star_score: number | string = 0;
</script>

<div class="flex justify-start">
	<TagStar bind:text={star_score}></TagStar>
</div>
<div class="flex w-full flex-wrap gap-2">
	{#if tagText}
		{#each tagText as { category_name }}
			<TagLabel canRemove={false} bind:text={category_name}></TagLabel>
		{/each}
	{/if}
</div>
