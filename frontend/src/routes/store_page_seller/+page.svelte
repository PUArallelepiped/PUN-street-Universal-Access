<script lang="ts">
	import watermelon from '$lib/assets/watermelon.png';
	import Transhcan from '$lib/assets/transhcan.svg';
	import { HashtagLabel, CategoryLabel, DiscountCard, StoreProductCard } from '$lib';
	// export let data: PageData;

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
		}
	];

	let hashtag_input_id: { id: number; inputText: string }[] = [];
	let hashtag_input_nextId = hashtag_input_id.length + 1;
	let hashtag_text: { label: string }[] = [
		{ label: 'free delivery' },
		{ label: 'free delivery' },
		{ label: 'free delivery' },
		{ label: 'free delivery' }
	];
	function removeInput_addHashtag(id: number) {
		let result = hashtag_input_id.find((item) => item.id === id);
		if (result !== undefined) {
			hashtag_text = [...hashtag_text, { label: `${result.inputText}` }];
			hashtag_input_id = hashtag_input_id.filter((cat) => cat.id !== id);
		}
	}

	function handleClick() {
		hashtag_input_id = [...hashtag_input_id, { id: hashtag_input_nextId, inputText: '' }];
		hashtag_input_nextId++;
	}

	function handleKeyPress(event: KeyboardEvent, id: number) {
		if (event.key === 'Enter') {
			removeInput_addHashtag(id);
		}
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
						on:keydown={(e) => handleKeyPress(e, id)}
						type={'input'}
						id={`input${id}`}
						bind:text={inputText}
					></HashtagLabel>
				{/each}
				<HashtagLabel on:click={handleClick} type={'add'}></HashtagLabel>
			</div>
		</div>
	</div>

	<CategoryLabel text={'Product List'} img_nedd={true}></CategoryLabel>

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
	<CategoryLabel text={'Shipping Discount List'}></CategoryLabel>

	<div class="relative mx-5 space-y-4">
		<div class="flex items-center gap-4">
			<DiscountCard type="in" text={'1000'}></DiscountCard>
			<button>
				<img src={Transhcan} alt="" class="h-6 w-6" />
			</button>
		</div>
		<div class="flex items-center gap-4">
			<DiscountCard></DiscountCard>
		</div>
	</div>
</div>
