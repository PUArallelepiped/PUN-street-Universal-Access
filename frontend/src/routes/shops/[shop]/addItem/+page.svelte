<script lang="ts">
	import { Input } from '$lib/components/ui/input';
	import { onMount } from 'svelte';
	import { backendPath } from '$lib/components/PUA/env';
	import { Textarea, DisCountArea, OkButton, StatusButton, ErrorMsg } from '$lib/index';
	import AddCategoryAndItemArea from '$lib/components/PUA/addCategoryAndItemArea.svelte';
	let showModel = false; //only 模擬
	let Status = [{ label: '上架中' }, { label: '已售完' }]; //1上架中  2已售完
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

	function addDiscountButton() {
		const newLabel = {
			discount_max_quantity: 0,
			product_id: 1,
			discount_name: '',
			discount_description: 'buy 2 get 1 free', //simulate
			discount_id: 1,
			status: 1
		};

		product_data = {
			...product_data,
			event_discount_array: [...product_data.event_discount_array, newLabel]
		};
		return null;
	}
	onMount(async () => {
		const Resp = await fetch(`/product.json`);
		product_data = await Resp.json();
		const backResp = await fetch(backendPath + `/product/2`);
		if (backResp.status == 200) {
			product_data = await backResp.json();
		}
	});
</script>

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
					class="mt-100 flex h-[250px] w-[250px] items-center justify-center rounded-lg bg-gray-300 shadow-inner"
				>
					{#if !product_data.picture}
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
					{:else}
						<img
							src={product_data.picture}
							alt=""
							class="mt-100 flex h-full w-full rounded-lg object-cover"
						/>
					{/if}
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
					{addDiscountButton}
					addSign={true}
					type={false}
				></DisCountArea>

				<div class="border-PUA-stone flex h-[30px] w-full items-center border-b-[1px] border-solid">
					<div class="text-PUA-stone font-bold">Set Status</div>
				</div>
				<div class="m-4 flex justify-center gap-10">
					{#each Status as { label }, index}
						<div class="flex items-center justify-center">
							{#if index + 1 == product_data.status}
								<StatusButton text={label} id={label} checked={true}></StatusButton>
							{:else}
								<StatusButton text={label} id={label}></StatusButton>
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
					></OkButton>
				</div>
			</div>
		</div>
	</div>
</div>
