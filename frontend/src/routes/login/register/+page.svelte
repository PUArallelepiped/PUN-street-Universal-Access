<script lang="ts">
	import ProgressBar from '$lib/components/PUA/progressBar.svelte';
	import Wallet from '$lib/assets/wallet.svg';
	import Casher from '$lib/assets/casher.svg';
	import InputBox from '$lib/components/PUA/InputBox.svelte';
	import OkButton from '$lib/components/PUA/OkButton.svelte';
	import CheckBox from '$lib/components/PUA/CheckBox.svelte';
	let context: { text: string; status: boolean }[] = [
		{ text: 'Choose User Type', status: false },
		{ text: 'Compelete basic information', status: false },
		{ text: 'Check email', status: false },
		{ text: 'Compelete!', status: false }
	];
	function NextStep(): null {
		for (let index = 0; index < context.length; index++) {
			if (!context[index].status) {
				context[index].status = !context[index].status;
				return null;
			}
		}
		return null;
	}
	function ClickWallet() {
		NextStep();
	}
	function ClickCasher() {
		context = context
			.slice(0, 2)
			.concat({ text: 'Compelete Store Info', status: false })
			.concat(context.slice(2, 5));
		NextStep();
	}
</script>

<div class="flex flex-col gap-9 py-10">
	<ProgressBar {context}></ProgressBar>

	{#if !context[0].status}
		<div class="flex justify-center gap-28">
			<button
				on:click={ClickWallet}
				class="flex h-[295px] w-[295px] flex-col items-center justify-center gap-7 rounded-[20px] bg-white shadow hover:bg-neutral-200"
			>
				<img src={Wallet} alt="" class="h-40 w-40" />
				<div class="text-center text-4xl font-bold leading-8 text-red-900">Customer</div>
			</button>
			<button
				on:click={ClickCasher}
				class="flex h-[295px] w-[295px] flex-col items-center justify-center gap-7 rounded-[20px] bg-white shadow hover:bg-neutral-200"
			>
				<img src={Casher} alt="" class="h-40 w-40" />
				<div class="text-center text-4xl font-bold leading-8 text-red-900">Store</div>
			</button>
		</div>
	{:else if !context[1].status}
		<div class="flex justify-center">
			<div class="flex flex-col items-center gap-10 rounded-lg bg-white p-12">
				<InputBox value="" type="" label="Name" />
				<InputBox value="" type="" label="Email" />
				<InputBox value="" type="" label="Password" />
				<InputBox value="" type="" label="Password Check" />
				<InputBox value="" type="" label="Phone Number" />
				<InputBox value="" type="" label="Birthday" />
				<InputBox value="" type="" label="Address" />
				<CheckBox value="si" id="id" text="Do you be a good PUA user?"></CheckBox>
				<OkButton onclick={NextStep} text="Next Step"></OkButton>
			</div>
		</div>
	{:else if !context[2].status && context[2].text === 'Compelete Store Info'}
		<div class="flex justify-center">
			<div class="flex flex-col items-center gap-10 rounded-lg bg-white p-12">
				<InputBox value="" type="" label="Store Name" />
				<InputBox value="" type="" label="Store Description" />
				<InputBox value="" type="" label="Address" />
				<InputBox value="" type="" label="Shipping Fee" />
				<CheckBox value="si" id="id" text="Do you be a good PUA user?"></CheckBox>
				<OkButton onclick={NextStep} text="Next Step"></OkButton>
			</div>
		</div>
	{:else if !context[2].status}
		<div class="flex flex-col items-center">
			<div class="text-center text-4xl font-bold leading-8 text-PUA-dark-red">Compelete!</div>
			<div class="text-center text-xl font-bold leading-8 text-PUA-dark-red">
				you r a PUA member now
			</div>
		</div>
	{/if}
</div>
