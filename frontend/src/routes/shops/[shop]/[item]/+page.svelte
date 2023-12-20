<script lang="ts">
	import watermelon from '$lib/assets/watermelon.png';
	import { Counter } from '$lib';
	import { DiscountArea, Checkcontainer, OkButton, NeedChooseLabel } from '$lib/index';
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
</script>

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
					<Counter />

					<OkButton
						onclick={() => {
							return null;
						}}
						text="Add Cart"
					></OkButton>
				</div>
			</div>
		</div>
	</div>
</div>
