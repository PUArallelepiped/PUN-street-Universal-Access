<script lang="ts">
	import transhcan from '$lib/assets/transhcan.svg';
	import { Redradiobox, Textcontainer, NeedChooseLabel } from '$lib/index';
	import RequiredArea from '$lib/components/PUA/requiredArea.svelte';

	export let category_item: {
		item_array: {
			name: string;
		}[];
		product_id: number;
		label_name: string;
		required: boolean;
	}[];

	export let product_id = 0;

	function addNewCategory() {
		const newLabel = {
			item_array: [{ name: '' }],
			product_id: product_id,
			label_name: '',
			required: true
		};

		category_item = [...category_item, newLabel];
	}

	function addNewSubcategory(index: number) {
		const newItem = { name: '' };
		category_item[index].item_array = [...category_item[index].item_array, newItem];
	}

	function removeCategory(index: number) {
		if (index >= 0 && index < category_item.length) {
			const newArray = category_item.filter((_, i) => i !== index);
			category_item = [...newArray];
		}
	}

	function removeNewSubcategory(index: number) {
		if (index >= 0 && index < category_item.length) {
			const newArray = [...category_item];
			if (newArray[index].item_array.length > 0) {
				newArray[index].item_array.pop();
				category_item = [...newArray];
			}
		}
	}
</script>

<div class="relative mb-[10px] w-full">
	{#each category_item as item, index}
		<div class="mb-[15px]">
			<div
				class="flex h-[30px] items-center justify-between border-b-[1px] border-solid border-red-950"
			>
				<div class="flex items-center">
					<Textcontainer
						width="100"
						min_width="100"
						max_Width="200"
						id={index.toString()}
						text_size="14"
						bind:inputValue={item.label_name}
					/>

					<NeedChooseLabel></NeedChooseLabel>
					<RequiredArea
						bind:product_required={item.required}
						nameYes={item.label_name + 'Choice' + index.toString()}
						idYes={item.label_name + 'yes' + index.toString()}
						nameNo={item.label_name + 'Choice' + index.toString()}
						idNo={item.label_name + 'No' + index.toString()}
						bind:group={item.required}
					></RequiredArea>
				</div>
				<div class="flex items-center justify-end">
					<button
						on:click={() => removeCategory(index)}
						class="flex duration-150 hover:scale-125"
						type="button"
					>
						<img src={transhcan} alt="Trash Can" class="object-cover" />
					</button>
				</div>
			</div>
			<div class="relative flex-col items-start">
				{#each item.item_array as { name }, subIndex}
					<div class="flex w-full items-center space-x-2">
						<div class="flex h-8 w-full items-center px-5">
							<Redradiobox name={item.label_name} id={item.label_name + name} />

							<div class="flex w-full justify-end">
								<input
									class="w-full border-b-[1px] border-solid border-gray-300 border-transparent bg-transparent text-end font-bold text-red-950 underline outline-none"
									type="text"
									placeholder="item{subIndex + 1}"
									bind:value={name}
								/>
							</div>
						</div>
					</div>
				{/each}
				<div class="flex h-[30px] w-full items-center justify-end gap-2">
					<button
						type="button"
						on:click={() => removeNewSubcategory(index)}
						class="m-[3px] flex h-[20px] w-[20px] items-center justify-center rounded-[10px] bg-red-900 px-[0px] py-[0px] text-sm font-bold text-white duration-150 hover:scale-125 hover:border-2 hover:border-PUA-dark-red hover:bg-transparent hover:text-PUA-dark-red"
						>-</button
					>
					<button
						type="button"
						on:click={() => addNewSubcategory(index)}
						class="m-[3px] flex h-[20px] w-[20px] items-center justify-center rounded-[10px] bg-red-900 px-[0px] py-[0px] text-sm font-bold text-white duration-150 hover:scale-125 hover:border-2 hover:border-PUA-dark-red hover:bg-transparent hover:text-PUA-dark-red"
						>+</button
					>
				</div>
			</div>
		</div>
	{/each}
</div>
<div class="flex h-full items-center justify-center">
	<button
		type="button"
		on:click={addNewCategory}
		class="h-[30px] w-[30px] rounded-[15px] bg-red-900 font-bold text-white duration-150 hover:scale-125 hover:border-2 hover:border-PUA-dark-red hover:bg-transparent hover:text-PUA-dark-red"
		>+</button
	>
</div>
