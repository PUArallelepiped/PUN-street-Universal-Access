<script lang="ts">
	import watermelon from '$lib/assets/watermelon.png';
	import { CategoryLabel, DiscountCard } from '$lib';
	import ChangeDiscountPage from '$lib/components/PUA/changeDiscountPage.svelte';
	import StoreProductCardArea from '$lib/components/PUA/store_page_seller/storeProductCardArea.svelte';
	import TagLabelArea from '$lib/components/PUA/store_page_seller/tagLabelArea.svelte';
	import { onMount } from 'svelte';

	interface productsType {
		product_id: number;
		name: string;
		store_id: number;
		description: string;
		picture: string;
		price: number;
		stock: number;
		status: number;
	}
	interface categoriesType {
		category_id: number;
		category_name: string;
	}
	interface shippingDiscountType {
		discount_name: string;
		discount_description: string;
		discount_max_price: number;
		discount_id: number;
		status: number;
	}
	let productsList: productsType[] = [
		{
			product_id: 0,
			name: '',
			store_id: 0,
			description: '',
			picture: '',
			price: 0,
			stock: 0,
			status: 0
		}
	];
	let tagList: categoriesType[] = [
		{
			category_name: 'drink',
			category_id: 1
		}
	];

	let shippingList: shippingDiscountType = {
		discount_name: 'free shipping',
		discount_description: 'free shipping when total price over 1000',
		discount_max_price: 1000,
		discount_id: 1,
		status: 1
	};

	let showProductCard = true;
	let showModel = false;
	let dis_haved: boolean;

	function deleteDiscountCard() {
		shippingList = {
			...shippingList,
			discount_name: '',
			discount_description: '',
			discount_max_price: 0,
			discount_id: 0,
			status: 0
		};
		dis_haved = !dis_haved;
	}

	let myElement: HTMLDivElement | null = null;
	let k = 0;

	$: {
		k = 20 * productsList.length + 50;
		showProductCard;
		toggleHeight();
	}

	const toggleHeight = () => {
		if (myElement) {
			if (showProductCard) {
				myElement.style.height = `${k}vh`;
			} else {
				myElement.style.height = `100vh`;
			}
		}
	};

	onMount(() => {
		toggleHeight();
	});
	onMount(async () => {
		const init_product = await fetch(`/products.json`);
		productsList = await init_product.json();
		const tag_category = await fetch(`/categories.json`);
		tagList = await tag_category.json();
		const init_shipping_discount = await fetch(`/shipping-discount.json`);
		shippingList = await init_shipping_discount.json();
		dis_haved = shippingList ? true : false;
	});
</script>

<div class="h-48 w-full overflow-hidden">
	<img src={watermelon} alt="" class="w-full object-cover" />
</div>

<div class="mt-10 lg:px-40">
	<div class="mx-5 space-y-2">
		<div class="text-5xl font-bold text-PUA-stone">銀記手稈刀切牛肉麵</div>
		<div class="font-bold text-red-950">100台灣台北市中正區八德路一段82巷9弄17號</div>
		<div class="flex w-full justify-start gap-6">
			<TagLabelArea bind:tagText={tagList}></TagLabelArea>
		</div>
	</div>

	<div bind:this={myElement} class="min-h-screen">
		<CategoryLabel
			on:click={() => (showProductCard = !showProductCard)}
			text={'Product List'}
			img_need={true}
			bind:dropdown={showProductCard}
		></CategoryLabel>

		<div
			class={` ${
				showProductCard ? 'max-h-full' : 'max-h-0'
			}   overflow-hidden transition-all duration-[1300ms] ease-in-out`}
		>
			<div
				class={` ${
					showProductCard ? ' translate-y-0 ' : 'translate-y-[-100%]'
				}   transition-all duration-[1500ms] ease-in-out `}
			>
				<StoreProductCardArea bind:productListResponse={productsList}></StoreProductCardArea>
			</div>
		</div>
		<CategoryLabel text={'Shipping Discount List'}></CategoryLabel>
		<div class="relative mx-5 space-y-4">
			<div class="flex items-center gap-4">
				<DiscountCard
					bind:discountCardData={shippingList}
					bind:dis_haved
					on:click={() => (showModel = !showModel)}
					{deleteDiscountCard}
				></DiscountCard>
			</div>
		</div>
	</div>
</div>
<ChangeDiscountPage bind:changePageData={shippingList} bind:showModel bind:dis_haved
></ChangeDiscountPage>
