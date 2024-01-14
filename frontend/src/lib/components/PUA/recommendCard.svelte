<script lang="ts">
	import { goto } from '$app/navigation';

	export let recommend: {
		store_id: number;
		product_id: number;
		picture: string;
		product_name: string;
	};
	async function gotoproduct() {
		await goto('/shops/' + recommend.store_id.toString() + '/' + recommend.product_id.toString());
		location.reload();
	}
</script>

{#if recommend}
	<button
		on:click={gotoproduct}
		class="h-52 w-36 rounded-2xl bg-white p-4 shadow hover:bg-zinc-300 hover:shadow-lg hover:shadow-zinc-400"
	>
		<div class="flex h-full flex-col items-center gap-2">
			<img src={recommend.picture} alt="" class=" h-28 w-28 rounded-lg object-cover" />
			<div
				class="line-clamp-2 w-full text-ellipsis break-words text-center font-bold text-PUA-dark-red"
			>
				{recommend.product_name}
			</div>
		</div>
	</button>
{/if}
