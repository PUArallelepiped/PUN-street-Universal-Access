<script lang="ts">
	const routes = [
		{ path: '/', title: 'My profile' }, // TODO
		{ path: '/shops', title: 'ShopList' },
		{ path: '/', title: 'Check Order' }, // TODO
		{ path: '/', title: 'History' } // TODO
	];
</script>

<div class="flex h-20 place-content-between bg-white shadow-md">
	<a
		href="/"
		class="flex h-full items-center justify-start pl-2 font-serif text-4xl font-medium lg:px-11"
	>
		<div class="text-red-950">PU</div>
		<div class="text-orange-500">N street</div>
		<div class="text-red-950">&nbspa</div>
		<div class="text-orange-500">ccess</div>
	</a>
	<div class="flex h-full">
		<div
			class="invisible flex h-full w-0 items-center justify-end gap-10 lg:visible lg:w-max lg:px-10"
		>
			{#each routes as r}
				<a href={r.path} class="block p-2">{r.title}</a>
			{/each}
		</div>
		<a href="/cart" class="m-3 flex items-center rounded-xl bg-red-800 p-3 text-white">
			Cart
			<svg
				width="25"
				height="20"
				viewBox="0 0 25 20"
				fill="none"
				xmlns="http://www.w3.org/2000/svg"
			>
				<path
					d="M1.33263 0.00676144C0.918149 0.0774619 0.546992 0.320802 0.300805 0.68325C0.0546178 1.0457 -0.0464324 1.49756 0.019884 1.93944C0.0862003 2.38132 0.314451 2.77701 0.654423 3.03948C0.994395 3.30194 1.41824 3.40967 1.83272 3.33897H6.52108L6.80238 4.17202L8.08387 8.33728L9.36536 12.5025C9.49038 12.9357 10.0217 13.3356 10.4281 13.3356H21.3676C21.8051 13.3356 22.3052 12.9357 22.4303 12.5025L24.962 4.17202C25.087 3.73883 24.8995 3.33897 24.4619 3.33897H10.5843L9.39661 0.939779C9.27011 0.664403 9.07478 0.431586 8.83213 0.266965C8.58948 0.102344 8.30892 0.0123046 8.02136 0.00676144L1.77021 0.00676144C1.67663 -0.00225381 1.58248 -0.00225381 1.4889 0.00676144C1.42645 0.00275918 1.36383 0.00275918 1.30137 0.00676144L1.33263 0.00676144ZM11.2094 16.6678C10.3343 16.6678 9.64666 17.4009 9.64666 18.3339C9.64666 19.2669 10.3343 20 11.2094 20C12.0846 20 12.7722 19.2669 12.7722 18.3339C12.7722 17.4009 12.0846 16.6678 11.2094 16.6678ZM20.5862 16.6678C19.711 16.6678 19.0234 17.4009 19.0234 18.3339C19.0234 19.2669 19.711 20 20.5862 20C21.4613 20 22.149 19.2669 22.149 18.3339C22.149 17.4009 21.4613 16.6678 20.5862 16.6678Z"
					fill="white"
					fill-opacity="0.94"
				/>
			</svg>
		</a>
	</div>
</div>
