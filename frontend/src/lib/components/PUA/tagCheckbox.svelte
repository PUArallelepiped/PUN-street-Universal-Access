<script lang="ts">
	import Checkbox from '../ui/checkbox/checkbox.svelte';
	import Label from '../ui/label/label.svelte';

	export let id: string;
</script>

<div class="flex items-center justify-start gap-2">
	<Checkbox {id} class="h-6 w-6 rounded-sm border-2 border-PUA-dark-red" />
	<Label for={id} class="truncate text-xl font-bold leading-relaxed text-PUA-dark-red">{id}</Label>
</div>
