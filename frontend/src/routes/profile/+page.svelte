<script lang="ts">
	import wallet from '$lib/assets/profile_wallet.svg';

	let currentTab = 0;
	function switchProfile() {
		if (currentTab == 1) {
			document.getElementById('account-tab')?.classList.add('hidden');
			document.getElementById('profile-tab')?.classList.remove('hidden');
			document.getElementById('tab-underline')?.classList.replace('right-0', 'left-0');
			document.getElementById('profileButton')?.classList.replace('bg-gray-300', 'bg-white');
			document.getElementById('accountButton')?.classList.replace('bg-white', 'bg-gray-300');
		}
		currentTab = 0;
	}
	function switchAccount() {
		if (currentTab == 0) {
			document.getElementById('profile-tab')?.classList.add('hidden');
			document.getElementById('account-tab')?.classList.remove('hidden');
			document.getElementById('tab-underline')?.classList.replace('left-0', 'right-0');
			document.getElementById('accountButton')?.classList.replace('bg-gray-300', 'bg-white');
			document.getElementById('profileButton')?.classList.replace('bg-white', 'bg-gray-300');
		}
		currentTab = 1;
	}
</script>

<div class="flex justify-center">
	<div class="flex w-3/5 flex-col">
		<div class=" bg-PUA-stone">
			<div class="flex">
				<img src={wallet} alt="" class="my-6 ml-10" />
				<div class="flex-col pl-5 pt-7">
					<div class="font-['Inter'] text-2xl font-bold text-PUA-gray">Usernamenamename</div>
					<div class="font-['Inter'] text-base font-bold text-PUA-gray">t110595959@rice.org.tw</div>
					<div class="pt-5 font-['Inter'] text-xl font-bold text-PUA-gray">
						100台灣台北市中正區八德路一段82巷9弄17號
					</div>
				</div>
			</div>
		</div>

		<div class="relative bg-PUA-dark-gray">
			<div class="flex">
				<button
					id="profileButton"
					class="w-1/2 bg-white py-3 text-center font-['Inter'] text-2xl font-bold text-PUA-red"
					on:click={switchProfile}>Profile</button
				>
				<button
					id="accountButton"
					class="w-1/2 bg-gray-300 py-3 text-center font-['Inter'] text-2xl font-bold text-PUA-red"
					on:click={switchAccount}>Account</button
				>
			</div>
			<div class="absolute bottom-0 left-0 h-1 w-1/2 bg-PUA-red" id="tab-underline"></div>
		</div>

		<div class="min-h-screen bg-white">
			<div class="flex-col pl-28 pt-10" id="profile-tab">
				<div class="py-5">
					<div class="inline-flex items-center justify-center gap-5 rounded-full bg-gray-300 px-5">
						<div class=" font-['Inter'] text-base font-bold leading-tight text-PUA-dark-orange">
							User Name
						</div>
					</div>
					<input
						type="text"
						value="Usernamename"
						class="flex bg-inherit font-['Inter'] text-2xl font-bold text-PUA-dark-orange underline focus:outline-none"
					/>
				</div>
				<div class="py-5">
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
				<div class="py-5">
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

			<div class="hidden flex-col pl-28 pt-10" id="account-tab">
				<div class="py-5">
					<div class="inline-flex items-center justify-center gap-5 rounded-full bg-gray-300 px-5">
						<div class=" font-['Inter'] text-base font-bold leading-tight text-PUA-dark-orange">
							Email
						</div>
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
		</div>
	</div>
</div>
