<script lang="ts">
	import { backendPath } from '$lib/components/PUA/env';

	import order_icon from '$lib/assets/order_icon.svg';
	import customer_icon from '$lib/assets/customer_icon.svg';
	import seller_icon from '$lib/assets/seller_icon.svg';

	export let firstCol: string;
	export let secondCol: string;
	export let type: string;
	export let userID: number = 0;

	export let ban: boolean = true;

	async function banUserById(userID: number) {
		const resp = await fetch(backendPath + '/admin/ban-user/' + userID.toString(), {
			method: 'PUT'
		});
		const json = await resp.json();
		ban = json.status;
	}

	async function UnbanUserById(userID: number) {
		const resp = await fetch(backendPath + '/admin/unban-user/' + userID.toString(), {
			method: 'PUT'
		});
		const json = await resp.json();
		ban = json.status;
	}

	async function changeButtonStatus() {
		if (ban) {
			await banUserById(userID);
		} else {
			await UnbanUserById(userID);
		}
	}
</script>

<div
	class:border-PUA-dark-gray={ban}
	class:border-PUA-red={!ban}
	class="mx-12 my-3 flex justify-between rounded-2xl border-2 hover:bg-slate-200"
>
	<button class="w-full">
		<div class="flex items-center justify-between">
			<div class="flex items-center">
				{#if type === 'customer'}
					<img src={customer_icon} alt="" class="my-4 ml-6 flex h-16 w-16" />
				{:else if type === 'store'}
					<img src={seller_icon} alt="" class="my-4 ml-6 flex h-16 w-16" />
				{:else}
					<img src={order_icon} alt="" class="my-4 ml-6 flex h-16 w-16" />
				{/if}
				<div class="m-4 flex text-left text-PUA-dark-red">
					{firstCol}<br />
					Order user : {secondCol}
				</div>
			</div>
			<div class="m-6">
				{#if type === 'order'}
					<button
						class="border-2 border-PUA-dark-red px-7 py-0 text-PUA-dark-red hover:bg-PUA-dark-red hover:text-white"
						>Detail</button
					>
				{:else if ban}
					<button
						on:click={changeButtonStatus}
						class="w-24 border-2 border-PUA-dark-red px-7 py-0 text-PUA-dark-red">Ban</button
					>
				{:else}
					<button
						on:click={changeButtonStatus}
						class="flex w-24 justify-center border-2 border-PUA-dark-red bg-PUA-dark-red px-7 py-0 text-white"
						>UnBan</button
					>
				{/if}
			</div>
		</div>
	</button>
</div>
