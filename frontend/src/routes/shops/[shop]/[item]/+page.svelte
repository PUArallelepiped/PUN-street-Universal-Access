<script lang="ts">
	import watermelon from '$lib/assets/watermelon.png';
	import { Counter } from '$lib';
	import { DiscountArea, Checkcontainer, OkButton, NeedChooseLabel } from '$lib/index';

	let product: {
		title: string;
		picture: string;
		price: number;
		content: string;
		choose: {
			id: number;
			need_choose: boolean;
			category: string;
			subcategories: string[];
		}[];
		discount: {
			id: string;
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
				need_choose: true,
				category: '調整甜度',
				subcategories: ['正常糖', '半糖', '微糖', '少糖', '無糖']
			},
			{ id: 1, need_choose: false, category: '加牛奶', subcategories: ['是', '否'] }
		],
		discount: [
			{ id: '1', label: '買二送一' },
			{ id: '2', label: '買二送一' },
			{ id: '3', label: '買二送一' },
			{ id: '4', label: '買二送一' },
			{ id: '5', label: '買二送一' },
			{ id: '6', label: '買二送一' }
		]
	};
</script>

<div class="flex justify-center">
	<div class="my-6 flex h-full w-4/5 flex-col gap-8">
		<div class=" flex w-full items-center text-4xl text-PUA-dark-red">
			{product.title}
		</div>

		<div class="flex gap-16">
			<div class="">
				<img src={product.picture} alt="" class="mt-100 flex h-60 w-60 rounded-lg object-cover" />
				<div class="flex items-baseline gap-3 py-5 font-bold text-PUA-dark-red">
					<p class="text-2xl">NT$</p>
					<p class="text-4xl">{product.price}</p>
				</div>
				<div class="w-[250px] text-justify text-base text-gray-600">
					{product.content}
				</div>
			</div>
			<div class=" flex w-full flex-col gap-4">
				<div class=" flex w-full flex-col gap-4">
					{#each product.choose as { need_choose, category, subcategories }}
						<div class="">
							<div class="flex items-center">
								<div class="font-bold text-PUA-stone">{category}</div>
								{#if need_choose}
									<NeedChooseLabel></NeedChooseLabel>
								{/if}
							</div>
							<div class="h-[1px] bg-PUA-dark-red"></div>
							<div class="flex flex-col">
								{#each subcategories as subcategory}
									<Checkcontainer {category} {subcategory}></Checkcontainer>
								{/each}
							</div>
						</div>
					{/each}
				</div>
				<div>
					<DiscountArea
						discount={product.discount}
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
