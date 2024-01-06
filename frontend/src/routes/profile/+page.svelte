<script lang="ts">
	import Casher from '$lib/assets/picturecasher_head.svg';
	import wallet from '$lib/assets/profile_wallet.svg';
	import DenyButton from '$lib/components/PUA/denyButton.svelte';
	import Chart, { type ChartItem } from 'chart.js/auto';
	import { onMount } from 'svelte';
	import { deserialize } from '$app/forms';
	import type { ActionResult } from '@sveltejs/kit';
	import type { PageData } from './$types';
	import { goto } from '$app/navigation';
	import OkButton from '$lib/components/PUA/OkButton.svelte';
	import { logout } from '$lib/components/PUA/logout';
	export let data: PageData;
	console.log(data);

	let choosingMonth = 1;
	let choosingYear = 2020;
	let currentTab = 0;
	function switchProfile() {
		currentTab = 0;
	}
	function switchAccount() {
		currentTab = 1;
	}
	type monthChartType = { product_name: number; product_quantity: number };
	type yearChartType = { price: number };
	let yearChartData: yearChartType[] = [{ price: 100 }];
	let monthChartData: monthChartType[] = [{ product_name: 2010, product_quantity: 10 }];
	async function handleSubmit(event: Event) {
		event.preventDefault();
		await UpdateMonthData();
		await UpdateYearData();
	}
	async function UpdateYearData() {
		const formData = new FormData(document.getElementById('stasticForm') as HTMLFormElement);
		const resp = await fetch('?/yearStatistic', {
			method: 'POST',
			body: formData
		});
		const data: ActionResult = deserialize(await resp.text());

		if (data.type === 'success') {
			yearChartData = data.data as yearChartType[];
			yearChart.data.datasets = [
				{
					label: 'product quantity',
					data: yearChartData.map((row) => row.price),
					//set color
					borderColor: '#822E2E',
					backgroundColor: '#822E2E'
				}
			];
			yearChart.data.labels = monthText;
			yearChart.update();
		}
	}
	async function UpdateMonthData() {
		const formData = new FormData(document.getElementById('stasticForm') as HTMLFormElement);
		const resp = await fetch('?/monthStatistic', {
			method: 'POST',
			body: formData
		});
		const data: ActionResult = deserialize(await resp.text());

		if (data.type === 'success') {
			monthChartData = data.data as monthChartType[];
			monthChart.data.datasets = [
				{
					label: 'product quantity',
					data: monthChartData.map((row) => row.product_quantity),
					borderColor: '#822E2E',
					backgroundColor: '#822E2E',
					maxBarThickness: 50
				}
			];
			monthChart.data.labels = monthChartData.map((row) => row.product_name);
			monthChart.update();
		}
	}
	let year = [2020, 2021, 2022, 2023];
	let month = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12];
	let monthText = [
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
		'十二月'
	];
	let monthChart: Chart;
	let yearChart: Chart;
	function initChart() {
		const mChart = document.getElementById('monthChart') as ChartItem;
		const yChart = document.getElementById('yearChart') as ChartItem;

		if (monthChart) {
			monthChart.destroy();
		}
		monthChart = new Chart(mChart, {
			type: 'bar',
			data: {
				labels: monthChartData.map((row) => row.product_name),
				datasets: [
					{
						label: 'product quantity',
						data: monthChartData.map((row) => row.product_quantity)
					}
				]
			}
		});
		yearChart = new Chart(yChart, {
			type: 'line',
			data: {
				labels: monthText,
				datasets: [
					{
						label: 'product quantity',
						data: yearChartData.map((row) => row.price),
						borderColor: '#36A2EB',
						backgroundColor: '#FFFFFF'
					}
				]
			}
		});
	}

	type userShort = {
		user_id: number;
		user_name: string;
		user_email: string;
		address: string;
		phone: string;
		birthday: string;
		password: string;
		authority: string;
	};

	let userInfo: userShort = {
		user_id: 1,
		user_name: 'admin',
		user_email: 'aaa',
		address: 'aaa',
		phone: '0912345678',
		birthday: '2021-05-01',
		password: '123456',
		authority: 'customer'
	};

	onMount(async () => {
		userInfo = data.userInfo;
		initChart();
		UpdateMonthData();
		UpdateYearData();
	});
</script>

