<script lang="ts">
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	export let data: PageData;
	let shopName = data.shop;
	let response: {
		id: string;
		name: string;
		address: string;
		email: string;
		phone: string;
	};
	onMount(async () => {
		const resp = await fetch(`http://localhost:5000/api/v1/store/1`);
		response = await resp.json();
		console.log(response);
		return response;
	});
</script>

<div>
	{#if response}
		<div class="text-center text-2xl">
			{response.name}
		</div>
		<div class="text-center text-2xl">
			{response.address}
		</div>
		<div class="text-center text-2xl">
			{response.email}
		</div>
		<div class="text-center text-2xl">
			id: {response.id}
		</div>
		<div class="text-center text-2xl">
			{response.phone}
		</div>
	{:else}
		loading....
	{/if}
</div>
<h1>
	here is {shopName}
</h1>
<p>Visit <a href="./{shopName}/product1" class="text-blue-400"> Product1 </a></p>
<p>Visit <a href="./{shopName}/product2" class="text-blue-400"> Product2 </a></p>
<p>Visit <a href="./{shopName}/product3" class="text-blue-400"> Product3 </a></p>
