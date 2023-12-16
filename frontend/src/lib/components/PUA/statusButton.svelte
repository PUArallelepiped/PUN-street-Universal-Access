<script lang="ts">
	export let text: string;
	export let id: string;
	export let checked = false;
</script>

<div class=" w-28">
	<input type="radio" name="option" class="peer hidden" {id} {checked} />
	<label
		for={id}
		class="block cursor-pointer select-none rounded-[20px] border-2 border-PUA-stone px-2 py-0 text-center font-bold font-bold text-PUA-stone peer-checked:bg-PUA-stone peer-checked:font-bold peer-checked:text-white"
		>{text}</label
	>
</div>
