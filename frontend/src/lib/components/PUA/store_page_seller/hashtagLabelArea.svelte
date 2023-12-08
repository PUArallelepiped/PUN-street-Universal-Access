<script lang="ts">
	import { HashtagLabel } from '$lib';
	export let hashtag_text: { id: number; label: string }[];

	let hashtag_input_id: { id: number; inputText: string }[] = [];
	let hashtag_text_nextId = hashtag_text.length + 1;
	let hashtag_input_nextId = hashtag_input_id.length + 1;

	function removeInput_addHashtag(id: number) {
		let result = hashtag_input_id.find((item) => item.id === id);
		if (result !== undefined) {
			hashtag_text = [...hashtag_text, { id: hashtag_text_nextId, label: `${result.inputText}` }];
			hashtag_input_id = hashtag_input_id.filter((cat) => cat.id !== id);
			hashtag_text_nextId = hashtag_text_nextId + 1;
		}
	}

	function removedHashtag(id: number) {
		let result = hashtag_text.find((item) => item.id === id);
		if (result !== undefined) {
			hashtag_text = hashtag_text.filter((cat) => cat.id !== id);
		}
	}
	function addHashtagInput() {
		hashtag_input_id = [...hashtag_input_id, { id: hashtag_input_nextId, inputText: '' }];
		hashtag_input_nextId = hashtag_input_nextId + 1;
	}
	function setHashtag(event: KeyboardEvent, id: number) {
		if (event.key === 'Enter') {
			removeInput_addHashtag(id);
		}
	}
</script>

<div class="flex justify-start">
	<HashtagLabel type={'star'} text={'4.7'}></HashtagLabel>
</div>
<div class="flex w-full flex-wrap gap-2">
	{#each hashtag_text as { id, label }}
		<HashtagLabel type={'text'} text={label} on:click={() => removedHashtag(id)}></HashtagLabel>
	{/each}
	{#each hashtag_input_id as { id, inputText }}
		<HashtagLabel
			on:keydown={(e) => setHashtag(e, id)}
			type={'input'}
			id={`input${id}`}
			bind:text={inputText}
		></HashtagLabel>
	{/each}
	<HashtagLabel on:click={addHashtagInput} type={'add'}></HashtagLabel>
</div>
