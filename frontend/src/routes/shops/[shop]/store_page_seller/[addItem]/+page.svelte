<script lang="ts">
	import { Input } from '$lib/components/ui/input';
	import { onMount } from 'svelte';
	import { backendPath } from '$lib/components/PUA/env';
	import { Textarea, DisCountArea, OkButton, StatusButton, ErrorMsg } from '$lib/index';
	import AddCategoryAndItemArea from '$lib/components/PUA/addCategoryAndItemArea.svelte';
	import ChangeDiscountPage from '$lib/components/PUA/changeDiscountPage.svelte';
	import type { PageData } from './$types';
	export let data: PageData;
	let shop_id = data.shop;
	let item_id = data.item;
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
		store_id: Number(shop_id),
		product_label_array: [
			{
				item_array: [{ name: '' }],
				product_id: 0,
				label_name: '',
				required: true
			}
		],
		price: 0,
		product_id: 0,
		name: '',
		description: '',
		stock: 0,
		event_discount_array: [],
		picture: 'https://i.imgur.com/3i3tyXJ.gif',
		status: 1
	};
	let current_discount_array = {
		discount_max_quantity: 0,
		product_id: 0,
		discount_name: '',
		discount_description: '',
		discount_id: 0,
		status: 1
	};
	let showModel = false;
	let Status = [{ label: '上架中' }, { label: '已售完' }];
	let new_index = 0;
	let discountData = { kind: 'Event Discount', how: 'Get', way: 'FOR FREE ONE' };

	function addDiscountButton() {
		const newLabel = { ...current_discount_array };
		if (product_data.event_discount_array[new_index]) {
			product_data.event_discount_array[new_index] = current_discount_array;
		} else {
			product_data = {
				...product_data,
				event_discount_array: [...product_data.event_discount_array, newLabel]
			};
		}
		return null;
	}
	function delete_Discount() {
		if (new_index >= 0 && new_index < product_data.event_discount_array.length) {
			const newArray = product_data.event_discount_array.filter((_, i) => i !== new_index);
			product_data.event_discount_array = newArray;
		}
		return null;
	}
	function getDiscount(i: number) {
		if (product_data.event_discount_array[i]) {
			current_discount_array = product_data.event_discount_array[i];
		} else {
			current_discount_array = {
				discount_max_quantity: 0,
				product_id: 0,
				discount_name: '',
				discount_description: '',
				discount_id: 0,
				status: 0
			};
		}
		new_index = i;
		return null;
	}

	async function getProductResp() {
		if (item_id != '0') {
			const res = await fetch(backendPath + `/product/` + item_id);

			if (res.status == 200) {
				product_data = await res.json();
			}
			console.log(product_data);
		}
		return;
	}

	async function PostProductResp() {
		let post_status = fetch(backendPath + `/store/` + shop_id + `/add-product`, {
			method: 'POST',
			body: JSON.stringify(product_data)
		});
		window.location.href = './';
		console.log(post_status);
	}

	onMount(async () => {
		getProductResp();
	});
</script>

{#await getProductResp() then}
	<form on:submit={PostProductResp}>
		<div class="flex h-fit justify-start">
			<div class="relative left-1/2 mt-6 h-full w-4/5 -translate-x-1/2 transform">
				<div class="h-100 text-33 flex w-full flex-col justify-center text-PUA-dark-red">
					<Input
						required
						bind:value={product_data.name}
						type="text"
						placeholder="Enter Product Name"
						class="max-wxs peer w-full rounded-[0] border-b border-l-0 border-r-0 border-t-0 border-gray-400 text-3xl"
					/>
					<div class="invisible py-4 peer-invalid:visible">
						<ErrorMsg width={'30'} height={'30'} text={`CANNOT BE EMPTY`}></ErrorMsg>
					</div>
				</div>
				<div class="flex h-full w-full">
					<div class="relative h-full w-[500px]">
						<div class=" flex h-[250px] w-[250px] rounded-lg bg-gray-300 shadow-inner">
							{#if !product_data.picture}
								<div class="flex h-full w-full items-center justify-center">
									<label
										for="fileInput"
										class="h-11 cursor-pointer rounded-[20px] bg-gray-400 px-12 py-2 text-white shadow-md"
										>Upload Image</label
									>
									<input
										type="file"
										id="fileInput"
										accept="image/png, image/jpeg"
										class="absolute right-0 top-0 cursor-pointer font-bold opacity-0"
									/>
								</div>
							{:else}
								<div class="absolute z-10 h-[250px] w-[250px] bg-opacity-0">
									<label for="fileInput" class=" block h-full w-full cursor-pointer bg-opacity-0"
									></label>
									<input
										type="file"
										id="fileInput"
										accept="image/png, image/jpeg"
										class="absolute right-0 top-0 cursor-pointer font-bold opacity-0"
									/>
								</div>
								<img
									src={product_data.picture}
									alt=""
									class="flex h-full w-full rounded-lg object-cover"
								/>
							{/if}
						</div>
						<div class="flex w-64 flex-wrap items-center">
							<div class="gap-3 py-[20px] text-3xl font-bold">NT$</div>
							<input
								required
								bind:value={product_data.price}
								type="number"
								min="0"
								max="999999999"
								placeholder="Enter price"
								class="peer ml-[10px] w-[180px] max-w-xs rounded-[0px] border-b border-l-0 border-r-0 border-t-0 border-gray-400 bg-transparent text-[30px]"
							/>
							<div class="invisible h-8 w-64 peer-invalid:visible">
								{#if !product_data.price}
									<ErrorMsg width={'30'} height={'30'} text={'CANNOT BE EMPTY'}></ErrorMsg>
								{:else if product_data.price > 999999999 || product_data.price < 0}
									<ErrorMsg width={'30'} height={'30'} text={'EXCEEDS LIMIT'}></ErrorMsg>
								{/if}
							</div>
						</div>
						<div class="w-[250px] text-base text-gray-600">
							<Textarea width="64" bind:value={product_data.description} required={true} />
						</div>
					</div>
					<div class="relative h-fit w-full">
						<AddCategoryAndItemArea
							bind:category_item={product_data.product_label_array}
							bind:product_id={product_data.product_id}
						></AddCategoryAndItemArea>
						<DisCountArea
							toggleModel={() => {
								showModel = !showModel;
								return null;
							}}
							bind:discount={product_data.event_discount_array}
							addDiscountButton={getDiscount}
							addSign={true}
							type={false}
						></DisCountArea>

						<div
							class="flex h-[30px] w-full items-center border-b-[1px] border-solid border-PUA-stone"
						>
							<div class="font-bold text-PUA-stone">Set Status</div>
						</div>
						<div class="m-4 flex justify-center gap-10">
							{#each Status as { label }, index}
								<div class="flex items-center justify-center">
									{#if index + 1 == product_data.status}
										<StatusButton
											text={label}
											id={label}
											checked={true}
											value={index + 1}
											bind:group={product_data.status}
										></StatusButton>
									{:else}
										<StatusButton
											text={label}
											id={label}
											value={index + 1}
											bind:group={product_data.status}
										></StatusButton>
									{/if}
								</div>
							{/each}
						</div>
						<div class="mb-5 flex items-center justify-center">
							<OkButton
								onclick={() => {
									return null;
								}}
								text="Add Product"
								type={'submit'}
							></OkButton>
						</div>
					</div>
				</div>
			</div>
		</div>
	</form>
	<ChangeDiscountPage
		bind:changePageData={current_discount_array}
		{discountData}
		bind:showModel
		dis_haved={false}
		add_Discount={addDiscountButton}
		{delete_Discount}
	></ChangeDiscountPage>
{/await}
