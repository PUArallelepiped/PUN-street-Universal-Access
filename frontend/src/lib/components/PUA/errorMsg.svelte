<script lang="ts">
	import error from '$lib/assets/error.svg';

	export let height: string = '6';
	export let width: string = '6';
	export let text = 'Error message';
</script>

<div class="flex items-center gap-2">
	<img src={error} alt="Error" class="object-cover" {width} {height} />
	<p class="text-xl font-bold text-red-500">{text}</p>
</div>
