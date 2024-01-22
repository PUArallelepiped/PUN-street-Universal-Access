<script lang="ts">
	import '../app.css';
	import Navbar from '$lib/navbar.svelte';
	const routes = [
		{ path: '/profile', title: 'My profile' },
		{ path: '/shops', title: 'ShopList' },
		{ path: '/order/checkOrder', title: 'Check Order' },
		{ path: '/history', title: 'History' }
	];
</script>

<Navbar {routes} />
<slot />
