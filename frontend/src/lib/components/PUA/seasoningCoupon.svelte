<script lang="ts">
	export let used: boolean = true;
</script>

<div class="flex h-32 w-96 flex-col">
	<div
		class:text-PUA-gray={used}
		class:text-PUA-stone={!used}
		class:bg-PUA-stone={used}
		class:bg-PUA-gray={!used}
		class="
			 flex h-full items-center rounded-xl border-2 border-PUA-stone text-center font-semibold"
	>
		<div class=" px-6 py-5 text-base">Seasoning Discount</div>
		<div class:bg-PUA-stone={!used} class:bg-PUA-gray={used} class=" h-12 w-1"></div>
		<div class="w-full">
			<div class="text-3xl">80%</div>
			<div>2023/01/01~2023/12/12</div>
		</div>
	</div>
</div>
