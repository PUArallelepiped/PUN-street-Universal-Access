<script lang="ts">
	export let href: string = '#';
</script>

<a {href}>
	<button
		on:click
		class="bg-PUA-dark-red hover:border-PUA-dark-red hover:text-PUA-dark-red flex h-7 w-7 items-center justify-center rounded-full text-center text-xl font-bold text-white duration-200 hover:scale-125 hover:border-2 hover:bg-gray-200"
	>
		+
	</button>
</a>
