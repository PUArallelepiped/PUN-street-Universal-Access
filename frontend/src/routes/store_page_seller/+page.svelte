<script lang="ts">
	import watermelon from '$lib/assets/watermelon.png';
	import { HashtagLabel, CategoryLabel, DiscountCard, StoreProductCard } from '$lib';
	import ChangeDiscountPage from '$lib/components/PUA/changeDiscountPage.svelte';

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
			store_id: 1,
			name: 'watermelon',
			description: 'a game',
			price: 0,
			picture: 'https://i.imgur.com/3i3tyXJ.gif',
			product_id: 2
		},
		{
			status: 1,
			stock: 100,
			store_id: 1,
			name: 'watermelon',
			description: 'a game',
			price: 0,
			picture: 'https://i.imgur.com/3i3tyXJ.gif',
			product_id: 2
		}
	];

	let hashtag_input_id: { id: number; inputText: string }[] = [];
	let hashtag_text: { label: string }[] = [{ label: 'free delivery' }, { label: 'free delivery' }];

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

	function removeInput_addHashtag(id: number) {
		let result = hashtag_input_id.find((item) => item.id === id);
		if (result !== undefined) {
			hashtag_text = [...hashtag_text, { label: `${result.inputText}` }];
			hashtag_input_id = hashtag_input_id.filter((cat) => cat.id !== id);
		}
	}
	function addHashtagInput() {
		let hashtag_input_nextId = hashtag_input_id.length + 1;
		hashtag_input_id = [...hashtag_input_id, { id: hashtag_input_nextId, inputText: '' }];
	}
	function setHashtag(event: KeyboardEvent, id: number) {
		if (event.key === 'Enter') {
			removeInput_addHashtag(id);
		}
	}
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
		<div class="flex w-full justify-start gap-3">
			<div class="flex justify-start">
				<HashtagLabel type={'star'} text={'4.7'}></HashtagLabel>
			</div>
			<div class="flex w-full flex-wrap gap-2">
				{#each hashtag_text as { label }}
					<HashtagLabel type={'text'} text={label}></HashtagLabel>
				{/each}
				{#each hashtag_input_id as { id, inputText }}
					<HashtagLabel
						on:keydown={(e) => setHashtag(e, id)}
						type={'input'}
						id={`input${id}`}
						bind:text={inputText}
					></HashtagLabel>
				{/each}
				<HashtagLabel on:click={addHashtagInput} type={'add'}></HashtagLabel>
			</div>
		</div>
	</div>

	<CategoryLabel
		on:click={() => (showProductCard = toggleModel(showProductCard))}
		text={'Product List'}
		img_need={true}
	></CategoryLabel>

	{#if showProductCard}
		<div class="mx-5 flex-row space-y-2 p-2">
			{#each prodctListResponse as product}
				<StoreProductCard
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
	{/if}

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
</div>

<ChangeDiscountPage bind:changePageData bind:showModel></ChangeDiscountPage>
