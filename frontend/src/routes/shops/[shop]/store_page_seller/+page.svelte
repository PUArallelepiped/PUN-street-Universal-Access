<script lang="ts">
	import watermelon from '$lib/assets/watermelon.png';
	import { CategoryLabel, DiscountCard } from '$lib';
	import ChangeDiscountPage from '$lib/components/PUA/changeDiscountPage.svelte';
	import StoreProductCardArea from '$lib/components/PUA/store_page_seller/storeProductCardArea.svelte';
	import TagLabelArea from '$lib/components/PUA/store_page_seller/tagLabelArea.svelte';
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	import { backendPath } from '$lib/components/PUA/env';
	export let data: PageData;
	let shop_id = data.shop;

	interface productsType {
		product_id: number;
		name: string;
		store_id: number;
		description: string;
		picture: string;
		price: number;
		stock: number;
		status: number;
		event_discount_array: {
			discount_max_quantity: number;
			product_id: number;
			discount_name: string;
			discount_description: string;
			discount_id: number;
			status: number;
		}[];
	}
	interface categoriesType {
		category_id: number;
		category_name: string;
	}
	interface shippingDiscountType {
		discount_name: string;
		discount_description: string;
		discount_max_quantity: number;
		discount_id: number;
		status: number;
	}

	interface shippingDiscountRespType {
		discount_name: string;
		discount_description: string;
		discount_max_price: number;
		discount_id: number;
		status: number;
	}

	interface storeDataType {
		store_id: number;
		shipping_fee: number;
		address: string;
		rate_count: number;
		rate: number;
		category_array: {
			category_name: string;
			category_id: number;
		}[];
		name: string;
		description: string;
		picture: string;
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
			status: 0,
			event_discount_array: []
		}
	];
	let tagList: categoriesType[] = [
		{
			category_name: 'drink',
			category_id: 1
		}
	];

	let shippingList: shippingDiscountType = {
		discount_name: '',
		discount_description: '',
		discount_max_quantity: 0,
		discount_id: 0,
		status: 1
	};

	let shippingListResp: shippingDiscountRespType = {
		discount_name: '',
		discount_description: '',
		discount_max_price: 0,
		discount_id: 0,
		status: 1
	};

	let shopDataList: storeDataType = {
		store_id: 1,
		shipping_fee: 100,
		address: 'pun street',
		rate_count: 10,
		rate: 5,
		category_array: [
			{
				category_name: 'drink',
				category_id: 1
			},
			{
				category_name: 'drink',
				category_id: 2
			}
		],
		name: "i'm pasta",
		description: 'good pasta',
		picture: 'https://i.imgur.com/1.jpg',
		status: 1
	};

	let showProductCard = true;
	let showModel = false;
	let dis_haved: boolean;
	async function putDiscount(discount_id: number) {
		await fetch(backendPath + '/discount/' + discount_id + '/delete-discount', {
			method: 'PUT'
		});
	}

	async function postDiscount() {
		shippingListResp = {
			discount_name: shippingList.discount_name,
			discount_id: shippingList.discount_id,
			discount_description: shippingList.discount_description,
			discount_max_price: shippingList.discount_max_quantity,
			status: shippingListResp.status
		};
		const respJson = await fetch(backendPath + '/store/' + '1' + '/shipping-discount', {
			method: 'POST',
			body: JSON.stringify(shippingListResp)
		}).then((resp) => {
			return resp.json();
		});

		shippingList = {
			discount_name: respJson.discount_name,
			discount_id: respJson.discount_id,
			discount_description: respJson.discount_description,
			discount_max_quantity: respJson.discount_max_price,
			status: respJson.status
		};
		dis_haved = shippingList ? true : false;
	}

	function addDiscount() {
		postDiscount();
		return null;
	}
	function deleteDiscountCard() {
		if (dis_haved) {
			putDiscount(shippingListResp.discount_id);
			shippingList = {
				...shippingList,
				discount_name: '',
				discount_description: '',
				discount_max_quantity: 0,
				discount_id: 0,
				status: 0
			};
			dis_haved = !dis_haved;
		}
		return null;
	}

	let myElement: HTMLDivElement | null = null;
	let k = 0;

	const toggleHeight = () => {
		if (myElement) {
			if (showProductCard) {
				myElement.style.height = `${k}vh`;
			} else {
				myElement.style.height = `100vh`;
			}
		}
	};

	async function getProductResp() {
		const init_product = await fetch(backendPath + `/store/` + shop_id + `/products`);
		if (init_product.status == 200) {
			productsList = await init_product.json();
			k = 20 * productsList.length + 50;
			showProductCard;
			toggleHeight();
		}
		return;
	}

	async function getCategoryResp() {
		const tag_category = await fetch(backendPath + `/categories`);
		if (tag_category.status == 200) {
			tagList = await tag_category.json();
		}
		return;
	}

	async function getDiscountResp() {
		const init_shipping_discount = await fetch(
			backendPath + `/store/` + shop_id + `/shipping-discount`
		);
		if (init_shipping_discount.status == 200) {
			shippingListResp = await init_shipping_discount.json();
			if (shippingListResp) {
				shippingList = {
					discount_name: shippingListResp.discount_name,
					discount_id: shippingListResp.discount_id,
					discount_description: shippingListResp.discount_description,
					discount_max_quantity: shippingListResp.discount_max_price,
					status: shippingListResp.status
				};
			}
			dis_haved = shippingListResp ? true : false;
		}
		return;
	}

	async function getStoreResp() {
		const init_store_data = await fetch(backendPath + `/store/` + shop_id);
		if (init_store_data.status == 200) {
			shopDataList = await init_store_data.json();
		}
		return;
	}

	async function getData() {
		await getProductResp();
		await getCategoryResp();
		await getDiscountResp();
		await getStoreResp();
	}
	$: {
		k = 20 * productsList.length + 50;
		showProductCard;
		toggleHeight();
	}
	onMount(async () => {
		await getData();
	});
</script>

{#await getData() then}
	<div class="h-48 w-full overflow-hidden">
		<img src={watermelon} alt="" class="w-full object-cover" />
	</div>

	<div class="mt-10 lg:px-40">
		<div class="mx-5 space-y-2">
			<div class="text-5xl font-bold text-PUA-stone">{shopDataList.name}</div>
			<div class="font-bold text-red-950">{shopDataList.address}</div>
			<div class="flex w-full justify-start gap-6">
				<TagLabelArea
					bind:tagText={shopDataList.category_array}
					bind:tagText_all={tagList}
					star_score={shopDataList.rate.toString()}
					{shop_id}
				></TagLabelArea>
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
					<StoreProductCardArea bind:productListResponse={productsList} {shop_id}
					></StoreProductCardArea>
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
	<ChangeDiscountPage
		bind:changePageData={shippingList}
		bind:showModel
		bind:dis_haved
		add_Discount={addDiscount}
		delete_Discount={deleteDiscountCard}
	></ChangeDiscountPage>
{/await}
