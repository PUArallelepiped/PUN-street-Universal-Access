<script lang="ts">
	export let statusCardContent: {
		time: string;
		price: string;
		src: string;
		text: string;
		user: string;
	};
</script>

<div class=" w-3/7">
	<div class="flex flex-wrap items-center justify-center gap-8">
		<button class="flex items-center gap-5 rounded-xl bg-white p-4 shadow-xl">
			<img src={statusCardContent.src} alt="" />
			<div class="w-64 text-start font-bold">
				<div class=" text-xl font-bold leading-relaxed text-PUA-dark-orange">
					{statusCardContent.time}
				</div>
				<div
					class=" flex items-baseline gap-1 text-xl font-bold leading-relaxed text-PUA-dark-orange"
				>
					<p class=" text-base">NT$</p>
					{statusCardContent.price}
				</div>
				<div class=" text-xl font-bold leading-relaxed text-PUA-dark-orange">
					Order user : {statusCardContent.user}
				</div>
			</div>
			<div>
				<div class="absolute z-10">
					<button
						class=" flex h-20 w-40 items-center justify-center rounded-2xl bg-PUA-stone p-2 text-center font-bold leading-relaxed text-white hover:border-[3px] hover:border-PUA-stone hover:bg-white hover:text-PUA-stone"
					>
						{statusCardContent.text}
					</button>
				</div>

				<div class="h-20 w-40"></div>
			</div>
		</button>
	</div>
</div>
