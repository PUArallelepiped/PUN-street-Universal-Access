<script lang="ts">
	import watermelon from '$lib/assets/watermelon.png';
</script>

<div class="flex w-96 flex-col rounded-xl bg-white p-2.5">
	<div class="flex gap-4">
		<img class="h-20 w-20 rounded-xl shadow-inner" src={watermelon} alt="img" />
		<div class="flex flex-col gap-3">
			<div class="text-2xl font-normal text-orange-950">茶碗蒸</div>
			<div class="text-sm font-normal text-neutral-400">
				鮮甜柴魚高湯配上香菇、雞腿肉，口感滑潤好滋味。可能不太適合拿來...
			</div>
		</div>
	</div>
	<div class="flex items-center justify-between">
		<div class="flex items-center gap-4 px-3">
			<div class="">
				<svg width="16" height="18" fill="none" xmlns="http://www.w3.org/2000/svg">
					<path
						id="Vector"
						d="M6.85714 0C5.6 0 4.57143 1.0125 4.57143 2.25H2.28571C1.02857 2.25 0 3.2625 0 4.5H16C16 3.2625 14.9714 2.25 13.7143 2.25H11.4286C11.4286 1.0125 10.4 0 9.14286 0H6.85714ZM2.28571 6.75V17.5725C2.28571 17.82 2.46857 18 2.72 18H13.3029C13.5543 18 13.7371 17.82 13.7371 17.5725V6.75H11.4514V14.625C11.4514 15.255 10.9486 15.75 10.3086 15.75C9.66857 15.75 9.16572 15.255 9.16572 14.625V6.75H6.88V14.625C6.88 15.255 6.37714 15.75 5.73714 15.75C5.09714 15.75 4.59429 15.255 4.59429 14.625V6.75H2.30857H2.28571Z"
						fill="#755555"
					/>
				</svg>
			</div>
			<div class="flex items-end gap-4">
				<span class=" text-base font-semibold text-red-900">NT$160</span>
				<span class=" text-xl font-semibold text-red-900">x2</span>
			</div>
		</div>
		<div class="rounded-xl border-2 border-stone-600 p-2 text-sm font-semibold text-stone-600">
			買二送一
		</div>
	</div>
</div>
