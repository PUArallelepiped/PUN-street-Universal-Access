<script lang="ts">
	import watermelon from '$lib/assets/watermelon.png';
	import Transhcan from '$lib/assets/transhcan.svg';
	import { ProductCard } from '$lib';
	import HashtagLabel from '$lib/components/PUA/hashtagLabel.svelte';
	import CategoryLabel from '$lib/components/PUA/categoryLabel.svelte';
	import DiscountCard from '$lib/components/PUA/discountCard.svelte';
	// export let data: PageData;
	let shopName = 'hi';

	let prodctListResponse: {
		name: string;
		description: string;
		price: number;
		href: string;
		picture: string;
		product_id: number;
		status: number;
		stock: number;
		store_id: number;
	}[] = [
		{
			status: 1,
			name: 'TEA EGG',
			description:
				'EGG of tea\n expensive\n also call putting, egg,egg,egg,egg, egg, egg,limit, limit, limit, limit, limit, limit, limit, limit, limit, ',
			price: 180,
			href: './' + shopName + '/product1',
			picture: 'https://i.imgur.com/3i3tyXJ.gif',
			product_id: 1,
			store_id: 1,

			stock: 100
		},
		{
			status: 1,
			stock: 100,
			store_id: 1,
			name: 'watermelon',
			description: 'a game',
			price: 0,
			href: './' + shopName + '/product2',
			picture: 'https://i.imgur.com/3i3tyXJ.gif',
			product_id: 2
		}
	];
	let discountcard = [{ label: '1' }, { label: '2' }, { label: '3' }];
</script>

<div class="h-48 w-full overflow-hidden">
	<img src={watermelon} alt="" class="w-full object-cover" />
</div>

<div class="mt-10 lg:px-40">
	<div class="mx-5 space-y-2">
		<div class="text-PUA-dark-red font-bold">100台灣台北市中正區八德路一段82巷9弄17號</div>
		<div class="text-PUA-stone text-5xl font-bold">銀記手稈刀切牛肉麵</div>
		<div class="flex w-full gap-3">
			<HashtagLabel type={'start'} text={'4.7'}></HashtagLabel>
			<HashtagLabel type={'text'} text={'free delivery'}></HashtagLabel>
			<HashtagLabel type={'add'}></HashtagLabel>
		</div>
	</div>

	<CategoryLabel text={'Product List'}></CategoryLabel>

	<div class="flex-row space-y-2 p-2">
		{#each prodctListResponse as product}
			<ProductCard
				name={product.name}
				description={product.description}
				price={product.price}
				imgUrl={product.picture}
			/>
		{/each}
		<div class="flex h-20 items-center justify-center">
			<HashtagLabel type={'add'}></HashtagLabel>
		</div>
	</div>
	<CategoryLabel text={'Shipping Discount List'}></CategoryLabel>

	<div class="relative space-y-4">
		{#each discountcard as { label }}
			<div class="flex items-center gap-4">
				<DiscountCard type="in"></DiscountCard>
				<img src={Transhcan} alt="" class="h-6 w-6" />
			</div>
		{/each}
		<div class="flex items-center gap-4">
			<DiscountCard></DiscountCard>
			<img src={Transhcan} alt="" class="h-6 w-6" />
		</div>
	</div>
</div>

<br />
<br />
<br />
