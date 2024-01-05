<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	const dispatch = createEventDispatcher();
	export let statusCardContent: {
		cart_id: number;
		order_date: string;
		status: number;
		store_id: number;
		total_price: number;
		user_id: number;
		user_name: string;
	};

	export let src: string;
	export let text: string;
	let viewBoxValue = '';

	$: {
		statusCardContent.status;
		switch (statusCardContent.status) {
			case 1:
				viewBoxValue = '0 0 512 512';
				break;
			case 2:
				viewBoxValue = '0 0 580 512';
				break;
			case 3:
				viewBoxValue = '0 0 640 512';
				break;
			case 4:
				viewBoxValue = '0 0 640 512';
				break;
			default:
				viewBoxValue = '0 0 640 600';
		}
	}
	function gotoPage() {
		dispatch('gotoPage', {
			userId: statusCardContent.user_id,
			storeId: statusCardContent.store_id,
			cartId: statusCardContent.cart_id
		});
	}
</script>

<div class=" w-3/7">
	<div class=" flex flex-wrap items-center justify-center gap-8">
		<button
			on:click={gotoPage}
			class=" group flex items-center gap-5 rounded-xl bg-white p-4 shadow transition-all duration-300 hover:-translate-y-6 hover:transform hover:shadow-2xl hover:shadow-zinc-600"
		>
			<svg
				xmlns="http://www.w3.org/2000/svg"
				height="80"
				width="100"
				viewBox={viewBoxValue}
				class="fill-PUA-stone"><path d={src} /></svg
			>
			<div class="w-64 text-start font-bold">
				<div class=" text-PUA-dark-orange text-xl font-bold leading-relaxed">
					{statusCardContent.order_date}
				</div>
				<div
					class=" text-PUA-dark-orange flex items-baseline gap-1 text-xl font-bold leading-relaxed"
				>
					<p class=" text-base">NT$</p>
					<p>{statusCardContent.total_price}</p>
				</div>
				<div class=" text-PUA-dark-orange text-xl font-bold leading-relaxed">
					Order user : {statusCardContent.user_name}
				</div>
			</div>

			<button
				on:click|stopPropagation
				class=" bg-PUA-stone hover:border-PUA-stone hover:text-PUA-stone flex h-20 w-40 items-center justify-center rounded-2xl p-2 text-center font-bold leading-relaxed text-white hover:border-[3px] hover:bg-white"
			>
				{text}
			</button>
		</button>
	</div>
</div>
