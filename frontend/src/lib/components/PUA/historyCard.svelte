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
<!-- <div
	class="inline-flex h-[156px] w-[957px] flex-col items-start justify-start gap-2.5 rounded-[10px] p-2.5"
>
	<div class="relative h-[136px] w-[937px]">
		<div class="absolute left-0 top-0 h-[133px] w-[266px]">
			<img
				class="absolute left-0 top-0 h-[133px] w-[266px] rounded-[10px] shadow-inner"
				src="https://via.placeholder.com/266x133"
			/>
			<div
				class="absolute left-[184px] top-[10px] inline-flex h-[29px] w-[72px] items-start justify-center gap-0.5 rounded-[10px] border-2 border-white border-opacity-50 bg-red-900 px-[7px] py-[5px]"
			>
				<div class="font-['Inter'] text-base font-semibold leading-tight text-white">4.87</div>
			</div>
		</div>
		<div class="absolute left-[286px] top-[106px] h-[30px] w-[284px]">
			<div
				class="absolute left-0 top-[3px] font-['Inter'] text-xl font-semibold leading-normal text-red-900"
			>
				Rate Store
			</div>
			<div
				class="absolute left-[122px] top-0 inline-flex h-[30px] w-[162px] items-start justify-start gap-[3px]"
			></div>
		</div>
		<div
			class="absolute left-[298px] top-[58px] font-['Inter'] text-base font-bold leading-tight text-stone-600"
		>
			IM pasta<br />Shop name
		</div>
		<div class="absolute left-[739px] top-[50px] h-[35px] w-[198px]">
			<div class="absolute left-0 top-0 h-[35px] w-[198px] rounded-[20px] bg-orange-700"></div>
			<div
				class="absolute left-0 top-[7px] h-[21px] w-[198px] text-center font-['Inter'] text-xl font-bold leading-relaxed text-white"
			>
				Detail
			</div>
		</div>
		<div class="absolute left-[286px] top-[26px]">
			<span style="text-orange-950 text-base font-bold font-['Inter'] leading-tight">NT$</span><span
				style="text-orange-950 text-xl font-bold font-['Inter'] leading-relaxed"
			>
				9999</span
			>
		</div>
	</div>
</div> -->
