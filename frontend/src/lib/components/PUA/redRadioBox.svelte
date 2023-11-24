<script lang="ts">
	export let name: string = 'null';
	export let value: string = 'null';
</script>

<input type="radio" {name} {value} class="h-[20px] w-[20px]" />

<style>
	input[type='radio'] {
		accent-color: rgb(120, 40, 40);
	}
</style>
