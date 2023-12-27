<script lang="ts">
	import { goto, invalidateAll } from '$app/navigation';
	import OrderStatusCard from '$lib/components/PUA/orderStatusCard.svelte';
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import type { PageData } from './$types';
	import { backendPath } from '$lib/components/PUA/env';

	export let data: PageData;

	async function postAndChangeStatus(index: number) {
		let orderRespList = data.orderRespList;
		await fetch(
			backendPath +
				`/seller/update-order-status/customer/` +
				orderRespList[index].user_id +
				`/cart/` +
				orderRespList[index].cart_id +
				`/store/` +
				orderRespList[index].store_id,
			{
				method: 'PUT'
			}
		);
		await invalidateAll();

		await refresh();
		return;
	}
	let statusCardContent: {
		time: string;
		price: string;
		src: string;
		text: string;
		user: string;
		status: number;
	}[] = [];
	let Deliver: string =
		'M48 0C21.5 0 0 21.5 0 48V368c0 26.5 21.5 48 48 48H64c0 53 43 96 96 96s96-43 96-96H384c0 53 43 96 96 96s96-43 96-96h32c17.7 0 32-14.3 32-32s-14.3-32-32-32V288 256 237.3c0-17-6.7-33.3-18.7-45.3L512 114.7c-12-12-28.3-18.7-45.3-18.7H416V48c0-26.5-21.5-48-48-48H48zM416 160h50.7L544 237.3V256H416V160zM112 416a48 48 0 1 1 96 0 48 48 0 1 1 -96 0zm368-48a48 48 0 1 1 0 96 48 48 0 1 1 0-96z';
	let Accept: string =
		'M96 80c0-26.5 21.5-48 48-48H432c26.5 0 48 21.5 48 48V384H96V80zm313 47c-9.4-9.4-24.6-9.4-33.9 0l-111 111-47-47c-9.4-9.4-24.6-9.4-33.9 0s-9.4 24.6 0 33.9l64 64c9.4 9.4 24.6 9.4 33.9 0L409 161c9.4-9.4 9.4-24.6 0-33.9zM0 336c0-26.5 21.5-48 48-48H64V416H512V288h16c26.5 0 48 21.5 48 48v96c0 26.5-21.5 48-48 48H48c-26.5 0-48-21.5-48-48V336z';
	let Arrival: string =
		'M304 48a48 48 0 1 0 -96 0 48 48 0 1 0 96 0zm0 416a48 48 0 1 0 -96 0 48 48 0 1 0 96 0zM48 304a48 48 0 1 0 0-96 48 48 0 1 0 0 96zm464-48a48 48 0 1 0 -96 0 48 48 0 1 0 96 0zM142.9 437A48 48 0 1 0 75 369.1 48 48 0 1 0 142.9 437zm0-294.2A48 48 0 1 0 75 75a48 48 0 1 0 67.9 67.9zM369.1 437A48 48 0 1 0 437 369.1 48 48 0 1 0 369.1 437z';
	let Making: string =
		'M320 352c-88.4 0-160-71.6-160-160c0-76.7 62.5-144.7 107.2-179.4c5-3.9 10.9-5.8 16.8-5.8c7.9-.1 16 3.1 22 9.2l46 46 11.3-11.3c11.7-11.7 30.6-12.7 42.3-1C464.5 108 480 160.2 480 192c0 88.4-71.6 160-160 160zm64-111.8c0-36.5-37-73-54.8-88.4c-5.4-4.7-13.1-4.7-18.5 0C293 167.1 256 203.6 256 240.2c0 35.3 28.7 64 64 64s64-28.7 64-64zM32 288c0-17.7 14.3-32 32-32H96c17.7 0 32 14.3 32 32s-14.3 32-32 32v64H544V320c-17.7 0-32-14.3-32-32s14.3-32 32-32h32c17.7 0 32 14.3 32 32v96c17.7 0 32 14.3 32 32v64c0 17.7-14.3 32-32 32H32c-17.7 0-32-14.3-32-32V416c0-17.7 14.3-32 32-32V288zM320 480a32 32 0 1 0 0-64 32 32 0 1 0 0 64zm160-32a32 32 0 1 0 -64 0 32 32 0 1 0 64 0zM192 480a32 32 0 1 0 0-64 32 32 0 1 0 0 64z';
	let icon_array = [
		{ icon: Arrival, text: 'Accept' },
		{ icon: Accept, text: 'Making' },
		{ icon: Making, text: 'Deliver' },
		{ icon: Deliver, text: 'Arrival' }
	];
	async function refresh() {
		statusCardContent = [];
		let orderRespList = data.orderRespList;
		if (orderRespList) {
			for (let i = 0; i < orderRespList.length; i++) {
				statusCardContent = [
					...statusCardContent,
					{
						time: orderRespList[i].order_date,
						price: orderRespList[i].total_price.toString(),
						src: icon_array[orderRespList[i].status - 1].icon,
						text: icon_array[orderRespList[i].status - 1].text,
						user: orderRespList[i].user_name,
						status: orderRespList[i].status
					}
				];
			}
		}
	}
	onMount(async () => {
		await refresh();
	});
</script>

<div class=" flex h-full flex-wrap items-center justify-center gap-10 py-12">
	{#each statusCardContent as sub, index}
		<OrderStatusCard
			on:click={() => postAndChangeStatus(index)}
			on:gotoPage={(event) => {
				goto(
					$page.route.id +
						'/detail/' +
						event.detail.userId +
						'/' +
						event.detail.storeId +
						'/' +
						event.detail.cartId
				);
			}}
			storeId={data.orderRespList[index].store_id}
			cartId={data.orderRespList[index].cart_id}
			userId={data.orderRespList[index].user_id}
			statusCardContent={sub}
		></OrderStatusCard>
	{/each}
</div>
