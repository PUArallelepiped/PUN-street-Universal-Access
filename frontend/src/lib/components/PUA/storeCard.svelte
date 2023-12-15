<script lang="ts">
	import { page } from '$app/stores';
	import * as Card from '$lib/components/ui/card';
	import { Diff } from 'lucide-svelte';

	export let name: string = 'im pasta';
	export let description: string = 'you are not pasta';
	export let picture: string = 'https://i.imgur.com/T5zyE63.png';
	export let address: string = 'pun street';
	export let rate: number = 0;
</script>

<a href="{$page.route.id}/{name}" class="shrink">
	<div
		class="flex h-96 w-96 flex-col items-center justify-start gap-3 rounded-lg bg-white px-6 py-3 hover:bg-PUA-dark-gray"
	>
		<div class="relative">
			<img src={picture} alt="" class="h-48 w-80 rounded object-cover" />
			<div
				class="absolute right-2 top-2 flex h-[29px] w-16 items-center justify-center gap-1 rounded-lg border-2 border-white border-opacity-50 bg-PUA-dark-red"
			>
				<div class="inline text-2xl text-yellow-400">★</div>
				<div class="text-center text-base font-semibold leading-tight text-white">
					{rate}
				</div>
			</div>
		</div>
		<div class="w-full text-4xl font-semibold leading-relaxed text-PUA-dark-orange">{name}</div>
		<div>tag</div>
	</div>
</a>
