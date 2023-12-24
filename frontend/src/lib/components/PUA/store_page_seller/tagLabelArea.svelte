<script lang="ts">
	import TagStar from '../tag/tagStar.svelte';
	import TagAdd from '../tag/tagAdd.svelte';
	import TagLabel from '../tag/tagLabel.svelte';
	import TagMenu from '../tag/tagMenu.svelte';
	import { backendPath } from '../env';
	export let tagText: { category_id: number; category_name: string }[];
	export let tagText_all: { category_id: number; category_name: string }[];
	export let star_score = '';
	export let shop_id = '1';

	async function PostGetCategoryResp(id: number) {
		const resp = await fetch(backendPath + '/store/' + shop_id + '/add-category/' + id.toString(), {
			method: 'POST'
		}).then(() => {
			return fetch(backendPath + `/store/` + shop_id);
		});
		let respJson = await resp.json();
		tagText = respJson.category_array;
		console.log('tagText', tagText);
		menu_show = !menu_show;
		return null;
	}

	async function removedTag(id: number) {
		const resp = await fetch(
			backendPath + '/store/' + shop_id + '/remove-category/' + id.toString(),
			{
				method: 'DELETE'
			}
		).then(() => {
			return fetch(backendPath + `/store/` + shop_id);
		});
		let respJson = await resp.json();
		tagText = respJson.category_array;
		console.log('tagText', tagText);
	}

	let menu_show = false;

	let intersectionData: {
		category_id: number;
		category_name: string;
	}[];
	$: {
		intersectionData = tagText_all.filter(
			(tagItem) => !tagText?.some((tagText) => tagText.category_name === tagItem.category_name)
		);
	}
</script>

<div class="flex justify-start">
	<TagStar text={star_score}></TagStar>
</div>
<div class="flex w-full flex-wrap gap-2">
	{#if tagText}
		{#each tagText as { category_id, category_name }}
			<TagLabel canRemove={true} text={category_name} on:click={() => removedTag(category_id)}
			></TagLabel>
		{/each}
	{/if}
	{#if menu_show && intersectionData.length !== 0}
		<TagMenu
			bind:tagMenuData={tagText_all}
			bind:disabled_tag={tagText}
			click_function={PostGetCategoryResp}
		></TagMenu>
	{:else if !menu_show && intersectionData.length !== 0}
		<TagAdd on:click={() => (menu_show = !menu_show)}></TagAdd>
	{/if}
</div>
