<script lang="ts">
	import { backendPath } from '$lib/components/PUA/env';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import InputBox from '$lib/components/PUA/InputBox.svelte';
	import { ErrorMessage } from '$lib';

	let user_email = '';
	let password = '';
	let errorMsgVisible = false;

	onMount(() => {
		let errorMsg = window.document.getElementById('errorMsg')! as HTMLElement;
		errorMsg.style.visibility = 'hidden';
	});

	async function login() {
		const res = await fetch(backendPath + '/login', {
			method: 'POST',
			credentials: 'include',
			body: JSON.stringify({
				user_email,
				password
			})
		});
		if (res.status == 200) {
			const data = await res.json();
			goto('/shops');
			console.log(data);
		} else {
			console.log('error');
			errorMsgVisible = true;
		}
	}

	function handleInput() {
		errorMsgVisible = false;
	}
</script>

<div>
	<h1 class="mx-20 my-10 p-12 text-5xl font-medium leading-10 text-red-900">| Sign in to PUA</h1>

	<div class="absolute inset-x-0 bottom-0 h-96 bg-red-900"></div>
	<div class="absolute right-1">
		<form class=" mx-20 rounded-2xl bg-white px-8 py-5 shadow" on:submit|preventDefault={login}>
			<div class="flex flex-col items-center gap-11">
				<InputBox label="Email address" bind:value={user_email} onInput={handleInput} type=""
				></InputBox>
				<InputBox label="Password" bind:value={password} onInput={handleInput} type="password"
				></InputBox>
			</div>
			<ErrorMessage {errorMsgVisible} errorMsg="Email or Password ERROR"></ErrorMessage>
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
