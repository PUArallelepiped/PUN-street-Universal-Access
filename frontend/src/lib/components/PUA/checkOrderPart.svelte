<script lang="ts">
	import { goto, invalidateAll } from '$app/navigation';
	import drone from '$lib/assets/drone.png';
	import { PUBLIC_BACKEND_PATH as backendPath } from '$env/static/public';
	import OkButton from './OkButton.svelte';
	import { getId } from './getId';
	import ProgressBar from './progressBar.svelte';

	export let status: number;
	export let picture: string;
	export let shopName: string;
	export let cartID: number;
	export let storeID: number;
	let context: { text: string; status: boolean }[] = [
		{ text: 'Accept order', status: 2 <= status },
		{ text: 'Making order', status: 3 <= status },
		{ text: 'Deliver Order', status: 4 <= status },
		{ text: 'Arrival Address', status: 5 <= status }
	];

	async function updateStatus() {
		try {
			const userId = (await getId()).valueOf();
			await fetch(
				backendPath +
					'/seller/update-order-status/customer/' +
					userId +
					'/cart/' +
					cartID +
					'/store/' +
					storeID,
				{
					method: 'PUT',
					headers: {
						'Content-Type': 'application/json'
					},
					body: JSON.stringify({
						status: 6
					})
				}
			);
			invalidateAll();
			return null;
		} catch (e) {
			goto('/login');
		}
	}
</script>

<div class="flex flex-col items-center justify-center gap-8">
	<div>
		<ProgressBar {context}></ProgressBar>
	</div>
	<div class="flex gap-16 rounded-xl bg-white p-4">
		<div>
			<img src={picture} alt="" class="h-64 w-72 rounded-xl object-cover" />
		</div>
		<div class="flex flex-col justify-around">
			<div class="flex items-center justify-start gap-5">
				<div class=" text-5xl font-semibold text-orange-950">運送中...</div>

				{#if status == 5}
					<OkButton
						text="Check Order"
						onclick={async () => {
							updateStatus();
						}}
					></OkButton>
				{/if}
			</div>
			<div class=" text-2xl font-semibold text-orange-950">{shopName}</div>
			<div class="flex gap-5">
				<img src={drone} alt="" class="h-20 w-32 rounded-xl object-cover" />
				<img src={drone} alt="" class="h-20 w-32 rounded-xl object-cover" />
				<img src={drone} alt="" class="h-20 w-32 rounded-xl object-cover" />
			</div>
		</div>
	</div>
</div>
