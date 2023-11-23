<script lang="ts">
	import error from '$lib/assets/error.png';
	import { Input } from '$lib/components/ui/input';
	import transhcan from '$lib/assets/transhcan.png';
	import Counter from '$lib/components/PUA/counter.svelte';
	import Discountbutton from '$lib/components/PUA/discountButton.svelte';

	let data = [{ id: 0, category: '', subcategories: [1, 2, 3] }];

	//id 辨識唯一data

	let nextCategoryId = data.length + 1;

	function addNewCategory() {
		data = [
			...data,
			{ id: nextCategoryId, category: `New Category ${nextCategoryId}`, subcategories: [] }
		];
		nextCategoryId++;
	}

	function addNewSubcategory(index: number) {
		const currentData = data[index];
		data[index] = {
			...currentData,
			subcategories: [...currentData.subcategories, currentData.subcategories.length + 1]
		};
	}

	function removeCategory(id: number) {
		data = data.filter((cat) => cat.id !== id);
	}

	function removeNewSubcategory(index: number) {
		const currentData = data[index];

		// 如果 subcategories 不為空，則移除最後一個元素
		if (currentData.subcategories.length > 0) {
			currentData.subcategories.pop();
		}

		// 更新 data
		data[index] = { ...currentData };
	}

	let buttons = [
		{ id: 1, label: '買二送一' },
		{ id: 2, label: '買二送一' },
		{ id: 3, label: '買二送一' },
		{ id: 4, label: '買二送一' },
		{ id: 5, label: '買二送一' }
	];

	function addDiscountButton() {
		buttons = [...buttons, { id: 6, label: `買二送一` }];
	}
</script>

