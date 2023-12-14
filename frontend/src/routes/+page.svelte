<script lang="ts">
	import Chart, { type ChartItem } from 'chart.js/auto';
	import { onMount } from 'svelte';
	onMount(() => {
		const data = [
			{ year: 2010, count: 10 },
			{ year: 2011, count: 20 },
			{ year: 2012, count: 15 },
			{ year: 2013, count: 25 },
			{ year: 2014, count: 22 },
			{ year: 2015, count: 30 },
			{ year: 2016, count: 28 }
		];

		console.log(document.getElementById('acquisitions'));
		const test = document.getElementById('acquisitions') as ChartItem;

		new Chart(test, {
			type: 'bar',
			data: {
				labels: data.map((row) => row.year),
				datasets: [
					{
						label: 'Acquisitions by year',
						data: data.map((row) => row.count)
					}
				]
			}
		});
	});
</script>

<div style="width: 800px;"><canvas id="acquisitions"></canvas></div>
<a href="/cart" class="m-4 rounded-lg border-2 border-blue-600 p-1 text-lg text-blue-600">cart</a>
<a href="/history" class="m-4 rounded-lg border-2 border-blue-600 p-1 text-lg text-blue-600">
	history
</a>
<a
	href="/order/checkOrder"
	class="m-4 rounded-lg border-2 border-blue-600 p-1 text-lg text-blue-600"
>
	checkOrder
</a>
<a
	href="/order/updateOrderStatus"
	class="m-4 rounded-lg border-2 border-blue-600 p-1 text-lg text-blue-600"
>
	updateOrderStatus
</a>
<a href="/shops" class="m-4 rounded-lg border-2 border-blue-600 p-1 text-lg text-blue-600">
	shoplist
</a>

<a href="/shops/1" class="m-4 rounded-lg border-2 border-blue-600 p-1 text-lg text-blue-600">
	shop
</a>
<a href="/shops/1/1" class="m-4 rounded-lg border-2 border-blue-600 p-1 text-lg text-blue-600">
	item
</a>
<a
	href="/shops/1/addItem"
	class="m-4 rounded-lg border-2 border-blue-600 p-1 text-lg text-blue-600"
>
	add item
</a>
<a href="/login" class="m-4 rounded-lg border-2 border-blue-600 p-1 text-lg text-blue-600">
	login
</a>

<a href="/login/register" class="m-4 rounded-lg border-2 border-blue-600 p-1 text-lg text-blue-600">
	register
</a>
