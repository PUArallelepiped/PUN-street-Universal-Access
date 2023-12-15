<script lang="ts">
	export let label: string;
	export let value: string;
	export let type: string;
</script>

<div class="">
	<div class="text-left text-xl font-bold leading-relaxed text-orange-950">{label}</div>
	<input
		{type}
		class="h-10 w-96 rounded-lg border-2 bg-gray-200 px-2 text-xl font-medium leading-relaxed text-orange-950 shadow-inner"
		{value}
	/>
</div>
