<script lang="ts">
	import TagStar from '../tag/tagStar.svelte';
	import TagAdd from '../tag/tagAdd.svelte';
	import TagLabel from '../tag/tagLabel.svelte';
	import TagInput from '../tag/tagInput.svelte';
	export let tagText: { category_id: number; category_name: string }[];

	let tag_input_id: { id: number; inputText: string }[] = [];
	let tag_text_nextId = tagText.length + 1;
	let tag_input_nextId = tag_input_id.length + 1;

	function removeInput_addTag(id: number) {
		let result = tag_input_id.find((item) => item.id === id);
		if (result !== undefined) {
			tagText = [
				...tagText,
				{ category_id: tag_text_nextId, category_name: `${result.inputText}` }
			];
			tag_input_id = tag_input_id.filter((cat) => cat.id !== id);
			tag_text_nextId = tag_text_nextId + 1;
		}
	}

	function removedTag(id: number) {
		let result = tagText.find((item) => item.category_id === id);
		if (result !== undefined) {
			tagText = tagText.filter((cat) => cat.category_id !== id);
		}
	}
	function addTagInput() {
		tag_input_id = [...tag_input_id, { id: tag_input_nextId, inputText: '' }];
		tag_input_nextId = tag_input_nextId + 1;
	}
	function setTag(event: KeyboardEvent, id: number) {
		if (event.key === 'Enter') {
			removeInput_addTag(id);
		}
	}
</script>

<div class="flex justify-start">
	<TagStar text={'4.7'}></TagStar>
</div>
<div class="flex w-full flex-wrap gap-2">
	{#each tagText as { category_id, category_name }}
		<TagLabel canRemove={true} text={category_name} on:click={() => removedTag(category_id)}
		></TagLabel>
	{/each}
	{#each tag_input_id as { id, inputText }}
		<TagInput on:keydown={(e) => setTag(e, id)} id={`input${id}`} bind:text={inputText}></TagInput>
	{/each}
	<TagAdd on:click={addTagInput}></TagAdd>
</div>
