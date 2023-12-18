<script lang="ts">
	import wallet from '$lib/assets/profile_wallet.svg';
	import OkButton from '$lib/components/PUA/OkButton.svelte';
	import DenyButton from '$lib/components/PUA/denyButton.svelte';
	import Chart, { type ChartItem } from 'chart.js/auto';
	import { onMount } from 'svelte';
	import type { PageData } from '../$types';
	import type { ActionData } from './$types';
	import { applyAction, deserialize } from '$app/forms';
	import type { ActionResult } from '@sveltejs/kit';

	let choosingMonth = 1;
	let currentTab = 0;
	function switchProfile() {
		currentTab = 0;
	}
	function switchAccount() {
		currentTab = 1;
	}
	// export let data: PageData;
	export let form: ActionData;
	type chartType = { product_name: number; product_quantity: number };
	let chartData: chartType[] = [
		{ product_name: 2010, product_quantity: 10 },
		{ product_name: 2010, product_quantity: 10 },
		{ product_name: 2010, product_quantity: 10 },
		{ product_name: 2010, product_quantity: 10 },
		{ product_name: 2010, product_quantity: 10 },
		{ product_name: 2010, product_quantity: 10 },
		{ product_name: 2010, product_quantity: 10 },
		{ product_name: 2010, product_quantity: 10 },
		{ product_name: 2010, product_quantity: 10 }
	];
	async function handleSubmit(event) {
		event.preventDefault();
		const formData = new FormData(document.getElementById('monthRadio') as HTMLFormElement);
		const resp = await fetch('?/statistic', {
			method: 'POST',
			body: formData
		});
		const data: ActionResult = deserialize(await resp.text());

		if (data.type === 'success') {
			console.log(data.data);
			chartData = data.data as chartType[];
			chart.data.datasets = [
				{
					label: 'Acquisitions by year',
					data: chartData.map((row) => row.product_quantity)
				}
			];
			chart.data.labels = chartData.map((row) => row.product_name);
			chart.update();
		}
	}
	let month = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12];
	let monthText = [
		'for index align',
		'一月',
		'二月',
		'三月',
		'四月',
		'五月',
		'六月',
		'七月',
		'八月',
		'九月',
		'十月',
		'十一月',
		'十二月',
		'十三月'
	];
	let chart: Chart;
	function initChart() {
		console.log(document.getElementById('acquisitions'));
		const test = document.getElementById('acquisitions') as ChartItem;

		if (chart) {
			chart.destroy();
		}
		chart = new Chart(test, {
			type: 'bar',
			data: {
				labels: chartData.map((row) => row.product_name),
				datasets: [
					{
						label: 'Acquisitions by year',
						data: chartData.map((row) => row.product_quantity)
					}
				]
			}
		});
	}
	onMount(() => {
		initChart();
	});
</script>

