<script lang="ts">
	let statusList = ['making', 'Arriving', 'delivering'];
</script>

<div class="flex flex-wrap items-center gap-4 p-10">
	{#each statusList as status}
		<div class=" flex items-center rounded-xl bg-neutral-300 p-4">
			<div class="w-64">
				<div class=" text-xl font-bold leading-relaxed text-PUA-dark-orange">2020</div>
				<div class=" text-xl font-bold leading-relaxed text-PUA-dark-orange">NT9999</div>
				<div class=" text-xl font-bold leading-relaxed text-PUA-dark-orange">user ATM</div>
			</div>
			<div>
				<button
					class="h-20 w-20 rounded-lg bg-PUA-stone p-2 text-center font-bold leading-relaxed text-white"
					>{status}</button
				>
			</div>
		</div>
	{/each}
</div>
