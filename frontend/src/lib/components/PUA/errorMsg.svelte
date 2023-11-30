<script lang="ts">
	import error from '$lib/assets/error.png';

	export let height: string = '6';
	export let width: string = '6';
</script>

<div class="flex items-center">
	<img src={error} alt="Error" class="object-cover" {width} {height} />
	<p class="text-xl font-bold text-red-500">Error message</p>
</div>
