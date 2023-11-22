<script lang="ts">
	import error from '$lib/assets/error.png';
	import { Input } from '$lib/components/ui/input';

	let options = [
		{ id: 'option1', label: 'Vertical Option 1' },
		{ id: 'option1', label: 'Vertical Option 1' },
		{ id: 'option1', label: 'Vertical Option 1' }
	];

	let options_con = [{ val: 0, label: '調整甜度', group: 'op_1', nes_choose: true }];

	let opt = [options];

	// 改動態增加長度
	let buttons = [
		{ id: 1, label: '買二送一' },
		{ id: 2, label: '買二送一' },
		{ id: 3, label: '買二送一' },
		{ id: 4, label: '買二送一' },
		{ id: 5, label: '買二送一' }
	];

	let count = 0;

	function decrease() {
		count--;
	}

	function increase() {
		count++;
	}
</script>

<div class="flex justify-start">
	<div class="relative left-[170px] top-[25px] h-[100%] w-[826px]">
		<div class="relative h-[100px] w-[100%] items-center justify-center">
			<Input
				type="text"
				placeholder="Enter Product Name"
				class="w-[100%] rounded-[0px] border-b border-l-0 border-r-0 border-t-0 border-gray-400 text-[33px]"
			/>
			<div class="flex">
				<div class="flex items-center">
					<img src={error} alt="Error" class="h-[35px] w-[35px] object-cover object-cover" />

					<p class="text-[20px] font-bold text-red-500">Error message</p>
				</div>
			</div>
		</div>

		<div class="flex h-[100%] w-[100%]">
			<div class="relative h-[100%] w-[500px]">
				<div
					class="mt-100 flex h-[250px] w-[250px] items-center justify-center rounded-[10px] bg-gray-300"
				>
					<button
						class="h-[45px] rounded-[20px] bg-gray-400 px-[50px] py-[3px] text-white shadow-md"
						>Upload Image</button
					>
				</div>
				<div class="flex items-center">
					<img src={error} alt="Error" class="h-[35px] w-[35px] object-cover object-cover" />
					<p class="text-[20px] font-bold text-red-500">Error message</p>
				</div>

				<div class="flex items-center">
					<div class="py-[20px] text-[30px] font-bold text-red-950">NT$</div>
					<Input
						type="text"
						placeholder="Enter price"
						class="ml-[10px] w-[60%] max-w-xs rounded-[0px] border-b border-l-0 border-r-0 border-t-0 border-gray-400 text-[30px]"
					/>
				</div>
				<div class="flex items-center">
					<img src={error} alt="Error" class="h-[35px] w-[35px] object-cover object-cover" />
					<p class="text-[20px] font-bold text-red-500">Error message</p>
				</div>

				<div class="w-[250px] overflow-hidden text-justify text-[15px] text-gray-600">
					<Input
						type="text"
						placeholder="Enter text"
						class="max-w-xs rounded-[0px] border-b border-l-0 border-r-0 border-t-0 border-gray-400"
					/>
				</div>
				<div class="flex items-center">
					<img src={error} alt="Error" class="h-[35px] w-[35px] object-cover object-cover" />
					<p class="text-[20px] font-bold text-red-500">Error message</p>
				</div>
			</div>
			<div class="relative h-[100%] w-[100%]">
				<div class="relative mb-[10px] w-[100%]">
					{#each options_con as { val, group }}
						<div class=" mb-[15px]">
							<div
								class="flex h-[30px] w-[100%] items-center border-b-[1px] border-solid border-red-950"
							>
								<Input
									type="text"
									placeholder="Enter label"
									class="h-[20px] w-[30%] rounded-[0px] border-b border-l-0 border-r-0 border-t-0 border-gray-400"
								/>
								<div
									class="ml-[10px] rounded-[20px] border-[2px] border-solid border-orange-700 px-[10px] py-[0px] text-[13px] font-bold text-orange-700"
								>
									必填
								</div>
								<input
									type="radio"
									id="contactChoice1"
									name="choose_nesssary"
									value="Choice"
									class="m-[10px]"
								/>
								<label for="Choice">Yes</label>
								<input
									type="radio"
									id="contactChoice1"
									name="choose_nesssary"
									value="Choice"
									class="m-[10px]"
								/>
								<label for="Choice">No</label>
							</div>
							<div class="relative ml-[25px] w-[90%] flex-col items-start space-y-2">
								{#each opt[val] as { id }}
									<!--have some bug  -->
									<div
										class="m-1 flex w-[100%] items-center space-x-2 border-b-[1px] border-solid border-amber-900"
									>
										<input type="radio" name={group} class="form-radio text-primary" />
										<label for={id} class="ml-auto text-primary text-red-950"></label>
										<div class="flex h-[30px] w-[100%] items-center justify-end">
											<Input
												type="text"
												placeholder="item"
												class="h-[10px] w-[30%] rounded-[0px] border-b border-l-0 border-r-0 border-t-0 border-gray-400"
											/>
										</div>
									</div>
								{/each}
								<div class="flex h-[30px] w-[100%] items-center justify-end">
									<button
										class="m-[3px] h-[20x] w-[20px] rounded-[10px] bg-red-900 px-[0px] py-[0px] text-[13px] font-bold text-white"
										>-</button
									>
									<button
										class="m-[3px] h-[20x] w-[20px] rounded-[10px] bg-red-900 px-[0px] py-[0px] text-[13px] font-bold text-white"
										>+</button
									>
								</div>
							</div>
						</div>
					{/each}
				</div>
				<div class="flex h-[100%] h-[100%] items-center justify-center">
					<button class="h-[30px] w-[30px] rounded-[15px] bg-red-900 font-bold text-white">+</button
					>
				</div>
				<div class="relative h-[100%] w-[100%]">
					<div
						class="flex h-[30px] w-[100%] items-center border-b-[1px] border-solid border-red-950"
					>
						<div class=" font-bold text-amber-900">Choose Discount</div>
					</div>
					<div class="ml-[25px] mt-[10px] flex w-[90%] flex-wrap">
						{#each buttons as { label }}
							<button
								class="mb-4 mr-4 w-[90px] rounded rounded-[10px] border-[3px]
						border-lime-800 bg-transparent px-1 py-0 text-center
						font-bold text-lime-800 transition-all duration-300 hover:bg-lime-800 hover:text-white"
							>
								{label}
							</button>
						{/each}
						<div class="w-[90px]">
							<button
								class="h-[20px] w-[20px] rounded-[10px] bg-lime-800 text-[13px] font-bold text-white"
								>+</button
							>
						</div>
					</div>
				</div>

				<div class="flex h-[30px] w-[100%] items-center border-b-[1px] border-solid border-red-950">
					<div class=" font-bold text-amber-900">Set Stock</div>
				</div>
				<div class="flex items-center justify-center">
					<div class="relative ml-[40px] h-[100%]">
						<div class="mt-[10px] flex w-[200px] flex-1 items-center space-x-4">
							<button class="triangle-button-d" on:click={decrease}> </button>

							<div
								class="rounded-[20px] border-[3px] border-red-900 bg-transparent p-4
						px-[60px] py-0"
							>
								{count}
							</div>

							<button class="triangle-button-u" on:click={increase}> </button>
						</div>
						<button class="mt-[10px] rounded-[20px] bg-orange-700 p-2 px-[50px] text-white"
							>Add Product</button
						>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>

<style>
	.triangle-button-d {
		width: 0;
		height: 0;
		border-top: 10px solid transparent;
		border-bottom: 10px solid transparent;
		border-right: 15px solid rgb(87, 19, 11); /* 按钮颜色 */
		cursor: pointer;
	}
	.triangle-button-u {
		width: 0;
		height: 0;
		border-top: 10px solid transparent;
		border-bottom: 10px solid transparent;
		border-left: 15px solid rgb(87, 19, 11); /* 按钮颜色 */
		cursor: pointer;
	}
</style>
