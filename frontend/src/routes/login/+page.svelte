<script lang="ts">
	import { twpUrl } from '$lib/components/PUA/env';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import InputBox from '$lib/components/PUA/InputBox.svelte';
	import { ErrorMessage } from '$lib';
	import NisePanda from '$lib/assets/nise_panda.png';
	import { PUBLIC_BACKEND_PATH as backendPath } from '$env/static/public';

	let user_email = '';
	let password = '';
	let errorMsgVisible = false;
	let errorMsg = 'Email or Password ERROR';

	onMount(() => {
		let errorMsg = window.document.getElementById('errorMsg')! as HTMLElement;
		errorMsg.style.visibility = 'hidden';
	});

	async function login() {
		let maxAge = 60 * 60 * 24; // 1 day
		const res = await fetch(backendPath + '/login', {
			method: 'POST',
			credentials: 'include',
			body: JSON.stringify({
				user_email,
				password
			})
		});
		if (res.status == 200) {
			await res
				.json()
				.then(
					(data) =>
						(document.cookie =
							'jwttoken=' + data + '; path=/' + '; samesite=strict' + '; max-age=' + maxAge)
				);
			// console.log((await getId()).valueOf());
			goto('/shops');
		} else if (res.status == 403) {
			errorMsg = 'You got banned, haha';
			errorMsgVisible = true;
		} else {
			errorMsg = 'Email or Password ERROR';
			errorMsgVisible = true;
		}
	}

	function handleInput() {
		errorMsgVisible = false;
		return;
	}
	const randomString = (length: number) => {
		const array = new Uint32Array(length);
		window.crypto.getRandomValues(array);
		return btoa(array.join('')) //
			.replace(/\+/g, '-')
			.replace(/\//g, '_')
			.replace(/=+$/, '');
	};
	const generateVerifier = () => {
		return randomString(4);
	};
	const generateChallenge = async (verifier: string) => {
		const encoder = new TextEncoder();
		const digest = await window.crypto.subtle.digest('SHA-256', encoder.encode(verifier));
		const array = Array.from(new Uint8Array(digest));
		return btoa(String.fromCharCode.apply(null, array)) //
			.replace(/\+/g, '-')
			.replace(/\//g, '_')
			.replace(/=+$/, '');
	};
	async function loginWithTwp() {
		//const { isSuccess } = await refetch();
		const authUrl = twpUrl + '/authorize';

		// if (isSuccess) {
		// 	goto('/');
		// }

		const state = randomString(8);
		const verifier = generateVerifier();
		const challenge = await generateChallenge(verifier);

		localStorage.setItem('state', state);
		localStorage.setItem('verifier', verifier);

		const searchParams = new URLSearchParams({
			client_id: 'twp',
			code_challenge: challenge,
			code_challenge_method: 'S256',
			redirect_uri: `${location.origin}/login/twpCallBack`,
			response_type: 'code',
			state: state
		});

		const url = new URL(authUrl);
		url.search = searchParams.toString();
		window.location.href = url.toString();
	}
</script>

<div>
	<h1 class="mx-20 my-10 p-12 text-5xl font-medium leading-10 text-red-900">| Sign in to PUA</h1>

	<div>
		<img class="absolute bottom-64 left-44 scale-125" src={NisePanda} alt="panda" />
		<img class="absolute bottom-48 left-108 scale-125" src={NisePanda} alt="panda" />
		<img class="absolute bottom-32 left-172 scale-125" src={NisePanda} alt="panda" />
	</div>
	<div class="absolute inset-x-0 bottom-0 h-60 bg-red-900"></div>
	<div class="absolute right-1">
		<div class=" mx-20 rounded-2xl bg-white px-8 py-5 shadow">
			<form on:submit|preventDefault={login}>
				<div class="flex flex-col items-center gap-11">
					<InputBox label="Email address" bind:value={user_email} onInput={handleInput} type=""
					></InputBox>
					<InputBox label="Password" bind:value={password} onInput={handleInput} type="password"
					></InputBox>
				</div>
				<ErrorMessage {errorMsgVisible} {errorMsg}></ErrorMessage>
				<div class="m-3 flex h-16 items-end justify-between gap-10 text-center">
					<div class="w-48">
						<div class="text-center text-base font-bold leading-tight text-stone-400">
							Not a PUA member?
						</div>
						<button
							class="h-9 w-full rounded-2xl bg-neutral-200"
							type="button"
							on:click={() => goto('/login/register')}
						>
							<div class=" text-xl font-bold text-stone-600">Sign up</div>
						</button>
					</div>
					<div class="w-48">
						<button class=" h-9 w-full rounded-2xl bg-orange-700" type="submit">
							<div class="text-xl font-bold text-white">Sign in</div>
						</button>
					</div>
				</div>
			</form>
			<div class="flex w-full items-center">
				<div class="h-[2px] grow bg-PUA-dark-orange"></div>
				<div class="p-4 text-center text-2xl font-bold text-PUA-dark-orange">Or With</div>
				<div class="h-[2px] grow bg-PUA-dark-orange"></div>
			</div>

			<button
				class="flex w-full justify-center gap-2 rounded-full bg-[#145142] p-2 text-xl font-bold text-white"
				on:click={loginWithTwp}
			>
				<img src="https://noobdy.asuscomm.com/assets/logo-28e599d4.png" class="h-7" alt="" />
				<div>Continue with TWP</div>
			</button>
		</div>
	</div>
</div>
