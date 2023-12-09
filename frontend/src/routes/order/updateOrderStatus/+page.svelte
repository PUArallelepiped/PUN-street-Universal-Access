<script lang="ts">
	import Deliver from '$lib/assets/deliver.svg';
	import Accept from '$lib/assets/accept.svg';
	import Arrival from '$lib/assets/arrival.svg';
	import Making from '$lib/assets/making.svg';
	import OrderStatusCard from '$lib/components/PUA/orderStatusCard.svelte';

	let statusCardContent: { src: string; text: string; user: string }[] = [
		{ src: Making, text: 'Making', user: 'user01' },
		{ src: Deliver, text: 'Deliver', user: 'user01' },
		{ src: Accept, text: 'Accept', user: 'user01' },
		{ src: Arrival, text: 'Arrival', user: 'user01' },
		{ src: Accept, text: 'Accept', user: 'user01' },
		{ src: Making, text: 'Making', user: 'user01' },
		{ src: Deliver, text: 'Deliver', user: 'user01' },
		{ src: Arrival, text: 'Arrival', user: 'user01' }
	];
</script>

<div class=" flex flex-wrap items-center justify-center gap-10 py-12">
	{#each statusCardContent as sub}
		<OrderStatusCard statusCardContent={sub}></OrderStatusCard>
	{/each}
</div>
