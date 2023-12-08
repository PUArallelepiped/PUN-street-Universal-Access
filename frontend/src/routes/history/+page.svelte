<script lang="ts">
	import HistoryCard from '$lib/components/PUA/historyCard.svelte';
</script>

<!-- customer/{userID}/get-history -->
<div class="m-2 flex flex-col items-center justify-center gap-3">
	<h1 class="text-5xl font-bold text-PUA-dark-red">History</h1>

	<HistoryCard></HistoryCard>
	<HistoryCard></HistoryCard>
</div>
