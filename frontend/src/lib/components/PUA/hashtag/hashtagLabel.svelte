<script lang="ts">
	import hashtagTranshcan from '$lib/assets/hashtagTrashcan.svg';
	export let text = 'null';
	export let canRemove = false;
</script>

<div class=" relative">
	{#if canRemove}
		<div
			class=" hover:border-PUA-dark-red hover:bg-PUA-gray group group absolute left-0 top-0 z-10 flex h-full w-full items-center justify-center rounded-full hover:border-2"
		>
			<button
				on:click
				class="flex w-full items-center justify-center gap-2 opacity-0 duration-300 group-hover:opacity-100"
			>
				<img
					src={hashtagTranshcan}
					alt=""
					class="h-5 w-5 bg-transparent opacity-0 transition-opacity duration-300 group-hover:opacity-100"
				/>
				<p class=" text-PUA-dark-red font-bold">Delete</p>
			</button>
		</div>
	{/if}
	<div
		class="bg-PUA-dark-red flex h-7 min-w-[142px] items-center justify-center rounded-[20px] px-4 py-0 text-center text-base font-bold text-white"
	>
		{text}
	</div>
</div>
