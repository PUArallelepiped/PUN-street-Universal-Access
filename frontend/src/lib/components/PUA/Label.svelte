<script lang="ts">
	export let labelName: string = 'null';
	export let value: string = 'null';
	export let left: boolean = false;
</script>

<div
	class:justify-between={!left}
	class:justify-start={left}
	class="flex gap-4 text-xl font-semibold text-stone-600"
>
	<div>
		{labelName}
	</div>
	<div>{value}</div>
</div>
