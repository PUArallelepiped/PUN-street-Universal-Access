<script lang="ts">
	import drink from '$lib/assets/watermelon.png';
	import { page } from '$app/stores';

	let options = [
		{ id: 'option1', label: 'Vertical Option 1' },
		{ id: 'option1', label: 'Vertical Option 2' },
		{ id: 'option1', label: 'Vertical Option 3' },
		{ id: 'option1', label: 'Vertical Option 4' }
	];
	let options2 = [
		{ id: 'option2', label: 'Vertical happy 1' },
		{ id: 'option2', label: 'Vertical happy 4' }
	];

	let options_con = [
		{ val: 0, label: '調整甜度',group:'op_1',nes_choose:true },
		{ val: 1, label: '加牛奶' ,group:'op_2',nes_choose:false},
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

	let product_title = "奶油龍蝦腳加水餃 放在一起烤 嘿~它為什麼要這樣叫阿";

	let product_contect = "1. 茶葉原料產地為斯里蘭卡。 2. 茶葉原料產地為斯里蘭卡, \
	印度及印尼。 冰飲。總熱量為 372.1大卡。標準糖量為 83.7 克。含咖啡因 59.5 毫克。"

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
				<img src={drink} alt="" class="mt-100 flex h-[250px] w-[250px] rounded-[10px] object-cover object-cover" />
				
				<div class="text-[30px] font-bold text-red-950 py-[20px]">NT $70</div>
				<div class="w-[250px] overflow-hidden text-justify text-gray-600 text-[15px]">
					{product_contect}
				
				</div>
			</div>
			<div class="relative h-[100%] w-[100%] ">
				<div class="relative w-[100%] mb-[10px] ">
					{#each options_con as { val, label ,group}}
					<div class=" mb-[15px]">
						<div class="flex h-[30px] w-[100%] items-center border-b-[1px] border-solid border-red-950 ">
							<div class=" font-bold text-amber-900">{label}</div>

						</div>
						<div class="relative ml-[25px] w-[90%] flex-col items-start space-y-2 ">
							{#each opt[val] as { id, label }}
								<div class="flex w-[100%] m-1 items-center space-x-2 border-b-[1px] border-solid border-amber-900">
									<input type="radio" name={group} class="form-radio text-primary" />
									<label for={id} class="text-primary ml-auto text-red-950">{label}</label>
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
						{#each buttons as { id, label }, index (id)}
							<button
								class="mb-4 mr-4 w-[90px] rounded rounded-[10px] border-[3px]
						border-lime-800 bg-transparent px-1 py-0 text-center
						font-bold text-lime-800 transition-all duration-300 hover:bg-lime-800 hover:text-white"
							>
								{label}
							</button>
						{/each}
					</div>
				</div>

				<div class="flex ml-[40px] h-[100%] w-[85%]">
					<div class=" flex-1 flex items-center space-x-4">
						<button class="triangle-button-d" on:click={decrease}> </button>

						<div class="rounded-[20px] p-4 px-[60px] py-0 border-[3px]
						border-red-900 bg-transparent">
							{count}
						</div>

						<button class="triangle-button-u" on:click={increase}> </button>
					</div>
					<button class="bg-orange-700 text-white p-2 rounded-[20px] px-[70px] ">Add cart</button>
				</div>
			</div>
		</div>
	</div>
</div>

<h1>here is shop list</h1>

<p>Visit <a href="{$page.route.id}/impasta" class="text-blue-400"> Shop </a></p>

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
