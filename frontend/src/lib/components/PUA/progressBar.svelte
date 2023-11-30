<script lang="ts">
	export let context: { text: string; status: Boolean }[] = [
		{ text: 'Accept order', status: true },
		{ text: 'Making order', status: true },
		{ text: 'Deliver Order', status: false },
		{ text: 'Arrival Address', status: false }
	];
</script>

<div class="flex w-screen justify-around">
	{#each context as c}
		<div class="flex grow flex-col gap-4">
			<div class="text-center text-lg font-bold leading-tight text-red-900">{c.text}</div>
			<div>
				<div class="h-1 w-full bg-PUA-dark-red"></div>
				<div
					class:bg-PUA-dark-red={!c.status}
					class:bg-PUA-pink={c.status}
					class="flex h-1 w-full"
				></div>
				<div class="h-1 w-full bg-PUA-dark-red"></div>
			</div>
		</div>
	{/each}
</div>
