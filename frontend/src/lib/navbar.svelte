<script lang="ts">
	// import{store} from "$app/stores"
	const routes = [
		{ path: '/', title: '🏠' },
		{ path: '/food', title: '🍔' },
		{ path: '/shoppingCart', title: '🛒' }
	];
</script>

<div class=" m-2 flex justify-center space-x-20 bg-sky-900 p-7 text-xl text-gray-50">
	{#each routes as r}
		<a href={r.path}>{r.title}</a>
	{/each}
</div>
