<script lang="ts">
	import { backendPath } from '$lib/components/PUA/env';

	import admin_icon from '$lib/assets/admin_icon.svg';
	import AdminDataCard from '$lib/components/PUA/adminDataCard.svelte';

	async function getOrders() {
		const resp = await fetch(backendPath + '/admin/get-all-orders');
		const json = await resp.json();
		return json;
	}

	async function getUsers() {
		const resp = await fetch(backendPath + '/admin/get-all-users');
		const json = await resp.json();
		return json;
	}

	async function getUserInfo() {
		const resp = await fetch(backendPath + '/user/get-info/1');
		const json = await resp.json();
		return json;
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
				{#await getUserInfo() then userInfo}
					<div>
						<div class="text-2xl text-PUA-gray">{userInfo.user_name}</div>
						<div class="text-gray-300">{userInfo.user_email}</div>
					</div>
					<div class="text-xl text-PUA-gray">{userInfo.address}</div>
				{/await}
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
			{#await getUsers() then users}
				{#each users as user}
					<AdminDataCard
						firstCol={user.user_name}
						secondCol={user.user_email}
						type={user.authority}
						userID={user.user_id}
						ban={user.status}
					></AdminDataCard>
				{/each}
			{/await}
		</div>
		<div class:hidden={profileTab != 1} class="hidden bg-white">
			{#await getOrders() then orders}
				{#each orders as order}
					<AdminDataCard firstCol={order.order_date} secondCol={order.user_name} type="0"
					></AdminDataCard>
				{/each}
			{/await}
		</div>
	</div>
</div>
