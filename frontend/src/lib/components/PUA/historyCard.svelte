<script>
	import OkButton from './OkButton.svelte';
</script>

<div class="flex gap-5 rounded-xl bg-white p-2">
	<div>
		<img
			class="h-32 w-64 rounded-xl shadow-inner"
			src="https://via.placeholder.com/256x128"
			alt="no pic"
		/>
	</div>
	<div class="flex w-96 flex-col">
		<div class=" text-2xl font-bold leading-relaxed text-PUA-dark-orange">Shop name</div>
		<div class=" text-base font-bold leading-relaxed text-PUA-dark-orange">2020-10-11</div>
		<div class=" text-base font-bold leading-relaxed text-PUA-dark-orange">NT$999</div>
		<div class="font-semibold leading-normal text-PUA-dark-red">Rate store</div>
	</div>
	<div class="flex items-center justify-center">
		<OkButton onclick={() => null} text="Detail"></OkButton>
	</div>
</div>
