<script lang="ts">
	export let text: string = 'null';
	export let id: number = 0;
	export let value: number = 0;
	export let group: number = 0;
</script>

<div class="">
	<input type="radio" name="option" class="peer hidden" id={id.toString()} {value} bind:group />
	<label
		for={id.toString()}
		class="block cursor-pointer select-none rounded-[10px] border-2 border-lime-800 px-2 py-0 text-center font-bold text-lime-800 peer-checked:bg-lime-800 peer-checked:font-bold peer-checked:text-white"
		>{text}</label
	>
</div>
