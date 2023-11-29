<script lang="ts">
	import watermelon from '$lib/assets/watermelon.png';
	import { Counter } from '$lib';
	import type { PageData } from './$types';
	export let data: PageData;

	import Redradiobox from '$lib/components/PUA/redRadioBox.svelte';
	import OkButton from '$lib/components/PUA/OkButton.svelte';
	import DiscountArea from '$lib/components/PUA/discountArea.svelte';

	let product: {
		title: string;
		picture: string;
		price: number;
		content: string;
		choose: {
			id: number;
			need_choose: number;
			category: string;
			subcategories: string[];
		}[];
		discount: {
			id: number;
			label: string;
		}[];
	} = {
		title: '奶油龍蝦腳加水餃 放在一起烤 嘿~它為什麼要這樣叫阿',
		picture: watermelon,
		price: 70,
		content:
			'1. 茶葉原料產地為斯里蘭卡。 2. 茶葉原料產地為斯里蘭卡, 印度及印尼。 冰飲。總熱量為 372.1大卡。標準糖量為 83.7 克。含咖啡因 59.5 毫克。',
		choose: [
			{
				id: 0,
				need_choose: 1,
				category: '調整甜度',
				subcategories: ['正常糖', '半糖', '微糖', '少糖', '無糖']
			},
			{ id: 1, need_choose: 2, category: '加牛奶', subcategories: ['是', '否'] }
		],
		discount: [
			{ id: 1, label: '買二送一' },
			{ id: 2, label: '買二送一' },
			{ id: 3, label: '買二送一' },
			{ id: 4, label: '買二送一' },
			{ id: 5, label: '買二送一' },
			{ id: 6, label: '買二送一' }
		]
	};
</script>

<h1>
	shop: {data.shop}
</h1>
<h2>
	item: {data.item}
</h2>
<div class="flex justify-start">
	<div class="relative left-1/2 top-[25px] h-full w-4/5 -translate-x-1/2 transform">
		<div class="h-100 mb-8 flex w-full items-center text-4xl text-red-950">
			{product.title}
		</div>

		<div class="flex h-full w-full">
			<div class="relative h-full w-[500px]">
				<img
					src={product.picture}
					alt=""
					class="mt-100 flex h-[250px] w-[250px] rounded-[10px] object-cover"
				/>
				<div class="flex items-baseline py-[20px] font-bold text-red-950">
					<p class="text-[28px]">NT$</p>
					<p>&nbsp</p>
					<p>&nbsp</p>
					<p>&nbsp</p>
					<p class="text-[40px]">{product.price}</p>
				</div>
				<div class="w-[250px] overflow-hidden text-justify text-[15px] text-gray-600">
					{product.content}
				</div>
			</div>
			<div class="relative h-full w-full">
				<div class="relative mb-[10px] w-full">
					{#each product.choose as { need_choose, category, subcategories }}
						<div class=" mb-[15px]">
							<div
								class="flex h-[30px] w-full items-center border-b-[1px] border-solid border-PUA-stone"
							>
								<div class="font-bold text-PUA-stone">{category}</div>
								{#if need_choose === 1}
									<div
										class="ml-2 w-[40px] rounded-[20px] border-[2px] border-solid border-orange-700 px-[8px] py-[0px] text-center text-[9px] font-bold text-orange-700"
									>
										必填
									</div>
								{/if}
							</div>
							<div class="w-9/10 relative ml-[25px] mt-[10px] flex-col items-start">
								{#each subcategories as subcategory}
									<div
										class="flex w-full items-center space-x-2 border-b-[1px] border-solid border-amber-900"
									>
										<Redradiobox name={category} id={category + subcategory} />

										<div class="flex h-[30px] w-full items-center justify-end">
											<label for={category} class="font-bold text-primary text-red-950"
												>{subcategory}</label
											>
										</div>
									</div>
								{/each}
							</div>
						</div>
					{/each}
				</div>
				<DiscountArea
					discount={product.discount}
					toggleModal={() => {
						return null;
					}}
					addDiscountButton={() => {
						return null;
					}}
					addSign={false}
					type={true}
				></DiscountArea>

				<div class="ml-[40px] flex h-full items-center space-x-4">
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
<br />
<br />
<br />
