<script lang="ts">
	// import{store} from "$app/stores"
	const routes = [
		{ path: '/', title: 'My profile' },
		{ path: '/food', title: 'ShopList' },
		{ path: '/shoppingCart', title: 'Check Order' },
		{ path: '/shoppingCart', title: 'History' }
	];
</script>

<div class="flex h-20 place-content-between bg-white">
	<div class="m-5 inline-block w-1/3 px-11 text-3xl">PUN street access</div>
	<div class="flex h-full items-center justify-end gap-10 px-10">
		{#each routes as r}
			<a href={r.path} class="block p-2">{r.title}</a>
		{/each}
	</div>
</div>
