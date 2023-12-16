<script lang="ts">
	import { Input } from '$lib/components/ui/input';
	import transhcan from '$lib/assets/transhcan.svg';
	import close from '$lib/assets/close.svg';
	import { onMount } from 'svelte';
	import { backendPath } from '$lib/components/PUA/env';
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

	let showModel = false; //only 模擬
	let discount_name_Input = '';
	let discount_maxquantity_Input = '';
	let discount_description_Input = '';

	let Status = [{ label: '下架中' }, { label: '已售完' }];

	let buttons = [
		{ id: '1', label: '買二送一' },
		{ id: '2', label: '買二送一' },
		{ id: '3', label: '買二送一' },
		{ id: '4', label: '買二送一' },
		{ id: '5', label: '買二送一' }
	];

	let product_data = {
		store_id: 1,
		product_label_array: [
			{
				item_array: [
					{
						name: 'have'
					},
					{
						name: 'no'
					}
				],
				product_id: 1,
				label_name: 'milk',
				required: true
			},
			{
				item_array: [
					{
						name: 'have'
					},
					{
						name: 'no'
					}
				],
				product_id: 1,
				label_name: 'milk22',
				required: true
			}
		],
		price: 9999,
		product_id: 1,
		name: 'pastasfkjl',
		description: 'tasty pasta ,342424l;',
		stock: 100,
		event_discount_array: [
			{
				discount_max_quantity: 2,
				product_id: 1,
				discount_name: 'new year discount',
				discount_description: 'black tea get two for one free',
				discount_id: 1,
				status: 1
			},
			{
				discount_max_quantity: 2,
				product_id: 1,
				discount_name: 'new year discount',
				discount_description: 'black tea get two for one free',
				discount_id: 1,
				status: 1
			}
		],
		picture: 'https://i.imgur.com/1.jpg',
		status: 1
	};
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

	function addDiscountButton() {
		buttons = [...buttons, { id: '6', label: `買二送一` }];
		return null;
	}

	function toggleModel() {
		showModel = !showModel;
	}
	let k = 0;
	let u = 'j3io';
	onMount(async () => {
		const resp = await fetch(backendPath + '/product/2');
		product_data = await resp.json();
		k = resp.status;
		u = product_data.product_label_array[0].label_name;
	});
	$: product_data;
</script>

