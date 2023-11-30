<script lang="ts">
	// import error from '$lib/assets/error.png';
	import { Input } from '$lib/components/ui/input';
	import transhcan from '$lib/assets/transhcan.png';
	import close from '$lib/assets/close.png';

	import {
		Redradiobox,
		Textcontainer,
		Textarea,
		DisCountArea,
		OkButton,
		StatusButton,
		ErrorMsg,
		NeedChooseLabel
	} from '$lib/index';

	let data = [{ id: 0, category: '0', subcategories: ['1', '2', '3'] }];

	let nextCategoryId = data.length + 1;

	function addNewCategory() {
		data = [...data, { id: nextCategoryId, category: `${nextCategoryId}`, subcategories: [] }];
		nextCategoryId++;
	}

	function addNewSubcategory(index: number) {
		const currentData = data[index];
		data[index] = {
			...currentData,
			subcategories: [
				...currentData.subcategories,
				(currentData.subcategories.length + 1).toString()
			]
		};
	}

	function removeCategory(id: number) {
		data = data.filter((cat) => cat.id !== id);
	}

	function removeNewSubcategory(index: number) {
		const currentData = data[index];

		if (currentData.subcategories.length > 0) {
			currentData.subcategories.pop();
		}

		data[index] = { ...currentData };
	}

	let buttons = [
		{ id: 1, label: '買二送一' },
		{ id: 2, label: '買二送一' },
		{ id: 3, label: '買二送一' },
		{ id: 4, label: '買二送一' },
		{ id: 5, label: '買二送一' }
	];

	let Status = [{ label: '下架中' }, { label: '已售完' }];

	function addDiscountButton() {
		buttons = [...buttons, { id: 6, label: `買二送一` }];
		return null;
	}

	let showModel = false; //only 模擬
	let discount_name_Input = '';
	let discount_maxquantity_Input = '';
	let discount_description_Input = '';
	function toggleModel() {
		showModel = !showModel;
	}
</script>

