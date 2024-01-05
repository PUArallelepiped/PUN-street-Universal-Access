<script lang="ts">
	import { Counter } from '$lib';
	import { DiscountArea, Checkcontainer, OkButton, NeedChooseLabel } from '$lib/index';
	import type { PageData } from './$types';
	import { onMount } from 'svelte';
	import { backendPath } from '$lib/components/PUA/env';
	import { goto } from '$app/navigation';
	import { getId } from '$lib/components/PUA/getId';
	import right_allow from '$lib/assets/right_allow.svg';
	import left_allow from '$lib/assets/left_allow.svg';
	import RecommandCard from '$lib/components/PUA/recommandCard.svelte';
	export let data: PageData;
	let shop_id = data.shop;
	let item_id = data.item;
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
		return;
	}
	async function PostProductResp() {
		try {
			const user_id = (await getId()).valueOf();
			productCartPost = {
				cart_id: 0,
				product_id: product.product_id,
				customer_id: Number(user_id),
				product_quantity: Number(productCartPost.product_quantity),
				discount_id: get_event_discount_id,
				store_id: product.store_id
			};
			let post = fetch(backendPath + `/customer/` + user_id + `/cart`, {
				method: 'POST',
				body: JSON.stringify(productCartPost)
			});
			if ((await post).status == 200) {
				goto('/shops/' + shop_id);
			}
		} catch (e) {
			goto('/login');
		}
	}

	let screenWidth = 0;

	onMount(async () => {
		getProductResp();
		screenWidth = window.innerWidth - 94;
	});

	let k = 0;

	let myElement: HTMLDivElement | null = null;
	let myElementWidth = 0;
	$: {
		if (myElement) {
			myElementWidth = myElement.clientWidth;
		}
	}
</script>

{#await getProductResp() then}
	<div class="flex justify-center">
		<div class="my-6 mb-10 flex h-full w-4/5 flex-col gap-8">
			<div class="rounded-lg bg-white p-4 shadow">
				<div class=" flex w-full items-center text-4xl text-PUA-dark-red">
					{product.name}
				</div>
			</div>

			<div class="flex gap-16">
				<div class="flex h-fit justify-center rounded-lg bg-white p-4 shadow">
					<div>
						<img
							src={product.picture}
							alt=""
							class="mt-100 flex h-60 w-60 rounded-lg object-cover"
						/>
						<div class="flex items-baseline gap-3 py-5 font-bold text-PUA-dark-red">
							<p class="text-2xl">NT$</p>
							<p class="text-4xl">{product.price}</p>
						</div>
						<div class="w-[250px] text-justify text-base text-gray-600">
							{product.description}
						</div>
					</div>
				</div>
				<div class=" flex w-full flex-col gap-4 rounded-lg bg-white p-4 shadow">
					<div class=" flex w-full flex-col gap-4">
						{#each product.product_label_array as { required, label_name, item_array }}
							<div class="">
								<div class="flex items-center">
									<div class="font-bold text-PUA-stone">{label_name}</div>
									{#if required}
										<NeedChooseLabel></NeedChooseLabel>
									{/if}
								</div>
								<div class="h-[0.04rem] bg-PUA-dark-red"></div>
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

	<p class="p-3 text-2xl font-bold text-PUA-dark-red">Recommand</p>
	<div class=" flex h-60 w-full justify-center">
		<div class=" absolute z-10 h-52 w-full">
			<div class="flex h-full w-full justify-between">
				<div>
					<button
						on:click={() => {
							if (k != 0) {
								k = k + 144 + 8;
							}
						}}
						disabled={k === 0}
						class="flex h-full w-10 items-center justify-center bg-gray-300 shadow-lg shadow-zinc-400 hover:bg-zinc-400 hover:disabled:bg-gray-300"
					>
						<img src={left_allow} alt="" class=" h-7 w-7" />
					</button>
				</div>

				<button
					on:click={() => {
						if (k - 144 - 8 >= -myElementWidth) {
							k = k - 144 - 8;
						}
					}}
					disabled={k - 144 - 8 <= -myElementWidth}
					class="group flex h-full w-10 items-center justify-center bg-gray-300 shadow-lg shadow-zinc-400 hover:bg-zinc-400 hover:disabled:bg-gray-300"
				>
					<img src={right_allow} alt="" class=" h-7 w-7" />
				</button>
			</div>
		</div>

		<div class="absolute z-20 flex h-56" style={`width: ${screenWidth}px`}>
			<div class="mx-2 h-full w-full overflow-hidden">
				<div
					class="flex w-fit gap-2 duration-300"
					style={`transform: translateX(${k}px);`}
					bind:this={myElement}
				>
					<RecommandCard></RecommandCard>
					<RecommandCard></RecommandCard>
					<RecommandCard></RecommandCard>
					<RecommandCard></RecommandCard>
					<RecommandCard></RecommandCard>
					<RecommandCard></RecommandCard>
					<RecommandCard></RecommandCard>
					<RecommandCard></RecommandCard>
					<RecommandCard></RecommandCard>
					<RecommandCard></RecommandCard>
					<RecommandCard></RecommandCard>
					<RecommandCard></RecommandCard>
					<RecommandCard></RecommandCard>
					<RecommandCard></RecommandCard>
					<RecommandCard></RecommandCard>
					<RecommandCard></RecommandCard>
				</div>
			</div>
		</div>
	</div>
{/await}
