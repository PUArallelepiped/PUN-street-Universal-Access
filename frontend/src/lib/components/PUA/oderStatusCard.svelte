<script lang="ts">
    export statusCardContent: { src: string; text: string; user: string } = {src:'',text:'',user:''};

</script>

{statusCardContent}
<div class="flex flex-wrap items-center justify-center gap-8 p-10">
	{#each statusCardTent as { src, text, user }}
		<div class=" w-3/7 flex items-center gap-5 rounded-xl bg-white p-4 shadow-xl">
			<img {src} alt="" />
			<div class="w-64 font-bold">
				<div class=" text-PUA-dark-orange text-xl font-bold leading-relaxed">
					2020-10-11 19:20:50
				</div>
				<div
					class=" text-PUA-dark-orange flex items-baseline gap-1 text-xl font-bold leading-relaxed"
				>
					<p class=" text-base">NT$</p>
					9999
				</div>
				<div class=" text-PUA-dark-orange text-xl font-bold leading-relaxed">
					Order user : {user}
				</div>
			</div>
			<div>
				<button
					class="bg-PUA-stone h-20 w-40 rounded-2xl p-2 text-center font-bold leading-relaxed text-white"
					>{text}</button
				>
			</div>
		</div>
	{/each}
</div>
