<script lang="ts">
	import close from '$lib/assets/close.svg';
	import ChangeDiscountInput from './changeDiscountInput.svelte';
	export let showModel = false;

	export let changePageData: {
		haved: boolean;
		name: string;
		maxquantity: string;
		description: string;
		kind: string;
		how: string;
		way: string;
	} = {
		haved: false,
		name: 'null',
		maxquantity: 'null',
		description: 'null',
		kind: 'null',
		how: 'null',
		way: 'null'
	};

	let currentData: {
		haved: boolean;
		name: string;
		maxquantity: string;
		description: string;
		kind: string;
		how: string;
		way: string;
	};

	currentData = { ...changePageData };

	$: {
		if (!changePageData.haved) {
			initPage();
		}
	}

	function initPage() {
		dataError = { maxquantity_err: false, name_err: false, description_err: false };
		currentData = { ...changePageData };
	}

	function toggleModel() {
		showModel = !showModel;
	}
	let dataError: { maxquantity_err: boolean; name_err: boolean; description_err: boolean } = {
		maxquantity_err: false,
		name_err: false,
		description_err: false
	};
	function saveCurrentData() {
		if (
			currentData.maxquantity !== '' &&
			currentData.name !== '' &&
			currentData.description !== ''
		) {
			changePageData = { ...currentData };
			changePageData.haved = true;
			dataError = { maxquantity_err: false, name_err: false, description_err: false };
		}
		if (currentData.maxquantity === '') {
			dataError.maxquantity_err = true;
		} else if (currentData.maxquantity !== '') {
			dataError.maxquantity_err = false;
		}
		if (currentData.name === '') {
			dataError.name_err = true;
		} else if (currentData.name !== '') {
			dataError.name_err = false;
		}
		if (currentData.description === '') {
			dataError.description_err = true;
		} else if (currentData.description !== '') {
			dataError.description_err = false;
		}
	}
	function closePage() {
		initPage();
		toggleModel();
	}
	function deletePage() {
		currentData.name = '';
		currentData.maxquantity = '';
		currentData.description = '';
	}
</script>

{#if showModel}
	<div class="absolute z-20">
		<div
			class=" fixed left-0 top-0 flex h-full w-full items-center justify-center bg-black bg-opacity-50 backdrop-blur"
		>
			<div class="rounded bg-white p-5">
				<div class="flex">
					<div class="w-1/2 text-left text-xl font-bold text-PUA-stone">
						<h2>Add A Discount</h2>
					</div>
					<div class="flex w-1/2 justify-end">
						<button on:click={closePage}>
							<img src={close} alt="" class="h-[25px] object-cover" />
						</button>
					</div>
				</div>
				<div class="relative ml-16 mr-16">
					<div class="flex w-full items-center justify-center pb-2 pt-2">
						<div class="flex w-64 rounded-[10px] border-[3px] border-red-900 p-3 text-red-900">
							<div class="w-2/5 text-center font-bold">{currentData.kind}</div>
							<div class="ml-1 mr-1 border-r-[2px] border-red-900"></div>
							<div class="flex w-3/5 flex-wrap items-center justify-center text-center font-bold">
								<p>{currentData.how}</p>
								<p class="ml-1 mr-1 text-xl">{currentData.maxquantity}</p>
								<p>{currentData.way}</p>
							</div>
						</div>
					</div>

					<ChangeDiscountInput
						title={'Max Quantity'}
						bind:value={currentData.maxquantity}
						bind:error={dataError.maxquantity_err}
						text={' Enter Max Quantity'}
					></ChangeDiscountInput>

					<ChangeDiscountInput
						title={'Discount Name'}
						bind:value={currentData.name}
						bind:error={dataError.name_err}
						text={' Enter Discount Name'}
					></ChangeDiscountInput>
					<ChangeDiscountInput
						title={'Description'}
						bind:value={currentData.description}
						bind:error={dataError.description_err}
						text={' Enter Discount Description'}
					></ChangeDiscountInput>
				</div>

				<div class="mt-5 flex items-center justify-between gap-5 text-center">
					<button
						on:click={deletePage}
						class="w-4/5 rounded-[20px] bg-gray-200 font-bold text-red-900"
					>
						Delete
					</button>

					<button
						on:click={saveCurrentData}
						class="w-4/5 rounded-[20px] bg-orange-700 font-bold text-white"
					>
						Save
					</button>
				</div>
			</div>
		</div>
	</div>
{/if}
