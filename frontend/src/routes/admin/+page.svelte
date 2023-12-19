<script lang="ts">
	import { backendPath } from '$lib/components/PUA/env';

	import admin_icon from '$lib/assets/admin_icon.svg';
	import AdminDataCard from '$lib/components/PUA/adminDataCard.svelte';
	import { onMount } from 'svelte';

	type userShort = {
		user_id: number;
		user_name: string;
		user_email: string;
		address: string;
		authority: string;
		status: number;
	}
	type orderShort = {
		store_id: number;
		cart_id: number;
		user_id: number;
		user_name: string;
		order_date: string;
	}


	let userInfo : userShort = {
		user_id : 1,
		user_name : "admin",
		user_email:"aaa",
		address:"aaa",
		authority: 'admin',
		status: 1
	}


	let users : userShort[] = [{
		user_id : 1,
		user_name : "admin",
		user_email:"aaa",
		address:"aaa",
		authority: 'admin',
		status: 1
	}]
	let orders : orderShort[] = [{
		store_id: 1,
		cart_id: 1,
		user_id: 1,
		user_name: "admin",
		order_date: "aa",
	}]

	onMount(async () => {
		getUserInfo()
		getUsers()
		getOrders()
	});

	async function getOrders() {
		const resp = await fetch(backendPath + '/admin/get-all-orders');
		const json = await resp.json();
		orders = json
	}

	async function getUsers() {
		const resp = await fetch(backendPath + '/admin/get-all-users');
		const json = await resp.json();
		users = json
	}

	async function getUserInfo() {
		const resp = await fetch(backendPath + '/user/get-info/1');
		const json = await resp.json();
		userInfo = json
	}

	function switchProfileTab(tabNumber: number) {
		profileTab = tabNumber;
	}

	let profileTab = 0;
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
				class:bg-white={profileTab == 0}
				class:bg-gray-200={profileTab != 0}
				class:border-b-4={profileTab == 0}
				on:click={() => {
					switchProfileTab(0);
				}}
				class="w-full border-PUA-dark-red text-xl text-PUA-dark-red">User List</button
			>
			<button
				class:bg-white={profileTab == 1}
				class:bg-gray-200={profileTab != 1}
				class:border-b-4={profileTab == 1}
				on:click={() => {
					switchProfileTab(1);
				}}
				class="w-full border-PUA-dark-red bg-gray-200 text-xl text-PUA-dark-red">Order List</button
			>
		</div>
		<div class:hidden={profileTab != 0} class="bg-white">
			{#each users as user}
				<AdminDataCard
					firstCol={user.user_name}
					secondCol={user.user_email}
					type={user.authority}
					userID={user.user_id}
					ban={Boolean(user.status)}
				></AdminDataCard>
			{/each}
		</div>
		<div class:hidden={profileTab != 1} class="hidden bg-white">
			{#each orders as order}
				<AdminDataCard firstCol={order.order_date} secondCol={order.user_name} type="order"
				></AdminDataCard>
			{/each}
		</div>
	</div>
</div>
