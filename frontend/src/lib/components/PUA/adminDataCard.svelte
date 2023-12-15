<script lang="ts">
	import order_icon from '$lib/assets/order_icon.svg';
	import customer_icon from '$lib/assets/customer_icon.svg';
	import seller_icon from '$lib/assets/seller_icon.svg';
	import { onMount } from 'svelte';

	export let firstCol: string;
	export let secondCol: string;
	export let type: string;

	export let ben: boolean = false;
	let self: HTMLDivElement;

	onMount(() => {
		if (ben) {
			self.classList.remove('border-PUA-dark-gray');
			self.classList.add('border-PUA-red');
		} else {
			self.classList.add('border-PUA-dark-gray');
			self.classList.remove('border-PUA-red');
		}
	});

	function changeButtonStatus() {
		ben = !ben;
		if (ben) {
			self.classList.remove('border-PUA-dark-gray');
			self.classList.add('border-PUA-red');
		} else {
			self.classList.add('border-PUA-dark-gray');
			self.classList.remove('border-PUA-red');
		}
	}
</script>

<div
	bind:this={self}
	class="mx-12 my-3 flex justify-between rounded-2xl border-2 border-PUA-dark-gray hover:bg-slate-200"
>
	<button class="w-full">
		<div class="flex items-center justify-between">
			<div class="flex items-center">
				{#if type == '0'}
					<img src={order_icon} alt="" class="my-4 ml-6 flex h-16 w-16" />
				{:else if type == '1'}
					<img src={customer_icon} alt="" class="my-4 ml-6 flex h-16 w-16" />
				{:else}
					<img src={seller_icon} alt="" class="my-4 ml-6 flex h-16 w-16" />
				{/if}
				<div class="m-4 flex text-left text-PUA-dark-red">
					{firstCol}<br />
					Order user : {secondCol}
				</div>
			</div>
			<div class="m-6">
				{#if type == '0'}
					<button
						class="border-2 border-PUA-dark-red px-7 py-0 text-PUA-dark-red hover:bg-PUA-dark-red hover:text-white"
						>Detail</button
					>
				{:else if ben}
					<button
						on:click={changeButtonStatus}
						class="flex w-24 justify-center border-2 border-PUA-dark-red bg-PUA-dark-red px-7 py-0 text-white"
						>UnBen</button
					>
				{:else}
					<button
						on:click={changeButtonStatus}
						class="w-24 border-2 border-PUA-dark-red px-7 py-0 text-PUA-dark-red">Ben</button
					>
				{/if}
			</div>
		</div>
	</button>
</div>
