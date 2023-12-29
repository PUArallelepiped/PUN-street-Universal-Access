<script lang="ts">
	export let href: string = '#';
</script>

<a {href}>
	<button
		on:click
		class="flex h-7 w-7 items-center justify-center rounded-full bg-PUA-dark-red text-center text-xl font-bold text-white duration-200 hover:scale-125 hover:border-2 hover:border-PUA-dark-red hover:bg-gray-200 hover:text-PUA-dark-red"
	>
		+
	</button>
</a>
