<script lang="ts">
	import TagStar from '../tag/tagStar.svelte';
	import TagAdd from '../tag/tagAdd.svelte';
	import TagLabel from '../tag/tagLabel.svelte';
	import TagMenu from '../tag/tagMenu.svelte';
	export let tagText: { category_id: number; category_name: string }[];
	export let tagText_all: { category_id: number; category_name: string }[];
	export let star_score = '';
	function removeInput_addTag(id: number) {
		let result = tagText_all.find((item) => item.category_id === id);
		if (result !== undefined) {
			tagText = [
				...tagText,
				{ category_id: result.category_id, category_name: `${result.category_name}` }
			];
		}
	}

	function removedTag(id: number) {
		let result = tagText.find((item) => item.category_id === id);
		if (result !== undefined) {
			tagText = tagText.filter((cat) => cat.category_id !== id);
		}
	}
	function setTag(id: number) {
		removeInput_addTag(id);
		menu_valus.show = !menu_valus.show;
	}
	let menu_valus = {
		show: false,
		return_id: -1
	};
</script>

<div class="flex justify-start">
	<TagStar text={star_score}></TagStar>
</div>
<div class="flex w-full flex-wrap gap-2">
	{#each tagText as { category_id, category_name }}
		<TagLabel canRemove={true} text={category_name} on:click={() => removedTag(category_id)}
		></TagLabel>
	{/each}

	{#if menu_valus.show}
		<TagMenu bind:tagMenuData={tagText_all} bind:disabled_tag={tagText} click_function={setTag}
		></TagMenu>
	{:else}
		<TagAdd on:click={() => (menu_valus.show = !menu_valus.show)}></TagAdd>
	{/if}
</div>
