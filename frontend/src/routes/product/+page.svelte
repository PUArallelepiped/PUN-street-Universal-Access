<script lang="ts">
	import drink from '$lib/assets/watermelon.png';

	let options = [
		{ id: 'option1', label: '正常糖' },
		{ id: 'option1', label: '半糖' },
		{ id: 'option1', label: '微糖' },
		{ id: 'option1', label: '少糖' },
		{ id: 'option1', label: '無糖' }
	];
	let options2 = [
		{ id: 'option2', label: '是' },
		{ id: 'option2', label: '否' }
	];

	let options_con = [
		{ val: 0, label: '調整甜度', group: 'op_1', nes_choose: true },
		{ val: 1, label: '加牛奶', group: 'op_2', nes_choose: false }
	];

	let opt = [options, options2];

	// 改動態增加長度
	let buttons = [
		{ id: 1, label: '買二送一' },
		{ id: 2, label: '買二送一' },
		{ id: 3, label: '買二送一' },
		{ id: 4, label: '買二送一' },
		{ id: 5, label: '買二送一' },
		{ id: 6, label: '買二送一' }
	];

	let product_title = '奶油龍蝦腳加水餃 放在一起烤 嘿~它為什麼要這樣叫阿';

	let product_contect =
		'1. 茶葉原料產地為斯里蘭卡。 2. 茶葉原料產地為斯里蘭卡, \
	印度及印尼。 冰飲。總熱量為 372.1大卡。標準糖量為 83.7 克。含咖啡因 59.5 毫克。';

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
		<div class="flex h-[100px] w-[100%] items-center justify-center text-[33px] text-red-950">
			{product_title}
		</div>

		<div class="flex h-[100%] w-[100%]">
			<div class="relative h-[100%] w-[500px]">
				<img
					src={drink}
					alt=""
					class="mt-100 flex h-[250px] w-[250px] rounded-[10px] object-cover object-cover"
				/>

				<div class="py-[20px] text-[30px] font-bold text-red-950">NT $70</div>
				<div class="w-[250px] overflow-hidden text-justify text-[15px] text-gray-600">
					{product_contect}
				</div>
			</div>
			<div class="relative h-[100%] w-[100%]">
				<div class="relative mb-[10px] w-[100%]">
					{#each options_con as { val, label, group }}
						<div class=" mb-[15px]">
							<div
								class="flex h-[30px] w-[100%] items-center border-b-[1px] border-solid border-red-950"
							>
								<div class="font-bold text-amber-900">{label}</div>
								<div
									class="ml-[10px] rounded-[20px] border-[2px] border-solid border-orange-700 px-[8px] py-[0px] text-[9px] font-bold text-orange-700"
								>
									必填
								</div>
							</div>
							<div class="relative ml-[25px] mt-[10px] w-[90%] flex-col items-start">
								{#each opt[val] as { id, label }}
									<div
										class="flex w-[100%] items-center space-x-2 border-b-[1px] border-solid border-amber-900"
									>
										<input type="radio" name={group} class="h-[20px] w-[20px]" />
										<div class="flex h-[30px] w-[100%] items-center justify-end">
											<label for={id} class="font-bold text-primary text-red-950">{label}</label>
										</div>
									</div>
								{/each}
							</div>
						</div>
					{/each}
				</div>
				<div class="relative h-[100%] w-[100%]">
					<div
						class="flex h-[30px] w-[100%] items-center border-b-[1px] border-solid border-red-950"
					>
						<div class=" font-bold text-amber-900">Choose Discount</div>
					</div>
					<div class="ml-[25px] mt-[10px] flex w-[90%] flex-wrap">
						{#each buttons as { label }}
							<div class="ml-[5px] flex justify-center">
								<button
									class="m-[10px] w-[90px] rounded rounded-[10px] border-[3px]
						border-lime-800 bg-transparent px-1 py-0 text-center
						font-bold text-lime-800 transition-all duration-300 hover:bg-lime-800 hover:text-white"
								>
									{label}
								</button>
							</div>
						{/each}
					</div>
				</div>

				<div class="ml-[40px] flex h-[100%] w-[85%]">
					<div class=" flex flex-1 items-center space-x-4">
						<button class="triangle-button-d" on:click={decrease}> </button>

						<div
							class="rounded-[20px] border-[3px] border-red-900 bg-transparent p-4
						px-[60px] py-0"
						>
							{count}
						</div>

						<button class="triangle-button-u" on:click={increase}> </button>
					</div>
					<button class="rounded-[20px] bg-orange-700 p-2 px-[70px] text-white">Add cart</button>
				</div>
			</div>
		</div>
	</div>
</div>
<br />
<br />
<br />

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
	input[type='radio'] {
		accent-color: rgb(120, 40, 40);
	}
</style>