<p>{k}</p>
<p>{product_data.product_label_array}</p>
<p>{u}</p>
<div class="flex justify-start">
	<div class="relative left-1/2 top-6 h-full w-4/5 -translate-x-1/2 transform">
		<div class="h-100 text-33 text-PUA-dark-red flex w-full flex-col justify-center">
			<Input
				bind:value={product_data.name}
				type="text"
				placeholder="Enter Product Name"
				class="max-wxs w-full rounded-[0] border-b border-l-0 border-r-0 border-t-0 border-gray-400 text-3xl"
			/>

			<ErrorMsg width={'30'} height={'30'}></ErrorMsg>
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

				<ErrorMsg width={'30'} height={'30'}></ErrorMsg>

				<div class="flex items-center">
					<div class="gap-3 py-[20px] text-3xl font-bold">NT$</div>
					<Input
						value={product_data.price}
						type="text"
						placeholder="Enter price"
						class="ml-[10px] w-[180px] max-w-xs rounded-[0px] border-b border-l-0 border-r-0 border-t-0 border-gray-400 text-[30px]"
					/>
				</div>

				<ErrorMsg width={'30'} height={'30'}></ErrorMsg>
				<div class="w-[250px] text-base text-gray-600">
					<Textarea width="64" bind:value={product_data.description} />
				</div>

				<ErrorMsg width={'30'} height={'30'}></ErrorMsg>
			</div>
			<div class="relative h-full w-full">
				<div class="relative mb-[10px] w-full">
					<!-- <p>{product_data.product_label_array[0].item_array}</p> -->
					<!-- {#each product_label_array as { id, category, subcategories }, index (category)} -->
					{#if product_data}
						{#each product_data.product_label_array as item, index}
							<div class="mb-[15px]">
								<div class="flex h-[30px] items-center border-b-[1px] border-solid border-red-950">
									<div class="flex w-[80%] items-center">
										<!-- <p>{item_array}</p> -->
										<!-- <p>{item.product_id}{item.label_name}{item.required}</p> -->
										<Textcontainer
											width="100"
											min_width="100"
											max_Width="200"
											id={index.toString()}
											text_size="14"
											bind:inputValue={item.label_name}
										/>

										<NeedChooseLabel></NeedChooseLabel>

										<div class="ml-2 mr-2 flex items-center justify-center">
											<Redradiobox
												name={item.label_name + 'Choice'}
												id={item.label_name + 'yes'}
												checked={true}
											/>
										</div>

										<label for="Choice" class="text-red-950">Yes</label>
										<div class="items-cener ml-2 mr-2 flex justify-center">
											<Redradiobox name={item.label_name + 'Choice'} id={item.label_name + 'No'} />
										</div>
										<label for="Choice" class="text-red-950">No</label>
									</div>
									<div class="flex w-1/3 items-center justify-end">
										<button on:click={() => removeCategory(index)} class="flex">
											<img src={transhcan} alt="Trash Can" class="object-cover" />
										</button>
									</div>
								</div>
								<div class="relative ml-[25px] mt-[10px] w-[90%] flex-col items-start">
									<!-- {#each subcategories as subcategory, subIndex (subcategory)} -->
									{#each item.item_array as { name }, subIndex}
										<div
											class="flex w-full items-center space-x-2 border-b-[1px] border-solid border-red-900"
										>
											<Redradiobox name={item.label_name} id={item.label_name + name} />

											<div class="mx-auto flex w-full justify-end">
												<input
													class="w-full border-b-[1px] border-solid border-gray-300 border-transparent bg-transparent text-end underline outline-none"
													type="text"
													placeholder="item{subIndex + 1}"
													value={name}
												/>
											</div>
										</div>
									{/each}
									<div class="flex h-[30px] w-full items-center justify-end">
										<button
											on:click={() => removeNewSubcategory(0)}
											class="m-[3px] h-[20x] w-[20px] rounded-[10px] bg-red-900 px-[0px] py-[0px] text-sm font-bold text-white"
											>-</button
										>
										<button
											on:click={() => addNewSubcategory(0)}
											class="m-[3px] h-[20x] w-[20px] rounded-[10px] bg-red-900 px-[0px] py-[0px] text-sm font-bold text-white"
											>+</button
										>
									</div>
								</div>
							</div>
						{/each}
					{/if}
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

				<div class="border-PUA-stone flex h-[30px] w-full items-center border-b-[1px] border-solid">
					<div class="text-PUA-stone font-bold">Set Status</div>
				</div>
				<div class="m-4 flex justify-center gap-10">
					{#each Status as { label }}
						<div class="flex items-center justify-center">
							<StatusButton text={label} id={label}></StatusButton>
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
				<div class="text-PUA-stone w-1/2 text-left text-xl font-bold">
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

					<ErrorMsg width={'24'} height={'24'}></ErrorMsg>
				</div>
				<div class="mt-2">
					<p class="text-3 font-bold text-red-900">Discount Name</p>
					<input
						type="text"
						class="w-full border-b-[2px] border-red-900 font-bold"
						bind:value={discount_name_Input}
						placeholder=" Enter Discount Name"
					/>

					<ErrorMsg width={'24'} height={'24'}></ErrorMsg>
				</div>
				<div class="mt-2">
					<p class="text-3 font-bold text-red-900">Description</p>
					<input
						type="text"
						class="w-full border-b-[2px] border-red-900 font-bold"
						bind:value={discount_description_Input}
						placeholder=" Enter Discount Description"
					/>
					<ErrorMsg width={'24'} height={'24'}></ErrorMsg>
				</div>
			</div>
			<div class="mt-5 flex items-center justify-between gap-5 text-center">
				<button class="w-4/5 rounded-[20px] bg-gray-200 font-bold text-red-900"> Delete </button>

				<button class="w-4/5 rounded-[20px] bg-orange-700 font-bold text-white"> Save </button>
			</div>
		</div>
	</div>
{/if}
