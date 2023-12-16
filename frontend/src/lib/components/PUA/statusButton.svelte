<script lang="ts">
	export let text: string;
	export let id: string;
	export let checked = false;
</script>

<div class=" w-28">
	<input type="radio" name="option" class="peer hidden" {id} {checked} />
	<label
		for={id}
		class="border-PUA-stone text-PUA-stone peer-checked:bg-PUA-stone block cursor-pointer select-none rounded-[20px] border-2 px-2 py-0 text-center font-bold font-bold peer-checked:font-bold peer-checked:text-white"
		>{text}</label
	>
</div>
