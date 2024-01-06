<script lang="ts">
	import ProgressBar from '$lib/components/PUA/progressBar.svelte';
	import Wallet from '$lib/assets/wallet.svg';
	import Casher from '$lib/assets/casher.svg';
	import InputBox from '$lib/components/PUA/InputBox.svelte';
	import OkButton from '$lib/components/PUA/OkButton.svelte';
	import { backendPath } from '$lib/components/PUA/env';
	import ErrorMessage from '$lib/components/PUA/ErrorMessage.svelte';
	import { goto } from '$app/navigation';
	import { CheckBox } from '$lib';

	let context: { text: string; status: boolean }[] = [
		{ text: 'Choose User Type', status: true },
		{ text: 'Complete basic information', status: false },
		// { text: 'Check email', status: false }, // dlc
		{ text: 'Complete!', status: false }
	];
	interface StoreInfo {
		name: string;
		description: string;
		address: string;
		shipping_fee: number | null;
		picture: string;
	}
	interface UserInfo {
		user_name: string;
		user_email: string;
		password: string;
		phone: string;
		address: string;
		birthday: string;
		StoreRegisterInfo: StoreInfo | null;
	}
	let storeInfo: StoreInfo = {
		name: '',
		description: '',
		address: '',
		shipping_fee: null,
		picture: 'https://i.imgur.com/3i3tyXJ.gif'
	};
	let userInfo: UserInfo = {
		user_name: '',
		user_email: '',
		password: '',
		phone: '',
		address: '',
		birthday: '2002-01-01',
		StoreRegisterInfo: null
	};
	let checkPassword = '';
	let errorMsgVisible: boolean = false;
	let errorMsg: string = '';
	let goodPUA = false;
	let goodPUAStore = false;
	async function NextStep() {
		for (let index = 0; index < context.length; index++) {
			if (!context[index + 1].status) {
				// check null & pwd
				HandleError(context[index].text);
				if (errorMsgVisible) {
					return;
				}
				// check email
				if (context[index].text === 'Complete basic information') {
					const exists = (await CheckEmailExist()).valueOf();
					if (exists === true) {
						errorMsg = 'Email already exists';
						errorMsgVisible = true;
						return;
					}
				}
				if (context[index].text === 'Complete Store Info') {
					userInfo.StoreRegisterInfo = storeInfo;
					// console.log(userInfo);
				}
				if (context[index + 1].text === 'Complete!') {
					Register();
				}
				context[index + 1].status = !context[index + 1].status;
				break;
			}
		}
		return;
	}
	function ClickWallet() {
		NextStep();
	}
	function ClickCasher() {
		context = context
			.slice(0, 2)
			.concat({ text: 'Complete Store Info', status: false })
			.concat(context.slice(2, 4));
		NextStep();
	}
	function CheckPassword() {
		if (userInfo.password === checkPassword) {
			return false;
		} else {
			return true;
		}
	}
	function CheckPhone(value: string) {
		let regex = new RegExp('^[0-9]{10}$');
		return !regex.test(value);
	}
	function CheckEmail(value: string) {
		let regex = new RegExp('^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+.[a-zA-Z0-9-.]+$');
		return !regex.test(value);
	}
	function CheckGoodPUA() {
		goodPUA = !goodPUA;
		return;
	}
	function CheckGoodPUAStore() {
		goodPUAStore = !goodPUAStore;
		return;
	}
	function CheckUserInfoNull() {
		if (
			userInfo.user_name === '' ||
			userInfo.user_email === '' ||
			userInfo.password === '' ||
			userInfo.phone === '' ||
			userInfo.address === ''
		) {
			return true;
		} else {
			return false;
		}
	}
	function CheckStoreInfoNull() {
		if (
			storeInfo.name === '' ||
			storeInfo.description === '' ||
			storeInfo.address === '' ||
			storeInfo.shipping_fee === null ||
			storeInfo.shipping_fee === ''
		) {
			return true;
		} else {
			return false;
		}
	}
	function HandleError(context: string) {
		errorMsgVisible = false;
		if (context === 'Complete basic information') {
			errorMsg = 'Please fill in all the blanks';
			if (CheckUserInfoNull()) {
				errorMsg = 'Please fill in all the blanks';
				errorMsgVisible = true;
			} else if (CheckPassword()) {
				errorMsg = 'Password not match';
				errorMsgVisible = true;
			} else if (CheckPhone(userInfo.phone)) {
				errorMsg = 'Phone number must be 10 digits';
				errorMsgVisible = true;
			} else if (CheckEmail(userInfo.user_email)) {
				errorMsg = 'Email format is incorrect';
				errorMsgVisible = true;
			}
		} else if (context === 'Complete Store Info') {
			errorMsg = 'Please fill in all the blanks';
			if (CheckStoreInfoNull()) {
				errorMsg = 'Please fill in all the blanks';
				errorMsgVisible = true;
			} else if (storeInfo.shipping_fee !== null && isNaN(storeInfo.shipping_fee) === true) {
				errorMsg = 'Shipping fee must be a number';
				errorMsgVisible = true;
			}
		}
	}
	function HandleInput() {
		if (context[2].status) {
			HandleError(context[2].text);
		} else if (context[1].status) {
			HandleError(context[1].text);
		}
		// errorMsgVisible = false;
	}
	function GotoLogin() {
		goto('/login');
	}
	async function Register() {
		if (imageFile) {
			const formData = new FormData();
			formData.append('file', imageFile[0]);
			let url = await fetch(backendPath + `/upload`, {
				method: 'POST',
				body: formData
			});
			storeInfo.picture = await url.json();
		}
		storeInfo.shipping_fee = Number(storeInfo.shipping_fee);
		const res = await fetch(backendPath + '/register', {
			method: 'POST',
			body: JSON.stringify(userInfo)
		});
		if (res.status != 200) {
			errorMsg = 'Register failed';
			errorMsgVisible = true;
		}
	}
	async function CheckEmailExist() {
		const res = await fetch(backendPath + '/check-email', {
			method: 'POST',
			body: JSON.stringify({ user_email: userInfo.user_email })
		});
		const exists = await res.json();
		// console.log(exists)
		if (res.status == 200 && exists === true) {
			return true;
		} else {
			return false;
		}
	}

	const onFileSelected = (e: Event) => {
		// ts too hard i give up
		if (e.target == null) return;
		let image = (e.target as HTMLInputElement).files[0];
		let reader = new FileReader();
		reader.readAsDataURL(image);
		reader.onload = (e) => {
			avatar = e.target.result;
		};
	};

	let imageFile: FileList;
	let avatar: string;
