<script lang="ts">
	// import{store} from "$app/stores"
	const routes = [
		{ path: '/', title: 'My profile' }, // TODO
		{ path: '/shops', title: 'ShopList' },
		{ path: '/', title: 'Check Order' }, // TODO
		{ path: '/', title: 'History' } // TODO
	];
</script>

<div class="flex h-20 place-content-between bg-white">
	<a href="/" class="flex h-full items-center justify-start px-11 text-3xl">PUN street access</a>
	<div class="flex h-full items-center justify-end gap-10 px-10">
		{#each routes as r}
			<a href={r.path} class="block p-2">{r.title}</a>
		{/each}
	</div>
</div>
