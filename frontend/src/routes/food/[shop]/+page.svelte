<script lang="ts">
	import type { PageData } from './$types';
	export let data: PageData;
	let shopName = data.shop;
</script>

<h1>
	here is {shopName}
</h1>
<p>Visit <a href="./{shopName}/product1" class="text-blue-400"> Product1 </a></p>
<p>Visit <a href="./{shopName}/product2" class="text-blue-400"> Product2 </a></p>
<p>Visit <a href="./{shopName}/product3" class="text-blue-400"> Product3 </a></p>
