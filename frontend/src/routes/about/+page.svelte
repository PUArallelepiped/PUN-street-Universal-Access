<script lang="ts">
	import watermelon from '$lib/assets/watermelon.png';
</script>

<!-- <img src={watermelon} alt="" /> -->
<div class="flex justify-items-start">
	<div class="flex w-full justify-center">
		<div class="flex w-2/3">
			<div class="wrap relative mt-12 w-full space-y-8">
				<!-- <div class=" peer flex justify-start"> -->
				<div
					class=" hover:w-144 group peer flex h-40 w-40 translate-x-60 justify-between overflow-hidden rounded-xl bg-white p-4 shadow-2xl duration-500 ease-in-out hover:translate-x-0 hover:scale-125 hover:shadow-zinc-500"
				>
					<img
						src={watermelon}
						alt=""
						width="100"
						height="100"
						class="mr-4 h-32 w-32 rounded-full shadow"
					/>
					<div
						class=" group-hover:text-PUA-stone w-96 text-xl font-bold text-transparent transition-all duration-[2000] ease-linear"
					>
						<p>+ Name : green</p>

						<div class="bg-PUA-stone my-2 h-[2px] w-full"></div>
						<p>+ Boss</p>
						<p>+ Front Holder</p>
						<p>+ ✇</p>
					</div>
				</div>
				<!-- </div> -->
				<div class=" flex justify-end duration-500 peer-hover:translate-y-20">
					<div
						class=" w-144 flex justify-between rounded-xl bg-white p-4 shadow-2xl duration-500 hover:translate-y-10 hover:scale-125 hover:shadow-zinc-500"
					>
						<div class="text-PUA-stone mr-4 w-full text-xl font-bold">
							<p>+ Name : wj4wj4</p>

							<div class="bg-PUA-stone my-2 h-[2px] w-full"></div>
							<p>+ Artist Holder</p>
							<p>+ Worker</p>
							<p>+ Sleep Every Day</p>
						</div>
						<img
							src={watermelon}
							alt=""
							width="100"
							height="100"
							class="h-32 w-32 rounded-full shadow"
						/>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>
