<script lang="ts">
	import HistoryCard from '$lib/components/PUA/historyCard.svelte';
	import type { PageData } from './$types';
	export let data: PageData;
	console.log(data);
</script>

<!-- customer/{userID}/get-history -->
<div class="m-2 flex flex-col items-center justify-center gap-3">
	<h1 class="text-5xl font-bold text-PUA-dark-red">History</h1>

	{#each data.history as history}
		<HistoryCard
			avgRate={history.store_rate}
			cost={history.total_price}
			date={history.order_date}
			shopName={history.store_name}
			img={history.store_picture}
		></HistoryCard>
	{/each}
</div>
