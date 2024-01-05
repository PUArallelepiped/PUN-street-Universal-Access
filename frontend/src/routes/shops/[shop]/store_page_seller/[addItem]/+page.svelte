<script lang="ts">
	import { onMount } from 'svelte';
	import { backendPath } from '$lib/components/PUA/env';
	import { Textarea, DisCountArea, OkButton, StatusButton, ErrorMsg } from '$lib/index';
	import AddCategoryAndItemArea from '$lib/components/PUA/addCategoryAndItemArea.svelte';
	import ChangeDiscountPage from '$lib/components/PUA/changeDiscountPage.svelte';
	import type { PageData } from './$types';
	import { goto } from '$app/navigation';
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
		picture: '',
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
		if (imageFile) {
			const formData = new FormData();
			formData.append('file', imageFile[0]);
			let url = await fetch(backendPath + `/upload`, {
				method: 'POST',
				body: formData
			});
			product_data.picture = await url.json();
		}
		let post_status = await fetch(backendPath + `/store/` + shop_id + `/add-product`, {
			method: 'POST',
			body: JSON.stringify(product_data)
		});
		goto('/shops/' + shop_id + '/store_page_seller');
		console.log(post_status);
	}

	const onFileSelected = (e: Event) => {
		// ts too hard i give up
		if (e.target == null) return;
		let image = (e.target as HTMLInputElement).files[0];
		let reader = new FileReader();
		reader.readAsDataURL(image);
		reader.onload = (e) => {
			avatar = e.target.result;
		};
	};

	let imageFile: FileList;
	let avatar: string;

	onMount(async () => {
		getProductResp();
	});
</script>

{#await getProductResp() then}
	<form on:submit={PostProductResp}>
		<div class="flex h-fit justify-start pb-10">
			<div class="relative left-1/2 mt-6 h-full w-4/5 -translate-x-1/2 transform space-y-8">
				<div
					class=" text-PUA-dark-red flex w-full flex-col justify-center rounded-lg bg-white px-4 pb-0 pt-4 shadow"
				>
					<input
						required
						bind:value={product_data.name}
						type="text"
						placeholder="Enter Product Name"
						class="max-wxs peer w-full rounded-[0] border-b border-l-0 border-r-0 border-t-0 border-gray-400 text-4xl outline-none"
					/>
					<div class="invisible py-4 peer-invalid:visible">
						<ErrorMsg width={'28'} height={'28'} text={`CANNOT BE EMPTY`}></ErrorMsg>
					</div>
				</div>
				<div class="flex h-full w-full gap-16">
					<div class="relative h-full rounded-lg bg-white p-4 shadow">
						<div class=" flex h-60 w-60 rounded-lg bg-gray-300 shadow-inner">
							<div class="absolute z-10 h-[250px] w-[250px] bg-opacity-0">
								<label for="fileInput" class=" block h-full w-full cursor-pointer bg-opacity-0"
								></label>
								<input
									type="file"
									id="fileInput"
									accept="image/png, image/jpeg, image/jpg, image/gif"
									bind:files={imageFile}
									on:change={(e) => {
										onFileSelected(e);
									}}
									class="absolute right-0 top-0 cursor-pointer font-bold opacity-0"
								/>
							</div>
							{#if avatar}
								<img src={avatar} alt="" class="flex h-full w-full rounded-lg object-cover" />
							{:else}
								<img
									src={product_data.picture}
									alt=""
									class="flex h-full w-full rounded-lg object-cover"
								/>
							{/if}
						</div>
						<div class="text-PUA-dark-red flex w-64 flex-wrap items-baseline gap-3 pt-5 font-bold">
							<div class="text-2xl">NT$</div>
							<input
								required
								bind:value={product_data.price}
								type="number"
								min="0"
								max="999999999"
								placeholder="Enter price"
								class="peer w-48 rounded-[0px] border-b border-l-0 border-r-0 border-t-0 border-gray-400 bg-transparent text-4xl outline-none placeholder:text-3xl"
							/>
							<div class="invisible h-8 w-64 peer-invalid:visible">
								{#if !product_data.price}
									<ErrorMsg width={'30'} height={'30'} text={'CANNOT BE EMPTY'}></ErrorMsg>
								{:else if product_data.price > 999999999 || product_data.price < 0}
									<ErrorMsg width={'30'} height={'30'} text={'EXCEEDS LIMIT'}></ErrorMsg>
								{/if}
							</div>
						</div>
						<div class=" text-base text-gray-600">
							<Textarea width="64" bind:value={product_data.description} required={true} />
						</div>
					</div>
					<div class="relative h-fit w-full rounded-lg bg-white p-4 shadow">
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
							class="border-PUA-stone flex h-[30px] w-full items-center border-b-[1px] border-solid"
						>
							<div class="text-PUA-stone font-bold">Set Status</div>
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
								text="Save"
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
		add_Discount={addDiscountButton}
		{delete_Discount}
	></ChangeDiscountPage>
{/await}
