<script lang="ts">
	import drink from '$lib/assets/watermelon.png';
	import error from '$lib/assets/error.png';
	import { page } from '$app/stores';

	//import * as Menubar from "$lib/components/ui/menubar";
	import * as Card from '$lib/components/ui/card';
	import { Select } from '$lib/components/ui/select';
	import { Input } from '$lib/components/ui/input';
	import { Checkbox } from '$lib/components/ui/checkbox';
	import { Label } from '$lib/components/ui/label';
	//import { Slider } from '$lib/components/ui/slider';
	//import { Button } from "$lib/components/ui/button";
	//let checked = false;

	//引入所需的UI元件
	//$app/stores 這個不知道

	let options = [
		{ id: 'option1', label: 'Vertical Option 1' },
		// { id: 'option1', label: 'Vertical Option 2' },
		// { id: 'option1', label: 'Vertical Option 3' },
		// { id: 'option1', label: 'Vertical Option 4' }
	];
	// let options2 = [
	// 	{ id: 'option2', label: 'Vertical happy 1' },
	// 	{ id: 'option2', label: 'Vertical happy 4' }
	// ];
	// let options3 = [
	// 	{ id: 'option2', label: 'Vertical happy 1' },
	// 	{ id: 'option2', label: 'Vertical happy 4' }
	// ];
	let options_con = [
		{ val: 0, label: '調整甜度',group:'op_1',nes_choose:true },
		// { val: 1, label: '加牛奶' ,group:'op_2',nes_choose:false},
		// { val: 2, label: '加牛奶' ,group:'op_3',nes_choose:false}
	];

	let opt = [options]; 

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
	印度及印尼。 冰飲。總熱量為 372.1大卡。標準糖量為 83.7 克。含咖啡因 59.5 毫克。 1. 茶葉原料產地為斯里蘭卡。"

	let count = 0;

	function decrease() {
		count--;
	}

	function increase() {
		count++;
	}

</script>

<!-- <div class="h-[50px] w-[40%] overflow-hidden">  -->
<!-- overflow-hidden => 隱藏超過文字元素的內容 -->
<!-- <div class="float-left text-2xl">PUN Street Access</div> -->
<!-- </div> -->

<!--flex 彈性布局 justify-start左上對齊  -->
<div class="flex justify-start">
	<!--  bg-red-500  -->
	<div class="relative left-[170px] top-[25px] h-[100%] w-[826px]">
		<div class="relative h-[100px] w-[100%] items-center justify-center">
			<!-- <div class="text-[33px] text-red-950">{product_title}</div> -->
            <Input type="text" placeholder="Enter Product Name" class="w-[100%] border-b border-gray-400 border-t-0 border-l-0 border-r-0 rounded-[0px] text-[33px]" />
			<div class="flex">
				<div class="flex items-center">
					<img src={error} alt="Error" class="w-[40px] h-[40px] object-cover object-cover"/>

					<p class="text-[20px] font-bold text-red-500">Error message</p>
				</div>
			</div>
			<!-- <img src={error} alt="" class="w-full rounded-[10px] object-cover object-cover"/> -->
            <!-- <div class="flex bg-red-400">error message</div> -->
		</div>

		<div class="flex h-[100%] w-[100%]">
			<div class="relative h-[100%] w-[500px]">
                <!-- <div class = "flex"> -->

                <div class="flex mt-100 h-[250px] w-[250px] rounded-[10px] bg-gray-300 items-center justify-center">
					<!-- <div class="w-full rounded-[10px] bg-gray-300" /> -->
					<button class="h-[45px] py-[3px] rounded-[20px] px-[50px] bg-gray-400 text-white shadow-md">Upload Image</button>

                </div>
				<div class="flex items-center">
					<img src={error} alt="Error" class="w-[40px] h-[40px] object-cover object-cover"/>

					<p class="text-[20px] font-bold text-red-500">Error message</p>
				</div>

				<div class ="flex items-center">
					
					<div class="text-[30px] font-bold text-red-950 py-[20px]">NT$</div>
					<Input type="text" placeholder="Enter price" class="ml-[10px] max-w-xs w-[60%] border-b border-gray-400 border-t-0 border-l-0 border-r-0 rounded-[0px] text-[30px]" />

				</div>
				<div class="w-[250px] overflow-hidden text-justify text-gray-600 text-[15px]">
					<!-- {product_contect}
					 -->
					<Input type="text" placeholder="Enter text" class="max-w-xs border-b border-gray-400 border-t-0 border-l-0 border-r-0 rounded-[0px]" />

					<br />
					<br />
					<br />
					<!-- 暫且先放一點br，不然版面好怪 -->
				</div>
			</div>
			<div class="relative h-[100%] w-[100%] ">
				<div class="relative w-[100%] mb-[10px] ">
					{#each options_con as { val, label ,group}}
					<div class=" mb-[15px]">
						<div
							class="flex h-[30px] w-[100%] items-center border-b-[1px] border-solid border-red-950">
							<div class=" font-bold text-amber-900">{label}</div>
						</div>
						<div class="relative ml-[25px] w-[90%] flex-col items-start space-y-2 ">
							{#each opt[val] as { id, label }}
								<!--have some bug  -->
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
								class="mb-4 mr-4 rounded rounded-[10px] border-[3px]
						border-lime-800 bg-transparent px-1 py-0 text-center
						font-bold text-lime-800 transition-all duration-300 hover:bg-lime-800 hover:text-white"
							>
								{label}
							</button>

							{#if (index + 1) % 5 === 0}
								<br />
							{/if}
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

<!-- <Select/> -->
<!-- <Checkbox/> -->
<h1>here is shop list</h1>

<p>Visit <a href="{$page.route.id}/impasta" class="text-blue-400"> Shop </a></p>

<!-- href =>連結 -->

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
