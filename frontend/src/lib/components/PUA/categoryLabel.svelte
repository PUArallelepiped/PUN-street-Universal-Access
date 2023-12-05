<script lang="ts">
	import Dropdown from '$lib/assets/dropdown.svg';
	export let text = 'null';
</script>

<div class="boder-b-1 mb-4 mt-4 flex justify-between border-b-[1px] border-b-PUA-stone">
	<div class="mx-5 font-bold text-PUA-stone">{text}</div>
	<img src={Dropdown} alt="" />
</div>
