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
			<svg
				xmlns="http://www.w3.org/2000/svg"
				height="80"
				width="100"
				viewBox="0 0 640 512"
				class="fill-PUA-stone"><path d={statusCardContent.src} /></svg
			>
			<div class="w-64 text-start font-bold">
				<div class=" text-xl font-bold leading-relaxed text-PUA-dark-orange">
					{statusCardContent.time}
				</div>
				<div
					class=" flex items-baseline gap-1 text-xl font-bold leading-relaxed text-PUA-dark-orange"
				>
					<p class=" text-base">NT$</p>
					<p>{statusCardContent.price}</p>
				</div>
				<div class=" text-xl font-bold leading-relaxed text-PUA-dark-orange">
					Order user : {statusCardContent.user}
				</div>
			</div>

			<a
				class=" flex h-20 w-40 items-center justify-center rounded-2xl bg-PUA-stone p-2 text-center font-bold leading-relaxed text-white hover:border-[3px] hover:border-PUA-stone hover:bg-white hover:text-PUA-stone"
			>
				{statusCardContent.text}
			</a>
		</button>
	</div>
</div>
