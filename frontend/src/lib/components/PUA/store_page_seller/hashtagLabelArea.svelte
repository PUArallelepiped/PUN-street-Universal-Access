<script lang="ts">
	import { HashtagLabel } from '$lib';
	export let hashtag_text: { label: string }[];

	let hashtag_input_id: { id: number; inputText: string }[] = [];

	function removeInput_addHashtag(id: number) {
		let result = hashtag_input_id.find((item) => item.id === id);
		if (result !== undefined) {
			hashtag_text = [...hashtag_text, { label: `${result.inputText}` }];
			hashtag_input_id = hashtag_input_id.filter((cat) => cat.id !== id);
		}
	}
	function addHashtagInput() {
		let hashtag_input_nextId = hashtag_input_id.length;
		hashtag_input_id = [...hashtag_input_id, { id: hashtag_input_nextId, inputText: '' }];
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
	{#each hashtag_text as { label }}
		<HashtagLabel type={'text'} text={label}></HashtagLabel>
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
