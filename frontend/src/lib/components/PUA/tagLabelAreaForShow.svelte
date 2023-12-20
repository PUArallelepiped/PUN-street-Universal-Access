<script lang="ts">
	import TagStar from './tag/tagStar.svelte';
	import TagLabel from './tag/tagLabel.svelte';
	export let tagText: { category_id: number; category_name: string }[];
	export let star_score = '';
</script>

<div class="flex justify-start">
	<TagStar text={star_score}></TagStar>
</div>
<div class="flex w-full flex-wrap gap-2">
	{#each tagText as { category_id, category_name }}
		<TagLabel canRemove={true} text={category_name}></TagLabel>
	{/each}
</div>
