<script>
	import { page } from '$app/stores';
	import '../app.css';
	import Navbar from '$lib/navbar.svelte';
	import { validateToken } from '$lib/components/PUA/validation';
	import { browser } from '$app/environment';
  
	$: if(browser && $page) {
			validateToken().catch(error => {
				console.log(error);
			});
	}
  </script>
  
  <Navbar />
  <slot />