<script lang="ts">
	import watermelon from '$lib/assets/watermelon.png';
</script>

<button
	class="h-52 w-36 rounded-2xl bg-white p-4 shadow hover:bg-zinc-300 hover:shadow-lg hover:shadow-zinc-400"
>
	<div class="flex h-full flex-col items-center gap-2">
		<img src={watermelon} alt="" class="  h-28 w-28 rounded-lg" />
		<div
			class="text-PUA-dark-red line-clamp-2 w-full text-ellipsis break-words text-center font-bold"
		>
			查碗蒸
		</div>
	</div>
</button>
