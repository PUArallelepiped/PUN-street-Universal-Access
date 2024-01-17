<script lang="ts">
	import CheckOrderPart from '$lib/components/PUA/checkOrderPart.svelte';
	import type { PageData } from './$types';
	import { onMount } from 'svelte';
	export let data: PageData;

	let socket;

	onMount(() => {
		socket = new WebSocket('ws://localhost:5000/socket');
		socket.addEventListener('message', function (event) {
			// console.log(event);
			if (event.data == 'update') {
				location.reload();
			}
		});
	});
</script>

<div class="flex flex-col gap-24 py-5">
	{#if data.orderInfoList.length == 0}
		<div class="text-center text-7xl font-bold leading-9 text-PUA-dark-red">No Order Now</div>
	{:else}
		{#each data.orderInfoList as orderInfo}
			<CheckOrderPart
				shopName={orderInfo.store_name}
				picture={orderInfo.store_picture}
				status={orderInfo.status}
				cartID={orderInfo.cart_id}
				storeID={orderInfo.store_id}
			></CheckOrderPart>
		{/each}
	{/if}
</div>
