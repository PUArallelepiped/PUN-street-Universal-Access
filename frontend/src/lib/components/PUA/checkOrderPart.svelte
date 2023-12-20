<script lang="ts">
	import { PUALabel } from '$lib';
	import drone from '$lib/assets/drone.png';
	import OkButton from './OkButton.svelte';
	import { backendPath } from './env';
	import ProgressBar from './progressBar.svelte';

	export let status: number;
	export let picture: string;
	export let shopName: string;
	export let cartID: number;
	export let storeID: number;
	let context: { text: string; status: boolean }[] = [
		{ text: 'Accept order', status: 0 <= status },
		{ text: 'Making order', status: 1 <= status },
		{ text: 'Deliver Order', status: 2 <= status },
		{ text: 'Arrival Address', status: 3 <= status }
	];
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

				<OkButton
					text="Check Order"
					onclick={async () => {
						let userId = 1;
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
									status: 4
								})
							}
						);
						return null;
					}}
				></OkButton>
			</div>
			<div class=" text-2xl font-semibold text-orange-950">{shopName}</div>
			<div class="flex gap-5">
				<img src={drone} alt="" class="h-20 w-32 rounded-xl object-cover" />
				<img src={drone} alt="" class="h-20 w-32 rounded-xl object-cover" />
				<img src={drone} alt="" class="h-20 w-32 rounded-xl object-cover" />
			</div>
			<PUALabel
				labelName="Taking Address"
				value="a street a street asterrt a street  bbbbbbbbbbbbbb"
				left={true}
			></PUALabel>
		</div>
	</div>
</div>
