<script lang="ts">
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	import { goto } from '$app/navigation';
	import type { ActionResult } from '@sveltejs/kit';
	import { deserialize } from '$app/forms';

	export let data: PageData;

	onMount(async () => {
		const stateStr = 'state';
		const state = localStorage.getItem(stateStr) as string;
		const veriStr = 'verifier';
		const verifier = localStorage.getItem(veriStr) as string;
		if (state !== data.state) {
			goto('/login', { replaceState: true });
		}
		let formData = new FormData();
		let code = data.code as string;
		formData.append('code', code);
		formData.append(veriStr, verifier);
		const resp = await fetch('?/getToken', {
			method: 'POST',
			body: formData
		});
		const actionResult: ActionResult = deserialize(await resp.text());
		if (actionResult.type === 'success') {
			goto(actionResult.data as unknown as string);
		}
	});
</script>

<div class="flex h-full w-full items-center justify-center text-4xl font-bold text-PUA-dark-orange">
	loading
</div>
