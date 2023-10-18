<script lang="ts">
	import watermelon from '$lib/assets/watermelon.png';
	import Counter from '$lib/components/counter.svelte';
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	export let data: PageData;

	const information: Array<Array<string>> = [[watermelon, 'this is watermelon', '100', '10']];

	enum informationContent {
		picture,
		text,
		price,
		stock
	}
</script>

<h1>
	shop: {data.shop}
</h1>
<h2>
	item: {data.item}
</h2>

{#each information as item}
	<div>
		<img src={item[informationContent.picture]} alt="watermelon" class="h-8" />
		<p>{item[informationContent.text]}</p>
		<p>{item[informationContent.price]}</p>
		<p>{item[informationContent.stock]}</p>
		<p>add to cart</p>
		<p>view seller shop</p>

		<p>buy counter</p>
		<Counter />
	</div>
{/each}
