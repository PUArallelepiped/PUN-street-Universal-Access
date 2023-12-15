<script lang="ts">
	export let tagMenuData: {
		category_id: number;
		category_name: string;
	}[] = [{ category_id: 0, category_name: 'free delivery' }];

	export let disabled_tag: {
		category_id: number;
		category_name: string;
	}[] = [{ category_id: 0, category_name: 'free delivery' }];

	export let click_function: (category_id: number) => void;

	let intersectionData: {
		category_id: number;
		category_name: string;
	}[];
	$: {
		intersectionData = tagMenuData.filter(
			(tagItem) =>
				!disabled_tag.some((disabledItem) => disabledItem.category_id === tagItem.category_id)
		);
	}
</script>

<div class=" flex h-7 w-fit items-center justify-center">
	<div class=" group relative block w-fit text-transparent">
		<button
			class="group inline-flex min-w-[142px] items-center justify-center rounded-full border-2 border-PUA-dark-red bg-PUA-dark-red px-4 py-0 text-center text-base font-bold text-white group-hover:border-2 group-hover:border-PUA-dark-red group-hover:bg-transparent group-hover:text-PUA-dark-red"
		>
			<span class="mr-1">Add Tag</span>
			<svg
				class="h-4 w-4 rotate-180 fill-current group-hover:rotate-0"
				xmlns="http://www.w3.org/2000/svg"
				viewBox="0 0 20 20"
				><path d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z" />
			</svg>
		</button>
		<ul
			class="absolute left-1/2 z-20 hidden w-fit -translate-x-1/2 transform pt-1 text-gray-700 group-hover:block"
		>
			<div class=" flex w-fit flex-col items-center justify-center gap-1 bg-white p-5 shadow-xl">
				{#each intersectionData as { category_id, category_name }}
					<li class="block w-fit cursor-pointer">
						<button
							class="min-w-[142px] whitespace-nowrap rounded-full border-2 border-PUA-dark-red bg-PUA-dark-red px-4 py-0 text-center font-bold text-white hover:border-2 hover:border-PUA-dark-red hover:bg-white hover:text-PUA-dark-red disabled:cursor-not-allowed disabled:border-gray-300 disabled:bg-gray-300 disabled:text-white hover:disabled:text-white"
							on:click={() => click_function(category_id)}
						>
							{category_name}
						</button>
					</li>
				{/each}
			</div>
		</ul>
	</div>
</div>
