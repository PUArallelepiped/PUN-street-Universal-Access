<script lang="ts">
	import watermelon from '$lib/assets/watermelon.png';
	import { CategoryLabel, DiscountCard } from '$lib';
	import ChangeDiscountPage from '$lib/components/PUA/changeDiscountPage.svelte';
	import StoreProducrCardArea from '$lib/components/PUA/store_page_seller/storeProducrCardArea.svelte';
	import HashtagLabelArea from '$lib/components/PUA/store_page_seller/hashtagLabelArea.svelte';
	let prodctListResponse: {
		name: string;
		description: string;
		price: number;
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
				'EGG of tea\n expensive\n also call putting, egg,egg,egg,egg, egg, egg,limit, limit, limit, limit, ',
			price: 180,
			picture: 'https://i.imgur.com/3i3tyXJ.gif',
			product_id: 1,
			store_id: 1,
			stock: 100
		},
		{
			status: 1,
			stock: 100,
			store_id: 2,
			name: 'watermelon',
			description: 'a game',
			price: 0,
			picture: 'https://i.imgur.com/3i3tyXJ.gif',
			product_id: 2
		},
		{
			status: 1,
			stock: 100,
			store_id: 3,
			name: 'watermelon',
			description: 'a game',
			price: 0,
			picture: 'https://i.imgur.com/3i3tyXJ.gif',
			product_id: 3
		}
	];

	let hashtag_text: { id: number; label: string }[] = [
		{ id: 0, label: 'free delivery' },
		{ id: 1, label: 'free delivery' }
	];

	let changePageData: {
		haved: boolean;
		name: string;
		maxquantity: string;
		description: string;
		kind: string;
		how: string;
		way: string;
	} = {
		haved: false,
		name: '',
		maxquantity: '',
		description: '',
		kind: 'Shipping Discount',
		how: 'NT$',
		way: 'free shipping'
	};

	let showProductCard = true;
	let showModel = false;

	function toggleModel(model: boolean) {
		return (model = !model);
	}
	function deleteDiscountCard() {
		changePageData = {
			...changePageData,
			haved: false,
			name: '',
			maxquantity: '',
			description: ''
		};
	}
</script>

<div class="h-48 w-full overflow-hidden">
	<img src={watermelon} alt="" class="w-full object-cover" />
</div>

<div class="mt-10 lg:px-40">
	<div class="mx-5 space-y-2">
		<div class="text-5xl font-bold text-PUA-stone">銀記手稈刀切牛肉麵</div>
		<div class="font-bold text-red-950">100台灣台北市中正區八德路一段82巷9弄17號</div>
		<div class="flex w-full justify-start gap-6">
			<HashtagLabelArea bind:hashtag_text></HashtagLabelArea>
		</div>
	</div>
	<div class="h-full min-h-screen">
		<CategoryLabel
			on:click={() => (showProductCard = toggleModel(showProductCard))}
			text={'Product List'}
			img_need={true}
			bind:dropdown={showProductCard}
		></CategoryLabel>

		<div class={` ${showProductCard ? 'visible h-full' : 'invisible h-0'}`}>
			<StoreProducrCardArea bind:prodctListResponse></StoreProducrCardArea>
		</div>

		<CategoryLabel text={'Shipping Discount List'}></CategoryLabel>
		<div class="relative mx-5 space-y-4">
			<div class="flex items-center gap-4">
				<DiscountCard
					discountCardData={changePageData}
					on:click={() => (showModel = toggleModel(showModel))}
					{deleteDiscountCard}
				></DiscountCard>
			</div>
		</div>
		<div class="relative mx-5 h-6 space-y-4"></div>
	</div>
</div>
<ChangeDiscountPage bind:changePageData bind:showModel></ChangeDiscountPage>
