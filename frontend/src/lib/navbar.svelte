<script lang="ts">
	// import{store} from "$app/stores"
	const routes = [
		{ path: '/', title: '🏠' },
		{ path: '/food', title: '🍔' },
		{ path: '/shoppingCart', title: '🛒' }
	];
</script>

<div class=" flex justify-center space-x-20 text-gray-50 p-7 m-2 text-xl bg-sky-900">
	{#each routes as r}
		<a href={r.path}>{r.title}</a>
	{/each}
</div>
