<script lang="ts">
	import admin_icon from '$lib/assets/admin_icon.svg';
	import AdminDataCard from '$lib/components/PUA/adminDataCard.svelte';

	let orders: {
		date: string;
		user: string;
	}[] = [
		{
			date: '2020-10-11 19:20:50',
			user: 'user01'
		},
		{
			date: '2020-10-11 19:20:50',
			user: 'user02'
		},
		{
			date: '2020-10-11 19:20:50',
			user: 'user03'
		},
		{
			date: '2020-10-11 19:20:50',
			user: 'user04'
		}
	];

	let users: {
		name: string;
		email: string;
		type: string;
		status: boolean;
	}[] = [
		{
			name: 'sakana',
			email: 'r1111111',
			type: '1',
			status: false
		},
		{
			name: 'sakana dewanai',
			email: 't1111111',
			type: '2',
			status: false
		},
		{
			name: 'sakana desu',
			email: 'fsddsdss',
			type: '2',
			status: true
		}
	];

	let userInfo = {
		user_name: 'sakana san',
		user_email: 't111059001',
		address: 'dsdsadsadasdsadasd',
		authority: 100
	};

	function changeDisplay(display_id_name: string, invisible_id_name: string) {
		if (display_id_name == 'userData') {
			document.getElementById('userButton')?.classList.add('bg-white');
			document.getElementById('userButton')?.classList.add('border-b-4');
			document.getElementById('orderButton')?.classList.remove('bg-white');
			document.getElementById('orderButton')?.classList.remove('border-b-4');
		} else {
			document.getElementById('orderButton')?.classList.add('bg-white');
			document.getElementById('orderButton')?.classList.add('border-b-4');
			document.getElementById('userButton')?.classList.remove('bg-white');
			document.getElementById('userButton')?.classList.remove('border-b-4');
		}

		document.getElementById(display_id_name)?.classList.remove('hidden');
		document.getElementById(invisible_id_name)?.classList.add('hidden');
	}
</script>

<div class="flex justify-center font-bold">
	<div class="flex w-3/5 flex-col">
		<div class="flex bg-PUA-stone">
			<img src={admin_icon} alt="" class="my-6 ml-10 flex h-28 w-28" />
			<div class="m-7 flex flex-col justify-between">
				<div>
					<div class="text-2xl text-PUA-gray">{userInfo.user_name}</div>
					<div class="text-gray-300">{userInfo.user_email}</div>
				</div>
				<div class="text-xl text-PUA-gray">{userInfo.address}</div>
			</div>
		</div>
		<div class="flex h-10">
			<button
				id="userButton"
				on:click={() => changeDisplay('userData', 'orderData')}
				class="w-full border-b-4 border-PUA-dark-red bg-gray-200 bg-white text-xl text-PUA-dark-red"
				>User List</button
			>
			<button
				id="orderButton"
				on:click={() => changeDisplay('orderData', 'userData')}
				class="w-full border-PUA-dark-red bg-gray-200 text-xl text-PUA-dark-red">Order List</button
			>
		</div>
		<div id="userData" class="bg-white">
			{#each users as user}
				<AdminDataCard
					firstCol={user.name}
					secondCol={user.email}
					type={user.type}
					ben={user.status}
				></AdminDataCard>
			{/each}
		</div>
		<div id="orderData" class="hidden bg-white">
			{#each orders as order}
				<AdminDataCard firstCol={order.date} secondCol={order.user} type="0"></AdminDataCard>
			{/each}
		</div>
	</div>
</div>