<div class="flex justify-start">
	<div class="relative left-1/2 top-[25px] h-full w-[80%] -translate-x-1/2 transform">
		<div class="flex h-[100px] w-full flex-col justify-center text-[33px] text-red-950">
			<Input
				type="text"
				placeholder="Enter Product Name"
				class="w-full rounded-[0px] border-b border-l-0 border-r-0 border-t-0 border-gray-400 text-[33px]"
			/>
			<div class="flex items-center">
				<img src={error} alt="Error" class="h-[35px] w-[35px] object-cover object-cover" />

				<p class="text-[20px] font-bold text-red-500">Error message</p>
			</div>
		</div>
		<div class="flex h-full w-full">
			<div class="relative h-full w-[500px]">
				<div
					class="mt-100 flex h-[250px] w-[250px] items-center justify-center rounded-[10px] bg-gray-300 shadow-inner"
				>
					<button
						class="h-[45px] rounded-[20px] bg-gray-400 px-[50px] py-[3px] text-white shadow-md"
						>Upload Image</button
					>
				</div>
				<div class="flex items-center">
					<img src={error} alt="Error" class="h-[35px] w-[35px] object-cover object-cover" />
					<p class="text-[20px] font-bold text-red-500">Error message</p>
				</div>

				<div class="flex items-center">
					<div class="py-[20px] text-[30px] font-bold text-red-950">NT$</div>
					<Input
						type="text"
						placeholder="Enter price"
						class="ml-[10px] w-[60%] max-w-xs rounded-[0px] border-b border-l-0 border-r-0 border-t-0 border-gray-400 text-[30px]"
					/>
				</div>
				<div class="flex items-center">
					<img src={error} alt="Error" class="h-[35px] w-[35px] object-cover object-cover" />
					<p class="text-[20px] font-bold text-red-500">Error message</p>
				</div>

				<div class="w-[250px] overflow-hidden text-justify text-[15px] text-gray-600">
					<Input
						type="text"
						placeholder="Enter text"
						class="max-w-xs rounded-[0px] border-b border-l-0 border-r-0 border-t-0 border-gray-400"
					/>
				</div>
				<div class="flex items-center">
					<img src={error} alt="Error" class="h-[35px] w-[35px] object-cover object-cover" />
					<p class="text-[20px] font-bold text-red-500">Error message</p>
				</div>
			</div>
			<div class="relative h-full w-full">
				<div class="relative mb-[10px] w-full">
					{#each data as { id, category, subcategories }, index (category)}
						<div class=" mb-[15px]">
							<div class="w- flex h-[30px] items-center border-b-[1px] border-solid border-red-950">
								<Input
									type="text"
									placeholder="Enter label"
									class="h-[20px] w-[30%] rounded-[0px] border-b border-l-0 border-r-0 border-t-0 border-gray-400"
								/>
								<div
									class="ml-[10px] rounded-[20px] border-[2px] border-solid border-orange-700 px-[8px] py-[0px] text-[9px] font-bold text-orange-700"
								>
									必填
								</div>
								<input
									type="radio"
									id="contactChoice1"
									name={category}
									value="Yes"
									class="m-[10px] h-[20px] w-[20px]"
								/>
								<label for="Choice">Yes</label>
								<input
									type="radio"
									id="contactChoice1"
									name={category}
									value="No"
									class="m-[10px] h-[20px] w-[20px]"
								/>
								<label for="Choice">No</label>

								<button
									on:click={() => removeCategory(id)}
									class="flex w-[35%] items-center justify-end"
								>
									<img src={transhcan} alt="Trash Can" class="object-cover" />
								</button>
							</div>
							<div class="relative ml-[25px] mt-[10px] w-[90%] flex-col items-start">
								{#each subcategories as subcategory, subIndex (subcategory)}
									<div
										class="m-1 flex w-full items-center space-x-2 border-b-[1px] border-solid border-amber-900"
									>
										<input type="radio" name={category} class="h-[20px] w-[20px]" />
										<div class="flex h-[30px] w-full items-center justify-end">
											<Input
												type="text"
												placeholder="item{subIndex + 1}"
												class="h-[10px] w-[15%] rounded-[0px] border-b border-l-0 border-r-0 border-t-0 border-gray-400"
											/>
										</div>
									</div>
								{/each}
								<div class="flex h-[30px] w-full items-center justify-end">
									<button
										on:click={() => removeNewSubcategory(index)}
										class="m-[3px] h-[20x] w-[20px] rounded-[10px] bg-red-900 px-[0px] py-[0px] text-[13px] font-bold text-white"
										>-</button
									>
									<button
										on:click={() => addNewSubcategory(index)}
										class="m-[3px] h-[20x] w-[20px] rounded-[10px] bg-red-900 px-[0px] py-[0px] text-[13px] font-bold text-white"
										>+</button
									>
								</div>
							</div>
						</div>
					{/each}
				</div>
				<div class="flex h-full items-center justify-center">
					<button
						on:click={addNewCategory}
						class="h-[30px] w-[30px] rounded-[15px] bg-red-900 font-bold text-white">+</button
					>
				</div>
				<div class="relative h-full w-full">
					<div class="flex h-[30px] w-full items-center border-b-[1px] border-solid border-red-950">
						<div class=" font-bold text-amber-900">Add Discount</div>
					</div>
					<div class="flex items-center justify-center">
						<div class="flex w-[63%] flex-wrap">
							{#each buttons as { label }}
								<div class="mt-[10px] flex justify-center">
									<Discountbutton text={label} />
								</div>
							{/each}
							<div class="flex w-[90px] items-center">
								<button
									on:click={addDiscountButton}
									class="m-[3px] h-[20px] w-[20px] rounded-[10px] bg-lime-800 text-center text-[13px] font-bold text-white"
									>+</button
								>
							</div>
						</div>
					</div>
				</div>

				<div class="flex h-[30px] w-full items-center border-b-[1px] border-solid border-red-950">
					<div class=" font-bold text-amber-900">Set Stock</div>
				</div>
				<div class="flex items-center justify-center">
					<div class="mt-[20px] flex flex-col">
						<Counter />
						<button class="mt-[20px] rounded-[20px] bg-orange-700 p-2 text-white"
							>Add Product</button
						>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>

<br />
<br />
<br />

<style>
	input[type='radio'] {
		accent-color: rgb(120, 40, 40);
	}
</style>
