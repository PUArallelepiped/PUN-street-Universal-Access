<script lang="ts">
	import { backendPath } from '$lib/components/PUA/env';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import InputBox from '$lib/components/PUA/InputBox.svelte';
	import { ErrorMessage } from '$lib';
	import NisePanda from '$lib/assets/nise_panda.png';

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
				.then((data) => (document.cookie = 'jwttoken=' + data + '; path=/' + '; samesite=strict' + '; max-age=' + maxAge));
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
		<form class=" mx-20 rounded-2xl bg-white px-8 py-5 shadow" on:submit|preventDefault={login}>
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
	</div>
</div>
