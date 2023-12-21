<script lang="ts">
	import { Counter } from '$lib';
	import { DiscountArea, Checkcontainer, OkButton, NeedChooseLabel } from '$lib/index';
	import type { PageData } from './$types';
	import { onMount } from 'svelte';
	import { backendPath } from '$lib/components/PUA/env';
	export let data: PageData;
	let item_id = data.item;
	let userID = '1';
	let get_event_discount_id = 1;
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
	type productCartPostType = {
		cart_id: number;
		product_id: number;
		customer_id: number;
		product_quantity: number;
		discount_id: number | null;
		store_id: number;
	};
	let product: productRespType = {
		store_id: 1,
		product_label_array: [
			{
				item_array: [
					{
						name: 'have'
					},
					{
						name: 'false'
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
						name: 'false'
					}
				],
				product_id: 1,
				label_name: 'milk22',
				required: true
			}
		],
		price: 9999,
		product_id: 1,
		name: 'pasta',
		description: 'tasty pasta',
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

	let productCartPost: productCartPostType = {
		cart_id: 0,
		product_id: 1,
		customer_id: 1,
		product_quantity: 1,
		discount_id: 1,
		store_id: 1
	};
	async function getProductResp() {
		const res = await fetch(backendPath + `/product/` + item_id);

		if (res.status == 200) {
			product = await res.json();
		}
		console.log(product);
		return;
	}
	async function PostProductResp() {
		productCartPost = {
			cart_id: 0,
			product_id: product.product_id,
			customer_id: Number(userID),
			product_quantity: Number(productCartPost.product_quantity),
			discount_id: get_event_discount_id,
			store_id: product.store_id
		};
		console.log(productCartPost);
		let post_status = fetch(backendPath + `/customer/` + userID + `/cart`, {
			method: 'POST',
			body: JSON.stringify(productCartPost)
		});
		console.log(post_status);
	}
	onMount(async () => {
		getProductResp();
	});
</script>

{#await getProductResp() then}
	<div class="flex justify-center">
		<div class="my-6 flex h-full w-4/5 flex-col gap-8">
			<div class=" text-PUA-dark-red flex w-full items-center text-4xl">
				{product.name}
			</div>

			<div class="flex gap-16">
				<div class="">
					<img src={product.picture} alt="" class="mt-100 flex h-60 w-60 rounded-lg object-cover" />
					<div class="text-PUA-dark-red flex items-baseline gap-3 py-5 font-bold">
						<p class="text-2xl">NT$</p>
						<p class="text-4xl">{product.price}</p>
					</div>
					<div class="w-[250px] text-justify text-base text-gray-600">
						{product.description}
					</div>
				</div>
				<div class=" flex w-full flex-col gap-4">
					<div class=" flex w-full flex-col gap-4">
						{#each product.product_label_array as { required, label_name, item_array }}
							<div class="">
								<div class="flex items-center">
									<div class="text-PUA-stone font-bold">{label_name}</div>
									{#if required}
										<NeedChooseLabel></NeedChooseLabel>
									{/if}
								</div>
								<div class="bg-PUA-dark-red h-[1px]"></div>
								<div class="flex flex-col">
									{#each item_array as { name }}
										<Checkcontainer category={label_name} subcategory={name}></Checkcontainer>
									{/each}
								</div>
							</div>
						{/each}
					</div>
					<div>
						<DiscountArea
							bind:group={get_event_discount_id}
							discount={product.event_discount_array}
							toggleModel={() => {
								return null;
							}}
							addDiscountButton={() => {
								return null;
							}}
							addSign={false}
							type={true}
						></DiscountArea>
					</div>
					<div class="  flex flex-wrap items-center justify-around gap-4">
						<Counter bind:count={productCartPost.product_quantity} />

						<OkButton onclick={PostProductResp} text="Add Cart"></OkButton>
					</div>
				</div>
			</div>
		</div>
	</div>
{/await}
