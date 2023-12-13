<script lang="ts">
	import ProgressBar from '$lib/components/PUA/progressBar.svelte';
	import Wallet from '$lib/assets/wallet.svg';
	import Casher from '$lib/assets/casher.svg';
	import InputBox from '$lib/components/PUA/InputBox.svelte';
	import OkButton from '$lib/components/PUA/OkButton.svelte';
	import CheckBox from '$lib/components/PUA/CheckBox.svelte';
	import { StoreIcon, UserSquare } from 'lucide-svelte';
	import { backendPath } from '$lib/components/PUA/env';
	let context: { text: string; status: boolean }[] = [
		{ text: 'Choose User Type', status: false },
		{ text: 'Compelete basic information', status: false },
		// { text: 'Check email', status: false }, // dlc
		{ text: 'Compelete!', status: false }
	];
	let storeInfo = {
		name: '',
		description: '',
		address: '',
		shipping_fee: '',
		picture: 'https://i.imgur.com/3i3tyXJ.gif'
	};
	let userInfo = {
		user_name: '',
		user_email: '',
		password: '',
		phone: '',
		address: '',
		birthday: '2002-01-01',
		storeInfo: null
	};
	let checkPassword = ''; 
	function NextStep(): null {
		for (let index = 0; index < context.length; index++) {
			if (!context[index].status) {
				context[index].status = !context[index].status;
				break;
			}
		}
		console.log(context);
		if (context[context.length - 2].status) {
			Register();
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
			.concat(context.slice(2, 4));
		NextStep();
	}
	function HandleInput() {
		console.log(userInfo);
	}
	async function Register() {
		const res = await fetch(backendPath + '/register', {
			method: 'POST',
			body: JSON.stringify(userInfo)
		});
		if (res.status == 200) {
			const data = await res.json();
			console.log(data);
		} else {
			console.log('error');
		}
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
				<InputBox bind:value={userInfo.user_name} type="" label="Name" />
				<InputBox bind:value={userInfo.user_email} type="" label="Email" />
				<InputBox bind:value={userInfo.password} type="" label="Password" />
				<InputBox bind:value={checkPassword} type="" label="Password Check" />
				<InputBox bind:value={userInfo.phone} type="" label="Phone Number" />
				<!-- <InputBox bind:value={userInfo.birthday} type="" label="Birthday" /> -->
				<InputBox bind:value={userInfo.address} type="" label="Address" />
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
	{:else if !context[context.length - 1].status}
		<div class="flex flex-col items-center">
			<div class="text-center text-4xl font-bold leading-8 text-PUA-dark-red">Compelete!</div>
			<div class="text-center text-xl font-bold leading-8 text-PUA-dark-red">
				you are a PUA member now
			</div>
		</div>
	{/if}
</div>