</script>

<div class="flex flex-col gap-9 py-10">
	<ProgressBar {context}></ProgressBar>

	{#if !context[1].status}
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
	{:else if !context[2].status}
		<div class="flex justify-center">
			<div class="flex flex-col items-center gap-10 rounded-lg bg-white p-12">
				<InputBox onInput={HandleInput} bind:value={userInfo.user_name} type="" label="Name" />
				<InputBox onInput={HandleInput} bind:value={userInfo.user_email} type="" label="Email" />
				<InputBox
					onInput={HandleInput}
					bind:value={userInfo.password}
					type="password"
					label="Password"
				/>
				<InputBox
					onInput={HandleInput}
					bind:value={checkPassword}
					type="password"
					label="Password Check"
				/>
				<InputBox onInput={HandleInput} bind:value={userInfo.phone} type="" label="Phone Number" />
				<InputBox bind:value={userInfo.birthday} type="date" label="Birthday" />
				<InputBox onInput={HandleInput} bind:value={userInfo.address} type="" label="Address" />
				<CheckBox on:click={CheckGoodPUA} value="si" id="goodPUA" text="Do you be a good PUA user?"
				></CheckBox>
				<ErrorMessage {errorMsgVisible} {errorMsg}></ErrorMessage>
				<OkButton onclick={NextStep} text="Next Step" disabled={!goodPUA}></OkButton>
			</div>
		</div>
	{:else if context[2].text === 'Complete Store Info' && !context[3].status}
		<div class="flex justify-center">
			<div class="flex flex-col items-center gap-10 rounded-lg bg-white p-12">
				<div class=" flex h-60 w-60 rounded-lg bg-gray-300 shadow-inner">
					<div class="absolute z-10 h-[250px] w-[250px] bg-opacity-0">
						<label for="fileInput" class=" block h-full w-full cursor-pointer bg-opacity-0"></label>
						<input
							type="file"
							id="fileInput"
							accept="image/png, image/jpeg, image/jpg, image/gif"
							bind:files={imageFile}
							on:change={(e) => {
								onFileSelected(e);
							}}
							class="absolute right-0 top-0 cursor-pointer font-bold opacity-0"
						/>
					</div>
					{#if avatar}
						<img src={avatar} alt="" class="flex h-full w-full rounded-lg object-cover" />
					{:else}
						<img
							src={storeInfo.picture}
							alt=""
							class="flex h-full w-full rounded-lg object-cover"
						/>
					{/if}
				</div>
				<InputBox onInput={HandleInput} bind:value={storeInfo.name} type="" label="Store Name" />
				<InputBox
					onInput={HandleInput}
					bind:value={storeInfo.description}
					type=""
					label="Store Description"
				/>
				<InputBox onInput={HandleInput} bind:value={storeInfo.address} type="" label="Address" />
				<InputBox
					onInput={HandleInput}
					bind:value={storeInfo.shipping_fee}
					type=""
					label="Shipping Fee"
				/>
				<CheckBox on:click={CheckGoodPUAStore} value="si" id="id" text="Do you be a good PUA store?"
				></CheckBox>
				<ErrorMessage {errorMsgVisible} {errorMsg}></ErrorMessage>
				<OkButton onclick={NextStep} text="Next Step" disabled={!goodPUAStore}></OkButton>
			</div>
		</div>
	{:else}
		<div class="flex flex-col items-center">
			<div class="text-center text-4xl font-bold leading-8 text-PUA-dark-red">Compelete!</div>
			<div class="text-center text-xl font-bold leading-8 text-PUA-dark-red">
				you are a PUA member now
			</div>
			<div class="flex h-9" />
			<OkButton onclick={GotoLogin} text="Go to Login"></OkButton>
		</div>
	{/if}
</div>