<div class="flex justify-center">
	<div class="flex min-h-screen w-full flex-col lg:w-3/5">
		<div class=" flex items-center gap-7 bg-PUA-stone px-11 py-6">
			{#if userInfo.authority === 'store'}
				<img src={Casher} alt="" class="" />
			{:else if userInfo.authority === 'customer'}
				<img src={wallet} alt="" class="" />
			{/if}
			<div class=" flex grow flex-col gap-6">
				<div class="flex flex-col">
					<div class=" text-2xl font-bold text-PUA-gray">{userInfo.user_name}</div>
					<div class=" text-base font-bold text-PUA-dark-gray">{userInfo.user_email}</div>
				</div>
				<div class=" text-xl font-bold text-PUA-gray">{userInfo.address}</div>
			</div>
			<div class="flex h-full w-full flex-wrap items-center justify-end gap-2">
				{#if userInfo.authority === 'store'}
					<div class="h-fit w-fit rounded-full bg-PUA-gray">
						<DenyButton
							onclick={() => {
								goto('/order/updateOrderStatus');

								return null;
							}}>process order</DenyButton
						>
					</div>

					<div class="h-fit w-fit rounded-full bg-PUA-gray">
						<DenyButton
							onclick={() => {
								goto('/shops/' + userInfo.user_id.toString() + '/store_page_seller');

								return null;
							}}>go my store</DenyButton
						>
					</div>
				{/if}
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
				{#if userInfo.authority === 'store'}
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
				{/if}
			</div>
		</div>

		<div class="grow bg-gray-100">
			<div class:hidden={currentTab != 0} class="flex w-full flex-col gap-7 px-24 py-10">
				<div class="">
					<div class="inline-flex items-center justify-center rounded-full bg-gray-300 px-5">
						<div class="text-base font-bold leading-tight text-PUA-dark-orange">User Name</div>
					</div>
					<div class="flex font-['Inter'] text-2xl font-bold text-PUA-dark-orange">
						{userInfo.user_name}
					</div>
				</div>
				<div class="">
					<div class="inline-flex items-center justify-center gap-5 rounded-full bg-gray-300 px-5">
						<div class=" font-['Inter'] text-base font-bold leading-tight text-PUA-dark-orange">
							Phone
						</div>
					</div>
					<div class="flex bg-inherit font-['Inter'] text-2xl font-bold text-PUA-dark-orange">
						{userInfo.phone}
					</div>
				</div>
				<div class="">
					<div class="inline-flex items-center justify-center gap-5 rounded-full bg-gray-300 px-5">
						<div class=" font-['Inter'] text-base font-bold leading-tight text-PUA-dark-orange">
							Birthday
						</div>
					</div>
					<div class="flex bg-inherit font-['Inter'] text-2xl font-bold text-PUA-dark-orange">
						{userInfo.birthday}
					</div>
				</div>
				<div class="flex justify-center gap-32 px-28 pt-20">
					<OkButton
						text="Logout"
						onclick={() => {
							logout();
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
					<div class="flex bg-inherit font-['Inter'] text-2xl font-bold text-PUA-dark-orange">
						{userInfo.user_email}
					</div>
				</div>
				<div class="py-5">
					<div class="inline-flex items-center justify-center gap-5 rounded-full bg-gray-300 px-5">
						<div class=" font-['Inter'] text-base font-bold leading-tight text-PUA-dark-orange">
							Password
						</div>
					</div>
					<div class="flex bg-inherit font-['Inter'] text-2xl font-bold text-PUA-dark-orange">
						{userInfo.password}
					</div>
				</div>
			</div>
			<div class:hidden={currentTab != 2} class="hidden w-full p-20">
				<form
					action="?/statistic"
					id="stasticForm"
					method="post"
					class="flex flex-col gap-2"
					on:submit|preventDefault={handleSubmit}
				>
					<div class="flex flex-wrap gap-2">
						{#each year as key}
							<input
								type="radio"
								bind:group={choosingYear}
								name={'year'}
								id={'year' + key}
								value={key}
								class="hidden"
								on:change={handleSubmit}
							/>
							<label
								for={'year' + key}
								class:bg-PUA-dark-red={choosingYear === key}
								class:bg-gray-100={!(choosingYear === key)}
								class:text-white={choosingYear === key}
								class:text-PUA-dark-red={!(choosingYear === key)}
								class="leading-relaxe h-10 rounded-full border-2 border-PUA-dark-red px-6 text-2xl font-bold"
							>
								{key}</label
							>
						{/each}
					</div>
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
								{monthText[key - 1]}</label
							>
						{/each}
					</div>
				</form>
				<canvas id="monthChart"></canvas>
				<canvas id="yearChart"></canvas>
			</div>
		</div>
	</div>
</div>
