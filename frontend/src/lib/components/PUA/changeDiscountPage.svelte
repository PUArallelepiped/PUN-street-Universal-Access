<script lang="ts">
	import close from '$lib/assets/close.svg';
	import ChangeDiscountInput from './changeDiscountInput.svelte';
	export let showModel = false;
	export let add_Discount: () => null;
	export let delete_Discount: () => null;
	export let changePageData: {
		discount_name: string;
		discount_description: string;
		discount_max_quantity: number;
		discount_id: number;
		status: number;
	};

	export let discountData = {
		kind: 'Shipping Discount',
		how: 'NT$',
		way: 'free shipping'
	};

	export let dis_haved: boolean;

	let currentData: {
		discount_name: string;
		discount_description: string;
		discount_max_quantity: number;
		discount_id: number;
		status: number;
	};
	currentData = { ...changePageData };

	$: {
		changePageData;
		initPage();
	}
	function initPage() {
		currentData = { ...changePageData };
	}
	function toggleModel() {
		showModel = !showModel;
	}

	function closePage() {
		initPage();
		toggleModel();
	}

	async function handleSubmit() {
		changePageData = { ...currentData };
		dis_haved = true;
		toggleModel();
		add_Discount();
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
							<img src={close} alt="" class="h-[25px] object-cover hover:scale-125" />
						</button>
					</div>
				</div>
				<form on:submit|preventDefault={handleSubmit}>
					<div class="relative mx-16 my-8">
						<div class="flex w-full items-center justify-center pb-2 pt-2">
							<div class="flex w-80 rounded-xl border-4 border-PUA-stone p-2 text-PUA-stone">
								<div class="  px-2 py-2 text-center text-base font-semibold">
									{discountData.kind}
								</div>
								<div class="mt-2 h-12 w-1 bg-PUA-stone"></div>
								<div class="w-full px-2 font-semibold">
									<div class="flex items-baseline justify-center">
										<span class="ml-1 mr-1 text-base">{discountData.how}</span><span
											class=" text-center text-2xl">{currentData.discount_max_quantity}</span
										>
									</div>
									<div class="text-center text-lg">{discountData.way}</div>
								</div>
							</div>
						</div>
						<ChangeDiscountInput
							title={'Max Quantity'}
							bind:value={currentData.discount_max_quantity}
							type={'number'}
							text={' Enter Max Quantity'}
							name={'MaxQuantity'}
						></ChangeDiscountInput>
						<ChangeDiscountInput
							title={'Discount Name'}
							bind:value={currentData.discount_name}
							text={' Enter Discount Name'}
							name={'DiscountName'}
						></ChangeDiscountInput>
						<ChangeDiscountInput
							title={'Description'}
							bind:value={currentData.discount_description}
							text={' Enter Discount Description'}
							name={'Description'}
						></ChangeDiscountInput>
					</div>

					<div class="mt-5 flex items-center justify-between gap-5 text-center">
						<button
							type="button"
							on:click={() => {
								delete_Discount(), toggleModel();
							}}
							class="w-4/5 rounded-[20px] bg-gray-200 font-bold text-red-900"
						>
							Delete
						</button>

						<button type="submit" class="w-4/5 rounded-[20px] bg-orange-700 font-bold text-white">
							Save
						</button>
					</div>
				</form>
			</div>
		</div>
	</div>
{/if}