<div class="flex justify-start">
	<div class="relative left-1/2 top-6 h-full w-4/5 -translate-x-1/2 transform">
		<div class="h-100 text-33 flex w-full flex-col justify-center text-red-950">
			<Input
				type="text"
				placeholder="Enter Product Name"
				class="max-wxs w-full rounded-[0] border-b border-l-0 border-r-0 border-t-0 border-gray-400 text-3xl"
			/>

			<ErrorMsg height={'10'} width={'10'}></ErrorMsg>
		</div>
		<div class="flex h-full w-full">
			<div class="relative h-full w-[500px]">
				<div
					class="mt-100 flex h-[250px] w-[250px] items-center justify-center rounded-[10px] bg-gray-300 shadow-inner"
				>
					<button class="h-11 rounded-[20px] bg-gray-400 px-[50px] py-[3px] text-white shadow-md"
						>Upload Image</button
					>
				</div>

				<ErrorMsg height={'10'} width={'10'}></ErrorMsg>

				<div class="flex items-center">
					<div class="py-[20px] text-3xl font-bold text-red-950">NT$</div>
					<Input
						type="text"
						placeholder="Enter price"
						class="ml-[10px] w-3/5 max-w-xs rounded-[0px] border-b border-l-0 border-r-0 border-t-0 border-gray-400 text-[30px]"
					/>
				</div>

				<ErrorMsg height={'10'} width={'10'}></ErrorMsg>

				<div class="w-[250px] text-base text-gray-600">
					<Textarea width="250" />
				</div>

				<ErrorMsg height={'10'} width={'10'}></ErrorMsg>
			</div>
			<div class="relative h-full w-full">
				<div class="relative mb-[10px] w-full">
					{#each data as { id, category, subcategories }, index (category)}
						<div class="mb-[15px]">
							<div class="flex h-[30px] items-center border-b-[1px] border-solid border-red-950">
								<div class="flex w-[80%] items-center">
									<Textcontainer
										width="100"
										min_width="100"
										max_Width="200"
										id={category + 'label'}
										text_size="14"
									/>
									<NeedChooseLabel></NeedChooseLabel>

									<div class="ml-2 mr-2 flex items-center justify-center">
										<Redradiobox name={category + 'Choice'} id={category + 'yes'} checked={true} />
									</div>

									<label for="Choice" class="text-red-950">Yes</label>
									<div class="ml-2 mr-2 flex items-center justify-center">
										<Redradiobox name={category + 'Choice'} id={category + 'No'} />
									</div>
									<label for="Choice" class="text-red-950">No</label>
								</div>
								<div class="flex w-1/3 items-center justify-end">
									<button on:click={() => removeCategory(id)} class="flex">
										<img src={transhcan} alt="Trash Can" class="object-cover" />
									</button>
								</div>
							</div>
							<div class="relative ml-[25px] mt-[10px] w-[90%] flex-col items-start">
								{#each subcategories as subcategory, subIndex (subcategory)}
									<div
										class="flex w-full items-center space-x-2 border-b-[1px] border-solid border-red-900"
									>
										<Redradiobox name={category} id={category + subcategory} />

										<div class="mx-auto flex w-full justify-end">
											<input
												class="w-full border-0 border-b-4 border-solid border-transparent bg-transparent text-end underline outline-none"
												type="text"
												placeholder="item{subIndex + 1}"
											/>
										</div>
									</div>
								{/each}
								<div class="flex h-[30px] w-full items-center justify-end">
									<button
										on:click={() => removeNewSubcategory(index)}
										class="m-[3px] h-[20x] w-[20px] rounded-[10px] bg-red-900 px-[0px] py-[0px] text-sm font-bold text-white"
										>-</button
									>
									<button
										on:click={() => addNewSubcategory(index)}
										class="m-[3px] h-[20x] w-[20px] rounded-[10px] bg-red-900 px-[0px] py-[0px] text-sm font-bold text-white"
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
				<DisCountArea
					toggleModel={() => {
						showModel = !showModel;
						return null;
					}}
					discount={buttons}
					{addDiscountButton}
					addSign={true}
					type={false}
				></DisCountArea>

				<div
					class="mt-4 flex h-[30px] w-full items-center border-b-[1px] border-solid border-PUA-stone"
				>
					<div class="font-bold text-PUA-stone">Set Status</div>
				</div>
				<div class="m-4 flex justify-between px-28">
					{#each Status as { label }}
						<div class="flex items-center justify-center">
							<StatusButton text={label}></StatusButton>
						</div>
					{/each}
				</div>
				<div class="flex items-center justify-center">
					<div class="flex flex-col">
						<OkButton
							onclick={() => {
								return null;
							}}
							text="Add Product"
						></OkButton>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>

{#if showModel}
	<div
		class="fixed left-0 top-0 flex h-full w-full items-center justify-center bg-black bg-opacity-50 backdrop-blur"
	>
		<div class="rounded bg-white p-5">
			<div class="flex">
				<div class="w-1/2 text-left text-xl font-bold text-PUA-stone">
					<h2>Add A Discount</h2>
				</div>
				<div class="flex w-1/2 justify-end">
					<button on:click={toggleModel}>
						<img src={close} alt="" class="h-[25px] object-cover" />
					</button>
				</div>
			</div>
			<div class="relative ml-16 mr-16">
				<div class="flex w-full items-center justify-center pb-2 pt-2">
					<div class="flex w-64 rounded-[10px] border-[3px] border-red-900 p-3 text-red-900">
						<div class="w-2/5 text-center font-bold">Event Discount</div>
						<div class="ml-1 mr-1 border-r-[2px] border-red-900"></div>
						<div class="flex w-3/5 flex-wrap items-center justify-center text-center font-bold">
							<p>Get</p>
							<p class="ml-1 mr-1 text-xl">{discount_maxquantity_Input}</p>
							<p>FOR FREE ONE</p>
						</div>
					</div>
				</div>
				<div class="mt-2">
					<p class="text-3 font-bold text-red-900">Enter Max Quantity</p>
					<input
						type="text"
						class="w-full border-b-[2px] border-red-900 font-bold"
						bind:value={discount_maxquantity_Input}
						placeholder=" Enter Max Quantity"
					/>

					<ErrorMsg height={'6'} width={'6'}></ErrorMsg>
				</div>
				<div class="mt-2">
					<p class="text-3 font-bold text-red-900">Discount Name</p>
					<input
						type="text"
						class="w-full border-b-[2px] border-red-900 font-bold"
						bind:value={discount_name_Input}
						placeholder=" Enter Discount Name"
					/>

					<ErrorMsg height={'6'} width={'6'}></ErrorMsg>
				</div>
				<div class="mt-2">
					<p class="text-3 font-bold text-red-900">Description</p>
					<input
						type="text"
						class="w-full border-b-[2px] border-red-900 font-bold"
						bind:value={discount_description_Input}
						placeholder=" Enter Discount Description"
					/>
					<ErrorMsg height={'6'} width={'6'}></ErrorMsg>
				</div>
			</div>
			<div class="mt-5 flex items-center justify-between gap-5 text-center">
				<button class="w-4/5 rounded-[20px] bg-gray-200 font-bold text-red-900"> Delete </button>

				<button class="w-4/5 rounded-[20px] bg-orange-700 font-bold text-white"> Save </button>
			</div>
		</div>
	</div>
{/if}
