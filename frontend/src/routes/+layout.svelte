<script>
	import { page } from '$app/stores';
	import '../app.css';
	import Navbar from '$lib/navbar.svelte';
	const routes = [
		{ path: '/profile', title: 'My profile' }, // TODO
		{ path: '/shops', title: 'ShopList' },
		{ path: '/order/checkOrder', title: 'Check Order' }, // TODO
		{ path: '/history', title: 'History' } // TODO
	];
</script>

<Navbar {routes} />
<slot />