<div class="flex justify-center">
	<div class="flex min-h-screen w-full flex-col lg:w-3/5">
		<div class=" flex items-center gap-7 bg-PUA-stone px-11 py-6">
			<img src={wallet} alt="" class="" />
			<div class=" flex grow flex-col gap-6">
				<div class="flex flex-col">
					<div class=" text-2xl font-bold text-PUA-gray">Usernamenamename</div>
					<div class=" text-base font-bold text-PUA-dark-gray">t110595959@rice.org.tw</div>
				</div>
				<div class=" text-xl font-bold text-PUA-gray">100台灣台北市中正區八德路一段82巷9弄17號</div>
			</div>
			<div class="h-fit w-fit rounded-full bg-PUA-gray">
				<DenyButton
					onclick={() => {
						return null;
					}}>go my store</DenyButton
				>
			</div>
		</div>
		<div class=" bg-PUA-dark-gray">
			<div class="flex">
				<div class="grow">
					<button
						class:bg-white={currentTab == 0}
						class:bg-gray-300={currentTab != 0}
						class=" w-full bg-white py-3 text-center text-2xl font-bold text-PUA-dark-red"
						on:click={switchProfile}>Profile</button
					>
					<div
						class:bg-PUA-dark-red={currentTab == 0}
						class:bg-gray-300={currentTab != 0}
						class=" h-1 w-full bg-gray-300"
					></div>
				</div>
				<div class="grow">
					<button
						class:bg-white={currentTab == 1}
						class:bg-gray-300={currentTab != 1}
						class="w-full bg-gray-300 py-3 text-center text-2xl font-bold text-PUA-dark-red"
						on:click={switchAccount}>Account</button
					>
					<div
						class:bg-PUA-dark-red={currentTab == 1}
						class:bg-gray-300={currentTab != 1}
						class=" h-1 w-full"
					></div>
				</div>
				<div class="grow">
					<button
						class:bg-white={currentTab == 2}
						class:bg-gray-300={currentTab != 2}
						class="w-full bg-gray-300 py-3 text-center text-2xl font-bold text-PUA-dark-red"
						on:click={() => {
							currentTab = 2;
						}}>Statistic</button
					>
					<div
						class:bg-PUA-dark-red={currentTab == 2}
						class:bg-gray-300={currentTab != 2}
						class=" h-1 w-full"
					></div>
				</div>
			</div>
		</div>

		<div class="grow bg-gray-100">
			<div class:hidden={currentTab != 0} class="flex w-full flex-col gap-7 px-24 py-10">
				<div class="">
					<div class="inline-flex items-center justify-center rounded-full bg-gray-300 px-5">
						<div class="text-base font-bold leading-tight text-PUA-dark-orange">User Name</div>
					</div>
					<input
						type="text"
						value="Usernamename"
						class="flex text-2xl font-bold text-PUA-dark-orange underline focus:outline-none"
					/>
				</div>
				<div class="">
					<div class="inline-flex items-center justify-center gap-5 rounded-full bg-gray-300 px-5">
						<div class=" font-['Inter'] text-base font-bold leading-tight text-PUA-dark-orange">
							Phone
						</div>
					</div>
					<input
						type="text"
						value="09-XXXX-XXXX"
						class="flex bg-inherit font-['Inter'] text-2xl font-bold text-PUA-dark-orange underline focus:outline-none"
					/>
				</div>
				<div class="">
					<div class="inline-flex items-center justify-center gap-5 rounded-full bg-gray-300 px-5">
						<div class=" font-['Inter'] text-base font-bold leading-tight text-PUA-dark-orange">
							Birthday
						</div>
					</div>
					<input
						type="text"
						value="2021-05-01"
						class="flex bg-inherit font-['Inter'] text-2xl font-bold text-PUA-dark-orange underline focus:outline-none"
					/>
				</div>
				<div class="flex justify-center gap-32">
					<DenyButton
						onclick={() => {
							return null;
						}}
					>
						Clear
					</DenyButton>
					<OkButton
						text="Save Change"
						onclick={() => {
							return null;
						}}
					></OkButton>
				</div>
			</div>

			<div class:hidden={currentTab != 1} class="flex w-full flex-col gap-7 px-24 py-10">
				<div class="py-5">
					<div class="inline-flex items-center justify-center gap-5 rounded-full bg-gray-300 px-5">
						<div class="  text-base font-bold leading-tight text-PUA-dark-orange">Email</div>
					</div>
					<input
						type="text"
						value="t110595959@rice.org.tw"
						class="flex bg-inherit font-['Inter'] text-2xl font-bold text-PUA-dark-orange underline focus:outline-none"
					/>
				</div>
				<div class="py-5">
					<div class="inline-flex items-center justify-center gap-5 rounded-full bg-gray-300 px-5">
						<div class=" font-['Inter'] text-base font-bold leading-tight text-PUA-dark-orange">
							Password
						</div>
					</div>
					<input
						type="text"
						value="●●●●●●●●●●"
						class="flex bg-inherit font-['Inter'] text-2xl font-bold text-PUA-dark-orange underline focus:outline-none"
					/>
				</div>
				<div class="flex gap-32 px-28 pt-5">
					<button
						class="rounded-full bg-gray-300 px-14 font-['Inter'] text-xl font-bold text-PUA-dark-orange"
						>Cancel</button
					>
					<button
						class="rounded-full bg-PUA-orange px-10 font-['Inter'] text-xl font-bold text-white"
						>Save change</button
					>
				</div>
			</div>
			<div class:hidden={currentTab != 2} class="hidden w-full p-20">
				<form
					action="?/statistic"
					id="monthRadio"
					method="post"
					on:submit|preventDefault={handleSubmit}
				>
					<div class="flex flex-wrap gap-2">
						{#each month as key}
							<input
								type="radio"
								bind:group={choosingMonth}
								name={'month'}
								id={'check' + key}
								value={key}
								class="hidden"
								on:change={handleSubmit}
							/>
							<label
								for={'check' + key}
								class:bg-PUA-dark-red={choosingMonth === key}
								class:bg-gray-100={!(choosingMonth === key)}
								class:text-white={choosingMonth === key}
								class:text-PUA-dark-red={!(choosingMonth === key)}
								class="leading-relaxe h-10 rounded-full border-2 border-PUA-dark-red px-3 text-2xl font-bold"
							>
								{monthText[key]}</label
							>
						{/each}
					</div>
				</form>
				<canvas id="acquisitions"></canvas>
			</div>
		</div>
	</div>
</div>
