<script lang="ts">
	import watermelon from '$lib/assets/watermelon.png';
	import { CategoryLabel, DiscountCard } from '$lib';
	import ChangeDiscountPage from '$lib/components/PUA/changeDiscountPage.svelte';
	import StoreProducrCardArea from '$lib/components/PUA/store_page_seller/storeProducrCardArea.svelte';
	import TagLabelArea from '$lib/components/PUA/store_page_seller/tagLabelArea.svelte';
	import { onMount } from 'svelte';
	import data from './data.json';

	// 使用讀取到的數據
	let productListResponse = data.productListResponse;
	let tagText = data.tagText;
	let changePageData = data.changePageData;

	let showProductCard = true;
	let showModel = false;

	function deleteDiscountCard() {
		changePageData = {
			...changePageData,
			haved: false,
			name: '',
			maxquantity: '',
			description: ''
		};
	}

	let myElement: HTMLDivElement | null = null;
	let k = 0;

	$: {
		k = 20 * productListResponse.length + 50;
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
</script>

<div class="h-48 w-full overflow-hidden">
	<img src={watermelon} alt="" class="w-full object-cover" />
</div>

<div class="mt-10 lg:px-40">
	<div class="mx-5 space-y-2">
		<div class="text-PUA-stone text-5xl font-bold">銀記手稈刀切牛肉麵</div>
		<div class="font-bold text-red-950">100台灣台北市中正區八德路一段82巷9弄17號</div>
		<div class="flex w-full justify-start gap-6">
			<TagLabelArea bind:tagText></TagLabelArea>
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
				<StoreProducrCardArea bind:productListResponse></StoreProducrCardArea>
			</div>
		</div>
		<CategoryLabel text={'Shipping Discount List'}></CategoryLabel>
		<div class="relative mx-5 space-y-4">
			<div class="flex items-center gap-4">
				<DiscountCard
					discountCardData={changePageData}
					on:click={() => (showModel = !showModel)}
					{deleteDiscountCard}
				></DiscountCard>
			</div>
		</div>
	</div>
</div>
<ChangeDiscountPage bind:changePageData bind:showModel></ChangeDiscountPage>
