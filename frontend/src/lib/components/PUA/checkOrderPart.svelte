<script lang="ts">
	import { PUALabel } from '$lib';
	import drone from '$lib/assets/drone.png';
	import ProgressBar from './progressBar.svelte';

	let context: { text: string; status: boolean }[] = [
		{ text: 'Accept order', status: true },
		{ text: 'Making order', status: true },
		{ text: 'Deliver Order', status: false },
		{ text: 'Arrival Address', status: false }
	];
</script>

<div class="flex flex-col items-center justify-center gap-8">
	<div>
		<ProgressBar {context}></ProgressBar>
	</div>
	<div class="flex gap-16 rounded-xl bg-white p-4">
		<div>
			<img src="https://i.imgur.com/3i3tyXJ.gif" alt="" class="h-64 w-72 rounded-xl object-cover" />
		</div>
		<div class="flex flex-col justify-around">
			<div class=" text-5xl font-semibold text-orange-950">運送中...</div>
			<div class="flex gap-5">
				<img src={drone} alt="" class="h-20 w-32 rounded-xl object-cover" />
				<img src={drone} alt="" class="h-20 w-32 rounded-xl object-cover" />
				<img src={drone} alt="" class="h-20 w-32 rounded-xl object-cover" />
			</div>
			<PUALabel
				labelName="Taking Address"
				value="a street a street asterrt a street  bbbbbbbbbbbbbb"
				left={true}
			></PUALabel>
		</div>
	</div>
</div>
