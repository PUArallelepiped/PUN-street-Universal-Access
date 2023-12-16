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
	type productRespType = {
		store_id: number;
		product_label_array: {
			item_array: {
				name: string;
			}[];
			product_id: number;
			label_name: string;
			required: boolean;
		}[];
		price: number;
		product_id: number;
		name: string;
		description: string;
		stock: number;
		event_discount_array: {
			discount_max_quantity: number;
			product_id: number;
			discount_name: string;
			discount_description: string;
			discount_id: number;
			status: number;
		}[];
		picture: string;
		status: number;
	};
	let product_data: productRespType = {
		store_id: 0,
		product_label_array: [
			{
				item_array: [{ name: '' }],
				product_id: 0,
				label_name: '',
				required: false
			}
		],
		price: 0,
		product_id: 0,
		name: '',
		description: '',
		stock: 0,
		event_discount_array: [
			{
				discount_max_quantity: 0,
				product_id: 0,
				discount_name: '',
				discount_description: '',
				discount_id: 0,
				status: 0
			}
		],
		picture: '',
		status: 0
	};
	function addNewCategory() {
		const newLabel = {
			item_array: [{ name: '' }],
			product_id: product_data.product_id,
			label_name: '',
			required: false
		};

		product_data = {
			...product_data,
			product_label_array: [...product_data.product_label_array, newLabel]
		};
	}

	function addNewSubcategory(index: number) {
		const newItem = { name: '' };
		if (index >= 0 && index < product_data.product_label_array.length) {
			const newArray = [...product_data.product_label_array];
			newArray[index].item_array.push(newItem);
			product_data = {
				...product_data,
				product_label_array: newArray
			};
		}
	}

	function removeCategory(index: number) {
		if (index >= 0 && index < product_data.product_label_array.length) {
			const newArray = product_data.product_label_array.filter((_, i) => i !== index);
			product_data = {
				...product_data,
				product_label_array: newArray
			};
		}
	}

	function removeNewSubcategory(index: number) {
		if (index >= 0 && index < product_data.product_label_array.length) {
			const newArray = [...product_data.product_label_array];

			if (newArray[index].item_array.length > 0) {
				newArray[index].item_array.pop();
				product_data = {
					...product_data,
					product_label_array: newArray
				};
			}
		}
	}

	function addDiscountButton() {
		buttons = [...buttons, { id: '6', label: `買二送一` }];
		return null;
	}
	let k = 0;
	onMount(async () => {
		const backResp = await fetch(backendPath + `/product/8`);
		if (backResp.status == 200) {
			product_data = await backResp.json();
		}
		k = backResp.status;
	});
</script>

<p>{k}</p>
<p>{product_data.description}</p>
<div class="flex h-fit justify-start">
	<div class="relative left-1/2 mt-6 h-full w-4/5 -translate-x-1/2 transform">
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
			<div class="relative h-fit w-full">
				<div class="relative mb-[10px] w-full">
					{#each product_data.product_label_array as item, index}
						<div class="mb-[15px]">
							<div class="flex h-[30px] items-center border-b-[1px] border-solid border-red-950">
								<div class="flex w-[80%] items-center">
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
